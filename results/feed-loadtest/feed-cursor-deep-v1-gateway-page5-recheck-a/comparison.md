# Feed 三策略综合对比：feed-cursor-deep-v1-gateway-page5-recheck-a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 1658.11 | 15.642 | 16 | - | 16 | - |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 1985.64 | 12.547 | 16 | - | 16 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 1658.11 | 1658.11 | 1658.11 | 0.00% | 15.642 | 15.642 | 15.642 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 1985.64 | 1985.64 | 1985.64 | 0.00% | 12.547 | 12.547 | 12.547 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 16591/16591 | 0 | 0 | 1658.11 | 13.271 | 15.642 | 19.175 | 29.486 | true | 0 | 140.25 | 3.00 | 32988 | 3 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 19868/19868 | 0 | 0 | 1985.64 | 11.520 | 12.547 | 14.967 | 21.988 | true | 0 | 212.64 | 2.53 | 14678 | 3 |
