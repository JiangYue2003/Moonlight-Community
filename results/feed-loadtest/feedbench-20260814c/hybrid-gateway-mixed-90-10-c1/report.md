# Feed 压测报告：hybrid / gateway / mixed-90-10-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:50:46+08:00
- 采样时长：21.6565767s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 125.61 | 7.805 | 9.897 | 11.059 |
| publish_total | 30 | 30 | 0 | 0 | 1.39 | 719.547 | 755.289 | 821.505 |
| publish_draft | 30 | 30 | 0 | 0 | 1.39 | 6.481 | 7.397 | 7.976 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.39 | 232.968 | 256.439 | 263.745 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.39 | 234.164 | 249.366 | 270.643 |
| publish_commit | 30 | 30 | 0 | 0 | 1.39 | 241.370 | 274.022 | 279.141 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 14.760 |
| docker:zg-counter | memory_percent | 0.220 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 18.310 |
| docker:zg-gateway | memory_percent | 0.170 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 51.930 |
| docker:zg-knowpost | memory_percent | 0.410 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 20.780 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 440780.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 106.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9761069.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 540051.000 |
| redis | keyspace_hits | 13732656.000 |
| redis | keyspace_misses | 124252.000 |
| redis | net_input_bytes | 1006010663.000 |
| redis | net_output_bytes | 2977918063.000 |
| redis | ops_per_sec | 5906.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92543728.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
