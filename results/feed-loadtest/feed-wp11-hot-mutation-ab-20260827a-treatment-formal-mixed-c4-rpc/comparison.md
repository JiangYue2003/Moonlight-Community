# Feed 三策略综合对比：feed-wp11-hot-mutation-ab-20260827a-treatment-formal-mixed-c4-rpc

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 6.01 | 828.745 | 4 | - | 4 | - |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 24.04 | 3.189 | 4 | - | 4 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 6.14 | 779.365 | 4 | - | 4 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 55.23 | 2.636 | 4 | - | 4 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 6.01 | 5.93 | 6.02 | 1.53% | 828.745 | 749.094 | 840.138 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 24.04 | 23.71 | 24.08 | 1.53% | 3.189 | 3.137 | 3.217 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 6.14 | 6.07 | 6.56 | 7.99% | 779.365 | 744.209 | 782.808 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 55.23 | 54.59 | 59.00 | 7.99% | 2.636 | 2.607 | 2.647 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 364/364 | 0 | 0 | 6.02 | 778.591 | 840.138 | 877.172 | 891.982 | true | 7 | 84.09 | 0.15 | 19118 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 356/356 | 0 | 0 | 5.93 | 771.486 | 828.745 | 858.716 | 862.173 | true | 8 | 86.24 | 0.17 | 19346 | 7 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 364/364 | 0 | 0 | 6.01 | 739.130 | 749.094 | 813.547 | 820.985 | true | 7 | 82.95 | 0.18 | 18646 | 5 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 1/3 | 1456/1456 | 0 | 0 | 24.08 | 2.669 | 3.189 | 4.557 | 6.185 | true | 7 | 84.09 | 0.15 | 19118 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 2/3 | 1424/1424 | 0 | 0 | 23.71 | 2.665 | 3.137 | 4.337 | 10.812 | true | 8 | 86.24 | 0.17 | 19346 | 7 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 1456/1456 | 0 | 0 | 24.04 | 2.684 | 3.217 | 4.416 | 6.315 | true | 7 | 82.95 | 0.18 | 18646 | 5 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 396/396 | 0 | 0 | 6.56 | 713.240 | 744.209 | 773.412 | 788.151 | true | 8 | 82.10 | 0.23 | 19885 | 7 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 372/372 | 0 | 0 | 6.14 | 745.706 | 782.808 | 904.284 | 918.729 | true | 8 | 90.52 | 0.19 | 18955 | 7 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 364/364 | 0 | 0 | 6.07 | 748.824 | 779.365 | 855.995 | 867.692 | true | 7 | 82.89 | 0.19 | 19307 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 1/3 | 3564/3564 | 0 | 0 | 59.00 | 2.093 | 2.647 | 3.805 | 9.153 | true | 8 | 82.10 | 0.23 | 19885 | 7 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 2/3 | 3348/3348 | 0 | 0 | 55.23 | 2.079 | 2.607 | 3.364 | 21.425 | true | 8 | 90.52 | 0.19 | 18955 | 7 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 3276/3276 | 0 | 0 | 54.59 | 2.097 | 2.636 | 3.745 | 23.944 | true | 7 | 82.89 | 0.19 | 19307 | 4 |
