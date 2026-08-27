# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:19:49+08:00
- 采样时长：1m0.0194075s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 176016 | 176016 | 0 | 0 | 2933.24 | 10.701 | 12.630 | 13.472 | 15.590 | 48.224 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 109 | 0.001 |
| redis | 528157 | 3.001 |
| relation | 176016 | 1.000 |

- Cold compute：176016（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 176016 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 176016 | 0.186 |
| counter | 240 | 0.875 |
| hydrate | 176016 | 0.288 |
| inbox | 176016 | 0.204 |
| merge_dedup | 176016 | 0.012 |
| relation | 176016 | 9.719 |
| route | 176016 | 0.004 |
| total | 176016 | 10.436 |

## Redis 本轮边界增量

- Commands：1473397；input：335622799 bytes；output：1533099938 bytes
- Hits/Misses：8278733/109；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26140；safety epoch：404 -> 404

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.618 |
| client:loadtest | cpu_percent_total | 73.882 |
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
| docker:zg-canal | cpu_percent | 2.360 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.280 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.860 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 179.700 |
| docker:zg-kafka | memory_percent | 8.380 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 62.760 |
| docker:zg-zk | memory_percent | 1.170 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 11919750.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 39.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 12.520 |
| process:counter | cpu_seconds_total | 70.359 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 42749952.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.312 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37392384.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 296.254 |
| process:knowpost | cpu_seconds_total | 2473.734 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 85106688.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 289.546 |
| process:relation | cpu_seconds_total | 2348.109 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 87728128.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.591 |
| process:search | cpu_seconds_total | 0.516 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37040128.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.766 |
| process:user-storage | cpu_seconds_total | 0.469 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35180544.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 96198599.000 |
| redis | connected_clients | 237.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 404.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698791.000 |
| redis | keyspace_hits | 551357036.000 |
| redis | keyspace_misses | 95761.000 |
| redis | net_input_bytes | 22426628180.000 |
| redis | net_output_bytes | 100682398237.000 |
| redis | ops_per_sec | 26140.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21628.000 |
| redis | used_memory_bytes | 107089992.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
