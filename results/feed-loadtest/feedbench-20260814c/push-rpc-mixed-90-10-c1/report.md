# Feed 压测报告：push / rpc / mixed-90-10-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:26:38+08:00
- 采样时长：19.2362359s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 216.95 | 4.687 | 5.578 | 6.462 |
| publish_total | 30 | 30 | 0 | 0 | 1.56 | 640.312 | 670.381 | 710.588 |
| publish_draft | 30 | 30 | 0 | 0 | 1.56 | 5.338 | 6.452 | 7.646 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.56 | 209.641 | 230.849 | 230.974 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.56 | 207.516 | 227.320 | 236.211 |
| publish_commit | 30 | 30 | 0 | 0 | 1.56 | 214.442 | 230.223 | 235.758 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.590 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.650 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 38.280 |
| docker:zg-knowpost | memory_percent | 0.460 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 21.620 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 2.000 |
| kafka | lag_total | 2.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 267902.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 131.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1861523.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 536457.000 |
| redis | keyspace_hits | 4631665.000 |
| redis | keyspace_misses | 17236.000 |
| redis | net_input_bytes | 272461309.000 |
| redis | net_output_bytes | 869113746.000 |
| redis | ops_per_sec | 6245.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 89123664.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
