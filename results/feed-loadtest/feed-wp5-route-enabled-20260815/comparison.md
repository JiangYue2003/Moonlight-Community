# Feed 三策略综合对比：feed-wp5-route-enabled-20260815

容量口径：`stable c` 是吞吐增幅首次低于 10% 且 P95 上升前的并发档；`knee c` 是对应拐点；失败并发始终作为硬边界。该口径用于定位性能曲线，不是 SLA 判定。

| strategy | entry | scenario | stage | peak no-error QPS | P95@peak(ms) | stable c | knee c | no-error upper c | first-error c |
|---|---|---|---|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 3978.74 | 22.054 | 64 | 128 | 256 | - |

## 重复轮次中位数与波动

| strategy | entry | scenario | stage | c | complete/runs | QPS median | min | max | spread | P95 median(ms) | min | max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 32 | 3/3 | 3572.68 | 3555.13 | 3590.73 | 1.00% | 12.050 | 11.932 | 12.082 |
| hybrid | rpc | distributed-read | read | 64 | 3/3 | 3978.74 | 3968.69 | 3981.83 | 0.33% | 22.054 | 22.002 | 22.058 |
| hybrid | rpc | distributed-read | read | 128 | 3/3 | 3896.52 | 3861.64 | 3989.05 | 3.27% | 49.301 | 47.190 | 49.414 |
| hybrid | rpc | distributed-read | read | 256 | 3/3 | 3884.14 | 3861.21 | 3887.13 | 0.67% | 88.936 | 87.092 | 89.979 |

## 原始指标

| strategy | entry | scenario | stage | c | trial/expected | success/total | failed | timeout | QPS | P90(ms) | P95(ms) | P99(ms) | Max(ms) | complete | Kafka lag max | KnowPost CPU max(%) | Redis ops/s max | MySQL running max |
|---|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|---:|---:|---:|---:|
| hybrid | rpc | distributed-read | read | 32 | 1/3 | 213333/213333 | 0 | 0 | 3555.13 | 11.171 | 12.082 | 14.171 | 30.754 | true | 0 | 396.43 | 25650 | 3 |
| hybrid | rpc | distributed-read | read | 32 | 2/3 | 214380/214380 | 0 | 0 | 3572.68 | 11.148 | 12.050 | 14.230 | 32.584 | true | 0 | 393.44 | 25328 | 3 |
| hybrid | rpc | distributed-read | read | 32 | 3/3 | 215470/215470 | 0 | 0 | 3590.73 | 11.086 | 11.932 | 13.989 | 30.987 | true | 0 | 394.58 | 25500 | 3 |
| hybrid | rpc | distributed-read | read | 64 | 1/3 | 238948/238948 | 0 | 0 | 3981.83 | 20.384 | 22.002 | 25.971 | 49.960 | true | 0 | 432.38 | 29011 | 3 |
| hybrid | rpc | distributed-read | read | 64 | 2/3 | 238163/238163 | 0 | 0 | 3968.69 | 20.462 | 22.054 | 25.897 | 56.250 | true | 0 | 426.43 | 29335 | 3 |
| hybrid | rpc | distributed-read | read | 64 | 3/3 | 238777/238777 | 0 | 0 | 3978.74 | 20.406 | 22.058 | 26.003 | 56.188 | true | 0 | 438.61 | 28955 | 3 |
| hybrid | rpc | distributed-read | read | 128 | 1/3 | 233866/233866 | 0 | 0 | 3896.52 | 45.413 | 49.414 | 56.144 | 100.754 | true | 0 | 410.34 | 30588 | 3 |
| hybrid | rpc | distributed-read | read | 128 | 2/3 | 231771/231771 | 0 | 0 | 3861.64 | 45.602 | 49.301 | 55.978 | 100.694 | true | 0 | 432.21 | 29910 | 3 |
| hybrid | rpc | distributed-read | read | 128 | 3/3 | 239419/239419 | 0 | 0 | 3989.05 | 43.554 | 47.190 | 53.937 | 98.391 | true | 0 | 419.51 | 30815 | 3 |
| hybrid | rpc | distributed-read | read | 256 | 1/3 | 233194/233194 | 0 | 0 | 3884.14 | 82.928 | 87.092 | 102.477 | 186.461 | true | 0 | 443.71 | 30464 | 3 |
| hybrid | rpc | distributed-read | read | 256 | 2/3 | 233437/233437 | 0 | 0 | 3887.13 | 84.013 | 88.936 | 102.037 | 190.773 | true | 0 | 445.47 | 30487 | 3 |
| hybrid | rpc | distributed-read | read | 256 | 3/3 | 231833/231833 | 0 | 0 | 3861.21 | 85.061 | 89.979 | 104.378 | 206.395 | true | 0 | 428.77 | 30317 | 3 |
