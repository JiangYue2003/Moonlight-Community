# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T20:05:56+08:00
- 采样时长：1m0.0188322s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 151879 | 151879 | 0 | 0 | 2530.96 | 12.396 | 14.920 | 15.972 | 18.624 | 30.157 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 113 | 0.001 |
| redis | 455750 | 3.001 |
| relation | 151879 | 1.000 |

- Cold compute：151879（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 151879 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 151879 | 0.239 |
| counter | 240 | 0.786 |
| hydrate | 151879 | 0.353 |
| inbox | 151879 | 0.263 |
| merge_dedup | 151879 | 0.012 |
| relation | 151879 | 10.092 |
| route | 151879 | 0.004 |
| total | 151879 | 10.988 |

## Redis 本轮边界增量

- Commands：1273645；input：289729274 bytes；output：1322778354 bytes
- Hits/Misses：6992411/151992；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22690；safety epoch：440 -> 440

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.646 |
| client:loadtest | cpu_percent_total | 58.341 |
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
| docker:zg-canal | cpu_percent | 0.860 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.970 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 7.820 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 343.050 |
| docker:zg-kafka | memory_percent | 8.510 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 78.550 |
| docker:zg-zk | memory_percent | 1.440 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 18752335.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 29.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.821 |
| process:counter | cpu_seconds_total | 674.844 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44240896.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 241.326 |
| process:gateway | cpu_seconds_total | 1814.219 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 53891072.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 264.444 |
| process:knowpost | cpu_seconds_total | 9141.656 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 97288192.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 258.622 |
| process:relation | cpu_seconds_total | 8289.609 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 86528000.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.769 |
| process:search | cpu_seconds_total | 5.828 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37281792.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 143.755 |
| process:user-storage | cpu_seconds_total | 953.516 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 57032704.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 163055377.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 440.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 676488.000 |
| redis | keyspace_hits | 911835043.000 |
| redis | keyspace_misses | 3937151.000 |
| redis | net_input_bytes | 37649919098.000 |
| redis | net_output_bytes | 167729644161.000 |
| redis | ops_per_sec | 22690.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35195.000 |
| redis | used_memory_bytes | 99943896.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
