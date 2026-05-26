# counter-aggregator

计数聚合 worker：消费 Kafka 事件 → 累加到 Redis Hash 聚合桶 → 每秒批量刷写 SDS。

## 数据流

```
counter-events (kafka topic)
    │
    ▼
[ AggregationConsumer ]  (group: counter-agg)
    │  HINCRBY agg:v1:{etype}:{eid} {idx} {delta}
    │  CommitMessage
    ▼
[ Flusher ticker 1s ]
    │  SCAN agg:v1:* → 读字段 →
    │  IncrFieldLua  cnt:v1:{etype}:{eid}  +delta
    │  DecrFieldLua  agg:v1:{etype}:{eid}  -delta
    ▼
SDS（cnt:v1:{etype}:{eid}）
```

## 启动
```
go run ./services/counter/aggregator
```

## 配置
见 `etc/aggregator.yaml`。
