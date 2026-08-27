# Feed 压测报告：hybrid / gateway / mixed-80-20-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:52:34+08:00
- 采样时长：6.9550378s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 240 | 240 | 0 | 0 | 706.16 | 116.810 | 176.268 | 195.349 |
| publish_total | 60 | 60 | 0 | 0 | 8.63 | 2777.545 | 3212.314 | 3219.971 |
| publish_draft | 60 | 60 | 0 | 0 | 8.63 | 12.205 | 78.038 | 90.678 |
| publish_metadata | 60 | 60 | 0 | 0 | 8.63 | 1035.073 | 1389.420 | 1401.821 |
| publish_confirm | 60 | 60 | 0 | 0 | 8.63 | 818.427 | 885.432 | 891.638 |
| publish_commit | 60 | 60 | 0 | 0 | 8.63 | 894.421 | 935.602 | 939.369 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.750 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 20.320 |
| docker:zg-gateway | memory_percent | 0.190 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 137.540 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 3.670 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 18.000 |
| kafka | lag_total | 18.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 445897.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 82.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 10160569.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 540618.000 |
| redis | keyspace_hits | 13823393.000 |
| redis | keyspace_misses | 129644.000 |
| redis | net_input_bytes | 1037092853.000 |
| redis | net_output_bytes | 3004507207.000 |
| redis | ops_per_sec | 13647.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 95753216.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
