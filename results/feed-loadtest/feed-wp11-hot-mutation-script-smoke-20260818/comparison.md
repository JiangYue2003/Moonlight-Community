# Feed 三策略综合对比：feed-wp11-hot-mutation-script-smoke-20260818

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 5.02 | 425.741 | 2 | - | 2 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 45.21 | 2.243 | 2 | - | 2 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 5.02 | 5.02 | 5.02 | 0.00% | 425.741 | 425.741 | 425.741 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 45.21 | 45.21 | 45.21 | 0.00% | 2.243 | 2.243 | 2.243 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 12/12 | 0 | 0 | 5.02 | 423.072 | 425.741 | 425.741 | 425.741 | true | 0 | 56.09 | 0.25 | 14758 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 108/108 | 0 | 0 | 45.21 | 2.091 | 2.243 | 5.918 | 5.918 | true | 0 | 56.09 | 0.25 | 14758 | 3 |
