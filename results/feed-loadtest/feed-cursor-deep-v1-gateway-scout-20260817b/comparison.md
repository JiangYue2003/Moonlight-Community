# Feed 三策略综合对比：feed-cursor-deep-v1-gateway-scout-20260817b

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 2070.68 | 12.414 | 16 | 32 | 32 | - |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 1135.41 | 45.684 | 16 | 32 | 32 | - |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 2253.28 | 11.142 | 16 | 32 | 32 | - |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 2718.30 | 9.396 | 16 | 32 | 32 | - |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 889.58 | 56.044 | 32 | - | 32 | - |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 1245.24 | 39.594 | 32 | - | 32 | - |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 438.12 | 57.347 | 16 | 32 | 32 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 2070.68 | 2070.68 | 2070.68 | 0.00% | 12.414 | 12.414 | 12.414 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 1764.53 | 1764.53 | 1764.53 | 0.00% | 26.752 | 26.752 | 26.752 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 1125.35 | 1125.35 | 1125.35 | 0.00% | 23.879 | 23.879 | 23.879 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 1135.41 | 1135.41 | 1135.41 | 0.00% | 45.684 | 45.684 | 45.684 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 2253.28 | 2253.28 | 2253.28 | 0.00% | 11.142 | 11.142 | 11.142 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 1798.07 | 1798.07 | 1798.07 | 0.00% | 27.291 | 27.291 | 27.291 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 2718.30 | 2718.30 | 2718.30 | 0.00% | 9.396 | 9.396 | 9.396 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 2103.51 | 2103.51 | 2103.51 | 0.00% | 21.687 | 21.687 | 21.687 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 758.55 | 758.55 | 758.55 | 0.00% | 38.478 | 38.478 | 38.478 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 889.58 | 889.58 | 889.58 | 0.00% | 56.044 | 56.044 | 56.044 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 1127.76 | 1127.76 | 1127.76 | 0.00% | 24.463 | 24.463 | 24.463 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 1245.24 | 1245.24 | 1245.24 | 0.00% | 39.594 | 39.594 | 39.594 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 438.12 | 438.12 | 438.12 | 0.00% | 57.347 | 57.347 | 57.347 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 398.56 | 398.56 | 398.56 | 0.00% | 108.569 | 108.569 | 108.569 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 20719/20719 | 0 | 0 | 2070.68 | 10.516 | 12.414 | 15.301 | 26.406 | true | 0 | 167.80 | 3.53 | 33535 | 3 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 17667/17667 | 0 | 0 | 1764.53 | 24.227 | 26.752 | 36.415 | 51.976 | true | 0 | 172.60 | 3.73 | 27536 | 3 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 11262/11262 | 0 | 0 | 1125.35 | 19.681 | 23.879 | 30.205 | 40.743 | true | 0 | 128.57 | 2.12 | 22226 | 3 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 11378/11378 | 0 | 0 | 1135.41 | 40.383 | 45.684 | 59.833 | 83.567 | true | 0 | 132.38 | 2.86 | 21871 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 22542/22542 | 0 | 0 | 2253.28 | 9.986 | 11.142 | 14.375 | 20.385 | true | 0 | 221.35 | 5.20 | 33739 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 18000/18000 | 0 | 0 | 1798.07 | 24.623 | 27.291 | 35.963 | 59.180 | true | 0 | 206.03 | 4.54 | 27526 | 3 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 27194/27194 | 0 | 0 | 2718.30 | 8.374 | 9.396 | 11.378 | 18.203 | true | 0 | 217.46 | 5.32 | 20643 | 3 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 21053/21053 | 0 | 0 | 2103.51 | 19.123 | 21.687 | 31.208 | 46.057 | true | 0 | 173.76 | 4.27 | 15752 | 3 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 7598/7598 | 0 | 0 | 758.55 | 29.756 | 38.478 | 48.810 | 77.862 | true | 0 | 171.05 | 1.18 | 6586 | 3 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 8921/8921 | 0 | 0 | 889.58 | 49.311 | 56.044 | 80.402 | 132.177 | true | 0 | 201.71 | 1.38 | 7162 | 3 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 11287/11287 | 0 | 0 | 1127.76 | 19.416 | 24.463 | 31.037 | 43.790 | true | 0 | 151.34 | 1.70 | 8733 | 3 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 12464/12464 | 0 | 0 | 1245.24 | 35.360 | 39.594 | 59.508 | 89.390 | true | 0 | 164.96 | 1.89 | 10873 | 3 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 4389/4389 | 0 | 0 | 438.12 | 53.059 | 57.347 | 68.248 | 84.465 | true | 0 | 201.33 | 0.75 | 3624 | 5 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 4001/4001 | 0 | 0 | 398.56 | 103.168 | 108.569 | 117.323 | 153.322 | true | 0 | 215.18 | 0.80 | 3576 | 3 |
