# Feed 压测报告：hybrid / rpc / publish-c48

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:58:04+08:00
- 采样时长：17.8103073s
- 并发：48
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 192 | 32 | 160 | 0 | 1.80 | 4924.236 | 4992.226 | 4993.002 |
| publish_draft | 192 | 192 | 0 | 0 | 10.78 | 12.448 | 54.422 | 64.459 |
| publish_metadata | 192 | 192 | 0 | 0 | 10.78 | 1664.015 | 1822.873 | 1861.292 |
| publish_confirm | 192 | 192 | 0 | 0 | 10.78 | 1691.530 | 1852.647 | 1878.337 |
| publish_commit | 192 | 32 | 160 | 0 | 1.80 | 1607.519 | 1693.781 | 1707.601 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 7.970 |
| docker:zg-counter | memory_percent | 0.200 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.310 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 133.940 |
| docker:zg-knowpost | memory_percent | 0.240 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 9.110 |
| docker:zg-relation | memory_percent | 0.250 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 42.000 |
| kafka | lag_total | 42.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 448476.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 52.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 10531072.000 |
| redis | connected_clients | 366.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 541027.000 |
| redis | keyspace_hits | 13824629.000 |
| redis | keyspace_misses | 132803.000 |
| redis | net_input_bytes | 1064616856.000 |
| redis | net_output_bytes | 3010196781.000 |
| redis | ops_per_sec | 27484.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 97483656.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
