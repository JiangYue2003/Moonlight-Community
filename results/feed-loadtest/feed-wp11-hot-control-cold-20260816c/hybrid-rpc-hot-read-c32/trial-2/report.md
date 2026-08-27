# Feed 压测报告：hybrid / rpc / hot-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:05:58+08:00
- 采样时长：1m0.0192412s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 179890 | 179890 | 0 | 0 | 2997.83 | 10.492 | 12.730 | 13.669 | 16.242 | 45.918 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 79 | 0.000 |
| redis | 539749 | 3.000 |
| relation | 179890 | 1.000 |

- Cold compute：179890（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 179890 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 179890 | 0.202 |
| counter | 12 | 1.761 |
| hydrate | 179890 | 0.298 |
| inbox | 179890 | 0.220 |
| merge_dedup | 179890 | 0.012 |
| relation | 179890 | 9.454 |
| route | 179890 | 0.003 |
| total | 179890 | 10.210 |

## Redis 本轮边界增量

- Commands：1498723；input：342752889 bytes；output：1566663125 bytes
- Hits/Misses：8455137/83；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27202；safety epoch：393 -> 393

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.627 |
| client:loadtest | cpu_percent_total | 74.039 |
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
| docker:zg-canal | cpu_percent | 2.900 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.670 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 6.010 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 211.010 |
| docker:zg-kafka | memory_percent | 8.390 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 57.770 |
| docker:zg-zk | memory_percent | 1.190 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 9551495.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 39.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.834 |
| process:counter | cpu_seconds_total | 17.734 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40534016.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.172 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37806080.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 281.698 |
| process:knowpost | cpu_seconds_total | 404.531 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 78258176.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 288.575 |
| process:relation | cpu_seconds_total | 371.156 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 67579904.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.777 |
| process:search | cpu_seconds_total | 0.219 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37683200.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.125 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35258368.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 76449697.000 |
| redis | connected_clients | 76.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 393.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698905.000 |
| redis | keyspace_hits | 440214081.000 |
| redis | keyspace_misses | 94262.000 |
| redis | net_input_bytes | 17918140777.000 |
| redis | net_output_bytes | 80088902113.000 |
| redis | ops_per_sec | 27202.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 20797.000 |
| redis | used_memory_bytes | 103524760.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
