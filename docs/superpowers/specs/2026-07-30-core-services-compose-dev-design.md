# 核心服务 Compose 开发栈设计

日期：2026-07-30

## 背景

现有 `deploy/compose/docker-compose.dev.yml` 只管理 ZooKeeper、Kafka、Canal 和
Elasticsearch 等开发依赖，Gateway 与各核心 Go 服务仍需在宿主机逐个启动。

仓库中另有 `docker-compose.full.yml` 与 all-in-one 草案，但它们不符合当前开发边界：

- `full` 会创建 MySQL、Redis、etcd，并包含已封存的 LLM、Agent；
- all-in-one 使用 Supervisor 在单容器中运行多个旧 RPC 入口，已落后于当前 merged
  服务结构；
- 现有 Dockerfile 使用 Go 1.23，而 `go.mod` 要求 Go 1.25.8；
- 容器内 RPC 服务默认向 etcd 注册容器 IP，Windows 宿主机上的压测程序无法直连该
  地址。

本机实测从现有 Kafka 容器可以访问宿主机的 `3306`、`6379`、`2379` 端口，因此可复用
已经启动的 MySQL、Redis、etcd。

## 目标

- 使用一个 Compose 命令构建并启动所有核心应用服务。
- 每个业务域使用独立容器，沿用当前 merged 服务入口。
- 复用宿主机 MySQL、Redis、etcd，不创建或管理它们。
- 保留容器内服务调用和宿主机 RPC 联调两条链路。
- 保证现有 `cmd/loadtest/feed_loadtest.exe` 可继续通过宿主机 etcd 发现 RPC。
- 通过健康检查和依赖条件消除“容器已运行但服务尚不可用”的启动竞态。
- 不构建、不启动 LLM 和 Agent。

## 容器拓扑

核心应用共用一个 `zhiguang-go-core:dev` 镜像，但分别执行独立命令：

| Compose 服务 | 进程内容 | 发布到宿主机的端口 |
| --- | --- | --- |
| `gateway` | Gateway HTTP | `127.0.0.1:8080` |
| `user-storage` | User RPC、Storage RPC | `127.0.0.1:9002`、`127.0.0.1:9013` |
| `counter` | Counter RPC、Aggregator | `127.0.0.1:9003` |
| `knowpost` | KnowPost RPC | `127.0.0.1:9004` |
| `relation` | Relation RPC、Syncer | `127.0.0.1:9006` |
| `search` | Search RPC、Indexer | `127.0.0.1:9017` |

ZooKeeper、Kafka、Canal、Elasticsearch 继续由现有开发 Compose 管理。Adminer 与
Kibana 保持工具属性，不作为核心服务健康的必要条件。

## 核心镜像

新增核心开发 Dockerfile，构建阶段使用满足 `go.mod` 最低要求的 Go 1.25
补丁版本，并只编译：

- `gateway`
- `storage-merged`
- `counter-merged`
- `knowpost-merged`
- `relation-merged`
- `search-merged`

运行阶段只包含这些二进制、对应配置以及健康检查需要的最小工具。LLM、Agent
及其服务端二进制不进入该镜像，本机 `certs/` 也不进入构建上下文或镜像层。

一次性 `jwt-cert-init` 容器在 Compose 命名卷中生成开发用 JWT 密钥。该卷只读挂载到
`user-storage`；Gateway、Counter、KnowPost、Relation、Search 均不会获得私钥。
命名卷不存在时自动生成，后续启动复用已有私钥。

所有应用服务复用相同的镜像定义，避免重复维护 Dockerfile。Compose 仍通过不同
`command` 启动各业务域进程。

## 宿主机基础设施访问

应用容器使用 `extra_hosts` 将下列既有配置主机名映射到 Docker host gateway：

```text
mysql -> host-gateway
redis -> host-gateway
etcd  -> host-gateway
```

因此现有 Docker 配置中的 `mysql:3306`、`redis:6379`、`etcd:2379` 无需复制为一套
Windows 专用配置。Kafka 与 Elasticsearch 仍通过 Compose 默认网络中的服务名访问：

```text
kafka:29092
elasticsearch:9200
```

一个一次性 `host-infra-check` 服务在应用启动前检查宿主机 MySQL、Redis、etcd
端口。检查超时会使 Compose 明确失败，而不是让所有应用进入反复重启。

ZooKeeper、Kafka、Canal、Elasticsearch 以及工具 profile 的宿主机端口统一绑定到
`127.0.0.1`。容器间调用仍走 Compose 网络，不向局域网暴露匿名或明文开发服务。

## RPC 双向可达性

go-zero 在 `ListenOn` 为 `0.0.0.0` 时会自动选择容器内部 IP 并注册到 etcd。该 IP
可供同一 Docker 网络使用，但 Windows 宿主机无法稳定直连，因而会破坏现有
`feed_loadtest.exe -> etcd -> RPC` 链路。

开发 Compose 采用双通道设计：

1. 容器内 RPC 客户端使用 Docker DNS 与固定容器端口直连，例如
   `knowpost:9004`、`counter:9003`；
