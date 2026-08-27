# Feed 三策略综合对比：feed-cursor-deep-v1-gateway-page5-recheck-c

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 1628.25 | 15.914 | 16 | - | 16 | - |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 1428.29 | 18.509 | 16 | - | 16 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 1628.25 | 1628.25 | 1628.25 | 0.00% | 15.914 | 15.914 | 15.914 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 1428.29 | 1428.29 | 1428.29 | 0.00% | 18.509 | 18.509 | 18.509 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 16290/16290 | 0 | 0 | 1628.25 | 13.517 | 15.914 | 19.495 | 27.720 | true | 0 | 158.19 | 2.47 | 31735 | 3 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 14291/14291 | 0 | 0 | 1428.29 | 15.388 | 18.509 | 26.361 | 39.240 | true | 0 | 167.31 | 1.99 | 11775 | 3 |
