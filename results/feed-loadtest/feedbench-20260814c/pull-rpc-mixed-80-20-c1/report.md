# Feed 压测报告：pull / rpc / mixed-80-20-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:37:45+08:00
- 采样时长：39.8383694s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 147.30 | 6.657 | 8.033 | 9.088 |
| publish_total | 60 | 60 | 0 | 0 | 1.51 | 659.972 | 688.840 | 800.794 |
| publish_draft | 60 | 60 | 0 | 0 | 1.51 | 5.295 | 6.448 | 10.585 |
| publish_metadata | 60 | 60 | 0 | 0 | 1.51 | 216.841 | 231.290 | 257.075 |
| publish_confirm | 60 | 60 | 0 | 0 | 1.51 | 216.046 | 230.343 | 255.418 |
| publish_commit | 60 | 60 | 0 | 0 | 1.51 | 222.447 | 238.229 | 283.186 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.740 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.920 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 79.660 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 12.270 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 341636.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 93.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5161372.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 537439.000 |
| redis | keyspace_hits | 7795128.000 |
| redis | keyspace_misses | 77811.000 |
| redis | net_input_bytes | 601043232.000 |
| redis | net_output_bytes | 1655048597.000 |
| redis | ops_per_sec | 2407.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 85746984.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
