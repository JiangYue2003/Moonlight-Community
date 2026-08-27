# Feed 三策略综合对比：feed-wp11-hot-mutation-ab-20260827a-control-formal-mixed-c4

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 6.85 | 712.669 | 4 | - | 4 | - |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 27.41 | 3.215 | 4 | - | 4 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 6.52 | 729.459 | 4 | - | 4 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 58.66 | 3.232 | 4 | - | 4 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 6.85 | 6.73 | 6.95 | 3.28% | 712.669 | 692.486 | 716.333 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 27.41 | 26.92 | 27.82 | 3.28% | 3.215 | 3.001 | 3.284 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 6.52 | 6.50 | 6.52 | 0.30% | 729.459 | 718.609 | 766.426 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 58.66 | 58.54 | 58.71 | 0.30% | 3.232 | 3.188 | 3.251 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 412/412 | 0 | 0 | 6.85 | 678.666 | 716.333 | 760.685 | 789.412 | true | 8 | 89.74 | 0.17 | 19203 | 7 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 420/420 | 0 | 0 | 6.95 | 665.631 | 712.669 | 807.461 | 828.561 | true | 7 | 90.63 | 0.17 | 20115 | 5 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 404/404 | 0 | 0 | 6.73 | 675.340 | 692.486 | 724.375 | 754.141 | true | 8 | 89.95 | 0.20 | 19856 | 4 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 1/3 | 1648/1648 | 0 | 0 | 27.41 | 2.686 | 3.001 | 3.702 | 5.899 | true | 8 | 89.74 | 0.17 | 19203 | 7 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 2/3 | 1680/1680 | 0 | 0 | 27.82 | 2.752 | 3.215 | 4.050 | 8.419 | true | 7 | 90.63 | 0.17 | 20115 | 5 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 1616/1616 | 0 | 0 | 26.92 | 3.130 | 3.284 | 4.252 | 9.488 | true | 8 | 89.95 | 0.20 | 19856 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 393/393 | 0 | 0 | 6.52 | 689.107 | 718.609 | 830.921 | 851.520 | true | 8 | 82.88 | 0.19 | 18965 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 392/392 | 0 | 0 | 6.52 | 725.080 | 766.426 | 932.375 | 952.010 | true | 8 | 86.59 | 0.28 | 19118 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 392/392 | 0 | 0 | 6.50 | 698.180 | 729.459 | 783.763 | 793.883 | true | 8 | 80.53 | 0.26 | 18761 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 1/3 | 3537/3537 | 0 | 0 | 58.71 | 2.820 | 3.232 | 4.299 | 10.683 | true | 8 | 82.88 | 0.19 | 18965 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 2/3 | 3528/3528 | 0 | 0 | 58.66 | 2.785 | 3.251 | 4.156 | 15.090 | true | 8 | 86.59 | 0.28 | 19118 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 3528/3528 | 0 | 0 | 58.54 | 2.772 | 3.188 | 3.954 | 5.735 | true | 8 | 80.53 | 0.26 | 18761 | 4 |
