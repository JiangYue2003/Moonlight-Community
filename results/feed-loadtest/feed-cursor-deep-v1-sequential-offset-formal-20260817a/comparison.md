# Feed 三策略综合对比：feed-cursor-deep-v1-sequential-offset-formal-20260817a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 776.33 | 37.328 | 16 | 32 | 32 | - |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 999.06 | 56.482 | 16 | 32 | 32 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 776.33 | 771.07 | 871.46 | 12.93% | 37.328 | 33.196 | 40.855 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 643.01 | 552.32 | 673.29 | 18.81% | 89.294 | 87.197 | 103.425 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 981.85 | 971.29 | 983.31 | 1.22% | 28.836 | 28.641 | 29.124 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 999.06 | 995.44 | 999.18 | 0.37% | 56.482 | 56.022 | 56.734 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 52750/52750 | 0 | 0 | 871.46 | 29.120 | 33.196 | 42.315 | 86.904 | true | 0 | 276.33 | 1.23 | 7343 | 4 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 46400/46400 | 0 | 0 | 771.07 | 35.102 | 40.855 | 52.790 | 113.726 | true | 0 | 215.49 | 1.19 | 7179 | 3 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 47200/47200 | 0 | 0 | 776.33 | 33.021 | 37.328 | 46.550 | 85.712 | true | 0 | 223.13 | 1.30 | 7067 | 4 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 41550/41550 | 0 | 0 | 673.29 | 77.387 | 87.197 | 105.857 | 163.159 | true | 0 | 220.09 | 1.40 | 6165 | 3 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 33600/33600 | 0 | 0 | 552.32 | 93.046 | 103.425 | 124.816 | 202.931 | true | 0 | 196.29 | 1.25 | 5537 | 4 |
| hybrid | gateway | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 39500/39500 | 0 | 0 | 643.01 | 80.116 | 89.294 | 106.436 | 155.606 | true | 0 | 205.76 | 1.31 | 5997 | 5 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 59200/59200 | 0 | 0 | 981.85 | 25.724 | 28.641 | 34.647 | 58.569 | true | 0 | 253.65 | 1.44 | 7902 | 5 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 59200/59200 | 0 | 0 | 983.31 | 25.913 | 28.836 | 34.638 | 53.194 | true | 0 | 250.23 | 1.36 | 8206 | 3 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 58550/58550 | 0 | 0 | 971.29 | 26.062 | 29.124 | 35.240 | 54.195 | true | 0 | 245.17 | 1.53 | 7878 | 3 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 60750/60750 | 0 | 0 | 995.44 | 50.654 | 56.482 | 68.448 | 118.172 | true | 0 | 294.88 | 1.32 | 8320 | 3 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 60800/60800 | 0 | 0 | 999.06 | 50.585 | 56.022 | 67.603 | 117.644 | true | 0 | 283.70 | 1.27 | 8297 | 3 |
| hybrid | rpc | sequential-page-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 60750/60750 | 0 | 0 | 999.18 | 50.462 | 56.734 | 69.489 | 118.059 | true | 0 | 269.54 | 1.34 | 8131 | 4 |
