# Feed 压测报告：push / rpc / mixed-80-20-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:27:47+08:00
- 采样时长：38.8385046s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 208.29 | 4.740 | 5.471 | 6.399 |
| publish_total | 60 | 60 | 0 | 0 | 1.54 | 645.677 | 673.586 | 733.975 |
| publish_draft | 60 | 60 | 0 | 0 | 1.54 | 5.284 | 6.755 | 8.006 |
| publish_metadata | 60 | 60 | 0 | 0 | 1.54 | 211.691 | 223.994 | 250.837 |
| publish_confirm | 60 | 60 | 0 | 0 | 1.54 | 211.044 | 224.338 | 236.089 |
| publish_commit | 60 | 60 | 0 | 0 | 1.54 | 215.960 | 226.859 | 239.043 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.770 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.870 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 38.040 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 21.430 |
| docker:zg-relation | memory_percent | 0.280 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 272239.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 112.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2249668.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 536886.000 |
| redis | keyspace_hits | 4684822.000 |
| redis | keyspace_misses | 24289.000 |
| redis | net_input_bytes | 302995856.000 |
| redis | net_output_bytes | 883905726.000 |
| redis | ops_per_sec | 6163.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 99762704.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
