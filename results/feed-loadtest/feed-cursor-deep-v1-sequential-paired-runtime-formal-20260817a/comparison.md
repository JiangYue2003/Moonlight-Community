# Feed 三策略综合对比：feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 1797.02 | 14.378 | 16 | - | 16 | - |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 801.84 | 36.980 | 16 | - | 16 | - |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 4974.63 | 5.425 | 16 | - | 16 | - |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 990.99 | 28.017 | 16 | - | 16 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1797.02 | 1788.37 | 2203.83 | 23.12% | 14.378 | 11.605 | 14.642 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 801.84 | 767.01 | 925.37 | 19.75% | 36.980 | 30.230 | 37.925 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 4974.63 | 4934.19 | 5028.57 | 1.90% | 5.425 | 5.376 | 5.472 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 990.99 | 988.78 | 998.96 | 1.03% | 28.017 | 27.914 | 28.651 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 132700/132700 | 0 | 0 | 2203.83 | 10.033 | 11.605 | 14.882 | 25.980 | true | 0 | 186.79 | 4.01 | 37442 | 4 |
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 108300/108300 | 0 | 0 | 1797.02 | 12.370 | 14.378 | 18.533 | 33.040 | true | 0 | 183.47 | 3.57 | 31175 | 3 |
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 107800/107800 | 0 | 0 | 1788.37 | 12.530 | 14.642 | 18.737 | 35.305 | true | 0 | 176.70 | 3.38 | 30420 | 3 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 56000/56000 | 0 | 0 | 925.37 | 27.042 | 30.230 | 36.684 | 61.041 | true | 0 | 245.62 | 1.37 | 7385 | 4 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 48700/48700 | 0 | 0 | 801.84 | 32.460 | 36.980 | 46.258 | 76.558 | true | 0 | 236.83 | 1.32 | 7365 | 4 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 46350/46350 | 0 | 0 | 767.01 | 33.358 | 37.925 | 47.448 | 78.739 | true | 0 | 220.97 | 1.26 | 7264 | 4 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 296350/296350 | 0 | 0 | 4934.19 | 4.782 | 5.472 | 7.076 | 14.629 | true | 0 | 247.03 | 4.26 | 83267 | 3 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 302000/302000 | 0 | 0 | 5028.57 | 4.716 | 5.376 | 6.950 | 14.150 | true | 0 | 231.29 | 4.33 | 83029 | 3 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 298750/298750 | 0 | 0 | 4974.63 | 4.743 | 5.425 | 7.013 | 16.260 | true | 0 | 258.33 | 4.31 | 83024 | 3 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 60200/60200 | 0 | 0 | 998.96 | 25.100 | 27.914 | 33.657 | 64.275 | true | 0 | 264.81 | 1.56 | 8276 | 3 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 60000/60000 | 0 | 0 | 988.78 | 25.618 | 28.651 | 34.488 | 59.508 | true | 0 | 260.87 | 1.48 | 8152 | 3 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 60000/60000 | 0 | 0 | 990.99 | 25.291 | 28.017 | 33.661 | 58.932 | true | 0 | 257.78 | 1.43 | 7935 | 3 |
