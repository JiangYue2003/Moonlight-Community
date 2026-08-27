# Feed 压测报告：push / rpc / mixed-80-20-c256

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:29:15+08:00
- 采样时长：5.773648s
- 并发：256
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 1563.70 | 105.870 | 135.326 | 144.686 |
| publish_total | 60 | 60 | 0 | 0 | 10.39 | 4832.303 | 4883.726 | 4893.036 |
| publish_draft | 60 | 60 | 0 | 0 | 10.39 | 44.355 | 95.419 | 111.159 |
| publish_metadata | 60 | 60 | 0 | 0 | 10.39 | 1703.924 | 1785.843 | 1800.946 |
| publish_confirm | 60 | 60 | 0 | 0 | 10.39 | 1456.741 | 1530.875 | 1546.553 |
| publish_commit | 60 | 60 | 0 | 0 | 10.39 | 1592.610 | 1651.350 | 1674.192 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.110 |
| docker:zg-counter | memory_percent | 0.130 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.310 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 117.470 |
| docker:zg-knowpost | memory_percent | 0.470 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 16.390 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 52.000 |
| kafka | lag_total | 52.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 276928.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 131.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 2708573.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 537262.000 |
| redis | keyspace_hits | 4722632.000 |
| redis | keyspace_misses | 32576.000 |
| redis | net_input_bytes | 338664072.000 |
| redis | net_output_bytes | 897009177.000 |
| redis | ops_per_sec | 13114.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 112966992.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
