# Feed 三策略综合对比：feed-cursor-deep-v1-gateway-formal-20260817a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 1879.23 | 13.626 | 16 | 32 | 32 | - |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 1259.23 | 39.120 | 16 | 32 | 32 | - |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 1972.86 | 12.624 | 16 | 32 | 32 | - |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 1896.99 | 14.052 | 16 | 32 | 32 | - |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 1037.57 | 25.289 | 16 | 32 | 32 | - |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 1568.51 | 17.294 | 16 | 32 | 32 | - |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 497.56 | 44.861 | 16 | 32 | 32 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1879.23 | 1490.25 | 1891.04 | 21.33% | 13.626 | 13.543 | 17.531 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 1676.71 | 1617.52 | 1713.84 | 5.74% | 29.534 | 28.330 | 30.225 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1232.79 | 1138.50 | 1308.25 | 13.77% | 21.483 | 20.254 | 23.305 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 1259.23 | 1239.14 | 1407.96 | 13.41% | 39.120 | 33.822 | 39.575 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1972.86 | 1825.68 | 2016.76 | 9.69% | 12.624 | 12.620 | 13.724 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 1893.96 | 1536.70 | 1992.18 | 24.05% | 27.050 | 24.311 | 31.085 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1896.99 | 1420.98 | 2143.32 | 38.08% | 14.052 | 12.470 | 19.432 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 1766.77 | 1538.93 | 1967.44 | 24.25% | 26.413 | 24.136 | 30.469 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1037.57 | 921.98 | 1048.94 | 12.24% | 25.289 | 25.100 | 29.970 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 908.44 | 888.66 | 965.87 | 8.50% | 58.437 | 54.638 | 62.178 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1568.51 | 1436.85 | 1589.95 | 9.76% | 17.294 | 17.146 | 19.928 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 1373.83 | 1325.45 | 1374.89 | 3.60% | 35.384 | 35.264 | 37.013 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 497.56 | 494.73 | 501.91 | 1.44% | 44.861 | 44.319 | 45.645 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 335.76 | 319.80 | 379.02 | 17.64% | 130.163 | 110.782 | 136.779 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 113472/113472 | 0 | 0 | 1891.04 | 11.784 | 13.626 | 17.300 | 32.714 | true | 0 | 176.72 | 3.80 | 33358 | 4 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 89425/89425 | 0 | 0 | 1490.25 | 15.036 | 17.531 | 22.409 | 43.486 | true | 0 | 178.35 | 3.16 | 26187 | 3 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 112764/112764 | 0 | 0 | 1879.23 | 11.558 | 13.543 | 16.990 | 27.302 | true | 0 | 170.55 | 3.71 | 32059 | 3 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 97072/97072 | 0 | 0 | 1617.52 | 26.935 | 30.225 | 39.664 | 84.654 | true | 0 | 166.41 | 3.39 | 28117 | 4 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 100623/100623 | 0 | 0 | 1676.71 | 26.058 | 29.534 | 37.892 | 63.454 | true | 0 | 176.55 | 3.58 | 29394 | 3 |
| hybrid | gateway | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 102856/102856 | 0 | 0 | 1713.84 | 25.087 | 28.330 | 37.065 | 70.457 | true | 0 | 166.11 | 3.80 | 28793 | 3 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 73975/73975 | 0 | 0 | 1232.79 | 18.218 | 21.483 | 27.300 | 49.639 | true | 0 | 140.76 | 2.48 | 25333 | 4 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 68316/68316 | 0 | 0 | 1138.50 | 19.736 | 23.305 | 29.816 | 61.449 | true | 0 | 129.19 | 2.30 | 23550 | 3 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 78500/78500 | 0 | 0 | 1308.25 | 17.270 | 20.254 | 25.700 | 51.968 | true | 0 | 138.05 | 2.38 | 27858 | 4 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 74369/74369 | 0 | 0 | 1239.14 | 35.046 | 39.575 | 52.900 | 84.170 | true | 0 | 133.93 | 2.46 | 24342 | 3 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 75569/75569 | 0 | 0 | 1259.23 | 34.556 | 39.120 | 52.116 | 88.725 | true | 0 | 159.69 | 2.51 | 26308 | 4 |
| hybrid | gateway | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 84494/84494 | 0 | 0 | 1407.96 | 30.062 | 33.822 | 45.816 | 78.187 | true | 0 | 156.24 | 2.74 | 28345 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 118381/118381 | 0 | 0 | 1972.86 | 11.031 | 12.624 | 15.729 | 24.667 | true | 0 | 214.24 | 4.62 | 31869 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 121019/121019 | 0 | 0 | 2016.76 | 11.121 | 12.620 | 16.009 | 29.058 | true | 0 | 211.60 | 4.83 | 33670 | 4 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 109551/109551 | 0 | 0 | 1825.68 | 12.040 | 13.724 | 17.387 | 31.654 | true | 0 | 205.06 | 4.32 | 29507 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 113651/113651 | 0 | 0 | 1893.96 | 23.985 | 27.050 | 34.059 | 69.084 | true | 0 | 218.17 | 4.74 | 33943 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 92224/92224 | 0 | 0 | 1536.70 | 28.085 | 31.085 | 41.119 | 75.763 | true | 0 | 202.14 | 3.99 | 24376 | 3 |
| hybrid | gateway | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 119554/119554 | 0 | 0 | 1992.18 | 21.687 | 24.311 | 31.492 | 62.876 | true | 0 | 230.74 | 4.90 | 33300 | 3 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 128608/128608 | 0 | 0 | 2143.32 | 10.419 | 12.470 | 16.067 | 31.598 | true | 0 | 169.89 | 3.89 | 19658 | 3 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 85267/85267 | 0 | 0 | 1420.98 | 15.293 | 19.432 | 24.354 | 47.808 | true | 0 | 154.93 | 3.29 | 11928 | 4 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 113824/113824 | 0 | 0 | 1896.99 | 11.685 | 14.052 | 17.777 | 34.547 | true | 0 | 181.96 | 3.82 | 16750 | 3 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 118068/118068 | 0 | 0 | 1967.44 | 21.265 | 24.136 | 34.673 | 70.748 | true | 0 | 201.30 | 3.61 | 17483 | 3 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 92361/92361 | 0 | 0 | 1538.93 | 26.645 | 30.469 | 44.109 | 78.518 | true | 0 | 181.95 | 3.19 | 12250 | 4 |
| hybrid | gateway | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 106016/106016 | 0 | 0 | 1766.77 | 22.525 | 26.413 | 37.505 | 74.063 | true | 0 | 181.05 | 3.39 | 14289 | 3 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 55328/55328 | 0 | 0 | 921.98 | 25.304 | 29.970 | 38.443 | 69.022 | true | 0 | 220.52 | 1.41 | 8227 | 3 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 62266/62266 | 0 | 0 | 1037.57 | 22.941 | 25.289 | 30.828 | 63.658 | true | 0 | 238.60 | 1.36 | 8163 | 3 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 62945/62945 | 0 | 0 | 1048.94 | 22.647 | 25.100 | 30.280 | 48.871 | true | 0 | 241.43 | 1.34 | 8433 | 5 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 57969/57969 | 0 | 0 | 965.87 | 46.267 | 54.638 | 73.449 | 129.102 | true | 0 | 271.66 | 1.45 | 8038 | 3 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 54524/54524 | 0 | 0 | 908.44 | 50.147 | 58.437 | 83.734 | 170.163 | true | 0 | 232.13 | 1.27 | 8199 | 5 |
| hybrid | gateway | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 53334/53334 | 0 | 0 | 888.66 | 51.873 | 62.178 | 85.236 | 181.684 | true | 0 | 216.07 | 1.37 | 7988 | 3 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 86219/86219 | 0 | 0 | 1436.85 | 15.400 | 19.928 | 25.897 | 42.444 | true | 0 | 184.31 | 2.21 | 14386 | 3 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 95405/95405 | 0 | 0 | 1589.95 | 13.722 | 17.294 | 22.442 | 42.449 | true | 0 | 189.76 | 2.19 | 13883 | 4 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 94122/94122 | 0 | 0 | 1568.51 | 14.183 | 17.146 | 24.079 | 46.040 | true | 0 | 224.63 | 2.19 | 15296 | 3 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 82459/82459 | 0 | 0 | 1373.83 | 31.312 | 35.264 | 52.959 | 106.278 | true | 0 | 174.11 | 2.08 | 11927 | 4 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 82518/82518 | 0 | 0 | 1374.89 | 31.578 | 35.384 | 53.425 | 97.679 | true | 0 | 202.01 | 2.15 | 11844 | 4 |
| hybrid | gateway | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 79549/79549 | 0 | 0 | 1325.45 | 32.634 | 37.013 | 55.262 | 96.661 | true | 0 | 180.37 | 2.12 | 12624 | 3 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 29695/29695 | 0 | 0 | 494.73 | 41.166 | 44.319 | 51.409 | 77.326 | true | 0 | 218.97 | 0.79 | 4437 | 4 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 30121/30121 | 0 | 0 | 501.91 | 41.351 | 44.861 | 52.091 | 75.279 | true | 0 | 249.54 | 0.88 | 4582 | 4 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 29859/29859 | 0 | 0 | 497.56 | 42.173 | 45.645 | 52.613 | 73.887 | true | 0 | 266.25 | 0.82 | 4486 | 3 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 20153/20153 | 0 | 0 | 335.76 | 120.304 | 130.163 | 147.422 | 205.751 | true | 0 | 188.21 | 0.92 | 3150 | 4 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 19202/19202 | 0 | 0 | 319.80 | 128.514 | 136.779 | 153.040 | 233.239 | true | 0 | 173.35 | 1.01 | 3071 | 3 |
| hybrid | gateway | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 22757/22757 | 0 | 0 | 379.02 | 103.925 | 110.782 | 125.337 | 169.100 | true | 0 | 195.90 | 1.22 | 3452 | 3 |
