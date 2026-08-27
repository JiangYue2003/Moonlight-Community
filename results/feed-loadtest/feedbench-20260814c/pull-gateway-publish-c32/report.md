# Feed 压测报告：pull / gateway / publish-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:39:46+08:00
- 采样时长：11.9756869s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 128 | 128 | 0 | 0 | 10.69 | 2964.819 | 3119.204 | 3129.095 |
| publish_draft | 128 | 128 | 0 | 0 | 10.69 | 11.504 | 20.593 | 21.121 |
| publish_metadata | 128 | 128 | 0 | 0 | 10.69 | 995.767 | 1110.229 | 1121.925 |
| publish_confirm | 128 | 128 | 0 | 0 | 10.69 | 921.490 | 1021.439 | 1023.070 |
| publish_commit | 128 | 128 | 0 | 0 | 10.69 | 952.822 | 1135.945 | 1144.983 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.460 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 8.130 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 143.980 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 3.200 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 348185.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 117.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5641693.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 537449.000 |
| redis | keyspace_hits | 7857799.000 |
| redis | keyspace_misses | 86348.000 |
| redis | net_input_bytes | 640503485.000 |
| redis | net_output_bytes | 1693980119.000 |
| redis | ops_per_sec | 13345.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87356192.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
