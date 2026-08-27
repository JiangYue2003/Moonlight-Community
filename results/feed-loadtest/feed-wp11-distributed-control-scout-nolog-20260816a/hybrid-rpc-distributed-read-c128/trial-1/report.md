# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp11-distributed-control-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:32:21+08:00
- 采样时长：10.0185397s
- 并发：128
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 64871 | 64871 | 0 | 0 | 6478.17 | 17.436 | 29.875 | 33.086 | 41.043 | 85.010 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 2400 | 0.037 |
| mysql | 119 | 0.002 |
| redis | 194732 | 3.002 |
| relation | 64871 | 1.000 |

- Cold compute：64871（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 64871 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 64871 | 4.564 |
| counter | 2400 | 5.066 |
| hydrate | 64871 | 4.510 |
| inbox | 64871 | 4.642 |
| merge_dedup | 64871 | 0.004 |
| relation | 64871 | 5.511 |
| route | 64871 | 0.193 |
| total | 64871 | 19.443 |

## Redis 本轮边界增量

- Commands：509771；input：59059996 bytes；output：177734101 bytes
- Hits/Misses：1203246/70414；run hit rate：94.47%
- Evicted/Rejected：0/0；ops/s max：57815；safety epoch：527 -> 527

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.761 |
| client:loadtest | cpu_percent_total | 92.173 |
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
| docker:zg-canal | cpu_percent | 2.540 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.880 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.800 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 281.690 |
| docker:zg-kafka | memory_percent | 8.690 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 36.870 |
| docker:zg-zk | memory_percent | 1.620 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23094235.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 76.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 17.810 |
| process:counter | cpu_seconds_total | 16.344 |
| process:counter | pid | 8320.000 |
| process:counter | process_start_ms | 1786894081426.000 |
| process:counter | rss_bytes | 53071872.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.016 |
| process:gateway | pid | 9852.000 |
| process:gateway | process_start_ms | 1786894099200.000 |
| process:gateway | rss_bytes | 36909056.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 366.940 |
| process:knowpost | cpu_seconds_total | 193.688 |
| process:knowpost | pid | 7360.000 |
| process:knowpost | process_start_ms | 1786894091315.000 |
| process:knowpost | rss_bytes | 92766208.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 291.435 |
| process:relation | cpu_seconds_total | 146.672 |
| process:relation | pid | 26448.000 |
| process:relation | process_start_ms | 1786894085233.000 |
| process:relation | rss_bytes | 63614976.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.141 |
| process:search | pid | 19932.000 |
| process:search | process_start_ms | 1786894095566.000 |
| process:search | rss_bytes | 35852288.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.125 |
| process:user-storage | pid | 7096.000 |
| process:user-storage | process_start_ms | 1786894077784.000 |
| process:user-storage | rss_bytes | 34758656.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 203537990.000 |
| redis | connected_clients | 257.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 527.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 677095.000 |
| redis | keyspace_hits | 1000131062.000 |
| redis | keyspace_misses | 5774630.000 |
| redis | net_input_bytes | 42565316480.000 |
| redis | net_output_bytes | 188163939165.000 |
| redis | ops_per_sec | 57815.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47530.000 |
| redis | used_memory_bytes | 108508552.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
