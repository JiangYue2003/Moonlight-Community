# Feed 压测报告：push / rpc / mixed-90-10-c8

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:27:00+08:00
- 采样时长：19.2072968s
- 并发：8
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 1019.55 | 6.404 | 9.234 | 11.186 |
| publish_total | 30 | 30 | 0 | 0 | 1.56 | 638.921 | 656.396 | 741.104 |
| publish_draft | 30 | 30 | 0 | 0 | 1.56 | 5.448 | 7.031 | 8.068 |
| publish_metadata | 30 | 30 | 0 | 0 | 1.56 | 210.999 | 220.639 | 306.538 |
| publish_confirm | 30 | 30 | 0 | 0 | 1.56 | 207.963 | 214.804 | 215.561 |
| publish_commit | 30 | 30 | 0 | 0 | 1.56 | 214.028 | 222.876 | 226.186 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.700 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 1.820 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 20.080 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 2.300 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 268631.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 131.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 1932194.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 536546.000 |
| redis | keyspace_hits | 4643086.000 |
| redis | keyspace_misses | 18001.000 |
| redis | net_input_bytes | 277946351.000 |
| redis | net_output_bytes | 872093943.000 |
| redis | ops_per_sec | 4180.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 90045696.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