2. RPC 服务设置 `POD_IP=127.0.0.1`，并将对应端口发布到宿主机；
3. 服务仍向宿主机 etcd 注册，注册值为 `127.0.0.1:<port>`；
4. 宿主机压测程序从 etcd 取得地址后通过已发布端口调用容器内 RPC。

宿主机 etcd 如果仍广播默认的 `http://localhost:2379`，go-zero 每分钟的成员自动同步
会记录一次连接告警；初始连接、租约续期和上述双通道调用不受影响。要消除该告警，
宿主机 etcd 的 `advertise-client-urls` 需配置为容器可达地址。

需要改为容器直连的活跃调用包括：

- Gateway 到 User、Storage、Counter、KnowPost、Relation、Search；
- KnowPost RPC 到 Counter、Relation；
- Relation RPC/Syncer 到 User、Counter；
- Search RPC/Indexer 到 Counter、KnowPost。

既有 LLM 兼容客户端保留非阻塞配置以避免修改封存业务代码，但不构建、不启动 LLM
服务端，也不注册 LLM 实例；LLM 路由不属于本次健康验收范围。

## 启动顺序与健康检查

依赖顺序如下：

1. `host-infra-check` 等待宿主机 MySQL、Redis、etcd 可连接，同时
   `jwt-cert-init` 准备开发 JWT 密钥；
2. ZooKeeper、Kafka、Canal 与 Elasticsearch 达到健康状态；
3. JWT 密钥准备完成后启动 User/Storage，同时启动 Counter；
4. 启动 Relation；
5. Relation 健康后启动 KnowPost，随后启动 Search；
6. 所有核心 RPC 端口健康后启动 Gateway。

应用健康检查使用容器内 TCP 探测：

- `user-storage` 同时检查 `9002` 与 `9013`；
- 其他 RPC 服务检查各自 RPC 端口；
- Gateway 检查 `8080`。
- Canal 在 Kafka 模式下不以 `11111` 是否监听作为就绪条件；检查 Java 主进程、
  Canal Server 启动完成日志和 `example` instance 启动成功日志。

所有长期运行服务使用 `restart: unless-stopped`。Compose 的
`depends_on.condition: service_healthy` 负责首次启动排序；应用自身仍需保留正常的
连接错误和重连行为。

标准启动命令：

```powershell
docker compose -f deploy/compose/docker-compose.dev.yml up -d --build --wait --wait-timeout 180
```

## 错误处理

- 宿主机端口不可达：`host-infra-check` 输出具体端口并失败。
- Kafka 或 Elasticsearch 未就绪：依赖服务不启动，Compose `--wait` 返回失败。
- 核心 RPC 主进程退出：`restart: unless-stopped` 负责重启。
- 主进程仍在但端口未监听：容器标记为 unhealthy，首次 `--wait` 启动会失败；Docker
  不会仅因 unhealthy 自动重启，需要查看日志后手动重启或修复。
- 宿主机端口被其他进程占用：Docker 在创建容器时直接报告端口冲突。
- LLM/Agent 缺失：不影响核心容器健康；相关封存接口不纳入验收。
- 停止开发栈：只停止 Compose 管理的容器，不操作宿主机 MySQL、Redis、etcd。

## 测试与验收

实施按以下顺序验证：

1. 为 Compose 服务集合、镜像构建范围、宿主机映射、RPC 直连与端口发布编写静态
   契约测试，并先确认测试因功能缺失而失败。
2. `docker compose config` 必须无错误，且默认服务中不包含 LLM、Agent、MySQL、
   Redis、etcd。
3. 核心镜像必须使用不低于 Go 1.25.8 的工具链，并成功构建六个目标二进制；本机
   `certs/` 必须被 `.dockerignore` 排除，最终镜像不得复制 JWT 私钥。
4. 使用标准启动命令等待所有默认服务就绪。
5. 从应用容器验证宿主机 MySQL、Redis、etcd 与容器内 Kafka、Elasticsearch 可达。
6. 检查宿主机 etcd 中的核心 RPC 注册值均为 `127.0.0.1:<port>`。
7. 通过 Gateway 执行核心接口冒烟测试。
8. 运行现有 `cmd/loadtest/feed_loadtest.exe`，确认发帖与 `GetUserFeed` 远程 RPC
   调用成功。
9. 重启 KnowPost 容器，等待恢复健康，再次验证 etcd 注册与 Feed RPC。
10. 扫描核心服务日志，排除 panic、`Unimplemented`、服务发现零地址和持续重启。

压测产生的业务数据继续保留供复核；本次不自动清理宿主机数据。

## 修改范围

- 新增核心开发 Dockerfile。
- 扩展 `deploy/compose/docker-compose.dev.yml`。
- 调整核心服务 Docker 配置中的活跃 RPC 客户端地址。
- 增加可重复运行的 Compose 静态契约测试。
- 更新开发容器启动文档。

## 非目标

- 不修改业务逻辑或 RPC 协议。
- 不构建、启动或修复 LLM、Agent。
- 不修复既有 full/all-in-one 部署方案。
- 不容器化 MySQL、Redis、etcd。
- 不引入 Kubernetes、Swarm 或生产部署能力。
- 不自动执行或回滚宿主机数据库迁移。
