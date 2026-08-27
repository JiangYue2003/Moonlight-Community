# Feed 压测报告：hybrid / rpc / mixed-80-20-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:49:04+08:00
- 采样时长：6.3619718s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 942.32 | 95.465 | 129.209 | 133.420 |
| publish_total | 60 | 60 | 0 | 0 | 9.43 | 2546.779 | 2839.019 | 2844.630 |
| publish_draft | 60 | 60 | 0 | 0 | 9.43 | 11.267 | 69.926 | 80.658 |
| publish_metadata | 60 | 60 | 0 | 0 | 9.43 | 842.520 | 1083.742 | 1089.902 |
| publish_confirm | 60 | 60 | 0 | 0 | 9.43 | 861.910 | 892.134 | 896.854 |
| publish_commit | 60 | 60 | 0 | 0 | 9.43 | 804.180 | 858.457 | 874.042 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.490 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.040 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 136.470 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 3.650 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 16.000 |
| kafka | lag_total | 16.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 422364.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 93.000 |
| mysql | threads_running | 5.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8880650.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539153.000 |
| redis | keyspace_hits | 12798525.000 |
| redis | keyspace_misses | 117503.000 |
| redis | net_input_bytes | 937849240.000 |
| redis | net_output_bytes | 2761949735.000 |
| redis | ops_per_sec | 16018.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91401648.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
