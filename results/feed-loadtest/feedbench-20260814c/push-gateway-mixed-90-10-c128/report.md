# Feed 压测报告：push / gateway / mixed-90-10-c128

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:31:26+08:00
- 采样时长：3.7768268s
- 并发：128
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 270 | 270 | 0 | 0 | 898.47 | 101.109 | 165.092 | 182.971 |
| publish_total | 30 | 30 | 0 | 0 | 7.94 | 1309.984 | 1653.588 | 1654.107 |
| publish_draft | 30 | 30 | 0 | 0 | 7.94 | 10.185 | 104.240 | 106.097 |
| publish_metadata | 30 | 30 | 0 | 0 | 7.94 | 408.981 | 765.471 | 768.482 |
| publish_confirm | 30 | 30 | 0 | 0 | 7.94 | 398.813 | 451.196 | 451.217 |
| publish_commit | 30 | 30 | 0 | 0 | 7.94 | 400.841 | 500.214 | 500.844 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 2.890 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 12.830 |
| docker:zg-gateway | memory_percent | 0.190 |
| docker:zg-gateway | pids | 23.000 |
| docker:zg-knowpost | cpu_percent | 100.900 |
| docker:zg-knowpost | memory_percent | 0.420 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 7.520 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 303684.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 95.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3406600.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 538263.000 |
| redis | keyspace_hits | 5605604.000 |
| redis | keyspace_misses | 38515.000 |
| redis | net_input_bytes | 416821736.000 |
| redis | net_output_bytes | 1067703680.000 |
| redis | ops_per_sec | 14743.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 123414688.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
