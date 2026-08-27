# Feed 压测报告：pull / rpc / mixed-90-10-c1

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:36:36+08:00
- 采样时长：20.2447526s
- 并发：1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 146.99 | 6.672 | 8.410 | 9.783 |
| publish_total | 30 | 30 | 0 | 0 | 1.48 | 660.756 | 800.253 | 807.883 |
| publish_draft | 30 | 30 | 0 | 0 | 1.48 | 5.302 | 6.395 | 8.040 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.48 | 215.071 | 257.670 | 262.534 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.48 | 213.829 | 321.754 | 338.041 |
| publish_commit | 30 | 30 | 0 | 0 | 1.48 | 223.501 | 244.581 | 266.500 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 4.240 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.530 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 78.440 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 24.200 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 337649.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 90.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4907751.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537386.000 |
| redis | keyspace_hits | 7709385.000 |
| redis | keyspace_misses | 71957.000 |
| redis | net_input_bytes | 579274626.000 |
| redis | net_output_bytes | 1620391735.000 |
| redis | ops_per_sec | 4101.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86053208.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
