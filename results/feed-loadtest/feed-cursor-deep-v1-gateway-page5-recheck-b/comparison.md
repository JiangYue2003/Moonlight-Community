# Feed 三策略综合对比：feed-cursor-deep-v1-gateway-page5-recheck-b

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 1437.08 | 18.250 | 16 | - | 16 | - |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 1603.49 | 17.195 | 16 | - | 16 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 1437.08 | 1437.08 | 1437.08 | 0.00% | 18.250 | 18.250 | 18.250 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 1603.49 | 1603.49 | 1603.49 | 0.00% | 17.195 | 17.195 | 17.195 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 14379/14379 | 0 | 0 | 1437.08 | 15.336 | 18.250 | 22.612 | 33.126 | true | 0 | 137.69 | 2.83 | 27930 | 3 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 16047/16047 | 0 | 0 | 1603.49 | 14.868 | 17.195 | 25.873 | 42.181 | true | 0 | 148.58 | 2.14 | 13108 | 3 |
