# Feed 压测报告：pull / rpc / publish-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:34:21+08:00
- 采样时长：19.7449511s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 30 | 30 | 0 | 0 | 1.52 | 657.804 | 692.240 | 702.435 |
| publish_draft | 30 | 30 | 0 | 0 | 1.52 | 5.437 | 7.984 | 23.690 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.52 | 216.011 | 223.587 | 230.985 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.52 | 212.532 | 221.741 | 223.337 |
| publish_commit | 30 | 30 | 0 | 0 | 1.52 | 222.502 | 241.743 | 254.714 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.680 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.640 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 19.660 |
| docker:zg-knowpost | memory_percent | 0.150 |
| docker:zg-knowpost | pids | 20.000 |
| docker:zg-relation | cpu_percent | 1.660 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 309551.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 4.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4004541.000 |
| redis | connected_clients | 230.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 537470.000 |
| redis | keyspace_hits | 5637284.000 |
| redis | keyspace_misses | 45121.000 |
| redis | net_input_bytes | 462631013.000 |
| redis | net_output_bytes | 1082288258.000 |
| redis | ops_per_sec | 1935.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 81741280.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
