# Feed 压测报告：hybrid / rpc / publish-c36

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T16:00:53+08:00
- 采样时长：15.1384813s
- 并发：36
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 144 | 144 | 0 | 0 | 9.51 | 3816.148 | 3879.471 | 3888.936 |
| publish_draft | 144 | 144 | 0 | 0 | 9.51 | 10.712 | 62.821 | 66.576 |
| publish_metadata | 144 | 144 | 0 | 0 | 9.51 | 1303.085 | 1344.518 | 1358.220 |
| publish_confirm | 144 | 144 | 0 | 0 | 9.51 | 1249.543 | 1301.728 | 1308.123 |
| publish_commit | 144 | 144 | 0 | 0 | 9.51 | 1226.236 | 1342.161 | 1353.461 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 5.680 |
| docker:zg-counter | memory_percent | 0.200 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.090 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 142.280 |
| docker:zg-knowpost | memory_percent | 0.210 |
| docker:zg-knowpost | pids | 24.000 |
| docker:zg-relation | cpu_percent | 7.900 |
| docker:zg-relation | memory_percent | 0.240 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 31.000 |
| kafka | lag_total | 31.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 451728.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 40.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 10801506.000 |
| redis | connected_clients | 353.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 541335.000 |
| redis | keyspace_hits | 13825376.000 |
| redis | keyspace_misses | 135744.000 |
| redis | net_input_bytes | 1085507678.000 |
| redis | net_output_bytes | 3014869736.000 |
| redis | ops_per_sec | 17532.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 98200840.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
