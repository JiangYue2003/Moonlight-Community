# Feed 压测报告：hybrid / gateway / distributed-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T20:16:00+08:00
- 采样时长：1m0.0800345s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 135590 | 135590 | 0 | 0 | 2257.21 | 93.139 | 191.003 | 238.347 | 375.343 | 1313.465 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 113 | 0.001 |
| redis | 406883 | 3.001 |
| relation | 135590 | 1.000 |

- Cold compute：135590（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 135590 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 135590 | 0.320 |
| counter | 240 | 1.020 |
| hydrate | 135590 | 0.441 |
| inbox | 135590 | 0.354 |
| merge_dedup | 135590 | 0.013 |
| relation | 135590 | 110.099 |
| route | 135590 | 0.005 |
| total | 135590 | 111.257 |

## Redis 本轮边界增量

- Commands：1143397；input：259077634 bytes；output：1181047482 bytes
- Hits/Misses：6243113/135707；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22251；safety epoch：448 -> 448

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.488 |
| client:loadtest | cpu_percent_total | 55.811 |
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
| docker:zg-canal | cpu_percent | 3.130 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.800 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 7.230 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 305.330 |
| docker:zg-kafka | memory_percent | 8.540 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 86.060 |
| docker:zg-zk | memory_percent | 1.410 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20153099.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 73.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.955 |
| process:counter | cpu_seconds_total | 710.828 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44810240.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 250.684 |
| process:gateway | cpu_seconds_total | 3001.844 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 70647808.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 274.379 |
| process:knowpost | cpu_seconds_total | 10486.031 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 118317056.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 272.307 |
| process:relation | cpu_seconds_total | 9551.156 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 103059456.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 6.078 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37298176.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 130.060 |
| process:user-storage | cpu_seconds_total | 1575.719 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 56373248.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 174827180.000 |
| redis | connected_clients | 153.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 448.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676622.000 |
| redis | keyspace_hits | 976201012.000 |
| redis | keyspace_misses | 5336249.000 |
| redis | net_input_bytes | 40320109659.000 |
| redis | net_output_bytes | 179906471328.000 |
| redis | ops_per_sec | 22251.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35799.000 |
| redis | used_memory_bytes | 100545296.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
