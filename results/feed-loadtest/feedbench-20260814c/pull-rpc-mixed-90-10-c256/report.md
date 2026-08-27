# Feed 压测报告：pull / rpc / mixed-90-10-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:37:38+08:00
- 采样时长：3.7192105s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 510.79 | 426.261 | 478.084 | 487.729 |
| publish_total | 30 | 30 | 0 | 0 | 8.07 | 2996.307 | 3008.193 | 3008.193 |
| publish_draft | 30 | 30 | 0 | 0 | 8.07 | 49.054 | 93.920 | 97.661 |
| publish_metadata | 30 | 30 | 0 | 0 | 8.07 | 1387.143 | 1430.475 | 1431.509 |
| publish_confirm | 30 | 30 | 0 | 0 | 8.07 | 750.515 | 761.229 | 763.160 |
| publish_commit | 30 | 30 | 0 | 0 | 8.07 | 786.377 | 791.101 | 792.612 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.390 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 0.340 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 111.660 |
| docker:zg-knowpost | memory_percent | 0.620 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 1.610 |
| docker:zg-relation | memory_percent | 0.320 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 340569.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 93.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5080488.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 537418.000 |
| redis | keyspace_hits | 7779157.000 |
| redis | keyspace_misses | 76633.000 |
| redis | net_input_bytes | 594399686.000 |
| redis | net_output_bytes | 1647398133.000 |
| redis | ops_per_sec | 12111.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86969848.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
