# Feed 端到端压测可靠性设计

日期：2026-07-29

## 背景

直接在 `cmd/loadtest` 中运行 `feed_loadtest.exe` 时，关注关系初始化和
`GetUserFeed` 读取成功，但 1000 次发帖全部失败。分阶段 RPC 诊断确认了三处独立问题：

1. 压测固定使用从 `100000` 开始的用户 ID，但这些用户不一定存在，
   `CreateDraft` 会触发 `know_posts.creator_id` 外键错误。
2. 压测调用 `ConfirmContent` 时没有传服务端要求的 `object_key`。
3. KnowPost 的 `FeedWriter` 被注入了 `nil` Kafka producer，普通用户发布进入推模式时
   会在发送 fanout 事件时空指针 panic。

此外，压测即使存在失败请求仍以退出码 0 结束，无法用于自动化验收。

## 目标

- 不依赖预先存在的固定用户 ID，压测可以自行创建有效测试用户。
- CreateDraft、PatchMetadata、ConfirmContent、Publish 四阶段发帖链路可完整执行。
- 推模式发布能够向 Kafka `feed-fanout` topic 发送事件，并由 fanout worker 消费。
- `GetUserFeed` 在并发读取下保持可远程调用。
- 任一初始化、发帖或读取请求失败时，压测进程返回非零退出码。
- 保持 `-skip-setup` 兼容：显式跳过初始化时仍按 `UserIDStart` 构造已有用户 ID。

## 方案

### 压测用户生命周期

在压测配置中增加 `UserRpc`。正常初始化时，通过 User RPC `Create` 创建
`UserCount` 个具有唯一邮箱的测试用户，保存服务端返回的真实 ID 列表。

后续创建关注关系、随机选择发帖作者和随机读取 Feed 都只从这份 ID 列表取值，
不再推导 `UserIDStart + offset`。

使用 `-skip-setup` 时不创建用户，继续根据 `UserIDStart` 和 `UserCount`
生成连续 ID，以兼容已有的预置数据压测方式。

### 发帖内容确认

每篇压测帖子使用其 ID 构造唯一测试对象键：

```text
loadtest/<post-id>/content.md
```

该对象键仅作为当前 KnowPost 元数据确认所需字段；服务端当前不向对象存储发起校验。
压测同时提供非零的测试内容大小。

### Kafka producer 与 fanout worker

KnowPost `ServiceContext` 使用配置中的 Kafka broker 创建：

- 绑定 `feed-fanout` topic 的 `kq.Pusher`；
- `FeedWriter` 使用的 `KafkaProducerAdapter`。

适配器使用 `PushWithKey` 保留 `postID` 消息键。`FeedWriter` 对缺失 producer
增加显式错误，确保错误配置返回可诊断错误而不是 panic。

KnowPost RPC 启动时创建并启动 `feed-fanout-group` consumer；RPC 上下文结束时停止
consumer。`ServiceContext` 持有 producer 并在服务退出时关闭它，以便刷新缓冲消息并
释放连接。独立 RPC 入口和合并服务入口使用同一生命周期逻辑。

### 压测结果语义

压测在最终报告后检查以下计数：

- 初始化失败数必须为 0；
- 发帖失败数必须为 0；
- 读取失败数必须为 0；
- 实际请求总数必须与配置一致。

任一条件不满足时打印阶段错误并以退出码 1 结束。每个阶段只输出前若干条详细错误，
避免压力场景刷屏，同时保留根因。

## 错误处理

- User RPC 创建失败：初始化立即失败，不使用不存在的 ID 继续压测。
- Kafka producer 未配置：`FeedWriter` 返回明确错误，不允许 nil 解引用。
- Kafka 发送失败：Publish 维持现有业务降级语义并记录错误；端到端测试额外通过
  Kafka/Redis 结果验证 fanout 是否完成。
- RPC 单请求失败：计入失败数并采样打印原始 gRPC 错误。
- 配置文件不存在：保留当前命令行默认配置回退行为，并补齐默认 `UserRpc`。

## 测试与验收

1. 单元测试验证 nil Kafka producer 返回错误且不会 panic。
2. 单元测试验证压测用户创建请求具有唯一标识，且发帖确认请求包含非空对象键。
3. 运行 KnowPost RPC、User RPC、Relation RPC、Counter RPC、Redis、MySQL、etcd
   和 Kafka 的真实链路。
4. 小规模冒烟压测必须达到发帖与读取 100% 成功。
5. 运行原始命令 `cmd/loadtest/feed_loadtest.exe`，默认 1000 次发帖和 5000 次读取
   均须成功，进程退出码为 0。
6. 对服务日志扫描 `panic`、`Unimplemented`、`produced zero addresses` 和 Kafka
   fanout 错误。

压测会保留测试用户、关系和帖子，便于复核结果；不自动删除业务数据。测试期间临时启动
的进程和 etcd 注册会在结束后清理，用户原先启动的服务不会被终止。

## 非目标

- 不在本次修改中实现对象存储上传。
- 不重构已有 Counter、Relation 或 User 服务。
- 不调整大 V 阈值或推拉算法。
- 不自动删除此前或本轮产生的压测数据。
