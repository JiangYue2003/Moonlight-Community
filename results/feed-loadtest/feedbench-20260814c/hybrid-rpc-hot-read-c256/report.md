# Feed 压测报告：hybrid / rpc / hot-read-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:45:34+08:00
- 采样时长：3.117941s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4723 | 4723 | 0 | 0 | 1514.78 | 165.050 | 201.642 | 268.804 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 80.170 |
| docker:zg-counter | memory_percent | 0.230 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.160 |
| docker:zg-gateway | memory_percent | 0.160 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 250.670 |
| docker:zg-knowpost | memory_percent | 0.540 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 146.060 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 384310.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 131.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7269028.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 538728.000 |
| redis | keyspace_hits | 9533704.000 |
| redis | keyspace_misses | 106787.000 |
| redis | net_input_bytes | 776565632.000 |
| redis | net_output_bytes | 2204519938.000 |
| redis | ops_per_sec | 50173.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 99579520.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
