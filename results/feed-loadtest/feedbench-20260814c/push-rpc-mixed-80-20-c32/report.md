# Feed 压测报告：push / rpc / mixed-80-20-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:28:54+08:00
- 采样时长：7.8323812s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 1469.08 | 14.325 | 36.337 | 39.577 |
| publish_total | 60 | 60 | 0 | 0 | 7.66 | 867.728 | 1012.335 | 1017.151 |
| publish_draft | 60 | 60 | 0 | 0 | 7.66 | 6.435 | 21.805 | 24.020 |
| publish_metadata | 60 | 60 | 0 | 0 | 7.66 | 283.979 | 375.868 | 382.434 |
| publish_confirm | 60 | 60 | 0 | 0 | 7.66 | 279.385 | 330.137 | 332.963 |
| publish_commit | 60 | 60 | 0 | 0 | 7.66 | 289.706 | 331.730 | 334.179 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.350 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.240 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 92.590 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 4.830 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 7.000 |
| kafka | lag_total | 7.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 274589.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 35.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2487422.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 537108.000 |
| redis | keyspace_hits | 4705520.000 |
| redis | keyspace_misses | 26745.000 |
| redis | net_input_bytes | 321113169.000 |
| redis | net_output_bytes | 890797214.000 |
| redis | ops_per_sec | 16680.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 107926920.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
