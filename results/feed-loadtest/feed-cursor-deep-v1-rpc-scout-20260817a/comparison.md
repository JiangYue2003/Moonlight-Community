# Feed 三策略综合对比：feed-cursor-deep-v1-rpc-scout-20260817a

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | cardinality | cache state | topology | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 4712.79 | 5.513 | 16 | 32 | 64 | - |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 2839.31 | 30.436 | 16 | 32 | 64 | - |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 7282.40 | 6.994 | 16 | 32 | 64 | - |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 5320.16 | 17.650 | 16 | 32 | 64 | - |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 1205.60 | 81.329 | 16 | 32 | 64 | - |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 2351.00 | 22.997 | 16 | 32 | 64 | - |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 591.04 | 139.912 | 16 | 32 | 64 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 4712.79 | 4712.79 | 4712.79 | 0.00% | 5.513 | 5.513 | 5.513 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 4656.45 | 4656.45 | 4656.45 | 0.00% | 10.795 | 10.795 | 10.795 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 4707.75 | 4707.75 | 4707.75 | 0.00% | 19.340 | 19.340 | 19.340 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 2728.95 | 2728.95 | 2728.95 | 0.00% | 9.708 | 9.708 | 9.708 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 2808.68 | 2808.68 | 2808.68 | 0.00% | 17.194 | 17.194 | 17.194 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 2839.31 | 2839.31 | 2839.31 | 0.00% | 30.436 | 30.436 | 30.436 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 7151.19 | 7151.19 | 7151.19 | 0.00% | 3.736 | 3.736 | 3.736 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 7282.40 | 7282.40 | 7282.40 | 0.00% | 6.994 | 6.994 | 6.994 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 7254.54 | 7254.54 | 7254.54 | 0.00% | 12.755 | 12.755 | 12.755 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 5214.52 | 5214.52 | 5214.52 | 0.00% | 5.076 | 5.076 | 5.076 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 5240.29 | 5240.29 | 5240.29 | 0.00% | 10.204 | 10.204 | 10.204 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 5320.16 | 5320.16 | 5320.16 | 0.00% | 17.650 | 17.650 | 17.650 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 1179.76 | 1179.76 | 1179.76 | 0.00% | 22.576 | 22.576 | 22.576 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 1202.67 | 1202.67 | 1202.67 | 0.00% | 44.167 | 44.167 | 44.167 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 1205.60 | 1205.60 | 1205.60 | 0.00% | 81.329 | 81.329 | 81.329 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 2306.35 | 2306.35 | 2306.35 | 0.00% | 11.694 | 11.694 | 11.694 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 2351.00 | 2351.00 | 2351.00 | 0.00% | 22.997 | 22.997 | 22.997 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 2299.85 | 2299.85 | 2299.85 | 0.00% | 36.832 | 36.832 | 36.832 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 582.04 | 582.04 | 582.04 | 0.00% | 35.626 | 35.626 | 35.626 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 575.74 | 575.74 | 575.74 | 0.00% | 72.409 | 72.409 | 72.409 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 591.04 | 591.04 | 591.04 | 0.00% | 139.912 | 139.912 | 139.912 |

## 原始指标

| strategy | entry | scenario | cardinality | cache state | topology | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Client CPU normalized max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|---:|
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 47133/47133 | 0 | 0 | 4712.79 | 4.808 | 5.513 | 6.926 | 12.197 | true | 0 | 211.60 | 4.62 | 76290 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 46598/46598 | 0 | 0 | 4656.45 | 9.071 | 10.795 | 13.781 | 23.427 | true | 0 | 219.20 | 4.84 | 73582 | 3 |
| hybrid | rpc | cursor-page20 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 47121/47121 | 0 | 0 | 4707.75 | 17.372 | 19.340 | 25.855 | 42.917 | true | 0 | 281.82 | 6.00 | 76847 | 3 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 27301/27301 | 0 | 0 | 2728.95 | 8.281 | 9.708 | 11.839 | 22.647 | true | 0 | 150.19 | 2.88 | 52860 | 3 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 28109/28109 | 0 | 0 | 2808.68 | 14.628 | 17.194 | 22.521 | 35.144 | true | 0 | 151.83 | 2.73 | 52560 | 3 |
| hybrid | rpc | cursor-page5 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 28430/28430 | 0 | 0 | 2839.31 | 27.860 | 30.436 | 42.694 | 62.323 | true | 0 | 177.44 | 3.16 | 55476 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 71532/71532 | 0 | 0 | 7151.19 | 3.209 | 3.736 | 4.830 | 8.990 | true | 0 | 296.28 | 10.23 | 110001 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 72841/72841 | 0 | 0 | 7282.40 | 5.985 | 6.994 | 8.791 | 16.042 | true | 0 | 286.83 | 9.96 | 106277 | 3 |
| hybrid | rpc | cursor-page50 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 72594/72594 | 0 | 0 | 7254.54 | 11.398 | 12.755 | 17.002 | 27.149 | true | 0 | 314.10 | 9.40 | 108922 | 3 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 52151/52151 | 0 | 0 | 5214.52 | 4.566 | 5.076 | 6.213 | 10.083 | true | 0 | 257.13 | 4.65 | 39501 | 3 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 52424/52424 | 0 | 0 | 5240.29 | 8.643 | 10.204 | 12.362 | 23.605 | true | 0 | 255.45 | 5.10 | 39299 | 3 |
| hybrid | rpc | page1 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 53225/53225 | 0 | 0 | 5320.16 | 14.536 | 17.650 | 23.933 | 51.344 | true | 0 | 247.55 | 4.74 | 40691 | 3 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 11811/11811 | 0 | 0 | 1179.76 | 20.535 | 22.576 | 26.855 | 38.609 | true | 0 | 217.95 | 1.95 | 8982 | 3 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 12045/12045 | 0 | 0 | 1202.67 | 35.435 | 44.167 | 55.380 | 80.958 | true | 0 | 224.36 | 1.61 | 9143 | 3 |
| hybrid | rpc | page20 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 12104/12104 | 0 | 0 | 1205.60 | 64.949 | 81.329 | 106.421 | 168.885 | true | 0 | 635.21 | 1.66 | 8802 | 3 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 23070/23070 | 0 | 0 | 2306.35 | 10.583 | 11.694 | 13.733 | 20.145 | true | 0 | 187.46 | 2.95 | 17117 | 3 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 23529/23529 | 0 | 0 | 2351.00 | 17.610 | 22.997 | 27.631 | 41.851 | true | 0 | 213.61 | 2.95 | 17206 | 4 |
| hybrid | rpc | page5 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 23029/23029 | 0 | 0 | 2299.85 | 33.225 | 36.832 | 54.661 | 77.350 | true | 0 | 243.83 | 3.05 | 17638 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 16 | 1/1 | 5827/5827 | 0 | 0 | 582.04 | 33.565 | 35.626 | 40.521 | 52.619 | true | 0 | 245.49 | 1.00 | 4722 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 32 | 1/1 | 5772/5772 | 0 | 0 | 575.74 | 67.998 | 72.409 | 82.763 | 103.148 | true | 0 | 247.87 | 0.63 | 4786 | 3 |
| hybrid | rpc | page50 | cursor-deep | cold | compose-middleware-local-services | read | 64 | 1/1 | 5954/5954 | 0 | 0 | 591.04 | 129.599 | 139.912 | 160.246 | 206.831 | true | 0 | 383.80 | 0.63 | 4759 | 3 |
