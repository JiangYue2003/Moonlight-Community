# Feed 压测报告：pull / rpc / steady-read-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T16:59:50+08:00
- 采样时长：30.098003s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 20193 | 20193 | 0 | 0 | 670.94 | 186.522 | 232.156 | 249.033 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 6.620 |
| docker:zg-counter | memory_percent | 0.190 |
| docker:zg-counter | pids | 24.000 |
| docker:zg-gateway | cpu_percent | 3.430 |
| docker:zg-gateway | memory_percent | 0.140 |
| docker:zg-gateway | pids | 24.000 |
| docker:zg-knowpost | cpu_percent | 318.150 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 104.990 |
| docker:zg-relation | memory_percent | 0.330 |
| docker:zg-relation | pids | 32.000 |
| kafka | current_offset_total | 2751.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2751.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 596037.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 69.000 |
| mysql | threads_running | 4.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 13346705.000 |
| redis | connected_clients | 419.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 537645.000 |
| redis | keyspace_hits | 20044130.000 |
| redis | keyspace_misses | 316041.000 |
| redis | net_input_bytes | 1448198656.000 |
| redis | net_output_bytes | 4248571078.000 |
| redis | ops_per_sec | 20397.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 92088888.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
