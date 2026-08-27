# Feed 压测报告：hybrid / gateway / mixed-80-20-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:51:32+08:00
- 采样时长：43.8097591s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 127.80 | 7.779 | 9.404 | 10.762 |
| publish_total | 60 | 60 | 0 | 0 | 1.37 | 725.965 | 779.748 | 814.446 |
| publish_draft | 60 | 60 | 0 | 0 | 1.37 | 6.661 | 7.914 | 68.480 |
| publish_metadata | 60 | 60 | 0 | 0 | 1.37 | 239.913 | 254.996 | 349.294 |
| publish_confirm | 60 | 60 | 0 | 0 | 1.37 | 234.916 | 252.755 | 260.152 |
| publish_commit | 60 | 60 | 0 | 0 | 1.37 | 243.153 | 258.321 | 267.387 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 10.800 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 18.260 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 51.970 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 15.280 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 443553.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9972905.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 540357.000 |
| redis | keyspace_hits | 13788809.000 |
| redis | keyspace_misses | 127095.000 |
| redis | net_input_bytes | 1022438194.000 |
| redis | net_output_bytes | 2993731917.000 |
| redis | ops_per_sec | 3643.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 93722136.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
