# Feed 压测报告：hybrid / rpc / mixed-80-20-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:47:44+08:00
- 采样时长：41.1041451s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 149.67 | 6.498 | 8.014 | 8.566 |
| publish_total | 60 | 60 | 0 | 0 | 1.46 | 680.477 | 722.663 | 765.069 |
| publish_draft | 60 | 60 | 0 | 0 | 1.46 | 5.301 | 6.197 | 8.197 |
| publish_metadata | 60 | 60 | 0 | 0 | 1.46 | 224.476 | 241.048 | 251.902 |
| publish_confirm | 60 | 60 | 0 | 0 | 1.46 | 222.611 | 238.622 | 280.768 |
| publish_commit | 60 | 60 | 0 | 0 | 1.46 | 228.371 | 241.802 | 260.453 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 17.120 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 1.670 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 55.640 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 12.300 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 418923.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 93.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8603631.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 538964.000 |
| redis | keyspace_hits | 12747528.000 |
| redis | keyspace_misses | 113080.000 |
| redis | net_input_bytes | 916186554.000 |
| redis | net_output_bytes | 2748006341.000 |
| redis | ops_per_sec | 2730.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89457240.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
