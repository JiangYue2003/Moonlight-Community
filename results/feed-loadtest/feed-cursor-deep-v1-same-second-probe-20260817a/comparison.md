# Feed 三策略综合对比：feed-cursor-deep-v1-same-second-probe-20260817a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | same-second-cursor | cursor-deep | cold | compose-middleware-local-services | read | 816.88 | 1.651 | 1 | - | 1 | - |
| hybrid | rpc | same-second-cursor | cursor-deep | cold | compose-middleware-local-services | read | 1122.57 | 1.172 | 1 | - | 1 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | same-second-cursor | cursor-deep | cold | compose-middleware-local-services | read | 1 | 1/1 | 816.88 | 816.88 | 816.88 | 0.00% | 1.651 | 1.651 | 1.651 |
| hybrid | rpc | same-second-cursor | cursor-deep | cold | compose-middleware-local-services | read | 1 | 1/1 | 1122.57 | 1122.57 | 1122.57 | 0.00% | 1.172 | 1.172 | 1.172 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | same-second-cursor | cursor-deep | cold | compose-middleware-local-services | read | 1 | 1/1 | 4085/4085 | 0 | 0 | 816.88 | 1.614 | 1.651 | 2.071 | 12.176 | true | 0 | 49.41 | 1.02 | 15512 | 3 |
| hybrid | rpc | same-second-cursor | cursor-deep | cold | compose-middleware-local-services | read | 1 | 1/1 | 5614/5614 | 0 | 0 | 1122.57 | 1.131 | 1.172 | 1.749 | 24.562 | true | 0 | 72.07 | 1.50 | 21586 | 3 |
