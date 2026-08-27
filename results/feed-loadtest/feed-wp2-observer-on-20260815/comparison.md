# Feed 三策略综合对比：feed-wp2-observer-on-20260815

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 2602.77 | 64.066 | 128 | - | 128 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 128 | 3/3 | 2602.77 | 2600.83 | 2630.54 | 1.14% | 64.066 | 62.454 | 64.123 |

## 原始指标

| strategy | entry | scenario | stage | c | trial | success/total | failed | timeout | QPS | P95(ms) | P99(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 128 | 1 | 52094/52094 | 0 | 0 | 2600.83 | 64.123 | 73.116 | true | 0 | 326.55 | 21189 | 5 |
| hybrid | rpc | distributed-read | read | 128 | 2 | 52705/52705 | 0 | 0 | 2630.54 | 62.454 | 70.572 | true | 0 | 324.00 | 21443 | 4 |
| hybrid | rpc | distributed-read | read | 128 | 3 | 52144/52144 | 0 | 0 | 2602.77 | 64.066 | 74.058 | true | 0 | 330.85 | 21218 | 4 |
