# Feed 三策略综合对比：feed-wp11-hot-control-formal-mixed-c4-c8-20260818

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 5.35 | 1642.853 | 4 | 8 | 8 | - |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 21.40 | 6.030 | 4 | 8 | 8 | - |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4.56 | 1019.499 | 4 | 8 | 8 | - |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 41.08 | 3.987 | 4 | 8 | 8 | - |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 5.02 | 931.307 | 4 | 8 | 8 | - |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 20.08 | 3.171 | 4 | 8 | 8 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4.81 | 1845.810 | 4 | 8 | 8 | - |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 43.28 | 5.285 | 4 | 8 | 8 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 4.97 | 4.43 | 4.97 | 10.74% | 922.041 | 920.483 | 1103.999 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 3/3 | 5.35 | 4.66 | 5.44 | 14.71% | 1642.853 | 1639.456 | 1923.610 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 19.87 | 17.73 | 19.87 | 10.74% | 3.920 | 3.701 | 4.288 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 3/3 | 21.40 | 18.65 | 21.78 | 14.63% | 6.030 | 6.018 | 7.157 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 4.56 | 3.92 | 4.58 | 14.43% | 1019.499 | 999.356 | 1273.300 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 3/3 | 4.38 | 4.15 | 4.55 | 9.24% | 2237.823 | 1874.412 | 2400.878 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 41.08 | 35.25 | 41.18 | 14.43% | 3.987 | 3.855 | 4.656 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 3/3 | 39.41 | 37.33 | 40.97 | 9.23% | 6.468 | 6.245 | 6.902 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 5.02 | 5.01 | 5.05 | 0.75% | 931.307 | 920.613 | 1009.308 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 3/3 | 4.92 | 4.91 | 4.92 | 0.21% | 1774.196 | 1772.457 | 1801.218 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 20.08 | 20.03 | 20.18 | 0.75% | 3.171 | 2.797 | 3.174 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 3/3 | 19.67 | 19.65 | 19.69 | 0.21% | 5.413 | 5.384 | 5.513 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 4.58 | 4.58 | 4.64 | 1.37% | 1009.855 | 994.897 | 1032.673 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 3/3 | 4.81 | 4.67 | 4.86 | 3.96% | 1845.810 | 1732.632 | 1848.754 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 41.25 | 41.22 | 41.78 | 1.37% | 3.161 | 2.824 | 3.196 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 3/3 | 43.28 | 42.08 | 43.78 | 3.92% | 5.285 | 5.035 | 5.377 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 268/268 | 0 | 0 | 4.43 | 1069.248 | 1103.999 | 1284.871 | 1286.430 | true | 7 | 79.73 | 0.13 | 17878 | 4 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 300/300 | 0 | 0 | 4.97 | 892.894 | 922.041 | 995.926 | 1009.968 | true | 6 | 75.19 | 0.14 | 17948 | 6 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 300/300 | 0 | 0 | 4.97 | 909.060 | 920.483 | 964.247 | 971.841 | true | 0 | 81.40 | 0.15 | 17749 | 5 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/3 | 328/328 | 0 | 0 | 5.44 | 1590.872 | 1642.853 | 1665.566 | 1675.392 | true | 8 | 89.14 | 0.14 | 18711 | 4 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 2/3 | 325/325 | 0 | 0 | 5.35 | 1609.963 | 1639.456 | 1664.088 | 1672.481 | true | 8 | 88.28 | 0.17 | 18228 | 4 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 3/3 | 280/280 | 0 | 0 | 4.66 | 1891.043 | 1923.610 | 1986.703 | 1992.830 | true | 8 | 82.75 | 0.15 | 17329 | 4 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 1/3 | 1072/1072 | 0 | 0 | 17.73 | 3.783 | 4.288 | 5.277 | 6.881 | true | 7 | 79.73 | 0.13 | 17878 | 4 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 2/3 | 1200/1200 | 0 | 0 | 19.87 | 3.213 | 3.701 | 4.401 | 5.836 | true | 6 | 75.19 | 0.14 | 17948 | 6 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 1200/1200 | 0 | 0 | 19.87 | 3.695 | 3.920 | 4.875 | 6.231 | true | 0 | 81.40 | 0.15 | 17749 | 5 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 1/3 | 1312/1312 | 0 | 0 | 21.78 | 5.378 | 6.030 | 7.500 | 9.148 | true | 8 | 89.14 | 0.14 | 18711 | 4 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 2/3 | 1300/1300 | 0 | 0 | 21.40 | 5.354 | 6.018 | 7.234 | 8.967 | true | 8 | 88.28 | 0.17 | 18228 | 4 |
| hybrid | gateway | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 3/3 | 1121/1121 | 0 | 0 | 18.65 | 6.427 | 7.157 | 8.509 | 17.110 | true | 8 | 82.75 | 0.15 | 17329 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 276/276 | 0 | 0 | 4.58 | 974.490 | 999.356 | 1053.109 | 1074.501 | true | 7 | 81.34 | 0.20 | 17352 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 276/276 | 0 | 0 | 4.56 | 986.873 | 1019.499 | 1081.027 | 1086.715 | true | 5 | 77.50 | 0.14 | 16951 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 236/236 | 0 | 0 | 3.92 | 1193.379 | 1273.300 | 1385.686 | 1422.710 | true | 4 | 71.97 | 0.19 | 15430 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/3 | 280/280 | 0 | 0 | 4.55 | 1842.884 | 1874.412 | 2029.760 | 2055.912 | true | 7 | 89.87 | 0.17 | 17665 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 2/3 | 256/256 | 0 | 0 | 4.15 | 2190.159 | 2400.878 | 2533.381 | 2560.996 | true | 7 | 89.10 | 0.18 | 16722 | 3 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 3/3 | 269/269 | 0 | 0 | 4.38 | 2033.273 | 2237.823 | 2307.443 | 2319.130 | true | 8 | 93.75 | 0.16 | 17145 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 1/3 | 2484/2484 | 0 | 0 | 41.18 | 3.706 | 3.987 | 5.186 | 13.785 | true | 7 | 81.34 | 0.20 | 17352 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 2/3 | 2484/2484 | 0 | 0 | 41.08 | 3.449 | 3.855 | 4.840 | 6.296 | true | 5 | 77.50 | 0.14 | 16951 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 2124/2124 | 0 | 0 | 35.25 | 3.920 | 4.656 | 5.992 | 8.877 | true | 4 | 71.97 | 0.19 | 15430 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 1/3 | 2520/2520 | 0 | 0 | 40.97 | 5.477 | 6.245 | 7.721 | 9.877 | true | 7 | 89.87 | 0.17 | 17665 | 4 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 2/3 | 2304/2304 | 0 | 0 | 37.33 | 5.971 | 6.902 | 8.388 | 12.716 | true | 7 | 89.10 | 0.18 | 16722 | 3 |
| hybrid | gateway | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 3/3 | 2422/2422 | 0 | 0 | 39.41 | 5.745 | 6.468 | 8.064 | 17.090 | true | 8 | 93.75 | 0.16 | 17145 | 4 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 304/304 | 0 | 0 | 5.05 | 875.915 | 1009.308 | 1099.251 | 1115.863 | true | 7 | 85.31 | 0.17 | 18979 | 5 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 304/304 | 0 | 0 | 5.02 | 886.317 | 920.613 | 964.899 | 1007.228 | true | 7 | 79.83 | 0.21 | 18396 | 4 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 304/304 | 0 | 0 | 5.01 | 895.441 | 931.307 | 985.042 | 997.378 | true | 8 | 73.54 | 0.18 | 18370 | 7 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/3 | 296/296 | 0 | 0 | 4.91 | 1722.983 | 1801.218 | 1887.389 | 1912.998 | true | 8 | 92.94 | 0.18 | 18853 | 4 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 2/3 | 300/300 | 0 | 0 | 4.92 | 1709.048 | 1772.457 | 1972.885 | 1994.160 | true | 8 | 87.06 | 0.14 | 17606 | 6 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | publish_total | 8 | 3/3 | 296/296 | 0 | 0 | 4.92 | 1745.919 | 1774.196 | 1814.785 | 1828.109 | true | 8 | 91.40 | 0.14 | 18044 | 11 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 1/3 | 1216/1216 | 0 | 0 | 20.18 | 2.648 | 2.797 | 3.690 | 6.308 | true | 7 | 85.31 | 0.17 | 18979 | 5 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 2/3 | 1216/1216 | 0 | 0 | 20.08 | 2.689 | 3.171 | 4.331 | 5.416 | true | 7 | 79.83 | 0.21 | 18396 | 4 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 1216/1216 | 0 | 0 | 20.03 | 2.672 | 3.174 | 3.837 | 12.053 | true | 8 | 73.54 | 0.18 | 18370 | 7 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 1/3 | 1184/1184 | 0 | 0 | 19.65 | 4.843 | 5.513 | 7.175 | 10.684 | true | 8 | 92.94 | 0.18 | 18853 | 4 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 2/3 | 1200/1200 | 0 | 0 | 19.69 | 4.805 | 5.384 | 6.895 | 7.530 | true | 8 | 87.06 | 0.14 | 17606 | 6 |
| hybrid | rpc | mixed-80-20 | hot | cold | compose-middleware-local-services | read | 8 | 3/3 | 1184/1184 | 0 | 0 | 19.67 | 4.871 | 5.413 | 6.885 | 8.508 | true | 8 | 91.40 | 0.14 | 18044 | 11 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 1/3 | 276/276 | 0 | 0 | 4.58 | 974.072 | 1009.855 | 1116.169 | 1130.105 | true | 7 | 78.26 | 0.27 | 17040 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 2/3 | 276/276 | 0 | 0 | 4.58 | 991.727 | 1032.673 | 1065.733 | 1078.668 | true | 4 | 78.35 | 0.19 | 17618 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 4 | 3/3 | 280/280 | 0 | 0 | 4.64 | 973.079 | 994.897 | 1086.919 | 1124.026 | true | 7 | 75.11 | 0.19 | 17613 | 7 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 1/3 | 285/285 | 0 | 0 | 4.67 | 1820.885 | 1848.754 | 1925.245 | 1965.853 | true | 5 | 87.47 | 0.18 | 17905 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 2/3 | 296/296 | 0 | 0 | 4.86 | 1783.547 | 1845.810 | 1913.899 | 1939.185 | true | 7 | 90.61 | 0.19 | 18300 | 11 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | publish_total | 8 | 3/3 | 293/293 | 0 | 0 | 4.81 | 1715.791 | 1732.632 | 1796.807 | 1807.307 | true | 8 | 90.68 | 0.13 | 17858 | 11 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 1/3 | 2484/2484 | 0 | 0 | 41.22 | 2.704 | 3.161 | 3.847 | 15.193 | true | 7 | 78.26 | 0.27 | 17040 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 2/3 | 2484/2484 | 0 | 0 | 41.25 | 2.726 | 3.196 | 3.835 | 12.488 | true | 4 | 78.35 | 0.19 | 17618 | 3 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 4 | 3/3 | 2520/2520 | 0 | 0 | 41.78 | 2.647 | 2.824 | 3.676 | 5.774 | true | 7 | 75.11 | 0.19 | 17613 | 7 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 1/3 | 2566/2566 | 0 | 0 | 42.08 | 4.482 | 5.035 | 6.206 | 11.392 | true | 5 | 87.47 | 0.18 | 17905 | 4 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 2/3 | 2664/2664 | 0 | 0 | 43.78 | 4.709 | 5.285 | 6.414 | 10.584 | true | 7 | 90.61 | 0.19 | 18300 | 11 |
| hybrid | rpc | mixed-90-10 | hot | cold | compose-middleware-local-services | read | 8 | 3/3 | 2637/2637 | 0 | 0 | 43.28 | 4.730 | 5.377 | 6.888 | 12.209 | true | 8 | 90.68 | 0.13 | 17858 | 11 |
