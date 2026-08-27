# Feed 三策略综合对比：feed-wp10-functional-20260816

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | hot-read | hot | l1-warm | compose-full | read | 32513.61 | 1.618 | 32 | - | 32 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | hot-read | hot | l1-warm | compose-full | read | 32 | 1/1 | 32513.61 | 32513.61 | 32513.61 | 0.00% | 1.618 | 1.618 | 1.618 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | rpc | hot-read | hot | l1-warm | compose-full | read | 32 | 1/1 | 162586/162586 | 0 | 0 | 32513.61 | 1.549 | 1.618 | 2.173 | 6.861 | true | 0 | 404.40 | 13.67 | 99 | 3 |
