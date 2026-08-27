# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-distributed-control-fixed-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:54:57+08:00
- 采样时长：10.0078308s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 64346 | 64346 | 0 | 0 | 6432.58 | 4.822 | 6.318 | 6.890 | 8.232 | 13.887 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.037 |
| mysql | 0 | 0.000 |
| redis | 193038 | 3.000 |
| relation | 64346 | 1.000 |

- Cold compute：64346（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 64346 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 64346 | 0.799 |
| counter | 2400 | 1.212 |
| hydrate | 64346 | 0.988 |
| inbox | 64346 | 1.024 |
| merge_dedup | 64346 | 0.004 |
| relation | 64346 | 1.835 |
| route | 64346 | 0.048 |
| total | 64346 | 4.716 |

## Redis 本轮边界增量

- Commands：514046；input：59154543 bytes；output：176491407 bytes
- Hits/Misses：1193834/69726；run hit rate：94.48%
- Evicted/Rejected：0/0；ops/s max：54609；safety epoch：579 -> 579

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.513 |
| client:loadtest | cpu_percent_total | 88.212 |
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
| docker:zg-canal | cpu_percent | 2.550 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.020 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.870 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 76.730 |
| docker:zg-kafka | memory_percent | 8.320 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 0.130 |
| docker:zg-zk | memory_percent | 1.520 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25428495.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 61.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 19.924 |
| process:counter | cpu_seconds_total | 5.438 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 47435776.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.546 |
| process:gateway | cpu_seconds_total | 0.469 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 37330944.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 319.347 |
| process:knowpost | cpu_seconds_total | 40.484 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 66334720.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 290.398 |
| process:relation | cpu_seconds_total | 31.000 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 53178368.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.785 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35418112.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.797 |
| process:user-storage | cpu_seconds_total | 0.047 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 34357248.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 256282842.000 |
| redis | connected_clients | 98.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 579.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 750042.000 |
| redis | keyspace_hits | 1067505380.000 |
| redis | keyspace_misses | 11157720.000 |
| redis | net_input_bytes | 49614894353.000 |
| redis | net_output_bytes | 199611444422.000 |
| redis | ops_per_sec | 54609.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52486.000 |
| redis | used_memory_bytes | 162657176.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
