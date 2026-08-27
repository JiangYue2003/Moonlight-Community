# Feed 三策略综合对比：feed-wp2-observer-off-20260815

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 0.00 | 0.000 | - | - | - | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 128 | 0/3 | 2636.07 | 2635.39 | 2640.07 | 0.18% | 62.935 | 62.709 | 63.564 |

## 原始指标

| strategy | entry | scenario | stage | c | trial | success/total | failed | timeout | QPS | P95(ms) | P99(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 128 | 1 | 52882/52882 | 0 | 0 | 2640.07 | 62.709 | 70.598 | false | 0 | 318.35 | 21388 | 5 |
| hybrid | rpc | distributed-read | read | 128 | 2 | 52804/52804 | 0 | 0 | 2636.07 | 62.935 | 72.339 | false | 0 | 327.60 | 21776 | 8 |
| hybrid | rpc | distributed-read | read | 128 | 3 | 52782/52782 | 0 | 0 | 2635.39 | 63.564 | 73.547 | false | 0 | 327.86 | 21771 | 5 |
