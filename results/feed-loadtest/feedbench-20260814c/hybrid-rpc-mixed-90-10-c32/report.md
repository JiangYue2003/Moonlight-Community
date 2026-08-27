# Feed 压测报告：hybrid / rpc / mixed-90-10-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:47:20+08:00
- 采样时长：5.9467665s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 1019.95 | 25.267 | 40.888 | 43.530 |
| publish_total | 30 | 30 | 0 | 0 | 5.04 | 701.643 | 980.779 | 980.779 |
| publish_draft | 30 | 30 | 0 | 0 | 5.04 | 6.844 | 15.919 | 15.930 |
| publish_metadata | 30 | 30 | 0 | 0 | 5.04 | 242.707 | 462.702 | 465.906 |
| publish_confirm | 30 | 30 | 0 | 0 | 5.04 | 225.969 | 258.542 | 258.551 |
| publish_commit | 30 | 30 | 0 | 0 | 5.04 | 239.039 | 258.950 | 259.514 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.930 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.260 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 89.190 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 3.230 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 416214.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 103.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8398621.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 538830.000 |
| redis | keyspace_hits | 12691918.000 |
| redis | keyspace_misses | 110037.000 |
| redis | net_input_bytes | 900273824.000 |
| redis | net_output_bytes | 2735486391.000 |
| redis | ops_per_sec | 8127.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89987008.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
