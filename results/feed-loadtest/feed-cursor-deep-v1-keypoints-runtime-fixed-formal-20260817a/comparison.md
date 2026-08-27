# Feed 三策略综合对比：feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 1698.16 | 15.719 | 16 | - | 16 | - |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 2180.45 | 11.369 | 16 | - | 16 | - |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 1045.19 | 25.413 | 16 | - | 16 | - |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 489.23 | 49.020 | 16 | - | 16 | - |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 4848.76 | 5.381 | 16 | - | 16 | - |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 7621.72 | 3.359 | 16 | - | 16 | - |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 1178.84 | 22.672 | 16 | - | 16 | - |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 592.37 | 35.660 | 16 | - | 16 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1698.16 | 1233.01 | 1788.88 | 32.73% | 15.719 | 14.393 | 22.187 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 2180.45 | 1575.93 | 2215.37 | 29.33% | 11.369 | 11.234 | 16.041 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1045.19 | 1030.87 | 1082.45 | 4.94% | 25.413 | 24.140 | 25.788 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 489.23 | 483.72 | 489.49 | 1.18% | 49.020 | 47.627 | 49.203 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 4848.76 | 4802.48 | 4849.21 | 0.96% | 5.381 | 5.379 | 5.420 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 7621.72 | 7580.98 | 7736.26 | 2.04% | 3.359 | 3.269 | 3.366 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1178.84 | 1172.87 | 1183.41 | 0.89% | 22.672 | 22.506 | 22.982 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 592.37 | 589.21 | 595.52 | 1.07% | 35.660 | 35.445 | 35.828 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 101899/101899 | 0 | 0 | 1698.16 | 13.693 | 15.719 | 20.340 | 35.884 | true | 0 | 164.00 | 3.25 | 31845 | 3 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 73986/73986 | 0 | 0 | 1233.01 | 19.222 | 22.187 | 28.588 | 56.590 | true | 0 | 161.14 | 2.98 | 23033 | 3 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 107342/107342 | 0 | 0 | 1788.88 | 12.423 | 14.393 | 18.254 | 31.427 | true | 0 | 170.33 | 3.61 | 31357 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 132931/132931 | 0 | 0 | 2215.37 | 9.905 | 11.234 | 14.292 | 25.229 | true | 0 | 221.92 | 4.99 | 36938 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 94571/94571 | 0 | 0 | 1575.93 | 14.243 | 16.041 | 20.424 | 37.860 | true | 0 | 202.21 | 4.00 | 26201 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 130839/130839 | 0 | 0 | 2180.45 | 9.947 | 11.369 | 14.233 | 30.630 | true | 0 | 225.00 | 4.55 | 34504 | 3 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 64960/64960 | 0 | 0 | 1082.45 | 21.925 | 24.140 | 28.458 | 45.992 | true | 0 | 252.32 | 1.56 | 8700 | 3 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 62719/62719 | 0 | 0 | 1045.19 | 22.486 | 25.413 | 31.215 | 55.559 | true | 0 | 241.75 | 1.52 | 8683 | 3 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 61862/61862 | 0 | 0 | 1030.87 | 22.652 | 25.788 | 32.813 | 59.012 | true | 0 | 238.23 | 1.39 | 8278 | 4 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 29361/29361 | 0 | 0 | 489.23 | 44.791 | 49.020 | 58.321 | 84.274 | true | 0 | 228.47 | 0.86 | 4317 | 5 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 29379/29379 | 0 | 0 | 489.49 | 43.790 | 47.627 | 55.380 | 80.744 | true | 0 | 222.98 | 0.80 | 4641 | 3 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 29036/29036 | 0 | 0 | 483.72 | 45.010 | 49.203 | 57.086 | 82.085 | true | 0 | 238.58 | 0.80 | 4287 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 290962/290962 | 0 | 0 | 4849.21 | 4.746 | 5.381 | 6.803 | 11.911 | true | 0 | 236.93 | 4.64 | 80035 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 288164/288164 | 0 | 0 | 4802.48 | 4.776 | 5.420 | 6.881 | 13.547 | true | 0 | 237.06 | 4.56 | 78173 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 290939/290939 | 0 | 0 | 4848.76 | 4.742 | 5.379 | 6.705 | 16.465 | true | 0 | 242.33 | 4.72 | 79246 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 454871/454871 | 0 | 0 | 7580.98 | 3.125 | 3.359 | 4.350 | 11.763 | true | 0 | 333.57 | 8.84 | 120914 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 457318/457318 | 0 | 0 | 7621.72 | 3.125 | 3.366 | 4.367 | 11.850 | true | 0 | 348.37 | 8.86 | 120356 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 464187/464187 | 0 | 0 | 7736.26 | 2.889 | 3.269 | 4.252 | 10.613 | true | 0 | 332.72 | 8.44 | 118001 | 4 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 71014/71014 | 0 | 0 | 1183.41 | 20.711 | 22.506 | 27.034 | 44.029 | true | 0 | 267.43 | 1.84 | 9107 | 4 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 70742/70742 | 0 | 0 | 1178.84 | 20.784 | 22.672 | 26.804 | 51.789 | true | 0 | 250.43 | 1.73 | 9214 | 3 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 70376/70376 | 0 | 0 | 1172.87 | 20.878 | 22.982 | 27.634 | 56.814 | true | 0 | 250.31 | 1.73 | 9155 | 4 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 35553/35553 | 0 | 0 | 592.37 | 33.081 | 35.660 | 41.018 | 57.388 | true | 0 | 230.49 | 0.89 | 5131 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 35739/35739 | 0 | 0 | 595.52 | 32.900 | 35.445 | 42.058 | 70.172 | true | 0 | 235.40 | 0.88 | 4945 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 35361/35361 | 0 | 0 | 589.21 | 33.554 | 35.828 | 40.942 | 57.317 | true | 0 | 253.24 | 0.84 | 5151 | 3 |
