# Feed 三策略综合对比：feed-wp11-hot-mutation-script-smoke-v2-20260818

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4.98 | 458.322 | 2 | - | 2 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 44.83 | 2.714 | 2 | - | 2 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 4.98 | 4.98 | 4.98 | 0.00% | 458.322 | 458.322 | 458.322 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 44.83 | 44.83 | 44.83 | 0.00% | 2.714 | 2.714 | 2.714 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 10/10 | 0 | 0 | 4.98 | 452.001 | 458.322 | 458.322 | 458.322 | true | 0 | 67.50 | 0.19 | 14655 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 90/90 | 0 | 0 | 44.83 | 2.246 | 2.714 | 6.371 | 6.371 | true | 0 | 67.50 | 0.19 | 14655 | 3 |
