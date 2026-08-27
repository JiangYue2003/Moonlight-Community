# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp2-phase0-enabled-20260815`
- 开始时间：2026-08-15T19:05:40+08:00
- 采样时长：1m0.0298524s
- 并发：64
- 重复轮次：2
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P95(ms) | P99(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 137140 | 137140 | 0 | 0 | 2284.94 | 27.592 | 36.108 | 40.566 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 631 | 0.005 |
| redis | 412051 | 3.005 |
| relation | 137140 | 1.000 |

- Cold compute：137140（1.000 / successful read）

| feed stage | calls | mean(ms) |
|---|---:|---:|
| bigv_pipeline | 137140 | 3.703 |
| counter | 240 | 5.235 |
| hydrate | 137140 | 5.295 |
| inbox | 137140 | 4.222 |
| merge_dedup | 137140 | 0.025 |
| relation | 137140 | 12.880 |
| route | 137140 | 0.019 |
| total | 137140 | 26.191 |

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| docker-state:zg-counter | healthy | 1.000 |
| docker-state:zg-counter | restart_count | 0.000 |
| docker-state:zg-counter | running | 1.000 |
| docker-state:zg-gateway | healthy | 1.000 |
| docker-state:zg-gateway | restart_count | 0.000 |
| docker-state:zg-gateway | running | 1.000 |
| docker-state:zg-knowpost | healthy | 1.000 |
| docker-state:zg-knowpost | restart_count | 0.000 |
| docker-state:zg-knowpost | running | 1.000 |
| docker-state:zg-relation | healthy | 1.000 |
| docker-state:zg-relation | restart_count | 0.000 |
| docker-state:zg-relation | running | 1.000 |
| docker:zg-counter | cpu_percent | 5.260 |
| docker:zg-counter | memory_percent | 0.110 |
| docker:zg-counter | pids | 18.000 |
| docker:zg-gateway | cpu_percent | 3.380 |
| docker:zg-gateway | memory_percent | 0.110 |
| docker:zg-gateway | pids | 19.000 |
| docker:zg-knowpost | cpu_percent | 320.880 |
| docker:zg-knowpost | memory_percent | 0.370 |
| docker:zg-knowpost | pids | 26.000 |
| docker:zg-relation | cpu_percent | 246.930 |
| docker:zg-relation | memory_percent | 0.310 |
| docker:zg-relation | pids | 28.000 |
| kafka | current_offset_total | 2962.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 2962.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 1302134.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 101.000 |
| mysql | threads_running | 6.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 10841613.000 |
| redis | connected_clients | 270.000 |
| redis | evicted_keys | 0.000 |
| redis | hit_rate | 0.893 |
| redis | keys | 519485.000 |
| redis | keyspace_hits | 54307664.000 |
| redis | keyspace_misses | 6481545.000 |
| redis | net_input_bytes | 2466950075.000 |
| redis | net_output_bytes | 11214062699.000 |
| redis | ops_per_sec | 18722.000 |
| redis | rejected_connections | 0.000 |
| redis | used_memory_bytes | 87987400.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
