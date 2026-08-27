# Feed 压测报告：hybrid / gateway / hot-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:57:32+08:00
- 采样时长：1m0.0456771s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 144872 | 144872 | 0 | 0 | 2413.10 | 51.401 | 61.635 | 66.467 | 81.004 | 226.163 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 158 | 0.001 |
| redis | 434774 | 3.001 |
| relation | 144872 | 1.000 |

- Cold compute：144872（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 144872 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 144872 | 0.372 |
| counter | 12 | 2.062 |
| hydrate | 144872 | 0.506 |
| inbox | 144872 | 0.415 |
| merge_dedup | 144872 | 0.013 |
| relation | 144872 | 49.855 |
| route | 144872 | 0.003 |
| total | 144872 | 51.189 |

## Redis 本轮边界增量

- Commands：1215129；input：276617125 bytes；output：1261711256 bytes
- Hits/Misses：6664344/145030；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22012；safety epoch：434 -> 434

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.633 |
| client:loadtest | cpu_percent_total | 58.133 |
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
| docker:zg-canal | cpu_percent | 2.320 |
| docker:zg-canal | memory_percent | 7.860 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 0.610 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 6.980 |
| docker:zg-etcd | memory_percent | 0.310 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 293.430 |
| docker:zg-kafka | memory_percent | 8.470 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 77.140 |
| docker:zg-zk | memory_percent | 1.320 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 17813127.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 75.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 13.134 |
| process:counter | cpu_seconds_total | 212.281 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 42778624.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 247.979 |
| process:gateway | cpu_seconds_total | 1008.656 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 59236352.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 264.520 |
| process:knowpost | cpu_seconds_total | 8027.938 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 106950656.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 257.426 |
| process:relation | cpu_seconds_total | 7436.266 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 91725824.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.545 |
| process:search | cpu_seconds_total | 1.562 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37076992.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 130.733 |
| process:user-storage | cpu_seconds_total | 532.844 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 64835584.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 145639283.000 |
| redis | connected_clients | 314.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 434.000 |
| redis | hit_rate | 0.997 |
| redis | keys | 687658.000 |
| redis | keyspace_hits | 868723719.000 |
| redis | keyspace_misses | 2998455.000 |
| redis | net_input_bytes | 35188076675.000 |
| redis | net_output_bytes | 159385935082.000 |
| redis | ops_per_sec | 22012.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23891.000 |
| redis | used_memory_bytes | 106430552.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
