# Feed 压测报告：push / gateway / mixed-80-20-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:32:32+08:00
- 采样时长：6.8236304s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 1066.94 | 77.490 | 135.654 | 148.188 |
| publish_total | 60 | 60 | 0 | 0 | 8.79 | 2909.471 | 2943.508 | 2946.813 |
| publish_draft | 60 | 60 | 0 | 0 | 8.79 | 38.095 | 133.487 | 159.456 |
| publish_metadata | 60 | 60 | 0 | 0 | 8.79 | 943.547 | 1147.393 | 1160.883 |
| publish_confirm | 60 | 60 | 0 | 0 | 8.79 | 874.192 | 899.270 | 913.151 |
| publish_commit | 60 | 60 | 0 | 0 | 8.79 | 894.169 | 974.896 | 993.321 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.060 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 15.760 |
| docker:zg-gateway | memory_percent | 0.190 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 119.880 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 10.910 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 307360.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 94.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3774058.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 538706.000 |
| redis | keyspace_hits | 5636557.000 |
| redis | keyspace_misses | 42637.000 |
| redis | net_input_bytes | 444929572.000 |
| redis | net_output_bytes | 1078435259.000 |
| redis | ops_per_sec | 15870.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 129252000.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
