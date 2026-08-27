# Feed 三策略综合对比：feed-wp11-hot-mutation-ab-20260827a-control-formal-publish-c2

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 6.30 | 421.279 | 2 | - | 2 | - |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 6.61 | 398.188 | 2 | - | 2 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 6.30 | 6.29 | 6.41 | 1.82% | 421.279 | 405.904 | 428.420 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 6.61 | 6.31 | 6.82 | 7.72% | 398.188 | 362.586 | 442.305 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/3 | 378/378 | 0 | 0 | 6.29 | 384.253 | 428.420 | 511.465 | 648.633 | true | 8 | 79.06 | 0.09 | 17676 | 5 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 2/3 | 386/386 | 0 | 0 | 6.41 | 354.325 | 405.904 | 513.322 | 514.613 | true | 7 | 71.94 | 0.08 | 16830 | 5 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 378/378 | 0 | 0 | 6.30 | 387.695 | 421.279 | 505.002 | 514.139 | true | 0 | 69.64 | 0.14 | 17155 | 5 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/3 | 380/380 | 0 | 0 | 6.31 | 387.332 | 442.305 | 597.141 | 695.421 | true | 8 | 75.87 | 0.18 | 16936 | 4 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 2/3 | 410/410 | 0 | 0 | 6.82 | 342.612 | 362.586 | 460.769 | 508.509 | true | 7 | 75.07 | 0.13 | 17957 | 4 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 398/398 | 0 | 0 | 6.61 | 358.179 | 398.188 | 487.588 | 497.418 | true | 6 | 74.36 | 0.15 | 17074 | 4 |
