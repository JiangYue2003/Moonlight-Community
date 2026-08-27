# Feed 三策略综合对比：feed-wp2-phase0-strict-20260815

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 2801.37 | 382.092 | 128 | 256 | 512 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 32 | 3/3 | 1945.29 | 1936.89 | 1971.36 | 1.77% | 20.847 | 20.547 | 21.305 |
| hybrid | rpc | distributed-read | read | 64 | 3/3 | 2327.03 | 2318.01 | 2333.48 | 0.66% | 35.336 | 35.302 | 35.703 |
| hybrid | rpc | distributed-read | read | 128 | 3/3 | 2590.31 | 2578.12 | 2609.04 | 1.19% | 64.343 | 63.659 | 65.004 |
| hybrid | rpc | distributed-read | read | 256 | 3/3 | 2793.47 | 2787.75 | 2808.98 | 0.76% | 123.603 | 121.805 | 125.255 |
| hybrid | rpc | distributed-read | read | 512 | 3/3 | 2801.37 | 2779.88 | 2805.25 | 0.91% | 382.092 | 378.928 | 383.916 |

## 原始指标

| strategy | entry | scenario | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 32 | 1/3 | 116236/116236 | 0 | 0 | 1936.89 | 19.913 | 21.305 | 24.874 | 37.218 | true | 0 | 294.33 | 16248 | 7 |
| hybrid | rpc | distributed-read | read | 32 | 2/3 | 116741/116741 | 0 | 0 | 1945.29 | 19.634 | 20.847 | 23.721 | 38.816 | true | 0 | 304.68 | 16288 | 8 |
| hybrid | rpc | distributed-read | read | 32 | 3/3 | 118304/118304 | 0 | 0 | 1971.36 | 19.374 | 20.547 | 23.044 | 37.331 | true | 0 | 297.07 | 16376 | 5 |
| hybrid | rpc | distributed-read | read | 64 | 1/3 | 139660/139660 | 0 | 0 | 2327.03 | 33.279 | 35.336 | 39.950 | 63.239 | true | 0 | 316.95 | 19253 | 5 |
| hybrid | rpc | distributed-read | read | 64 | 2/3 | 139134/139134 | 0 | 0 | 2318.01 | 33.508 | 35.703 | 40.730 | 80.785 | true | 0 | 318.66 | 19010 | 6 |
| hybrid | rpc | distributed-read | read | 64 | 3/3 | 140040/140040 | 0 | 0 | 2333.48 | 33.191 | 35.302 | 40.164 | 75.088 | true | 0 | 326.43 | 19515 | 6 |
| hybrid | rpc | distributed-read | read | 128 | 1/3 | 154788/154788 | 0 | 0 | 2578.12 | 60.835 | 65.004 | 74.692 | 145.664 | true | 0 | 342.90 | 21600 | 8 |
| hybrid | rpc | distributed-read | read | 128 | 2/3 | 156623/156623 | 0 | 0 | 2609.04 | 59.816 | 63.659 | 72.582 | 117.097 | true | 0 | 342.14 | 21365 | 7 |
| hybrid | rpc | distributed-read | read | 128 | 3/3 | 155523/155523 | 0 | 0 | 2590.31 | 60.319 | 64.343 | 73.725 | 134.126 | true | 0 | 333.87 | 22050 | 6 |
| hybrid | rpc | distributed-read | read | 256 | 1/3 | 167432/167432 | 0 | 0 | 2787.75 | 114.180 | 123.603 | 145.376 | 233.852 | true | 0 | 341.99 | 23102 | 9 |
| hybrid | rpc | distributed-read | read | 256 | 2/3 | 168714/168714 | 0 | 0 | 2808.98 | 112.734 | 121.805 | 144.012 | 251.108 | true | 0 | 337.02 | 23053 | 8 |
| hybrid | rpc | distributed-read | read | 256 | 3/3 | 167782/167782 | 0 | 0 | 2793.47 | 115.254 | 125.255 | 148.977 | 351.275 | true | 0 | 343.21 | 23394 | 6 |
| hybrid | rpc | distributed-read | read | 512 | 1/3 | 168718/168718 | 0 | 0 | 2805.25 | 312.925 | 382.092 | 544.774 | 1230.326 | true | 0 | 352.13 | 23316 | 7 |
| hybrid | rpc | distributed-read | read | 512 | 2/3 | 167191/167191 | 0 | 0 | 2779.88 | 314.994 | 383.916 | 544.696 | 1209.360 | true | 0 | 354.65 | 23293 | 5 |
| hybrid | rpc | distributed-read | read | 512 | 3/3 | 168527/168527 | 0 | 0 | 2801.37 | 309.743 | 378.928 | 536.935 | 1602.948 | true | 0 | 349.34 | 23445 | 6 |
