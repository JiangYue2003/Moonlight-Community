# Feed 三策略综合对比：feed-wp11-hot-mutation-ab-20260827a-treatment-formal-publish-c2-rpc

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 6.33 | 388.983 | 2 | - | 2 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 6.33 | 6.27 | 6.35 | 1.19% | 388.983 | 384.414 | 402.144 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/3 | 380/380 | 0 | 0 | 6.33 | 372.209 | 402.144 | 465.591 | 647.474 | true | 7 | 68.18 | 0.15 | 18000 | 5 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 2/3 | 378/378 | 0 | 0 | 6.27 | 366.328 | 388.983 | 513.722 | 593.180 | true | 7 | 66.57 | 0.16 | 17121 | 5 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 382/382 | 0 | 0 | 6.35 | 359.986 | 384.414 | 553.480 | 569.053 | true | 8 | 75.92 | 0.15 | 17316 | 5 |
