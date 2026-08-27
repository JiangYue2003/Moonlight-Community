# Feed 三策略综合对比：feed-wp11-hot-control-mutation-scout-low-20260818

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 5.40 | 3146.976 | 2 | 4 | 16 | - |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 21.83 | 11.132 | 2 | 4 | 16 | - |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 5.13 | 1655.002 | 2 | 4 | 16 | - |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 46.19 | 5.408 | 2 | 4 | 16 | - |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 5.46 | 803.041 | 2 | 4 | 16 | - |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 5.34 | 3192.447 | 4 | 8 | 16 | - |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 21.59 | 10.359 | 4 | 8 | 16 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 5.18 | 3175.319 | 2 | 4 | 16 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 46.64 | 11.134 | 2 | 4 | 16 | - |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 5.48 | 3062.065 | 2 | 4 | 16 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 4.75 | 4.75 | 4.75 | 0.00% | 514.039 | 514.039 | 514.039 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 4.94 | 4.94 | 4.94 | 0.00% | 1086.637 | 1086.637 | 1086.637 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 5.28 | 5.28 | 5.28 | 0.00% | 1632.262 | 1632.262 | 1632.262 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 5.40 | 5.40 | 5.40 | 0.00% | 3146.976 | 3146.976 | 3146.976 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 19.01 | 19.01 | 19.01 | 0.00% | 3.185 | 3.185 | 3.185 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 1/1 | 19.76 | 19.76 | 19.76 | 0.00% | 3.745 | 3.745 | 3.745 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 1/1 | 21.13 | 21.13 | 21.13 | 0.00% | 6.963 | 6.963 | 6.963 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 16 | 1/1 | 21.83 | 21.83 | 21.83 | 0.00% | 11.132 | 11.132 | 11.132 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 4.59 | 4.59 | 4.59 | 0.00% | 503.377 | 503.377 | 503.377 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 4.94 | 4.94 | 4.94 | 0.00% | 987.790 | 987.790 | 987.790 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 5.13 | 5.13 | 5.13 | 0.00% | 1655.002 | 1655.002 | 1655.002 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 5.10 | 5.10 | 5.10 | 0.00% | 3264.440 | 3264.440 | 3264.440 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 41.34 | 41.34 | 41.34 | 0.00% | 2.696 | 2.696 | 2.696 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 1/1 | 44.47 | 44.47 | 44.47 | 0.00% | 3.698 | 3.698 | 3.698 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 1/1 | 46.19 | 46.19 | 46.19 | 0.00% | 5.408 | 5.408 | 5.408 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 16 | 1/1 | 45.90 | 45.90 | 45.90 | 0.00% | 11.325 | 11.325 | 11.325 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 5.06 | 5.06 | 5.06 | 0.00% | 471.221 | 471.221 | 471.221 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 5.46 | 5.46 | 5.46 | 0.00% | 803.041 | 803.041 | 803.041 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 5.43 | 5.43 | 5.43 | 0.00% | 1615.759 | 1615.759 | 1615.759 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 5.32 | 5.32 | 5.32 | 0.00% | 3109.087 | 3109.087 | 3109.087 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 4.78 | 4.78 | 4.78 | 0.00% | 523.083 | 523.083 | 523.083 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 5.28 | 5.28 | 5.28 | 0.00% | 827.235 | 827.235 | 827.235 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 5.27 | 5.27 | 5.27 | 0.00% | 1638.584 | 1638.584 | 1638.584 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 5.34 | 5.34 | 5.34 | 0.00% | 3192.447 | 3192.447 | 3192.447 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 19.13 | 19.13 | 19.13 | 0.00% | 2.707 | 2.707 | 2.707 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 1/1 | 21.11 | 21.11 | 21.11 | 0.00% | 3.186 | 3.186 | 3.186 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 1/1 | 21.10 | 21.10 | 21.10 | 0.00% | 5.729 | 5.729 | 5.729 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 16 | 1/1 | 21.59 | 21.59 | 21.59 | 0.00% | 10.359 | 10.359 | 10.359 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 4.78 | 4.78 | 4.78 | 0.00% | 506.224 | 506.224 | 506.224 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 5.09 | 5.09 | 5.09 | 0.00% | 900.854 | 900.854 | 900.854 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 5.18 | 5.18 | 5.18 | 0.00% | 1599.773 | 1599.773 | 1599.773 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 5.18 | 5.18 | 5.18 | 0.00% | 3175.319 | 3175.319 | 3175.319 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 42.98 | 42.98 | 42.98 | 0.00% | 2.123 | 2.123 | 2.123 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 1/1 | 45.85 | 45.85 | 45.85 | 0.00% | 3.142 | 3.142 | 3.142 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 1/1 | 46.60 | 46.60 | 46.60 | 0.00% | 4.744 | 4.744 | 4.744 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 16 | 1/1 | 46.64 | 46.64 | 46.64 | 0.00% | 11.134 | 11.134 | 11.134 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 5.09 | 5.09 | 5.09 | 0.00% | 474.501 | 474.501 | 474.501 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 5.38 | 5.38 | 5.38 | 0.00% | 904.687 | 904.687 | 904.687 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 5.46 | 5.46 | 5.46 | 0.00% | 1581.694 | 1581.694 | 1581.694 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 5.48 | 5.48 | 5.48 | 0.00% | 3062.065 | 3062.065 | 3062.065 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 72/72 | 0 | 0 | 4.75 | 468.231 | 514.039 | 518.804 | 518.804 | true | 0 | 62.75 | 0.11 | 15878 | 3 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 76/76 | 0 | 0 | 4.94 | 938.800 | 1086.637 | 1093.618 | 1093.618 | true | 4 | 83.76 | 0.15 | 17271 | 3 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 80/80 | 0 | 0 | 5.28 | 1589.083 | 1632.262 | 1657.561 | 1657.561 | true | 8 | 95.25 | 0.14 | 17583 | 3 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 95/95 | 0 | 0 | 5.40 | 3130.911 | 3146.976 | 3186.670 | 3186.670 | true | 0 | 82.91 | 0.09 | 18580 | 3 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 288/288 | 0 | 0 | 19.01 | 2.724 | 3.185 | 4.354 | 5.937 | true | 0 | 62.75 | 0.11 | 15878 | 3 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 1/1 | 304/304 | 0 | 0 | 19.76 | 3.366 | 3.745 | 5.934 | 5.934 | true | 4 | 83.76 | 0.15 | 17271 | 3 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 1/1 | 320/320 | 0 | 0 | 21.13 | 6.313 | 6.963 | 8.415 | 11.207 | true | 8 | 95.25 | 0.14 | 17583 | 3 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 16 | 1/1 | 384/384 | 0 | 0 | 21.83 | 9.606 | 11.132 | 14.059 | 16.415 | true | 0 | 82.91 | 0.09 | 18580 | 3 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 70/70 | 0 | 0 | 4.59 | 458.860 | 503.377 | 533.968 | 533.968 | true | 5 | 68.94 | 0.13 | 15539 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 76/76 | 0 | 0 | 4.94 | 974.361 | 987.790 | 1000.087 | 1000.087 | true | 0 | 78.36 | 0.20 | 17167 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 80/80 | 0 | 0 | 5.13 | 1610.774 | 1655.002 | 1693.235 | 1693.235 | true | 7 | 84.35 | 0.26 | 19667 | 3 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 80/80 | 0 | 0 | 5.10 | 3235.773 | 3264.440 | 3346.674 | 3346.674 | true | 13 | 100.79 | 0.14 | 18443 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 630/630 | 0 | 0 | 41.34 | 2.592 | 2.696 | 3.249 | 6.328 | true | 5 | 68.94 | 0.13 | 15539 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 1/1 | 684/684 | 0 | 0 | 44.47 | 3.296 | 3.698 | 4.768 | 5.772 | true | 0 | 78.36 | 0.20 | 17167 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 1/1 | 720/720 | 0 | 0 | 46.19 | 4.918 | 5.408 | 6.771 | 7.956 | true | 7 | 84.35 | 0.26 | 19667 | 3 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 16 | 1/1 | 720/720 | 0 | 0 | 45.90 | 10.259 | 11.325 | 13.339 | 16.080 | true | 13 | 100.79 | 0.14 | 18443 | 4 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 76/76 | 0 | 0 | 5.06 | 443.800 | 471.221 | 498.744 | 498.744 | true | 6 | 64.29 | 0.04 | 17056 | 3 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 84/84 | 0 | 0 | 5.46 | 784.939 | 803.041 | 974.253 | 974.253 | true | 0 | 65.21 | 0.08 | 17901 | 3 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 88/88 | 0 | 0 | 5.43 | 1589.359 | 1615.759 | 1642.063 | 1642.063 | true | 0 | 89.91 | 0.08 | 17295 | 4 |
| hybrid | gateway | publish | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 89/89 | 0 | 0 | 5.32 | 3083.816 | 3109.087 | 3190.172 | 3190.172 | true | 0 | 72.00 | 0.12 | 18427 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 72/72 | 0 | 0 | 4.78 | 472.789 | 523.083 | 535.131 | 535.131 | true | 6 | 66.61 | 0.18 | 16121 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 80/80 | 0 | 0 | 5.28 | 822.492 | 827.235 | 831.685 | 831.685 | true | 7 | 85.19 | 0.14 | 17836 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 80/80 | 0 | 0 | 5.27 | 1619.244 | 1638.584 | 1668.481 | 1668.481 | true | 8 | 92.22 | 0.14 | 18313 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 86/86 | 0 | 0 | 5.34 | 3181.781 | 3192.447 | 3243.726 | 3243.726 | true | 13 | 77.46 | 0.21 | 18578 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 288/288 | 0 | 0 | 19.13 | 2.237 | 2.707 | 5.871 | 14.445 | true | 6 | 66.61 | 0.18 | 16121 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 1/1 | 320/320 | 0 | 0 | 21.11 | 2.713 | 3.186 | 5.288 | 5.814 | true | 7 | 85.19 | 0.14 | 17836 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 1/1 | 320/320 | 0 | 0 | 21.10 | 4.781 | 5.729 | 6.485 | 9.437 | true | 8 | 92.22 | 0.14 | 18313 | 3 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 16 | 1/1 | 348/348 | 0 | 0 | 21.59 | 9.555 | 10.359 | 13.289 | 14.195 | true | 13 | 77.46 | 0.21 | 18578 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 72/72 | 0 | 0 | 4.78 | 482.940 | 506.224 | 616.500 | 616.500 | true | 0 | 67.48 | 0.17 | 17205 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 80/80 | 0 | 0 | 5.09 | 877.317 | 900.854 | 1092.601 | 1092.601 | true | 0 | 74.40 | 0.16 | 17390 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 80/80 | 0 | 0 | 5.18 | 1585.339 | 1599.773 | 1679.477 | 1679.477 | true | 0 | 82.15 | 0.18 | 17710 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 80/80 | 0 | 0 | 5.18 | 3150.477 | 3175.319 | 3228.389 | 3228.389 | true | 12 | 84.42 | 0.20 | 17910 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 2 | 1/1 | 648/648 | 0 | 0 | 42.98 | 1.688 | 2.123 | 3.182 | 11.514 | true | 0 | 67.48 | 0.17 | 17205 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 1/1 | 720/720 | 0 | 0 | 45.85 | 2.654 | 3.142 | 4.697 | 10.383 | true | 0 | 74.40 | 0.16 | 17390 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 1/1 | 720/720 | 0 | 0 | 46.60 | 4.265 | 4.744 | 6.102 | 9.640 | true | 0 | 82.15 | 0.18 | 17710 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 16 | 1/1 | 720/720 | 0 | 0 | 46.64 | 10.141 | 11.134 | 13.428 | 16.055 | true | 12 | 84.42 | 0.20 | 17910 | 3 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 2 | 1/1 | 78/78 | 0 | 0 | 5.09 | 449.167 | 474.501 | 577.712 | 577.712 | true | 0 | 63.59 | 0.17 | 17168 | 4 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/1 | 84/84 | 0 | 0 | 5.38 | 890.372 | 904.687 | 907.843 | 907.843 | true | 0 | 73.61 | 0.13 | 17771 | 4 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/1 | 88/88 | 0 | 0 | 5.46 | 1577.726 | 1581.694 | 1613.084 | 1613.084 | true | 0 | 88.41 | 0.15 | 17625 | 3 |
| hybrid | rpc | publish | hot | cold | compose-middleware-local-services | publish_total | 16 | 1/1 | 96/96 | 0 | 0 | 5.48 | 3020.956 | 3062.065 | 3125.086 | 3125.086 | true | 14 | 96.97 | 0.11 | 18694 | 3 |
