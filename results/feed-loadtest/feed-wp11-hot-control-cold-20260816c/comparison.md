# Feed 三策略综合对比：feed-wp11-hot-control-cold-20260816c

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 2530.96 | 15.972 | 32 | 64 | 128 | - |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 2393.49 | 33.630 | 32 | 64 | 128 | - |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 2888.83 | 27.965 | 32 | 64 | 256 | - |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 2963.71 | 13.240 | 32 | 64 | 256 | - |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 3053.70 | 25.938 | 32 | 64 | 256 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 2530.96 | 2500.42 | 2542.15 | 1.65% | 15.972 | 15.794 | 16.200 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 2500.40 | 2432.45 | 2519.94 | 3.50% | 31.621 | 31.328 | 33.002 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 2471.18 | 2461.42 | 2472.74 | 0.46% | 63.110 | 62.819 | 63.766 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 256 | 1/1 | 2257.21 | 2257.21 | 2257.21 | 0.00% | 238.347 | 238.347 | 238.347 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 2379.29 | 2348.97 | 2513.10 | 6.90% | 17.029 | 16.021 | 17.345 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 2393.49 | 2389.49 | 2434.44 | 1.88% | 33.630 | 33.012 | 34.719 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 2389.34 | 2353.71 | 2413.10 | 2.49% | 67.923 | 66.467 | 68.651 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 256 | 1/1 | 2346.52 | 2346.52 | 2346.52 | 0.00% | 210.460 | 210.460 | 210.460 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 2829.27 | 2790.93 | 2934.47 | 5.07% | 13.900 | 13.689 | 14.221 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 2888.83 | 2832.45 | 2889.63 | 1.98% | 27.965 | 27.454 | 28.388 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 2843.69 | 2792.23 | 2901.13 | 3.83% | 57.450 | 54.026 | 57.650 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 256 | 3/3 | 2841.04 | 2829.02 | 2920.68 | 3.23% | 181.893 | 174.410 | 184.337 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 2963.71 | 2933.24 | 2977.89 | 1.51% | 13.240 | 13.133 | 13.472 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 2959.40 | 2956.11 | 2959.65 | 0.12% | 26.029 | 25.861 | 26.105 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 2928.31 | 2926.60 | 2961.24 | 1.18% | 51.328 | 51.017 | 51.611 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 256 | 3/3 | 2791.79 | 2675.85 | 2913.55 | 8.51% | 190.457 | 175.211 | 194.721 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 3028.91 | 2997.83 | 3041.62 | 1.45% | 13.401 | 13.290 | 13.669 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 3053.70 | 3046.77 | 3076.01 | 0.96% | 25.938 | 25.814 | 26.184 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 3045.85 | 3033.51 | 3074.93 | 1.36% | 51.132 | 50.893 | 51.170 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 256 | 3/3 | 3033.47 | 2991.85 | 3040.46 | 1.60% | 168.118 | 168.001 | 170.445 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 32 | 1/3 | 150049/150049 | 0 | 0 | 2500.42 | 15.062 | 16.200 | 19.384 | 36.937 | true | 0 | 256.40 | 3.60 | 22746 | 6 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 32 | 2/3 | 151879/151879 | 0 | 0 | 2530.96 | 14.920 | 15.972 | 18.624 | 30.157 | true | 0 | 264.44 | 3.65 | 22690 | 6 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 152548/152548 | 0 | 0 | 2542.15 | 14.771 | 15.794 | 18.454 | 30.940 | true | 0 | 274.74 | 3.72 | 23045 | 6 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 64 | 1/3 | 151240/151240 | 0 | 0 | 2519.94 | 29.418 | 31.328 | 35.994 | 55.526 | true | 0 | 267.43 | 3.64 | 22730 | 5 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 64 | 2/3 | 146011/146011 | 0 | 0 | 2432.45 | 30.864 | 33.002 | 38.483 | 63.194 | true | 0 | 261.02 | 3.76 | 22053 | 4 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 150063/150063 | 0 | 0 | 2500.40 | 29.602 | 31.621 | 36.922 | 64.272 | true | 0 | 268.34 | 3.88 | 22460 | 6 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 128 | 1/3 | 148455/148455 | 0 | 0 | 2472.74 | 59.608 | 63.110 | 73.891 | 114.908 | true | 0 | 265.88 | 3.82 | 22607 | 5 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 128 | 2/3 | 148360/148360 | 0 | 0 | 2471.18 | 59.006 | 62.819 | 71.167 | 108.257 | true | 0 | 274.75 | 3.94 | 22046 | 6 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 147780/147780 | 0 | 0 | 2461.42 | 60.085 | 63.766 | 72.737 | 105.325 | true | 0 | 279.54 | 3.80 | 22326 | 5 |
| hybrid | gateway | distributed-read | hot | cold | compose-middleware-local-services | read | 256 | 1/3 | 135590/135590 | 0 | 0 | 2257.21 | 191.003 | 238.347 | 375.343 | 1313.465 | true | 0 | 274.38 | 3.49 | 22251 | 5 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 32 | 1/3 | 142782/142782 | 0 | 0 | 2379.29 | 15.852 | 17.345 | 22.663 | 128.659 | true | 0 | 279.18 | 3.46 | 22270 | 5 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 32 | 2/3 | 140963/140963 | 0 | 0 | 2348.97 | 15.368 | 17.029 | 76.127 | 120.951 | true | 0 | 244.78 | 3.49 | 21290 | 5 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 150806/150806 | 0 | 0 | 2513.10 | 14.881 | 16.021 | 19.382 | 39.535 | true | 0 | 258.21 | 3.72 | 22332 | 5 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 64 | 1/3 | 143408/143408 | 0 | 0 | 2389.49 | 30.354 | 33.012 | 96.426 | 132.672 | true | 0 | 273.68 | 3.60 | 22040 | 6 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 64 | 2/3 | 146111/146111 | 0 | 0 | 2434.44 | 31.103 | 34.719 | 41.207 | 132.962 | true | 0 | 269.08 | 3.70 | 22233 | 5 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 143657/143657 | 0 | 0 | 2393.49 | 30.323 | 33.630 | 93.204 | 125.667 | true | 0 | 278.55 | 3.61 | 21549 | 5 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 128 | 1/3 | 144872/144872 | 0 | 0 | 2413.10 | 61.635 | 66.467 | 81.004 | 226.163 | true | 0 | 264.52 | 3.63 | 22012 | 6 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 128 | 2/3 | 143431/143431 | 0 | 0 | 2389.34 | 61.365 | 67.923 | 122.964 | 167.778 | true | 0 | 265.10 | 3.62 | 21944 | 8 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 141302/141302 | 0 | 0 | 2353.71 | 63.386 | 68.651 | 101.778 | 176.750 | true | 0 | 279.67 | 3.73 | 21665 | 6 |
| hybrid | gateway | hot-read | hot | cold | compose-middleware-local-services | read | 256 | 1/3 | 140981/140981 | 0 | 0 | 2346.52 | 175.218 | 210.460 | 294.660 | 696.864 | true | 0 | 264.33 | 3.67 | 21914 | 6 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 32 | 1/3 | 169782/169782 | 0 | 0 | 2829.27 | 12.774 | 13.689 | 23.246 | 100.472 | true | 0 | 306.90 | 4.08 | 25668 | 6 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 32 | 2/3 | 176091/176091 | 0 | 0 | 2934.47 | 12.744 | 13.900 | 16.756 | 41.380 | true | 0 | 334.93 | 4.14 | 26043 | 5 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 167476/167476 | 0 | 0 | 2790.93 | 12.965 | 14.221 | 19.828 | 106.972 | true | 0 | 334.68 | 4.05 | 25111 | 6 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 64 | 1/3 | 173418/173418 | 0 | 0 | 2889.63 | 25.291 | 28.388 | 37.624 | 103.865 | true | 0 | 316.94 | 4.22 | 25597 | 5 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 64 | 2/3 | 169988/169988 | 0 | 0 | 2832.45 | 25.722 | 27.965 | 84.272 | 118.788 | true | 0 | 299.99 | 4.06 | 25566 | 6 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 173375/173375 | 0 | 0 | 2888.83 | 25.585 | 27.454 | 33.073 | 64.428 | true | 0 | 316.46 | 4.08 | 26130 | 5 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 128 | 1/3 | 170714/170714 | 0 | 0 | 2843.69 | 51.036 | 57.450 | 108.314 | 153.384 | true | 0 | 320.06 | 4.22 | 26126 | 5 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 128 | 2/3 | 174147/174147 | 0 | 0 | 2901.13 | 50.614 | 54.026 | 64.796 | 106.980 | true | 0 | 312.11 | 4.15 | 25700 | 7 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 167606/167606 | 0 | 0 | 2792.23 | 51.215 | 57.650 | 119.365 | 162.990 | true | 0 | 296.68 | 4.22 | 25248 | 6 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 256 | 1/3 | 175429/175429 | 0 | 0 | 2920.68 | 144.342 | 174.410 | 244.042 | 515.292 | true | 0 | 325.68 | 4.41 | 25980 | 5 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 256 | 2/3 | 170674/170674 | 0 | 0 | 2841.04 | 150.248 | 181.893 | 256.286 | 572.930 | true | 0 | 320.56 | 4.11 | 26160 | 6 |
| hybrid | rpc | deep-page | hot | cold | compose-middleware-local-services | read | 256 | 3/3 | 170039/170039 | 0 | 0 | 2829.02 | 151.555 | 184.337 | 255.066 | 561.815 | true | 0 | 304.09 | 4.28 | 25388 | 6 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 32 | 1/3 | 176016/176016 | 0 | 0 | 2933.24 | 12.630 | 13.472 | 15.590 | 48.224 | true | 0 | 296.25 | 4.62 | 26140 | 6 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 32 | 2/3 | 177843/177843 | 0 | 0 | 2963.71 | 12.443 | 13.240 | 15.375 | 24.585 | true | 0 | 308.14 | 4.61 | 26178 | 6 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 178693/178693 | 0 | 0 | 2977.89 | 12.308 | 13.133 | 15.239 | 25.206 | true | 0 | 295.05 | 4.73 | 26238 | 6 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 64 | 1/3 | 177620/177620 | 0 | 0 | 2959.65 | 24.507 | 25.861 | 29.792 | 50.691 | true | 0 | 316.13 | 4.77 | 26161 | 6 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 64 | 2/3 | 177412/177412 | 0 | 0 | 2956.11 | 24.636 | 26.029 | 30.072 | 58.475 | true | 0 | 302.86 | 4.69 | 26373 | 6 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 177610/177610 | 0 | 0 | 2959.40 | 24.540 | 26.105 | 30.593 | 48.736 | true | 0 | 290.21 | 4.68 | 26211 | 5 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 128 | 1/3 | 175776/175776 | 0 | 0 | 2928.31 | 49.047 | 51.328 | 60.062 | 113.414 | true | 0 | 318.02 | 4.48 | 25697 | 6 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 128 | 2/3 | 177755/177755 | 0 | 0 | 2961.24 | 48.643 | 51.017 | 61.297 | 92.437 | true | 0 | 299.28 | 4.50 | 26226 | 5 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 175678/175678 | 0 | 0 | 2926.60 | 49.077 | 51.611 | 61.555 | 99.377 | true | 0 | 292.95 | 4.78 | 26335 | 6 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 256 | 1/3 | 160722/160722 | 0 | 0 | 2675.85 | 158.992 | 194.721 | 284.030 | 804.310 | true | 0 | 296.86 | 4.10 | 25623 | 6 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 256 | 2/3 | 175067/175067 | 0 | 0 | 2913.55 | 144.791 | 175.211 | 246.834 | 537.301 | true | 0 | 289.75 | 4.52 | 25701 | 5 |
| hybrid | rpc | distributed-read | hot | cold | compose-middleware-local-services | read | 256 | 3/3 | 167880/167880 | 0 | 0 | 2791.79 | 157.031 | 190.457 | 266.790 | 686.097 | true | 0 | 289.95 | 4.45 | 25352 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 32 | 1/3 | 182516/182516 | 0 | 0 | 3041.62 | 12.438 | 13.290 | 15.419 | 32.439 | true | 0 | 300.00 | 4.22 | 26996 | 6 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 32 | 2/3 | 179890/179890 | 0 | 0 | 2997.83 | 12.730 | 13.669 | 16.242 | 45.918 | true | 0 | 281.70 | 4.63 | 27202 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 32 | 3/3 | 181756/181756 | 0 | 0 | 3028.91 | 12.526 | 13.401 | 15.415 | 25.538 | true | 0 | 297.12 | 4.71 | 27927 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 64 | 1/3 | 182856/182856 | 0 | 0 | 3046.77 | 24.440 | 25.814 | 29.418 | 50.538 | true | 0 | 306.83 | 4.48 | 26898 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 64 | 2/3 | 183261/183261 | 0 | 0 | 3053.70 | 24.655 | 26.184 | 29.971 | 57.719 | true | 0 | 307.40 | 4.68 | 27456 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 64 | 3/3 | 184602/184602 | 0 | 0 | 3076.01 | 24.363 | 25.938 | 30.088 | 50.138 | true | 0 | 302.06 | 4.45 | 27892 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 128 | 1/3 | 182850/182850 | 0 | 0 | 3045.85 | 48.319 | 50.893 | 58.461 | 95.609 | true | 0 | 303.95 | 4.82 | 27616 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 128 | 2/3 | 184580/184580 | 0 | 0 | 3074.93 | 48.234 | 51.132 | 58.212 | 94.879 | true | 0 | 317.04 | 4.55 | 27405 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 128 | 3/3 | 182089/182089 | 0 | 0 | 3033.51 | 48.585 | 51.170 | 58.783 | 91.834 | true | 0 | 301.11 | 4.69 | 27870 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 256 | 1/3 | 182201/182201 | 0 | 0 | 3033.47 | 139.190 | 168.118 | 236.857 | 512.440 | true | 0 | 292.17 | 4.69 | 26916 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 256 | 2/3 | 182614/182614 | 0 | 0 | 3040.46 | 138.750 | 168.001 | 237.014 | 474.918 | true | 0 | 306.60 | 4.74 | 27208 | 5 |
| hybrid | rpc | hot-read | hot | cold | compose-middleware-local-services | read | 256 | 3/3 | 179733/179733 | 0 | 0 | 2991.85 | 141.096 | 170.445 | 240.071 | 559.201 | true | 0 | 295.32 | 4.53 | 27160 | 5 |
