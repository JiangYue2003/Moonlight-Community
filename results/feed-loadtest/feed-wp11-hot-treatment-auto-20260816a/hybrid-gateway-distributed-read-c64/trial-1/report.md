# Feed 压测报告：hybrid / gateway / distributed-read-c64

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T22:07:49+08:00
- 采样时长：1m0.0284794s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 307981 | 307981 | 0 | 0 | 5132.49 | 12.275 | 15.804 | 19.452 | 26.235 | 68.150 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2402 | 0.008 |
| counter | 177 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 2389 | 0.008 |
| relation | 177 | 0.001 |

- Cold compute：362（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 306665 |
| l1_stale | 0 |
| l2_fresh | 1316 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 362 | 0.213 |
| counter | 177 | 0.849 |
| hydrate | 362 | 0.284 |
| inbox | 362 | 0.213 |
| merge_dedup | 362 | 0.017 |
| relation | 177 | 1.860 |
| route | 362 | 1.337 |
| total | 307981 | 0.008 |

## Redis 本轮边界增量

- Commands：62781；input：5570874 bytes；output：6934247 bytes
- Hits/Misses：23710/177；run hit rate：99.26%
- Evicted/Rejected：0/0；ops/s max：1429；safety epoch：513 -> 513

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.870 |
| client:loadtest | cpu_percent_total | 109.922 |
| client:loadtest | logical_cpus | 16.000 |
| docker-state:zg-canal | health_configured | 1.000 |
| docker-state:zg-canal | healthy | 1.000 |
| docker-state:zg-canal | restart_count | 0.000 |
| docker-state:zg-canal | running | 1.000 |
| docker-state:zg-es | health_configured | 1.000 |
| docker-state:zg-es | healthy | 1.000 |
| docker-state:zg-es | restart_count | 0.000 |
| docker-state:zg-es | running | 1.000 |
| docker-state:zg-etcd | health_configured | 1.000 |
| docker-state:zg-etcd | healthy | 1.000 |
| docker-state:zg-etcd | restart_count | 0.000 |
| docker-state:zg-etcd | running | 1.000 |
| docker-state:zg-kafka | health_configured | 1.000 |
| docker-state:zg-kafka | healthy | 1.000 |
| docker-state:zg-kafka | restart_count | 0.000 |
| docker-state:zg-kafka | running | 1.000 |
| docker-state:zg-zk | health_configured | 1.000 |
| docker-state:zg-zk | healthy | 1.000 |
| docker-state:zg-zk | restart_count | 0.000 |
| docker-state:zg-zk | running | 1.000 |
| docker:zg-canal | cpu_percent | 2.680 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 15.370 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.340 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 18.000 |
| docker:zg-kafka | cpu_percent | 235.950 |
| docker:zg-kafka | memory_percent | 8.640 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 65.000 |
| docker:zg-zk | memory_percent | 1.290 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22732602.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.821 |
| process:counter | cpu_seconds_total | 223.750 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 44773376.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 422.190 |
| process:gateway | cpu_seconds_total | 3642.734 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 56029184.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 210.611 |
| process:knowpost | cpu_seconds_total | 8950.844 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 118788096.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.029 |
| process:relation | cpu_seconds_total | 10.641 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 48185344.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 2.000 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36986880.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 215.917 |
| process:user-storage | cpu_seconds_total | 1829.344 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 79544320.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 196477913.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 513.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677359.000 |
| redis | keyspace_hits | 993287698.000 |
| redis | keyspace_misses | 5382015.000 |
| redis | net_input_bytes | 41932918527.000 |
| redis | net_output_bytes | 187061348864.000 |
| redis | ops_per_sec | 1429.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 42508.000 |
| redis | used_memory_bytes | 100686464.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
