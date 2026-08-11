# Feed P1-1 批量扇出设计

## 背景

普通作者发布帖子后，KnowPost 的 Feed fanout consumer 会读取作者的全部粉丝，并为每个粉丝依次执行 `ZADD`、`ZREMRANGEBYRANK` 和 `EXPIRE`。当作者接近大 V 阈值时，一条事件会产生约三千次串行 Redis 往返。当前实现还会在部分粉丝写入失败后继续把事件标记为已处理，使失败粉丝无法通过 Kafka 重试补齐。

## 目标

- 将粉丝 inbox 写入按 100 个用户分块，并在每块内使用 Redis Pipeline。
- 保持至少一次投递语义：任何分块失败都让整条 Kafka 事件重试。
- 仅在所有分块成功后写入 processed 标记。
- processed 标记写入失败时同样重试整条事件。
- 保持现有 inbox 容量、TTL、Kafka topic 和 consumer group 不变。

## 非目标

- 不改变 Relation 当前一次加载全部粉丝的行为。
- 不增加按分块保存的 checkpoint。
- 不实现 Kafka DLQ、动态大 V 阈值、cursor 分页或 inbox 分片。
- 不清理未使用的旧 `FanoutWorker` 实现。

## 方案选择

采用“分块 Pipeline + 整事件重试”。与按分块保存 checkpoint 相比，它不引入新的持久状态；与大型 Lua 脚本相比，它不会形成长时间运行或跨 Redis Cluster slot 的脚本。重试可能重复执行已经成功的分块，但 inbox 使用 post ID 作为 ZSet member，相同事件重复写入不会产生重复帖子。

## 组件设计

### Fanout handler

从 Kafka 构造逻辑中提取一个可单元测试的 fanout handler。它负责：

1. 检查 processed 标记。
2. 从 Relation 获取粉丝 ID。
3. 以 100 个 ID 为一批调用 inbox batch writer。
4. 任意批次失败时立即返回错误。
5. 所有批次成功后写入 processed 标记。

Kafka consumer 只负责反序列化 `FeedEvent` 并调用 handler，现有无法解析的消息仍记录错误后跳过。由于项目锁定的 `go-queue v1.2.2` 不会自动重放返回错误的消息，consumer 使用单 processor，并在同一消息内以一秒间隔重试完整 handler；只有 handler 成功后才返回成功并提交 offset。服务停止时取消重试，保留未提交 offset 供重启后继续消费。

### Inbox batch writer

新增窄接口 `InboxBatchWriter`，方法接收一批 user ID、post ID 和时间戳。`RedisAdapter` 实现该接口，并在一个 Pipeline 中为每个用户排入：

- `ZADD feed:inbox:{userID} score postID`
- `ZREMRANGEBYRANK feed:inbox:{userID} 0 -(INBOX_MAX_SIZE+1)`
- `EXPIRE feed:inbox:{userID} INBOX_TTL_SECONDS`

Pipeline 返回错误时，调用方把整批视为失败。Redis 可能已经执行其中一部分命令，但下一次整事件重试会安全覆盖相同 ZSet member，并重新执行容量和 TTL 修复。

公开 Worker 构造函数继续接收原有 `RedisClient`，保持已有调用方和测试 mock 的源码兼容。运行时若实现同时提供批量接口（生产 `RedisAdapter` 即如此），自动选择 Pipeline；旧自定义实现则使用顺序兼容 writer。

## 错误与重试语义

- processed 标记已存在：直接返回成功，不执行 fanout。
- processed 查询失败：记录错误后继续 fanout；实际 Redis 写入失败仍会触发重试。
- Relation 查询失败：返回错误，由 Kafka 重试。
- 任意 Pipeline 失败：停止后续分块，不写 processed 标记，由 consumer 退避后重试整条事件。
- processed 标记写入失败：由 consumer 退避后重试整条事件。
- 空粉丝列表：不执行 Pipeline，直接写 processed 标记。

consumer 配置 `Processors=1`、`ForceCommit=false`，避免后续成功消息推进同一 partition 的 offset 并跳过失败事件。重试循环在服务停止时可取消；消费回调在解析消息前同步检查 Worker 生命周期，确保已经预取的下一条消息也返回错误，不会越过失败 offset。

该设计提供至少一次处理，不承诺恰好一次；用户可见结果依靠 ZSet member 幂等收敛。

## 测试设计

- 205 个粉丝被切分为 `100 + 100 + 5` 三次 batch writer 调用。
- 第二批写入失败时 handler 返回错误，第三批不执行，processed 标记不存在。
- 所有批次成功后 processed 标记存在。
- processed 标记写入失败时 handler 返回错误。
- 空粉丝列表不调用 batch writer，但写入 processed 标记。
- consumer 使用单 processor 且 handler 错误不强制提交 offset。
- 整事件重试在临时错误恢复后成功，并可在服务停止时取消。
- 服务停止后拒绝已经预取的下一条消息，即使该消息无法解析也不返回成功。
- 公开 Worker 构造函数保持原有 `RedisClient` 签名，旧实现走兼容 writer，生产实现走原生 batch writer。
- 使用 miniredis 验证 RedisAdapter 批量写入、容量裁剪和 TTL。
- 重复处理相同事件后，每个 inbox 中该 post ID 仍只有一个成员。
- KnowPost Feed 相关包及受影响服务回归测试保持通过。

## 验收标准

- 单条普通作者 fanout 的 Redis 网络往返由“每个粉丝三次”降为“每 100 个粉丝一次 Pipeline”。
- 部分失败不会留下错误的 processed 标记。
- 重试不会生成重复 Feed 项。
- 不改变大 V 拉模式和个人 Feed 读取路径。
