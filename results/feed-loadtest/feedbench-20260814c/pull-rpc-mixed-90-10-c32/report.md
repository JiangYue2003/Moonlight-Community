# Feed 压测报告：pull / rpc / mixed-90-10-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:37:22+08:00
- 采样时长：6.040666s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 555.39 | 48.815 | 60.987 | 66.775 |
| publish_total | 30 | 30 | 0 | 0 | 4.97 | 697.414 | 1165.956 | 1167.042 |
| publish_draft | 30 | 30 | 0 | 0 | 4.97 | 6.912 | 16.994 | 18.564 |
| publish_metadata | 30 | 30 | 0 | 0 | 4.97 | 233.879 | 662.331 | 662.891 |
| publish_confirm | 30 | 30 | 0 | 0 | 4.97 | 218.640 | 243.409 | 245.942 |
| publish_commit | 30 | 30 | 0 | 0 | 4.97 | 229.783 | 257.378 | 257.378 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 3.170 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 2.230 |
| docker:zg-gateway | memory_percent | 0.150 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 137.940 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 35.930 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 339044.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 78.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 4995463.000 |
| redis | connected_clients | 420.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 537382.000 |
| redis | keyspace_hits | 7744584.000 |
| redis | keyspace_misses | 73993.000 |
| redis | net_input_bytes | 586884182.000 |
| redis | net_output_bytes | 1633303027.000 |
| redis | ops_per_sec | 6760.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 86732872.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
