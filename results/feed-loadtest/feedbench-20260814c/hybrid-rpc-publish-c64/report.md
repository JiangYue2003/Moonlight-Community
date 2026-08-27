# Feed 压测报告：hybrid / rpc / publish-c64

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:58:28+08:00
- 采样时长：2.7910862s
- 并发：64
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 256 | 2 | 254 | 0 | 0.72 | 751.080 | 767.230 | 767.230 |
| publish_draft | 256 | 256 | 0 | 0 | 91.74 | 18.378 | 33.882 | 36.092 |
| publish_metadata | 256 | 2 | 254 | 0 | 0.72 | 261.579 | 272.010 | 272.010 |
| publish_confirm | 2 | 2 | 0 | 0 | 0.72 | 226.933 | 229.090 | 229.090 |
| publish_commit | 2 | 2 | 0 | 0 | 0.72 | 245.380 | 246.532 | 246.532 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.040 |
| docker:zg-counter | memory_percent | 0.190 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 0.240 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 115.980 |
| docker:zg-knowpost | memory_percent | 0.260 |
| docker:zg-knowpost | pids | 25.000 |
| docker:zg-relation | cpu_percent | 1.690 |
| docker:zg-relation | memory_percent | 0.260 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 449104.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 10558991.000 |
| redis | connected_clients | 401.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 541040.000 |
| redis | keyspace_hits | 13824645.000 |
| redis | keyspace_misses | 133169.000 |
| redis | net_input_bytes | 1066826484.000 |
| redis | net_output_bytes | 3010716073.000 |
| redis | ops_per_sec | 9695.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 98642456.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
