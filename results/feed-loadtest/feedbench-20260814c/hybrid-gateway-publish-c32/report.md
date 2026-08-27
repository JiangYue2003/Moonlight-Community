# Feed 压测报告：hybrid / gateway / publish-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:49:48+08:00
- 采样时长：12.7927635s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 128 | 128 | 0 | 0 | 10.01 | 3201.713 | 3240.688 | 3256.625 |
| publish_draft | 128 | 128 | 0 | 0 | 10.01 | 10.825 | 22.980 | 24.064 |
| publish_metadata | 128 | 128 | 0 | 0 | 10.01 | 1073.295 | 1093.647 | 1099.902 |
| publish_confirm | 128 | 128 | 0 | 0 | 10.01 | 1036.745 | 1144.239 | 1147.047 |
| publish_commit | 128 | 128 | 0 | 0 | 10.01 | 1049.523 | 1159.944 | 1169.280 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.920 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 10.420 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 143.870 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 6.480 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 17.000 |
| kafka | lag_total | 17.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 425822.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 116.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9232348.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539689.000 |
| redis | keyspace_hits | 12815078.000 |
| redis | keyspace_misses | 122919.000 |
| redis | net_input_bytes | 964818040.000 |
| redis | net_output_bytes | 2770676333.000 |
| redis | ops_per_sec | 16323.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93333880.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
