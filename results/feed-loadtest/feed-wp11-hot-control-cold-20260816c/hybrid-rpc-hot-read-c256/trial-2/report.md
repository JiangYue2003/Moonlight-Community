# Feed 压测报告：hybrid / rpc / hot-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:17:16+08:00
- 采样时长：1m0.0753872s
- 并发：256
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 182614 | 182614 | 0 | 0 | 3040.46 | 71.659 | 138.750 | 168.001 | 237.014 | 474.918 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 109 | 0.001 |
| redis | 547951 | 3.001 |
| relation | 182614 | 1.000 |

- Cold compute：182614（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 182614 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 182614 | 0.206 |
| counter | 12 | 1.899 |
| hydrate | 182614 | 0.299 |
| inbox | 182614 | 0.226 |
| merge_dedup | 182614 | 0.011 |
| relation | 182614 | 82.939 |
| route | 182614 | 0.003 |
| total | 182614 | 83.705 |

## Redis 本轮边界增量

- Commands：1520485；input：347882007 bytes；output：1590361758 bytes
- Hits/Misses：8583135/113；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27208；safety epoch：402 -> 402

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.735 |
| client:loadtest | cpu_percent_total | 75.764 |
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
| docker:zg-canal | cpu_percent | 2.340 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.690 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 6.310 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 216.280 |
| docker:zg-kafka | memory_percent | 8.420 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 52.360 |
| docker:zg-zk | memory_percent | 1.050 |
| docker:zg-zk | pids | 87.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 11499794.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 74.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 12.179 |
| process:counter | cpu_seconds_total | 60.391 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40882176.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 0.312 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 38260736.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 306.600 |
| process:knowpost | cpu_seconds_total | 2092.703 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 111112192.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 285.802 |
| process:relation | cpu_seconds_total | 1987.203 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 98836480.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.484 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37150720.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.375 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35512320.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 92685662.000 |
| redis | connected_clients | 218.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 402.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698805.000 |
| redis | keyspace_hits | 531642479.000 |
| redis | keyspace_misses | 95483.000 |
| redis | net_input_bytes | 21626600509.000 |
| redis | net_output_bytes | 97030279202.000 |
| redis | ops_per_sec | 27208.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21475.000 |
| redis | used_memory_bytes | 106285088.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
