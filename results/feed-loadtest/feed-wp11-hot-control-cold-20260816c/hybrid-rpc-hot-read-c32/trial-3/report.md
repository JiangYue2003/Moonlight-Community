# Feed 压测报告：hybrid / rpc / hot-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:07:12+08:00
- 采样时长：1m0.019546s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 181756 | 181756 | 0 | 0 | 3028.91 | 10.435 | 12.526 | 13.401 | 15.415 | 25.538 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 89 | 0.000 |
| redis | 545357 | 3.000 |
| relation | 181756 | 1.000 |

- Cold compute：181756（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 181756 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 181756 | 0.203 |
| counter | 12 | 1.884 |
| hydrate | 181756 | 0.299 |
| inbox | 181756 | 0.224 |
| merge_dedup | 181756 | 0.011 |
| relation | 181756 | 9.345 |
| route | 181756 | 0.003 |
| total | 181756 | 10.106 |

## Redis 本轮边界增量

- Commands：1513613；input：346265522 bytes；output：1582896573 bytes
- Hits/Misses：8542817/105；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27927；safety epoch：394 -> 394

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.712 |
| client:loadtest | cpu_percent_total | 75.392 |
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
| docker:zg-canal | cpu_percent | 0.240 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.870 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 6.920 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 213.300 |
| docker:zg-kafka | memory_percent | 8.410 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 61.470 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 9766573.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 38.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 12.357 |
| process:counter | cpu_seconds_total | 22.297 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40710144.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 0.188 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37916672.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 297.117 |
| process:knowpost | cpu_seconds_total | 592.750 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 79585280.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 280.641 |
| process:relation | cpu_seconds_total | 549.094 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 67043328.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.219 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37695488.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.141 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35151872.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 78242641.000 |
| redis | connected_clients | 76.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 394.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698889.000 |
| redis | keyspace_hits | 450310602.000 |
| redis | keyspace_misses | 94410.000 |
| redis | net_input_bytes | 18327672623.000 |
| redis | net_output_bytes | 81959754909.000 |
| redis | ops_per_sec | 27927.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 20871.000 |
| redis | used_memory_bytes | 103441104.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
