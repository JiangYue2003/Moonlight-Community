# Feed 压测报告：hybrid / gateway / hot-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:50:00+08:00
- 采样时长：1m0.0212765s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 142782 | 142782 | 0 | 0 | 2379.29 | 12.811 | 15.852 | 17.345 | 22.663 | 128.659 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 108 | 0.001 |
| redis | 428454 | 3.001 |
| relation | 142782 | 1.000 |

- Cold compute：142782（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 142782 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 142782 | 0.242 |
| counter | 12 | 1.888 |
| hydrate | 142782 | 0.357 |
| inbox | 142782 | 0.265 |
| merge_dedup | 142782 | 0.013 |
| relation | 142782 | 10.754 |
| route | 142782 | 0.003 |
| total | 142782 | 11.660 |

## Redis 本轮边界增量

- Commands：1200763；input：272844050 bytes；output：1243578866 bytes
- Hits/Misses：6568242/142902；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22270；safety epoch：428 -> 428

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.456 |
| client:loadtest | cpu_percent_total | 55.293 |
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
| docker:zg-canal | cpu_percent | 2.990 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.540 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 10.440 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 294.970 |
| docker:zg-kafka | memory_percent | 8.480 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 79.480 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 16787061.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 52.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 14.682 |
| process:counter | cpu_seconds_total | 184.594 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43970560.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 225.130 |
| process:gateway | cpu_seconds_total | 139.750 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 51146752.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 279.184 |
| process:knowpost | cpu_seconds_total | 7048.812 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 92798976.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 233.630 |
| process:relation | cpu_seconds_total | 6524.516 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 86179840.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.109 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37064704.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 125.330 |
| process:user-storage | cpu_seconds_total | 75.844 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 55197696.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 137015919.000 |
| redis | connected_clients | 314.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 428.000 |
| redis | hit_rate | 0.998 |
| redis | keys | 695660.000 |
| redis | keyspace_hits | 821613639.000 |
| redis | keyspace_misses | 1973441.000 |
| redis | net_input_bytes | 33230345846.000 |
| redis | net_output_bytes | 150466315528.000 |
| redis | ops_per_sec | 22270.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23439.000 |
| redis | used_memory_bytes | 107620768.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
