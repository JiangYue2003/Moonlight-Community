# Feed 三策略综合对比：feed-cursor-deep-v1-sequential-formal-20260817a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 1802.54 | 14.753 | 16 | 32 | 32 | - |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 4905.80 | 5.496 | 16 | 32 | 32 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1802.54 | 1264.93 | 1885.54 | 34.43% | 14.753 | 14.541 | 21.365 |
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 1562.01 | 1357.44 | 1618.88 | 16.74% | 32.370 | 30.915 | 38.084 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 4905.80 | 4338.04 | 4933.34 | 12.13% | 5.496 | 5.485 | 6.421 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 4881.87 | 4868.95 | 4922.10 | 1.09% | 10.418 | 10.280 | 10.503 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 113450/113450 | 0 | 0 | 1885.54 | 12.834 | 14.753 | 19.894 | 40.926 | true | 0 | 185.82 | 3.34 | 39737 | 4 |
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 76150/76150 | 0 | 0 | 1264.93 | 18.617 | 21.365 | 27.891 | 50.832 | true | 0 | 176.35 | 2.81 | 24057 | 4 |
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 108550/108550 | 0 | 0 | 1802.54 | 12.616 | 14.541 | 19.073 | 45.293 | true | 0 | 173.71 | 3.41 | 32837 | 3 |
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 97850/97850 | 0 | 0 | 1618.88 | 27.454 | 30.915 | 40.700 | 81.344 | true | 0 | 174.87 | 3.40 | 28851 | 4 |
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 82050/82050 | 0 | 0 | 1357.44 | 33.660 | 38.084 | 49.551 | 98.023 | true | 0 | 181.68 | 3.12 | 25133 | 4 |
| hybrid | gateway | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 94650/94650 | 0 | 0 | 1562.01 | 28.633 | 32.370 | 42.453 | 79.299 | true | 0 | 193.23 | 3.22 | 26984 | 4 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 260600/260600 | 0 | 0 | 4338.04 | 5.471 | 6.421 | 8.736 | 24.558 | true | 0 | 254.79 | 3.84 | 79553 | 3 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 294650/294650 | 0 | 0 | 4905.80 | 4.784 | 5.485 | 7.023 | 14.215 | true | 0 | 261.81 | 4.33 | 80108 | 4 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 296250/296250 | 0 | 0 | 4933.34 | 4.782 | 5.496 | 7.106 | 15.482 | true | 0 | 249.11 | 4.52 | 82133 | 3 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 292850/292850 | 0 | 0 | 4868.95 | 8.978 | 10.503 | 13.605 | 30.462 | true | 0 | 248.62 | 4.55 | 81114 | 3 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 295850/295850 | 0 | 0 | 4922.10 | 8.778 | 10.280 | 13.438 | 36.572 | true | 0 | 256.13 | 4.21 | 80286 | 3 |
| hybrid | rpc | sequential-50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 293750/293750 | 0 | 0 | 4881.87 | 8.894 | 10.418 | 13.507 | 26.857 | true | 0 | 282.64 | 4.37 | 80889 | 3 |
