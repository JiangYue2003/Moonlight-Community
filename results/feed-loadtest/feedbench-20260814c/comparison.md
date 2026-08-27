# Feed 三策略综合对比：feedbench-20260814c

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | distributed-read | read | 953.19 | 42.968 | 32 | 128 | 128 | - |
| hybrid | gateway | hot-read | read | 1005.29 | 160.612 | 32 | 128 | 128 | - |
| hybrid | gateway | mixed-80-20 | publish_total | 8.63 | 3212.314 | 128 | - | 128 | - |
| hybrid | gateway | mixed-80-20 | read | 706.16 | 176.268 | 32 | 128 | 128 | - |
| hybrid | gateway | mixed-90-10 | publish_total | 7.26 | 1928.023 | 128 | - | 128 | - |
| hybrid | gateway | mixed-90-10 | read | 711.26 | 48.349 | 32 | 128 | 128 | - |
| hybrid | gateway | publish | publish_total | 10.01 | 3240.688 | 32 | 128 | 32 | 128 |
| hybrid | rpc | burst-publish | publish_total | 10.13 | 3310.403 | 32 | - | 32 | - |
| hybrid | rpc | burst-read | read | 1544.12 | 106.303 | 128 | - | 128 | - |
| hybrid | rpc | deep-page | read | 1173.70 | 136.211 | 128 | 256 | 256 | - |
| hybrid | rpc | distributed-read | read | 1599.55 | 193.743 | 128 | 256 | 256 | - |
| hybrid | rpc | hot-read | read | 1536.62 | 103.197 | 128 | 256 | 256 | - |
| hybrid | rpc | mixed-80-20 | publish_total | 9.43 | 2839.019 | 128 | 256 | 128 | 256 |
| hybrid | rpc | mixed-80-20 | read | 942.32 | 129.209 | 128 | - | 128 | - |
| hybrid | rpc | mixed-90-10 | publish_total | 8.94 | 2605.246 | 1 | 8 | 256 | - |
| hybrid | rpc | mixed-90-10 | read | 1190.18 | 205.489 | 32 | 128 | 256 | - |
| hybrid | rpc | publish | publish_total | 10.09 | 3242.933 | 32 | 36 | 40 | 44 |
| hybrid | rpc | steady-read | read | 1590.01 | 103.534 | 128 | - | 128 | - |
| pull | gateway | distributed-read | read | 444.36 | 94.320 | 32 | 128 | 128 | - |
| pull | gateway | hot-read | read | 459.64 | 348.206 | 32 | 128 | 128 | - |
| pull | gateway | mixed-80-20 | publish_total | 8.94 | 3147.683 | 128 | - | 128 | - |
| pull | gateway | mixed-80-20 | read | 339.15 | 92.499 | 32 | 128 | 128 | - |
| pull | gateway | mixed-90-10 | publish_total | 7.40 | 1980.571 | 128 | - | 128 | - |
| pull | gateway | mixed-90-10 | read | 361.68 | 426.195 | 32 | 128 | 128 | - |
| pull | gateway | publish | publish_total | 10.69 | 3119.204 | 32 | 128 | 32 | 128 |
| pull | rpc | burst-publish | publish_total | 0.00 | 0.000 | - | 32 | - | 32 |
| pull | rpc | burst-read | read | 620.17 | 265.135 | 128 | - | 128 | - |
| pull | rpc | deep-page | read | 572.70 | 560.720 | 32 | 128 | 256 | - |
| pull | rpc | distributed-read | read | 621.91 | 471.097 | 32 | 128 | 256 | - |
| pull | rpc | hot-read | read | 603.45 | 512.097 | 32 | 128 | 256 | - |
| pull | rpc | mixed-80-20 | publish_total | 9.49 | 2879.121 | 128 | 256 | 128 | 256 |
| pull | rpc | mixed-80-20 | read | 426.61 | 300.058 | 8 | 32 | 128 | - |
| pull | rpc | mixed-90-10 | publish_total | 8.07 | 3008.193 | 128 | 256 | 256 | - |
| pull | rpc | mixed-90-10 | read | 555.39 | 60.987 | 32 | 128 | 256 | - |
| pull | rpc | publish | publish_total | 10.76 | 3029.476 | 32 | 128 | 32 | 128 |
| pull | rpc | steady-read | read | 670.94 | 232.156 | 128 | - | 128 | - |
| push | gateway | distributed-read | read | 1503.17 | 118.268 | 32 | 128 | 128 | - |
| push | gateway | hot-read | read | 1702.76 | 102.446 | 32 | 128 | 128 | - |
| push | gateway | mixed-80-20 | publish_total | 8.79 | 2943.508 | 128 | - | 128 | - |
| push | gateway | mixed-80-20 | read | 1066.94 | 135.654 | 128 | - | 128 | - |
| push | gateway | mixed-90-10 | publish_total | 7.94 | 1653.588 | 128 | - | 128 | - |
| push | gateway | mixed-90-10 | read | 914.25 | 53.178 | 32 | 128 | 128 | - |
| push | gateway | publish | publish_total | 10.02 | 3439.447 | 32 | 128 | 32 | 128 |
| push | rpc | burst-publish | publish_total | 0.00 | 0.000 | - | 32 | - | 32 |
| push | rpc | burst-read | read | 2695.27 | 62.438 | 128 | - | 128 | - |
| push | rpc | deep-page | read | 1782.58 | 224.324 | 128 | 256 | 256 | - |
| push | rpc | distributed-read | read | 2915.53 | 131.227 | 128 | 256 | 256 | - |
| push | rpc | hot-read | read | 2989.43 | 127.779 | 128 | 256 | 256 | - |
| push | rpc | mixed-80-20 | publish_total | 10.39 | 4883.726 | 128 | 256 | 256 | - |
| push | rpc | mixed-80-20 | read | 1563.70 | 135.326 | 32 | 128 | 256 | - |
| push | rpc | mixed-90-10 | publish_total | 9.73 | 2442.066 | 128 | 256 | 256 | - |
| push | rpc | mixed-90-10 | read | 1722.27 | 143.922 | 32 | 128 | 256 | - |
| push | rpc | publish | publish_total | 11.38 | 2985.753 | 32 | 128 | 32 | 128 |
| push | rpc | steady-read | read | 2744.23 | 61.066 | 128 | - | 128 | - |

