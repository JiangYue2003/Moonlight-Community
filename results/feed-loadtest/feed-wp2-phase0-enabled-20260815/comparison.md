# Feed 三策略综合对比：feed-wp2-phase0-enabled-20260815

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 2820.01 | 125.638 | 256 | 512 | 512 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 32 | 3/3 | 1963.39 | 1947.33 | 1964.65 | 0.88% | 20.738 | 20.724 | 20.894 |
| hybrid | rpc | distributed-read | read | 64 | 3/3 | 2284.94 | 2230.58 | 2292.16 | 2.70% | 36.185 | 36.108 | 37.341 |
| hybrid | rpc | distributed-read | read | 128 | 3/3 | 2559.22 | 2554.12 | 2573.84 | 0.77% | 64.815 | 64.609 | 65.044 |
| hybrid | rpc | distributed-read | read | 256 | 3/3 | 2820.01 | 2817.85 | 2829.55 | 0.41% | 125.638 | 123.626 | 126.886 |
| hybrid | rpc | distributed-read | read | 512 | 3/3 | 2793.31 | 2789.98 | 2808.89 | 0.68% | 390.378 | 389.646 | 395.408 |

## 原始指标

| strategy | entry | scenario | stage | c | trial | success/total | failed | timeout | QPS | P95(ms) | P99(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 32 | 1 | 117898/117898 | 0 | 0 | 1964.65 | 20.724 | 23.718 | true | 0 | 300.77 | 16366 | 5 |
| hybrid | rpc | distributed-read | read | 32 | 2 | 117822/117822 | 0 | 0 | 1963.39 | 20.738 | 23.413 | true | 0 | 298.44 | 16155 | 5 |
| hybrid | rpc | distributed-read | read | 32 | 3 | 116862/116862 | 0 | 0 | 1947.33 | 20.894 | 23.714 | true | 0 | 298.33 | 16263 | 5 |
| hybrid | rpc | distributed-read | read | 64 | 1 | 137575/137575 | 0 | 0 | 2292.16 | 36.185 | 40.861 | true | 0 | 317.06 | 18740 | 7 |
| hybrid | rpc | distributed-read | read | 64 | 2 | 137140/137140 | 0 | 0 | 2284.94 | 36.108 | 40.566 | true | 0 | 320.88 | 18722 | 6 |
| hybrid | rpc | distributed-read | read | 64 | 3 | 133886/133886 | 0 | 0 | 2230.58 | 37.341 | 42.600 | true | 0 | 325.21 | 18751 | 5 |
| hybrid | rpc | distributed-read | read | 128 | 1 | 153622/153622 | 0 | 0 | 2559.22 | 64.609 | 73.778 | true | 0 | 329.21 | 21031 | 5 |
| hybrid | rpc | distributed-read | read | 128 | 2 | 153313/153313 | 0 | 0 | 2554.12 | 65.044 | 74.852 | true | 0 | 328.64 | 21100 | 6 |
| hybrid | rpc | distributed-read | read | 128 | 3 | 154520/154520 | 0 | 0 | 2573.84 | 64.815 | 75.885 | true | 0 | 336.64 | 21345 | 6 |
| hybrid | rpc | distributed-read | read | 256 | 1 | 169243/169243 | 0 | 0 | 2817.85 | 126.886 | 151.972 | true | 0 | 344.77 | 23378 | 7 |
| hybrid | rpc | distributed-read | read | 256 | 2 | 169373/169373 | 0 | 0 | 2820.01 | 125.638 | 150.554 | true | 0 | 354.69 | 23116 | 4 |
| hybrid | rpc | distributed-read | read | 256 | 3 | 169942/169942 | 0 | 0 | 2829.55 | 123.626 | 147.420 | true | 0 | 345.96 | 23336 | 8 |
| hybrid | rpc | distributed-read | read | 512 | 1 | 168940/168940 | 0 | 0 | 2808.89 | 390.378 | 558.838 | true | 0 | 345.03 | 23418 | 7 |
| hybrid | rpc | distributed-read | read | 512 | 2 | 168043/168043 | 0 | 0 | 2793.31 | 395.408 | 562.532 | true | 0 | 347.00 | 23343 | 5 |
| hybrid | rpc | distributed-read | read | 512 | 3 | 167903/167903 | 0 | 0 | 2789.98 | 389.646 | 556.207 | true | 0 | 343.96 | 22982 | 6 |
