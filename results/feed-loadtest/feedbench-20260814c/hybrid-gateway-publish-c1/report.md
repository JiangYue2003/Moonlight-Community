# Feed 压测报告：hybrid / gateway / publish-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:49:23+08:00
- 采样时长：21.2220612s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 30 | 30 | 0 | 0 | 1.41 | 699.413 | 753.530 | 911.580 |
| publish_draft | 30 | 30 | 0 | 0 | 1.41 | 6.406 | 9.033 | 9.197 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.41 | 229.699 | 246.254 | 249.831 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.41 | 227.436 | 251.133 | 439.945 |
| publish_commit | 30 | 30 | 0 | 0 | 1.41 | 233.385 | 244.999 | 245.879 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.090 |
| docker:zg-counter | memory_percent | 0.210 |
| docker:zg-counter | pids | 23.000 |
| docker:zg-gateway | cpu_percent | 2.440 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 19.660 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 27.000 |
| docker:zg-relation | cpu_percent | 2.350 |
| docker:zg-relation | memory_percent | 0.290 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 1.000 |
| kafka | lag_total | 1.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 423993.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 119.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9055726.000 |
| redis | connected_clients | 514.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 539593.000 |
| redis | keyspace_hits | 12814484.000 |
| redis | keyspace_misses | 120922.000 |
| redis | net_input_bytes | 951114456.000 |
| redis | net_output_bytes | 2767671938.000 |
| redis | ops_per_sec | 2372.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 91135376.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