## 原始指标

| strategy | entry | scenario | stage | c | success/total | failed | timeout | QPS | P95(ms) | P99(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|
| hybrid | gateway | distributed-read | read | 1 | 385/385 | 0 | 0 | 128.18 | 9.557 | 10.588 | true | 0 | 37.62 | 4505 | 3 |
| hybrid | gateway | distributed-read | read | 32 | 2881/2881 | 0 | 0 | 953.19 | 42.968 | 48.129 | true | 0 | 201.01 | 31871 | 3 |
| hybrid | gateway | distributed-read | read | 128 | 2695/2695 | 0 | 0 | 869.06 | 180.420 | 200.639 | true | 0 | 203.85 | 28738 | 3 |
| hybrid | gateway | hot-read | read | 1 | 403/403 | 0 | 0 | 134.24 | 8.840 | 10.011 | true | 0 | 38.46 | 4819 | 3 |
| hybrid | gateway | hot-read | read | 32 | 2996/2996 | 0 | 0 | 986.51 | 41.432 | 48.175 | true | 0 | 209.29 | 33258 | 3 |
| hybrid | gateway | hot-read | read | 128 | 3101/3101 | 0 | 0 | 1005.29 | 160.612 | 190.657 | true | 0 | 205.98 | 34568 | 3 |
| hybrid | gateway | mixed-80-20 | publish_total | 1 | 60/60 | 0 | 0 | 1.37 | 779.748 | 814.446 | true | 2 | 51.97 | 3643 | 3 |
| hybrid | gateway | mixed-80-20 | publish_total | 32 | 60/60 | 0 | 0 | 6.75 | 1311.359 | 1320.570 | true | 6 | 85.84 | 10901 | 6 |
| hybrid | gateway | mixed-80-20 | publish_total | 128 | 60/60 | 0 | 0 | 8.63 | 3212.314 | 3219.971 | true | 18 | 137.54 | 13647 | 3 |
| hybrid | gateway | mixed-80-20 | read | 1 | 240/240 | 0 | 0 | 127.80 | 9.404 | 10.762 | true | 2 | 51.97 | 3643 | 3 |
| hybrid | gateway | mixed-80-20 | read | 32 | 240/240 | 0 | 0 | 650.88 | 52.443 | 61.873 | true | 6 | 85.84 | 10901 | 6 |
| hybrid | gateway | mixed-80-20 | read | 128 | 240/240 | 0 | 0 | 706.16 | 176.268 | 195.349 | true | 18 | 137.54 | 13647 | 3 |
| hybrid | gateway | mixed-90-10 | publish_total | 1 | 30/30 | 0 | 0 | 1.39 | 755.289 | 821.505 | true | 2 | 51.93 | 5906 | 3 |
| hybrid | gateway | mixed-90-10 | publish_total | 32 | 30/30 | 0 | 0 | 4.69 | 1148.126 | 1148.126 | true | 3 | 58.32 | 7383 | 3 |
| hybrid | gateway | mixed-90-10 | publish_total | 128 | 30/30 | 0 | 0 | 7.26 | 1928.023 | 1933.123 | true | 0 | 98.39 | 11424 | 3 |
| hybrid | gateway | mixed-90-10 | read | 1 | 270/270 | 0 | 0 | 125.61 | 9.897 | 11.059 | true | 2 | 51.93 | 5906 | 3 |
| hybrid | gateway | mixed-90-10 | read | 32 | 270/270 | 0 | 0 | 711.26 | 48.349 | 51.563 | true | 3 | 58.32 | 7383 | 3 |
| hybrid | gateway | mixed-90-10 | read | 128 | 270/270 | 0 | 0 | 670.63 | 216.553 | 223.962 | true | 0 | 98.39 | 11424 | 3 |
| hybrid | gateway | publish | publish_total | 1 | 30/30 | 0 | 0 | 1.41 | 753.530 | 911.580 | true | 1 | 19.66 | 2372 | 3 |
| hybrid | gateway | publish | publish_total | 32 | 128/128 | 0 | 0 | 10.01 | 3240.688 | 3256.625 | true | 17 | 143.87 | 16323 | 4 |
| hybrid | gateway | publish | publish_total | 128 | 0/512 | 512 | 0 | 0.00 | 0.000 | 0.000 | false | 0 | 126.01 | 11811 | 3 |
| hybrid | rpc | burst-publish | publish_total | 32 | 128/128 | 0 | 0 | 10.13 | 3310.403 | 3333.011 | true | 28 | 142.04 | 17527 | 7 |
| hybrid | rpc | burst-read | read | 128 | 15526/15526 | 0 | 0 | 1544.12 | 106.303 | 125.845 | true | 0 | 272.81 | 55203 | 5 |
| hybrid | rpc | deep-page | read | 1 | 427/427 | 0 | 0 | 142.19 | 8.616 | 9.800 | true | 0 | 41.34 | 5135 | 3 |
| hybrid | rpc | deep-page | read | 8 | 1984/1984 | 0 | 0 | 659.01 | 14.762 | 16.422 | true | 0 | 209.24 | 22639 | 4 |
| hybrid | rpc | deep-page | read | 32 | 2954/2954 | 0 | 0 | 976.26 | 41.969 | 49.040 | true | 0 | 263.45 | 33043 | 3 |
| hybrid | rpc | deep-page | read | 128 | 3604/3604 | 0 | 0 | 1173.70 | 136.211 | 151.540 | true | 0 | 290.93 | 39631 | 3 |
| hybrid | rpc | deep-page | read | 256 | 3647/3647 | 0 | 0 | 1150.92 | 276.740 | 305.585 | true | 0 | 282.75 | 35875 | 3 |
| hybrid | rpc | distributed-read | read | 1 | 476/476 | 0 | 0 | 158.60 | 7.390 | 8.230 | true | 0 | 41.71 | 5605 | 3 |
| hybrid | rpc | distributed-read | read | 8 | 2204/2204 | 0 | 0 | 732.63 | 13.471 | 14.677 | true | 0 | 193.54 | 23761 | 3 |
| hybrid | rpc | distributed-read | read | 32 | 3914/3914 | 0 | 0 | 1296.51 | 30.969 | 35.506 | true | 0 | 260.32 | 43382 | 3 |
| hybrid | rpc | distributed-read | read | 128 | 4865/4865 | 0 | 0 | 1592.80 | 101.032 | 111.292 | true | 0 | 273.16 | 53701 | 3 |
| hybrid | rpc | distributed-read | read | 256 | 4980/4980 | 0 | 0 | 1599.55 | 193.743 | 216.622 | true | 0 | 280.65 | 52178 | 3 |
| hybrid | rpc | hot-read | read | 1 | 481/481 | 0 | 0 | 160.14 | 7.390 | 8.070 | true | 0 | 40.43 | 5586 | 3 |
| hybrid | rpc | hot-read | read | 8 | 2334/2334 | 0 | 0 | 775.44 | 12.910 | 17.264 | true | 0 | 188.61 | 26418 | 3 |
| hybrid | rpc | hot-read | read | 32 | 3848/3848 | 0 | 0 | 1275.77 | 31.508 | 39.364 | true | 0 | 238.02 | 43032 | 3 |
| hybrid | rpc | hot-read | read | 128 | 4687/4687 | 0 | 0 | 1536.62 | 103.197 | 119.803 | true | 0 | 272.04 | 51057 | 3 |
| hybrid | rpc | hot-read | read | 256 | 4723/4723 | 0 | 0 | 1514.78 | 201.642 | 268.804 | true | 0 | 250.67 | 50173 | 3 |
| hybrid | rpc | mixed-80-20 | publish_total | 1 | 60/60 | 0 | 0 | 1.46 | 722.663 | 765.069 | true | 2 | 55.64 | 2730 | 3 |
| hybrid | rpc | mixed-80-20 | publish_total | 8 | 60/60 | 0 | 0 | 2.85 | 752.834 | 884.788 | true | 0 | 29.88 | 4736 | 3 |
| hybrid | rpc | mixed-80-20 | publish_total | 32 | 60/60 | 0 | 0 | 7.14 | 1246.946 | 1249.645 | true | 0 | 88.23 | 11633 | 3 |
| hybrid | rpc | mixed-80-20 | publish_total | 128 | 60/60 | 0 | 0 | 9.43 | 2839.019 | 2844.630 | true | 16 | 136.47 | 16018 | 5 |
| hybrid | rpc | mixed-80-20 | publish_total | 256 | 1/60 | 59 | 0 | 0.17 | 883.741 | 883.741 | false | 43 | 137.82 | 12617 | 3 |
| hybrid | rpc | mixed-80-20 | read | 1 | 240/240 | 0 | 0 | 149.67 | 8.014 | 8.566 | true | 2 | 55.64 | 2730 | 3 |
| hybrid | rpc | mixed-80-20 | read | 8 | 240/240 | 0 | 0 | 581.72 | 12.350 | 14.689 | true | 0 | 29.88 | 4736 | 3 |
| hybrid | rpc | mixed-80-20 | read | 32 | 240/240 | 0 | 0 | 850.17 | 49.707 | 51.267 | true | 0 | 88.23 | 11633 | 3 |
| hybrid | rpc | mixed-80-20 | read | 128 | 240/240 | 0 | 0 | 942.32 | 129.209 | 133.420 | true | 16 | 136.47 | 16018 | 5 |
| hybrid | rpc | mixed-80-20 | read | 256 | 240/240 | 0 | 0 | 989.05 | 208.419 | 214.934 | false | 43 | 137.82 | 12617 | 3 |
| hybrid | rpc | mixed-90-10 | publish_total | 1 | 30/30 | 0 | 0 | 1.46 | 727.028 | 756.386 | true | 0 | 55.25 | 4694 | 3 |
| hybrid | rpc | mixed-90-10 | publish_total | 8 | 30/30 | 0 | 0 | 1.45 | 840.583 | 878.433 | true | 1 | 19.84 | 2507 | 3 |
| hybrid | rpc | mixed-90-10 | publish_total | 32 | 30/30 | 0 | 0 | 5.04 | 980.779 | 980.779 | true | 0 | 89.19 | 8127 | 3 |
| hybrid | rpc | mixed-90-10 | publish_total | 128 | 30/30 | 0 | 0 | 8.40 | 1614.963 | 1616.592 | true | 0 | 116.99 | 13089 | 3 |
| hybrid | rpc | mixed-90-10 | publish_total | 256 | 30/30 | 0 | 0 | 8.94 | 2605.246 | 2607.372 | true | 0 | 125.62 | 13224 | 3 |
| hybrid | rpc | mixed-90-10 | read | 1 | 270/270 | 0 | 0 | 151.93 | 7.905 | 8.715 | true | 0 | 55.25 | 4694 | 3 |
| hybrid | rpc | mixed-90-10 | read | 8 | 270/270 | 0 | 0 | 683.28 | 12.417 | 14.871 | true | 1 | 19.84 | 2507 | 3 |
| hybrid | rpc | mixed-90-10 | read | 32 | 270/270 | 0 | 0 | 1019.95 | 40.888 | 43.530 | true | 0 | 89.19 | 8127 | 3 |
| hybrid | rpc | mixed-90-10 | read | 128 | 270/270 | 0 | 0 | 1083.29 | 133.619 | 138.995 | true | 0 | 116.99 | 13089 | 3 |
| hybrid | rpc | mixed-90-10 | read | 256 | 270/270 | 0 | 0 | 1190.18 | 205.489 | 214.862 | true | 0 | 125.62 | 13224 | 3 |
| hybrid | rpc | publish | publish_total | 1 | 30/30 | 0 | 0 | 1.47 | 710.227 | 736.669 | true | 2 | 20.34 | 2095 | 4 |
| hybrid | rpc | publish | publish_total | 8 | 32/32 | 0 | 0 | 8.14 | 1048.121 | 1060.732 | true | 7 | 88.97 | 11451 | 3 |
| hybrid | rpc | publish | publish_total | 32 | 128/128 | 0 | 0 | 10.09 | 3242.933 | 3247.612 | true | 0 | 147.16 | 16927 | 3 |
| hybrid | rpc | publish | publish_total | 36 | 144/144 | 0 | 0 | 9.51 | 3879.471 | 3888.936 | true | 31 | 142.28 | 17532 | 6 |
| hybrid | rpc | publish | publish_total | 40 | 160/160 | 0 | 0 | 9.47 | 4299.724 | 4318.382 | true | 35 | 137.61 | 18482 | 5 |
| hybrid | rpc | publish | publish_total | 44 | 34/176 | 142 | 0 | 8.51 | 3861.256 | 3868.764 | false | 0 | 134.15 | 12820 | 3 |
| hybrid | rpc | publish | publish_total | 48 | 32/192 | 160 | 0 | 1.80 | 4992.226 | 4993.002 | false | 42 | 133.94 | 27484 | 4 |
| hybrid | rpc | publish | publish_total | 64 | 2/256 | 254 | 0 | 0.72 | 767.230 | 767.230 | false | 0 | 115.98 | 9695 | 3 |
| hybrid | rpc | publish | publish_total | 96 | 6/384 | 378 | 0 | 5.68 | 1051.131 | 1051.131 | false | 0 | 1.29 | 4927 | 3 |
| hybrid | rpc | publish | publish_total | 128 | 1/512 | 511 | 0 | 0.36 | 721.608 | 721.608 | false | 0 | 115.77 | 7988 | 3 |
| hybrid | rpc | publish | publish_total | 256 | 1/1024 | 1023 | 0 | 1.10 | 904.588 | 904.588 | false | 0 | 0.41 | 1558 | 3 |
| hybrid | rpc | steady-read | read | 128 | 47777/47777 | 0 | 0 | 1590.01 | 103.534 | 115.576 | true | 0 | 274.96 | 57294 | 5 |
| pull | gateway | distributed-read | read | 1 | 388/388 | 0 | 0 | 129.31 | 9.202 | 10.611 | true | 0 | 64.87 | 3979 | 3 |
| pull | gateway | distributed-read | read | 32 | 1351/1351 | 0 | 0 | 444.36 | 94.320 | 101.973 | true | 0 | 212.87 | 13490 | 3 |
| pull | gateway | distributed-read | read | 128 | 1372/1372 | 0 | 0 | 431.69 | 367.743 | 399.614 | true | 0 | 209.06 | 13095 | 4 |
| pull | gateway | hot-read | read | 1 | 404/404 | 0 | 0 | 134.61 | 8.717 | 9.507 | true | 0 | 66.55 | 4049 | 3 |
| pull | gateway | hot-read | read | 32 | 1342/1342 | 0 | 0 | 440.58 | 94.896 | 105.336 | true | 0 | 210.21 | 13340 | 3 |
| pull | gateway | hot-read | read | 128 | 1452/1452 | 0 | 0 | 459.64 | 348.206 | 401.617 | true | 0 | 218.90 | 13720 | 4 |
| pull | gateway | mixed-80-20 | publish_total | 1 | 60/60 | 0 | 0 | 1.47 | 731.552 | 878.530 | true | 0 | 72.39 | 3103 | 3 |
| pull | gateway | mixed-80-20 | publish_total | 32 | 60/60 | 0 | 0 | 7.07 | 1561.544 | 1569.504 | true | 0 | 137.99 | 9746 | 3 |
| pull | gateway | mixed-80-20 | publish_total | 128 | 60/60 | 0 | 0 | 8.94 | 3147.683 | 3158.224 | true | 0 | 136.98 | 12806 | 5 |
| pull | gateway | mixed-80-20 | read | 1 | 240/240 | 0 | 0 | 114.85 | 10.588 | 11.648 | true | 0 | 72.39 | 3103 | 3 |
| pull | gateway | mixed-80-20 | read | 32 | 240/240 | 0 | 0 | 339.15 | 92.499 | 99.667 | true | 0 | 137.99 | 9746 | 3 |
| pull | gateway | mixed-80-20 | read | 128 | 240/240 | 0 | 0 | 332.60 | 364.520 | 381.596 | true | 0 | 136.98 | 12806 | 5 |
| pull | gateway | mixed-90-10 | publish_total | 1 | 30/30 | 0 | 0 | 1.44 | 865.465 | 932.359 | true | 0 | 72.28 | 4843 | 4 |
| pull | gateway | mixed-90-10 | publish_total | 32 | 30/30 | 0 | 0 | 4.65 | 1460.029 | 1460.546 | true | 0 | 144.48 | 6500 | 3 |
| pull | gateway | mixed-90-10 | publish_total | 128 | 30/30 | 0 | 0 | 7.40 | 1980.571 | 1980.571 | true | 0 | 118.32 | 12313 | 3 |
| pull | gateway | mixed-90-10 | read | 1 | 270/270 | 0 | 0 | 120.38 | 10.155 | 11.708 | true | 0 | 72.28 | 4843 | 4 |
| pull | gateway | mixed-90-10 | read | 32 | 270/270 | 0 | 0 | 343.73 | 98.224 | 99.990 | true | 0 | 144.48 | 6500 | 3 |
| pull | gateway | mixed-90-10 | read | 128 | 270/270 | 0 | 0 | 361.68 | 426.195 | 439.816 | true | 0 | 118.32 | 12313 | 3 |
| pull | gateway | publish | publish_total | 1 | 30/30 | 0 | 0 | 1.51 | 676.615 | 695.874 | true | 0 | 18.96 | 1923 | 3 |
| pull | gateway | publish | publish_total | 32 | 128/128 | 0 | 0 | 10.69 | 3119.204 | 3129.095 | true | 0 | 143.98 | 13345 | 3 |
| pull | gateway | publish | publish_total | 128 | 3/512 | 509 | 0 | 0.64 | 2459.919 | 2459.919 | false | 0 | 112.11 | 10918 | 3 |
| pull | rpc | burst-publish | publish_total | 32 | 127/128 | 1 | 0 | 10.54 | 3081.573 | 3095.037 | false | 0 | 145.83 | 13155 | 3 |
| pull | rpc | burst-read | read | 128 | 6291/6291 | 0 | 0 | 620.17 | 265.135 | 353.563 | true | 0 | 290.82 | 20331 | 3 |
| pull | rpc | deep-page | read | 1 | 447/447 | 0 | 0 | 148.93 | 7.880 | 9.015 | true | 0 | 67.86 | 4464 | 3 |
| pull | rpc | deep-page | read | 8 | 1434/1434 | 0 | 0 | 476.09 | 20.689 | 22.803 | true | 0 | 230.21 | 13911 | 3 |
| pull | rpc | deep-page | read | 32 | 1669/1669 | 0 | 0 | 549.98 | 75.564 | 82.649 | true | 0 | 242.22 | 14689 | 3 |
| pull | rpc | deep-page | read | 128 | 1740/1740 | 0 | 0 | 549.66 | 297.496 | 318.212 | true | 0 | 272.98 | 15707 | 3 |
| pull | rpc | deep-page | read | 256 | 1850/1850 | 0 | 0 | 572.70 | 560.720 | 598.776 | true | 0 | 288.97 | 16331 | 3 |
| pull | rpc | distributed-read | read | 1 | 483/483 | 0 | 0 | 160.78 | 7.427 | 8.361 | true | 0 | 65.90 | 4797 | 3 |
| pull | rpc | distributed-read | read | 8 | 1566/1566 | 0 | 0 | 520.26 | 18.494 | 19.623 | true | 0 | 238.90 | 15274 | 3 |
| pull | rpc | distributed-read | read | 32 | 1815/1815 | 0 | 0 | 599.06 | 67.970 | 74.027 | true | 0 | 273.58 | 17175 | 3 |
| pull | rpc | distributed-read | read | 128 | 1824/1824 | 0 | 0 | 585.63 | 243.794 | 256.473 | true | 0 | 265.08 | 16545 | 3 |
| pull | rpc | distributed-read | read | 256 | 2022/2022 | 0 | 0 | 621.91 | 471.097 | 513.875 | true | 0 | 276.17 | 17516 | 3 |
| pull | rpc | hot-read | read | 1 | 491/491 | 0 | 0 | 163.47 | 7.252 | 8.010 | true | 0 | 67.98 | 4992 | 3 |
| pull | rpc | hot-read | read | 8 | 1551/1551 | 0 | 0 | 515.34 | 18.473 | 21.043 | true | 0 | 225.93 | 14958 | 5 |
| pull | rpc | hot-read | read | 32 | 1769/1769 | 0 | 0 | 581.34 | 70.328 | 78.447 | true | 0 | 264.64 | 16662 | 3 |
| pull | rpc | hot-read | read | 128 | 1862/1862 | 0 | 0 | 598.19 | 247.701 | 275.320 | true | 0 | 268.68 | 17437 | 3 |
| pull | rpc | hot-read | read | 256 | 1951/1951 | 0 | 0 | 603.45 | 512.097 | 553.677 | true | 0 | 283.07 | 17582 | 3 |
| pull | rpc | mixed-80-20 | publish_total | 1 | 60/60 | 0 | 0 | 1.51 | 688.840 | 800.794 | true | 0 | 79.66 | 2407 | 3 |
| pull | rpc | mixed-80-20 | publish_total | 8 | 60/60 | 0 | 0 | 2.90 | 830.385 | 1066.660 | true | 0 | 54.79 | 3586 | 3 |
| pull | rpc | mixed-80-20 | publish_total | 32 | 60/60 | 0 | 0 | 7.34 | 1410.408 | 1422.177 | true | 0 | 88.72 | 9908 | 3 |
| pull | rpc | mixed-80-20 | publish_total | 128 | 60/60 | 0 | 0 | 9.49 | 2879.121 | 2883.635 | true | 0 | 138.05 | 12505 | 3 |
| pull | rpc | mixed-80-20 | publish_total | 256 | 1/60 | 59 | 0 | 0.17 | 837.847 | 837.847 | false | 0 | 133.90 | 13867 | 6 |
| pull | rpc | mixed-80-20 | read | 1 | 240/240 | 0 | 0 | 147.30 | 8.033 | 9.088 | true | 0 | 79.66 | 2407 | 3 |
| pull | rpc | mixed-80-20 | read | 8 | 240/240 | 0 | 0 | 396.00 | 18.468 | 24.102 | true | 0 | 54.79 | 3586 | 3 |
| pull | rpc | mixed-80-20 | read | 32 | 240/240 | 0 | 0 | 420.40 | 75.017 | 82.060 | true | 0 | 88.72 | 9908 | 3 |
| pull | rpc | mixed-80-20 | read | 128 | 240/240 | 0 | 0 | 426.61 | 300.058 | 308.674 | true | 0 | 138.05 | 12505 | 3 |
| pull | rpc | mixed-80-20 | read | 256 | 240/240 | 0 | 0 | 445.39 | 457.638 | 478.685 | false | 0 | 133.90 | 13867 | 6 |
| pull | rpc | mixed-90-10 | publish_total | 1 | 30/30 | 0 | 0 | 1.48 | 800.253 | 807.883 | true | 0 | 78.44 | 4101 | 3 |
| pull | rpc | mixed-90-10 | publish_total | 8 | 30/30 | 0 | 0 | 1.49 | 703.889 | 1056.186 | true | 0 | 19.68 | 1928 | 3 |
| pull | rpc | mixed-90-10 | publish_total | 32 | 30/30 | 0 | 0 | 4.97 | 1165.956 | 1167.042 | true | 0 | 137.94 | 6760 | 3 |
| pull | rpc | mixed-90-10 | publish_total | 128 | 30/30 | 0 | 0 | 7.92 | 1907.713 | 1907.713 | true | 0 | 116.82 | 12443 | 3 |
| pull | rpc | mixed-90-10 | publish_total | 256 | 30/30 | 0 | 0 | 8.07 | 3008.193 | 3008.193 | true | 0 | 111.66 | 12111 | 3 |
| pull | rpc | mixed-90-10 | read | 1 | 270/270 | 0 | 0 | 146.99 | 8.410 | 9.783 | true | 0 | 78.44 | 4101 | 3 |
| pull | rpc | mixed-90-10 | read | 8 | 270/270 | 0 | 0 | 449.35 | 18.876 | 23.909 | true | 0 | 19.68 | 1928 | 3 |
| pull | rpc | mixed-90-10 | read | 32 | 270/270 | 0 | 0 | 555.39 | 60.987 | 66.775 | true | 0 | 137.94 | 6760 | 3 |
| pull | rpc | mixed-90-10 | read | 128 | 270/270 | 0 | 0 | 477.20 | 270.700 | 282.411 | true | 0 | 116.82 | 12443 | 3 |
| pull | rpc | mixed-90-10 | read | 256 | 270/270 | 0 | 0 | 510.79 | 478.084 | 487.729 | true | 0 | 111.66 | 12111 | 3 |
| pull | rpc | publish | publish_total | 1 | 30/30 | 0 | 0 | 1.52 | 692.240 | 702.435 | true | 0 | 19.66 | 1935 | 3 |
| pull | rpc | publish | publish_total | 8 | 32/32 | 0 | 0 | 8.75 | 978.816 | 979.390 | true | 0 | 88.23 | 10541 | 3 |
| pull | rpc | publish | publish_total | 32 | 128/128 | 0 | 0 | 10.76 | 3029.476 | 3041.171 | true | 0 | 145.52 | 13214 | 4 |
| pull | rpc | publish | publish_total | 128 | 2/512 | 510 | 0 | 0.72 | 745.965 | 745.965 | false | 0 | 109.32 | 8878 | 3 |
| pull | rpc | publish | publish_total | 256 | 4/1024 | 1020 | 0 | 4.24 | 915.893 | 915.893 | false | 0 | 1.66 | 3367 | 9 |
| pull | rpc | steady-read | read | 128 | 20193/20193 | 0 | 0 | 670.94 | 232.156 | 249.033 | true | 0 | 318.15 | 20397 | 4 |
| push | gateway | distributed-read | read | 1 | 549/549 | 0 | 0 | 182.68 | 6.880 | 7.514 | true | 0 | 27.85 | 801 | 3 |
| push | gateway | distributed-read | read | 32 | 4372/4372 | 0 | 0 | 1447.86 | 28.922 | 33.100 | true | 0 | 164.84 | 4541 | 3 |
| push | gateway | distributed-read | read | 128 | 4601/4601 | 0 | 0 | 1503.17 | 118.268 | 133.783 | true | 0 | 165.88 | 4330 | 3 |
| push | gateway | hot-read | read | 1 | 553/553 | 0 | 0 | 184.19 | 6.418 | 7.011 | true | 0 | 27.73 | 791 | 3 |
| push | gateway | hot-read | read | 32 | 4873/4873 | 0 | 0 | 1616.45 | 26.135 | 30.763 | true | 0 | 175.13 | 5044 | 3 |
| push | gateway | hot-read | read | 128 | 5195/5195 | 0 | 0 | 1702.76 | 102.446 | 119.143 | true | 0 | 157.09 | 5408 | 3 |
| push | gateway | mixed-80-20 | publish_total | 1 | 60/60 | 0 | 0 | 1.42 | 780.361 | 879.609 | true | 2 | 31.04 | 6071 | 4 |
| push | gateway | mixed-80-20 | publish_total | 32 | 60/60 | 0 | 0 | 6.89 | 1210.701 | 1213.577 | true | 7 | 84.45 | 18037 | 4 |
| push | gateway | mixed-80-20 | publish_total | 128 | 60/60 | 0 | 0 | 8.79 | 2943.508 | 2946.813 | true | 0 | 119.88 | 15870 | 3 |
| push | gateway | mixed-80-20 | read | 1 | 240/240 | 0 | 0 | 151.58 | 7.953 | 9.025 | true | 2 | 31.04 | 6071 | 4 |
| push | gateway | mixed-80-20 | read | 32 | 240/240 | 0 | 0 | 927.34 | 57.530 | 63.353 | true | 7 | 84.45 | 18037 | 4 |
| push | gateway | mixed-80-20 | read | 128 | 240/240 | 0 | 0 | 1066.94 | 135.654 | 148.188 | true | 0 | 119.88 | 15870 | 3 |
| push | gateway | mixed-90-10 | publish_total | 1 | 30/30 | 0 | 0 | 1.42 | 787.377 | 873.434 | true | 2 | 42.87 | 6201 | 3 |
| push | gateway | mixed-90-10 | publish_total | 32 | 30/30 | 0 | 0 | 4.82 | 1063.727 | 1065.630 | true | 4 | 56.79 | 11182 | 4 |
| push | gateway | mixed-90-10 | publish_total | 128 | 30/30 | 0 | 0 | 7.94 | 1653.588 | 1654.107 | true | 0 | 100.90 | 14743 | 3 |
| push | gateway | mixed-90-10 | read | 1 | 270/270 | 0 | 0 | 171.55 | 7.353 | 8.112 | true | 2 | 42.87 | 6201 | 3 |
| push | gateway | mixed-90-10 | read | 32 | 270/270 | 0 | 0 | 914.25 | 53.178 | 57.883 | true | 4 | 56.79 | 11182 | 4 |
| push | gateway | mixed-90-10 | read | 128 | 270/270 | 0 | 0 | 898.47 | 165.092 | 182.971 | true | 0 | 100.90 | 14743 | 3 |
| push | gateway | publish | publish_total | 1 | 30/30 | 0 | 0 | 1.51 | 703.468 | 706.456 | true | 2 | 19.99 | 6114 | 4 |
| push | gateway | publish | publish_total | 32 | 128/128 | 0 | 0 | 10.02 | 3439.447 | 3448.266 | true | 32 | 141.40 | 27381 | 3 |
| push | gateway | publish | publish_total | 128 | 2/512 | 510 | 0 | 0.43 | 2458.191 | 2458.191 | false | 0 | 104.98 | 10874 | 3 |
| push | rpc | burst-publish | publish_total | 32 | 22/128 | 106 | 0 | 8.95 | 2336.120 | 2336.120 | false | 22 | 126.88 | 12550 | 3 |
| push | rpc | burst-read | read | 128 | 27020/27020 | 0 | 0 | 2695.27 | 62.438 | 73.917 | true | 0 | 278.76 | 8551 | 6 |
| push | rpc | deep-page | read | 1 | 569/569 | 0 | 0 | 189.28 | 6.466 | 7.422 | true | 0 | 33.95 | 800 | 3 |
| push | rpc | deep-page | read | 8 | 2678/2678 | 0 | 0 | 890.28 | 11.741 | 13.809 | true | 0 | 167.15 | 3046 | 3 |
| push | rpc | deep-page | read | 32 | 4085/4085 | 0 | 0 | 1355.38 | 31.813 | 37.983 | true | 0 | 246.73 | 4330 | 3 |
| push | rpc | deep-page | read | 128 | 5204/5204 | 0 | 0 | 1704.57 | 108.155 | 130.474 | true | 0 | 296.48 | 5483 | 3 |
| push | rpc | deep-page | read | 256 | 5528/5528 | 0 | 0 | 1782.58 | 224.324 | 296.030 | true | 0 | 315.38 | 5496 | 3 |
| push | rpc | distributed-read | read | 1 | 685/685 | 0 | 0 | 228.29 | 5.299 | 6.017 | true | 0 | 31.51 | 920 | 3 |
| push | rpc | distributed-read | read | 8 | 3459/3459 | 0 | 0 | 1150.88 | 9.000 | 10.910 | true | 0 | 156.80 | 3685 | 3 |
| push | rpc | distributed-read | read | 32 | 6280/6280 | 0 | 0 | 2085.53 | 20.361 | 23.655 | true | 0 | 211.51 | 6694 | 3 |
| push | rpc | distributed-read | read | 128 | 8379/8379 | 0 | 0 | 2759.77 | 60.037 | 66.048 | true | 0 | 276.49 | 8526 | 3 |
| push | rpc | distributed-read | read | 256 | 8918/8918 | 0 | 0 | 2915.53 | 131.227 | 167.052 | true | 0 | 269.59 | 8808 | 3 |
| push | rpc | hot-read | read | 1 | 693/693 | 0 | 0 | 230.97 | 5.174 | 5.833 | true | 0 | 32.75 | 931 | 3 |
| push | rpc | hot-read | read | 8 | 3599/3599 | 0 | 0 | 1197.64 | 8.125 | 9.583 | true | 0 | 157.53 | 3882 | 3 |
| push | rpc | hot-read | read | 32 | 6368/6368 | 0 | 0 | 2114.87 | 19.653 | 23.080 | true | 0 | 231.04 | 6656 | 3 |
| push | rpc | hot-read | read | 128 | 8245/8245 | 0 | 0 | 2720.97 | 62.708 | 91.483 | true | 0 | 246.26 | 8441 | 3 |
| push | rpc | hot-read | read | 256 | 9153/9153 | 0 | 0 | 2989.43 | 127.779 | 157.689 | true | 0 | 249.70 | 9373 | 3 |
| push | rpc | mixed-80-20 | publish_total | 1 | 60/60 | 0 | 0 | 1.54 | 673.586 | 733.975 | true | 2 | 38.04 | 6163 | 3 |
| push | rpc | mixed-80-20 | publish_total | 8 | 60/60 | 0 | 0 | 3.02 | 717.845 | 776.175 | true | 4 | 32.11 | 10445 | 3 |
| push | rpc | mixed-80-20 | publish_total | 32 | 60/60 | 0 | 0 | 7.66 | 1012.335 | 1017.151 | true | 7 | 92.59 | 16680 | 3 |
| push | rpc | mixed-80-20 | publish_total | 128 | 60/60 | 0 | 0 | 10.11 | 2471.319 | 2474.483 | true | 0 | 132.70 | 23667 | 5 |
| push | rpc | mixed-80-20 | publish_total | 256 | 60/60 | 0 | 0 | 10.39 | 4883.726 | 4893.036 | true | 52 | 117.47 | 13114 | 3 |
| push | rpc | mixed-80-20 | read | 1 | 240/240 | 0 | 0 | 208.29 | 5.471 | 6.399 | true | 2 | 38.04 | 6163 | 3 |
| push | rpc | mixed-80-20 | read | 8 | 240/240 | 0 | 0 | 899.92 | 7.948 | 12.434 | true | 4 | 32.11 | 10445 | 3 |
| push | rpc | mixed-80-20 | read | 32 | 240/240 | 0 | 0 | 1469.08 | 36.337 | 39.577 | true | 7 | 92.59 | 16680 | 3 |
| push | rpc | mixed-80-20 | read | 128 | 240/240 | 0 | 0 | 1552.23 | 97.922 | 105.400 | true | 0 | 132.70 | 23667 | 5 |
| push | rpc | mixed-80-20 | read | 256 | 240/240 | 0 | 0 | 1563.70 | 135.326 | 144.686 | true | 52 | 117.47 | 13114 | 3 |
| push | rpc | mixed-90-10 | publish_total | 1 | 30/30 | 0 | 0 | 1.56 | 670.381 | 710.588 | true | 2 | 38.28 | 6245 | 3 |
| push | rpc | mixed-90-10 | publish_total | 8 | 30/30 | 0 | 0 | 1.56 | 656.396 | 741.104 | true | 0 | 20.08 | 4180 | 4 |
| push | rpc | mixed-90-10 | publish_total | 32 | 30/30 | 0 | 0 | 5.41 | 842.370 | 842.922 | true | 0 | 57.62 | 11284 | 3 |
| push | rpc | mixed-90-10 | publish_total | 128 | 30/30 | 0 | 0 | 9.05 | 1487.335 | 1487.860 | true | 13 | 112.68 | 16970 | 3 |
| push | rpc | mixed-90-10 | publish_total | 256 | 30/30 | 0 | 0 | 9.73 | 2442.066 | 2443.650 | true | 0 | 143.06 | 11046 | 3 |
| push | rpc | mixed-90-10 | read | 1 | 270/270 | 0 | 0 | 216.95 | 5.578 | 6.462 | true | 2 | 38.28 | 6245 | 3 |
| push | rpc | mixed-90-10 | read | 8 | 270/270 | 0 | 0 | 1019.55 | 9.234 | 11.186 | true | 0 | 20.08 | 4180 | 4 |
| push | rpc | mixed-90-10 | read | 32 | 270/270 | 0 | 0 | 1562.32 | 30.986 | 35.172 | true | 0 | 57.62 | 11284 | 3 |
| push | rpc | mixed-90-10 | read | 128 | 270/270 | 0 | 0 | 1617.20 | 112.756 | 122.296 | true | 13 | 112.68 | 16970 | 3 |
| push | rpc | mixed-90-10 | read | 256 | 270/270 | 0 | 0 | 1722.27 | 143.922 | 148.836 | true | 0 | 143.06 | 11046 | 3 |
| push | rpc | publish | publish_total | 1 | 30/30 | 0 | 0 | 1.59 | 664.068 | 669.526 | true | 2 | 20.61 | 8301 | 3 |
| push | rpc | publish | publish_total | 8 | 32/32 | 0 | 0 | 8.97 | 1035.575 | 1042.546 | true | 8 | 94.31 | 13889 | 4 |
| push | rpc | publish | publish_total | 32 | 128/128 | 0 | 0 | 11.38 | 2985.753 | 2990.257 | true | 32 | 150.63 | 28000 | 6 |
| push | rpc | publish | publish_total | 128 | 3/512 | 509 | 0 | 1.09 | 705.127 | 705.127 | false | 0 | 103.15 | 12983 | 3 |
| push | rpc | publish | publish_total | 256 | 3/1024 | 1021 | 0 | 3.40 | 880.467 | 880.467 | false | 0 | 0.48 | 2834 | 3 |
| push | rpc | steady-read | read | 128 | 82405/82405 | 0 | 0 | 2744.23 | 61.066 | 68.412 | true | 0 | 289.79 | 8835 | 5 |
