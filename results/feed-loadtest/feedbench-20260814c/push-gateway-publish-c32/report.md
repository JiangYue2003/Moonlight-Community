# Feed 压测报告：push / gateway / publish-c32

- Run ID：`feedbench-20260814c`
- 开始时间：2026-08-14T15:29:49+08:00
- 采样时长：12.7698504s
- 并发：32
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| publish_total | 128 | 128 | 0 | 0 | 10.02 | 3186.429 | 3439.447 | 3448.266 |
| publish_draft | 128 | 128 | 0 | 0 | 10.02 | 10.329 | 47.745 | 48.275 |
| publish_metadata | 128 | 128 | 0 | 0 | 10.02 | 1040.614 | 1079.515 | 1091.590 |
| publish_confirm | 128 | 128 | 0 | 0 | 10.02 | 1009.323 | 1169.128 | 1175.351 |
| publish_commit | 128 | 128 | 0 | 0 | 10.02 | 1079.611 | 1226.513 | 1230.441 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker:zg-counter | cpu_percent | 7.280 |
| docker:zg-counter | memory_percent | 0.120 |
| docker:zg-counter | pids | 20.000 |
| docker:zg-gateway | cpu_percent | 8.930 |
| docker:zg-gateway | memory_percent | 0.130 |
| docker:zg-gateway | pids | 20.000 |
| docker:zg-knowpost | cpu_percent | 141.400 |
| docker:zg-knowpost | memory_percent | 0.430 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 13.300 |
| docker:zg-relation | memory_percent | 0.300 |
| docker:zg-relation | pids | 25.000 |
| kafka | lag_max | 32.000 |
| kafka | lag_total | 32.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 279348.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 129.000 |
| mysql | threads_running | 3.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 3080481.000 |
| redis | connected_clients | 449.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 537563.000 |
| redis | keyspace_hits | 4724607.000 |
| redis | keyspace_misses | 35195.000 |
| redis | net_input_bytes | 365403467.000 |
| redis | net_output_bytes | 901820356.000 |
| redis | ops_per_sec | 27381.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 120132704.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
