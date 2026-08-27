# Feed 三策略综合对比：feed-cursor-deep-v1-rpc-formal-20260817a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 4793.35 | 10.495 | 16 | 32 | 32 | - |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 2763.27 | 17.843 | 16 | 32 | 32 | - |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 7640.14 | 6.598 | 16 | 32 | 32 | - |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 5268.35 | 10.172 | 16 | 32 | 32 | - |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 1175.12 | 46.985 | 16 | 32 | 32 | - |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 2379.31 | 22.602 | 16 | 32 | 32 | - |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 586.55 | 71.412 | 16 | 32 | 32 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 4760.49 | 4744.64 | 4815.56 | 1.49% | 5.467 | 5.406 | 5.468 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 4793.35 | 4789.16 | 4811.06 | 0.46% | 10.495 | 10.445 | 10.537 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 2677.35 | 2624.75 | 2721.27 | 3.61% | 9.893 | 9.731 | 10.056 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 2763.27 | 2742.44 | 2771.99 | 1.07% | 17.843 | 17.803 | 17.925 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 7588.64 | 7569.25 | 7687.99 | 1.56% | 3.360 | 3.297 | 3.391 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 7640.14 | 7531.44 | 7730.01 | 2.60% | 6.598 | 6.509 | 6.678 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 5073.70 | 5054.89 | 5102.96 | 0.95% | 5.261 | 5.211 | 5.290 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 5268.35 | 5233.82 | 5334.89 | 1.92% | 10.172 | 9.995 | 10.203 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 1172.03 | 1160.03 | 1174.36 | 1.22% | 22.820 | 22.750 | 23.233 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 1175.12 | 1169.67 | 1190.83 | 1.80% | 46.985 | 46.148 | 47.090 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 2338.11 | 2324.45 | 2339.01 | 0.62% | 11.273 | 11.252 | 11.336 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 2379.31 | 2362.17 | 2382.44 | 0.85% | 22.602 | 22.454 | 22.744 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 581.12 | 579.73 | 588.31 | 1.48% | 35.945 | 35.472 | 36.742 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 586.55 | 583.54 | 590.99 | 1.27% | 71.412 | 70.333 | 72.171 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 285638/285638 | 0 | 0 | 4760.49 | 4.797 | 5.467 | 6.935 | 11.952 | true | 0 | 232.84 | 4.77 | 77670 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 284687/284687 | 0 | 0 | 4744.64 | 4.805 | 5.468 | 6.896 | 13.720 | true | 0 | 245.91 | 4.53 | 80505 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 288949/288949 | 0 | 0 | 4815.56 | 4.765 | 5.406 | 6.800 | 14.080 | true | 0 | 244.62 | 4.75 | 79969 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 288681/288681 | 0 | 0 | 4811.06 | 8.779 | 10.495 | 13.297 | 26.076 | true | 0 | 240.98 | 4.69 | 77171 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 287373/287373 | 0 | 0 | 4789.16 | 8.812 | 10.445 | 13.359 | 30.994 | true | 0 | 250.80 | 4.83 | 79083 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 287618/287618 | 0 | 0 | 4793.35 | 8.830 | 10.537 | 13.303 | 25.647 | true | 0 | 259.19 | 5.01 | 78851 | 3 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 157492/157492 | 0 | 0 | 2624.75 | 8.706 | 10.056 | 12.576 | 23.213 | true | 0 | 169.76 | 2.74 | 52435 | 3 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 160650/160650 | 0 | 0 | 2677.35 | 8.541 | 9.893 | 12.213 | 22.909 | true | 0 | 167.36 | 2.69 | 52974 | 3 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 163287/163287 | 0 | 0 | 2721.27 | 8.300 | 9.731 | 11.743 | 21.222 | true | 0 | 168.60 | 2.73 | 53048 | 4 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 164558/164558 | 0 | 0 | 2742.44 | 15.095 | 17.925 | 23.137 | 40.124 | true | 0 | 174.19 | 2.77 | 53819 | 3 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 165810/165810 | 0 | 0 | 2763.27 | 14.971 | 17.803 | 22.860 | 37.629 | true | 0 | 182.17 | 3.07 | 53599 | 3 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 166335/166335 | 0 | 0 | 2771.99 | 14.845 | 17.843 | 22.870 | 47.740 | true | 0 | 167.39 | 2.93 | 54656 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 455330/455330 | 0 | 0 | 7588.64 | 3.133 | 3.391 | 4.365 | 12.828 | true | 0 | 325.19 | 8.68 | 116993 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 454165/454165 | 0 | 0 | 7569.25 | 3.135 | 3.360 | 4.335 | 13.684 | true | 0 | 350.78 | 8.70 | 120161 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 461285/461285 | 0 | 0 | 7687.99 | 2.960 | 3.297 | 4.282 | 9.241 | true | 0 | 321.90 | 8.50 | 121319 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 463824/463824 | 0 | 0 | 7730.01 | 5.684 | 6.509 | 8.243 | 14.000 | true | 0 | 343.30 | 8.73 | 117228 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 458429/458429 | 0 | 0 | 7640.14 | 5.797 | 6.598 | 8.419 | 15.700 | true | 0 | 362.31 | 8.70 | 120220 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 451906/451906 | 0 | 0 | 7531.44 | 5.828 | 6.678 | 8.533 | 19.719 | true | 0 | 369.52 | 8.90 | 117062 | 3 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 306187/306187 | 0 | 0 | 5102.96 | 4.696 | 5.211 | 6.304 | 13.443 | true | 0 | 281.66 | 4.57 | 39150 | 3 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 304431/304431 | 0 | 0 | 5073.70 | 4.727 | 5.261 | 6.466 | 14.447 | true | 0 | 272.66 | 4.60 | 38854 | 3 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 303299/303299 | 0 | 0 | 5054.89 | 4.737 | 5.290 | 6.406 | 14.668 | true | 0 | 271.86 | 4.44 | 39422 | 3 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 314045/314045 | 0 | 0 | 5233.82 | 8.869 | 10.203 | 12.439 | 29.552 | true | 0 | 270.13 | 4.54 | 39972 | 3 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 316123/316123 | 0 | 0 | 5268.35 | 8.710 | 10.172 | 12.273 | 23.822 | true | 0 | 298.02 | 4.47 | 40445 | 3 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 320116/320116 | 0 | 0 | 5334.89 | 8.744 | 9.995 | 12.280 | 25.817 | true | 0 | 285.84 | 4.64 | 40661 | 3 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 70474/70474 | 0 | 0 | 1174.36 | 20.829 | 22.750 | 27.113 | 47.376 | true | 0 | 254.80 | 1.75 | 9200 | 3 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 70331/70331 | 0 | 0 | 1172.03 | 20.923 | 22.820 | 26.969 | 44.477 | true | 0 | 253.72 | 1.56 | 9298 | 3 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 69610/69610 | 0 | 0 | 1160.03 | 21.095 | 23.233 | 27.737 | 53.583 | true | 0 | 301.89 | 1.68 | 9084 | 4 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 70519/70519 | 0 | 0 | 1175.12 | 40.068 | 46.985 | 57.972 | 102.253 | true | 0 | 257.70 | 1.68 | 9146 | 4 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 70195/70195 | 0 | 0 | 1169.67 | 40.179 | 47.090 | 59.070 | 108.442 | true | 0 | 262.40 | 1.51 | 9141 | 4 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 71473/71473 | 0 | 0 | 1190.83 | 39.667 | 46.148 | 56.945 | 95.114 | true | 0 | 247.73 | 1.70 | 9194 | 3 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 140348/140348 | 0 | 0 | 2339.01 | 10.241 | 11.273 | 13.581 | 23.750 | true | 0 | 277.14 | 2.81 | 17892 | 3 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 140297/140297 | 0 | 0 | 2338.11 | 10.204 | 11.252 | 13.790 | 22.563 | true | 0 | 266.41 | 2.87 | 17904 | 4 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 139477/139477 | 0 | 0 | 2324.45 | 10.288 | 11.336 | 13.744 | 24.220 | true | 0 | 258.52 | 2.77 | 17836 | 3 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 141751/141751 | 0 | 0 | 2362.17 | 20.010 | 22.744 | 27.436 | 49.127 | true | 0 | 280.37 | 2.75 | 18320 | 3 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 142963/142963 | 0 | 0 | 2382.44 | 19.871 | 22.454 | 27.247 | 50.609 | true | 0 | 262.57 | 2.77 | 18255 | 5 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 142775/142775 | 0 | 0 | 2379.31 | 19.945 | 22.602 | 27.173 | 51.562 | true | 0 | 264.58 | 2.85 | 18308 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/3 | 35306/35306 | 0 | 0 | 588.31 | 33.178 | 35.472 | 40.818 | 55.717 | true | 0 | 250.87 | 0.85 | 4966 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 2/3 | 34876/34876 | 0 | 0 | 581.12 | 33.568 | 35.945 | 41.641 | 59.120 | true | 0 | 254.08 | 0.82 | 5004 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 3/3 | 34790/34790 | 0 | 0 | 579.73 | 34.114 | 36.742 | 42.858 | 64.377 | true | 0 | 250.50 | 0.89 | 4995 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/3 | 35025/35025 | 0 | 0 | 583.54 | 67.036 | 71.412 | 81.794 | 107.950 | true | 0 | 275.97 | 0.77 | 4895 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 2/3 | 35208/35208 | 0 | 0 | 586.55 | 66.803 | 72.171 | 82.549 | 113.055 | true | 0 | 281.43 | 0.71 | 5024 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 3/3 | 35474/35474 | 0 | 0 | 590.99 | 65.639 | 70.333 | 80.028 | 105.416 | true | 0 | 266.25 | 0.80 | 5116 | 3 |
