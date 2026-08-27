# Feed 三策略综合对比：feed-wp11-hot-control-formal-publish-c2-c4-20260818

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 4.98 | 980.207 | 2 | 4 | 4 | - |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 4.68 | 1041.125 | 2 | 4 | 4 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 4.68 | 4.33 | 4.69 | 7.61% | 525.875 | 520.583 | 632.475 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 4.98 | 4.60 | 4.99 | 7.87% | 980.207 | 924.459 | 1029.450 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 4.39 | 4.29 | 5.06 | 17.72% | 606.509 | 482.205 | 693.404 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 4.68 | 3.76 | 5.10 | 28.65% | 1041.125 | 972.972 | 1795.656 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/3 | 260/260 | 0 | 0 | 4.33 | 569.034 | 632.475 | 796.895 | 851.452 | true | 5 | 71.30 | 0.11 | 15748 | 4 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 2/3 | 282/282 | 0 | 0 | 4.68 | 494.551 | 520.583 | 614.298 | 665.105 | true | 5 | 71.24 | 0.09 | 16264 | 5 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 282/282 | 0 | 0 | 4.69 | 483.613 | 525.875 | 595.509 | 623.758 | true | 5 | 65.45 | 0.11 | 16443 | 5 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 300/300 | 0 | 0 | 4.99 | 951.716 | 980.207 | 1134.474 | 1168.330 | true | 6 | 75.93 | 0.10 | 17428 | 4 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 300/300 | 0 | 0 | 4.98 | 909.350 | 924.459 | 1099.071 | 1105.520 | true | 4 | 75.07 | 0.12 | 18178 | 4 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 276/276 | 0 | 0 | 4.60 | 985.722 | 1029.450 | 1163.361 | 1170.713 | true | 4 | 73.62 | 0.13 | 16791 | 4 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/3 | 304/304 | 0 | 0 | 5.06 | 454.403 | 482.205 | 541.457 | 553.874 | true | 6 | 71.29 | 0.16 | 16844 | 4 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 2/3 | 264/264 | 0 | 0 | 4.39 | 559.502 | 606.509 | 661.415 | 671.539 | true | 5 | 66.76 | 0.12 | 16319 | 4 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 3/3 | 258/258 | 0 | 0 | 4.29 | 593.300 | 693.404 | 744.339 | 748.859 | true | 9 | 64.31 | 0.13 | 16320 | 5 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 283/283 | 0 | 0 | 4.68 | 998.602 | 1041.125 | 1120.559 | 1177.768 | true | 7 | 70.87 | 0.15 | 17114 | 3 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 232/232 | 0 | 0 | 3.76 | 1728.105 | 1795.656 | 1945.970 | 1962.320 | true | 4 | 81.38 | 0.18 | 16430 | 4 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 308/308 | 0 | 0 | 5.10 | 908.122 | 972.972 | 1172.933 | 1176.034 | true | 7 | 78.15 | 0.15 | 17763 | 4 |
