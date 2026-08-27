# Feed 三策略综合对比：feed-wp6-combined-enabled-20260815

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 10661.31 | 18.112 | 64 | 128 | 256 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 32 | 3/3 | 7996.33 | 7980.49 | 7997.20 | 0.21% | 5.487 | 5.483 | 5.530 |
| hybrid | rpc | distributed-read | read | 64 | 3/3 | 9866.82 | 9763.14 | 9950.94 | 1.90% | 9.171 | 9.088 | 9.306 |
| hybrid | rpc | distributed-read | read | 128 | 3/3 | 10661.31 | 10638.67 | 10825.93 | 1.76% | 18.112 | 17.462 | 18.267 |
| hybrid | rpc | distributed-read | read | 256 | 3/3 | 10587.21 | 10489.76 | 10651.73 | 1.53% | 33.068 | 32.715 | 33.749 |

## 原始指标

| strategy | entry | scenario | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 32 | 1/3 | 478851/478851 | 0 | 0 | 7980.49 | 5.117 | 5.530 | 6.787 | 25.533 | true | 0 | 422.72 | 56577 | 3 |
| hybrid | rpc | distributed-read | read | 32 | 2/3 | 479808/479808 | 0 | 0 | 7996.33 | 5.087 | 5.483 | 6.604 | 20.547 | true | 0 | 420.10 | 56738 | 3 |
| hybrid | rpc | distributed-read | read | 32 | 3/3 | 479856/479856 | 0 | 0 | 7997.20 | 5.085 | 5.487 | 6.644 | 20.518 | true | 0 | 420.93 | 56177 | 3 |
| hybrid | rpc | distributed-read | read | 64 | 1/3 | 585843/585843 | 0 | 0 | 9763.14 | 8.530 | 9.306 | 11.131 | 41.823 | true | 0 | 470.99 | 70750 | 3 |
| hybrid | rpc | distributed-read | read | 64 | 2/3 | 592057/592057 | 0 | 0 | 9866.82 | 8.441 | 9.171 | 10.979 | 37.541 | true | 0 | 468.54 | 71566 | 3 |
| hybrid | rpc | distributed-read | read | 64 | 3/3 | 597103/597103 | 0 | 0 | 9950.94 | 8.379 | 9.088 | 10.855 | 37.865 | true | 0 | 468.63 | 70857 | 3 |
| hybrid | rpc | distributed-read | read | 128 | 1/3 | 638407/638407 | 0 | 0 | 10638.67 | 16.567 | 18.267 | 21.498 | 63.114 | true | 0 | 533.18 | 79643 | 3 |
| hybrid | rpc | distributed-read | read | 128 | 2/3 | 639778/639778 | 0 | 0 | 10661.31 | 16.501 | 18.112 | 21.190 | 60.715 | true | 0 | 487.67 | 79690 | 3 |
| hybrid | rpc | distributed-read | read | 128 | 3/3 | 649686/649686 | 0 | 0 | 10825.93 | 16.036 | 17.462 | 20.443 | 60.281 | true | 0 | 496.04 | 80510 | 3 |
| hybrid | rpc | distributed-read | read | 256 | 1/3 | 629556/629556 | 0 | 0 | 10489.76 | 31.933 | 33.749 | 36.574 | 133.309 | true | 0 | 476.33 | 80665 | 3 |
| hybrid | rpc | distributed-read | read | 256 | 2/3 | 639298/639298 | 0 | 0 | 10651.73 | 31.021 | 32.715 | 36.092 | 120.840 | true | 0 | 485.72 | 81160 | 3 |
| hybrid | rpc | distributed-read | read | 256 | 3/3 | 635415/635415 | 0 | 0 | 10587.21 | 31.473 | 33.068 | 35.944 | 132.075 | true | 0 | 476.54 | 81109 | 3 |
