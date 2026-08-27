# Feed 三策略综合对比：feed-wp11-hot-treatment-auto-20260816a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 5123.71 | 20.030 | 32 | 64 | 128 | - |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 4930.08 | 9.835 | 32 | 64 | 128 | - |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 3052.24 | 171.798 | 32 | 64 | 256 | - |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 8084.14 | 20.876 | 32 | 64 | 256 | - |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 8222.49 | 5.309 | 32 | 64 | 256 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 5105.15 | 5081.87 | 5194.44 | 2.21% | 9.274 | 9.013 | 9.301 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 5123.71 | 4991.78 | 5132.49 | 2.75% | 20.030 | 19.452 | 20.933 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 4890.87 | 4761.83 | 4923.98 | 3.32% | 46.780 | 45.209 | 49.939 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 4930.08 | 4759.23 | 4996.37 | 4.81% | 9.835 | 9.806 | 9.844 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 4811.63 | 4717.05 | 4999.71 | 5.87% | 21.265 | 21.098 | 21.917 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 4865.23 | 4626.71 | 4903.42 | 5.69% | 46.429 | 44.709 | 50.471 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 2932.98 | 2921.58 | 2948.23 | 0.91% | 13.990 | 13.590 | 14.122 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 2963.09 | 2954.41 | 2976.88 | 0.76% | 27.748 | 27.314 | 27.806 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 2999.85 | 2964.46 | 3029.43 | 2.17% | 54.124 | 53.959 | 55.787 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 256 | 3/3 | 3052.24 | 3047.04 | 3054.11 | 0.23% | 171.798 | 171.132 | 172.832 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 8047.05 | 8029.37 | 8159.55 | 1.62% | 5.365 | 5.306 | 5.423 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 8015.17 | 7938.39 | 8043.06 | 1.31% | 10.614 | 10.551 | 10.662 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 8084.14 | 8069.51 | 8088.64 | 0.24% | 20.876 | 20.722 | 21.153 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 256 | 3/3 | 8063.65 | 7894.27 | 8131.52 | 2.94% | 41.151 | 41.039 | 41.463 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 8222.49 | 8078.18 | 8237.59 | 1.94% | 5.309 | 5.308 | 5.388 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 8173.10 | 8148.58 | 8360.20 | 2.59% | 10.584 | 10.292 | 10.594 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 8101.45 | 8056.44 | 8117.70 | 0.76% | 21.052 | 20.535 | 21.099 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 256 | 3/3 | 8009.18 | 7716.11 | 8027.02 | 3.88% | 41.566 | 41.203 | 42.427 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 1/3 | 311688/311688 | 0 | 0 | 5194.44 | 7.879 | 9.013 | 11.710 | 25.350 | true | 0 | 217.44 | 6.82 | 1428 | 3 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 2/3 | 306335/306335 | 0 | 0 | 5105.15 | 8.124 | 9.274 | 12.210 | 27.290 | true | 0 | 199.83 | 7.23 | 1344 | 3 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 304930/304930 | 0 | 0 | 5081.87 | 8.144 | 9.301 | 12.294 | 42.193 | true | 0 | 199.55 | 7.07 | 1397 | 3 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 1/3 | 307981/307981 | 0 | 0 | 5132.49 | 15.804 | 19.452 | 26.235 | 68.150 | true | 0 | 210.61 | 6.87 | 1429 | 3 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 2/3 | 307458/307458 | 0 | 0 | 5123.71 | 16.092 | 20.030 | 26.733 | 67.610 | true | 0 | 200.95 | 7.05 | 1301 | 3 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 299555/299555 | 0 | 0 | 4991.78 | 16.914 | 20.933 | 28.468 | 101.456 | true | 0 | 210.89 | 7.37 | 1363 | 3 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 1/3 | 295563/295563 | 0 | 0 | 4923.98 | 35.160 | 45.209 | 66.224 | 160.621 | true | 0 | 196.94 | 7.07 | 1301 | 3 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 2/3 | 293521/293521 | 0 | 0 | 4890.87 | 34.480 | 46.780 | 84.332 | 176.191 | true | 0 | 204.22 | 6.98 | 1343 | 3 |
| hybrid | gateway | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 285808/285808 | 0 | 0 | 4761.83 | 35.289 | 49.939 | 91.441 | 177.717 | true | 0 | 191.27 | 6.63 | 1344 | 3 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 1/3 | 285574/285574 | 0 | 0 | 4759.23 | 8.279 | 9.844 | 15.100 | 122.887 | true | 0 | 185.99 | 6.59 | 1115 | 3 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 2/3 | 299804/299804 | 0 | 0 | 4996.37 | 8.359 | 9.835 | 15.395 | 176.453 | true | 0 | 198.07 | 6.95 | 1108 | 3 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 295830/295830 | 0 | 0 | 4930.08 | 8.266 | 9.806 | 15.076 | 180.783 | true | 0 | 204.84 | 6.82 | 1101 | 3 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 1/3 | 288765/288765 | 0 | 0 | 4811.63 | 17.053 | 21.265 | 38.707 | 198.531 | true | 0 | 204.57 | 6.62 | 1118 | 3 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 2/3 | 283053/283053 | 0 | 0 | 4717.05 | 17.430 | 21.917 | 87.410 | 142.474 | true | 0 | 191.06 | 6.54 | 1095 | 3 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 300031/300031 | 0 | 0 | 4999.71 | 17.049 | 21.098 | 28.412 | 78.913 | true | 0 | 200.53 | 7.30 | 1114 | 3 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 1/3 | 291984/291984 | 0 | 0 | 4865.23 | 33.770 | 46.429 | 96.077 | 401.259 | true | 0 | 210.98 | 6.79 | 1109 | 3 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 2/3 | 294296/294296 | 0 | 0 | 4903.42 | 35.775 | 44.709 | 65.750 | 295.149 | true | 0 | 192.98 | 6.94 | 1116 | 3 |
| hybrid | gateway | hot-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 277663/277663 | 0 | 0 | 4626.71 | 36.225 | 50.471 | 119.320 | 220.236 | true | 0 | 187.88 | 6.39 | 1101 | 3 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 32 | 1/3 | 175995/175995 | 0 | 0 | 2932.98 | 12.710 | 13.590 | 18.827 | 114.332 | true | 0 | 498.29 | 3.83 | 20932 | 6 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 32 | 2/3 | 176924/176924 | 0 | 0 | 2948.23 | 12.969 | 14.122 | 18.595 | 94.640 | true | 0 | 481.88 | 3.95 | 20439 | 6 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 175317/175317 | 0 | 0 | 2921.58 | 12.869 | 13.990 | 22.265 | 184.003 | true | 0 | 492.94 | 3.86 | 20632 | 6 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 64 | 1/3 | 178642/178642 | 0 | 0 | 2976.88 | 25.509 | 27.314 | 34.096 | 114.340 | true | 0 | 480.36 | 4.25 | 20294 | 8 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 64 | 2/3 | 177294/177294 | 0 | 0 | 2954.41 | 25.619 | 27.748 | 36.111 | 136.409 | true | 0 | 479.27 | 3.89 | 20651 | 7 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 178059/178059 | 0 | 0 | 2963.09 | 25.735 | 27.806 | 36.544 | 142.945 | true | 0 | 475.86 | 4.11 | 20581 | 6 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 128 | 1/3 | 180048/180048 | 0 | 0 | 2999.85 | 51.315 | 55.787 | 68.232 | 214.715 | true | 0 | 534.33 | 3.90 | 20753 | 7 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 128 | 2/3 | 181849/181849 | 0 | 0 | 3029.43 | 50.475 | 53.959 | 62.694 | 101.208 | true | 0 | 488.37 | 4.07 | 20936 | 6 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 177924/177924 | 0 | 0 | 2964.46 | 50.089 | 54.124 | 116.468 | 187.465 | true | 0 | 452.84 | 3.85 | 21042 | 6 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 256 | 1/3 | 182976/182976 | 0 | 0 | 3047.04 | 142.179 | 172.832 | 245.204 | 520.867 | true | 0 | 449.77 | 4.13 | 21202 | 6 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 256 | 2/3 | 183299/183299 | 0 | 0 | 3052.24 | 141.141 | 171.798 | 243.365 | 627.474 | true | 0 | 477.34 | 4.07 | 21355 | 6 |
| hybrid | rpc | deep-page | hot | l1-warm | compose-middleware-local-services | read | 256 | 3/3 | 183403/183403 | 0 | 0 | 3054.11 | 140.994 | 171.132 | 242.700 | 569.075 | true | 0 | 474.48 | 4.12 | 21128 | 6 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 1/3 | 482844/482844 | 0 | 0 | 8047.05 | 4.984 | 5.365 | 6.282 | 16.015 | true | 0 | 248.53 | 7.61 | 1427 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 2/3 | 489599/489599 | 0 | 0 | 8159.55 | 4.883 | 5.306 | 6.152 | 13.488 | true | 0 | 239.13 | 7.79 | 1295 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 481783/481783 | 0 | 0 | 8029.37 | 5.179 | 5.423 | 6.334 | 16.526 | true | 0 | 232.38 | 7.96 | 1426 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 1/3 | 480980/480980 | 0 | 0 | 8015.17 | 9.953 | 10.614 | 12.192 | 26.764 | true | 0 | 251.38 | 7.80 | 1314 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 2/3 | 476376/476376 | 0 | 0 | 7938.39 | 10.045 | 10.662 | 12.132 | 24.161 | true | 0 | 247.93 | 7.92 | 1372 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 482653/482653 | 0 | 0 | 8043.06 | 9.907 | 10.551 | 12.192 | 33.423 | true | 0 | 234.57 | 7.86 | 1335 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 1/3 | 484272/484272 | 0 | 0 | 8069.51 | 19.602 | 20.722 | 22.950 | 49.324 | true | 0 | 230.85 | 7.90 | 1297 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 2/3 | 485139/485139 | 0 | 0 | 8084.14 | 19.693 | 20.876 | 23.220 | 50.145 | true | 0 | 235.48 | 8.00 | 1382 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 485453/485453 | 0 | 0 | 8088.64 | 19.860 | 21.153 | 24.100 | 51.050 | true | 0 | 238.38 | 7.87 | 1298 | 4 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 256 | 1/3 | 488162/488162 | 0 | 0 | 8131.52 | 39.077 | 41.463 | 47.056 | 98.814 | true | 0 | 234.55 | 8.13 | 1322 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 256 | 2/3 | 484088/484088 | 0 | 0 | 8063.65 | 39.141 | 41.039 | 46.921 | 96.314 | true | 0 | 241.79 | 7.97 | 1322 | 3 |
| hybrid | rpc | distributed-read | hot | l1-warm | compose-middleware-local-services | read | 256 | 3/3 | 473927/473927 | 0 | 0 | 7894.27 | 38.962 | 41.151 | 46.043 | 231.844 | true | 0 | 247.93 | 7.55 | 1327 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 1/3 | 494281/494281 | 0 | 0 | 8237.59 | 4.871 | 5.309 | 6.201 | 13.843 | true | 0 | 254.50 | 7.81 | 1114 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 2/3 | 484725/484725 | 0 | 0 | 8078.18 | 5.008 | 5.388 | 6.371 | 20.070 | true | 0 | 246.85 | 7.90 | 1091 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 32 | 3/3 | 493380/493380 | 0 | 0 | 8222.49 | 4.886 | 5.308 | 6.151 | 16.329 | true | 0 | 250.78 | 7.77 | 1109 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 1/3 | 501683/501683 | 0 | 0 | 8360.20 | 9.581 | 10.292 | 11.720 | 26.754 | true | 0 | 251.75 | 7.70 | 1113 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 2/3 | 488973/488973 | 0 | 0 | 8148.58 | 9.947 | 10.584 | 12.181 | 26.861 | true | 0 | 235.11 | 7.83 | 1095 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 64 | 3/3 | 490433/490433 | 0 | 0 | 8173.10 | 10.009 | 10.594 | 11.870 | 24.290 | true | 0 | 235.44 | 7.85 | 1114 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 1/3 | 487169/487169 | 0 | 0 | 8117.70 | 19.423 | 20.535 | 22.812 | 49.869 | true | 0 | 243.18 | 8.00 | 1112 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 2/3 | 483477/483477 | 0 | 0 | 8056.44 | 19.894 | 21.099 | 23.732 | 53.183 | true | 0 | 236.81 | 7.90 | 1102 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 128 | 3/3 | 486215/486215 | 0 | 0 | 8101.45 | 19.674 | 21.052 | 23.629 | 51.533 | true | 0 | 231.36 | 7.98 | 1100 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 256 | 1/3 | 481796/481796 | 0 | 0 | 8027.02 | 39.660 | 41.566 | 46.702 | 98.594 | true | 0 | 236.25 | 8.00 | 1099 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 256 | 2/3 | 480779/480779 | 0 | 0 | 8009.18 | 39.088 | 41.203 | 46.754 | 118.546 | true | 0 | 235.58 | 8.17 | 1108 | 3 |
| hybrid | rpc | hot-read | hot | l1-warm | compose-middleware-local-services | read | 256 | 3/3 | 463195/463195 | 0 | 0 | 7716.11 | 39.804 | 42.427 | 88.424 | 236.793 | true | 0 | 230.05 | 7.52 | 1108 | 3 |
