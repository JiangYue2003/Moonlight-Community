# Feed 压测报告：hybrid / rpc / deep-page-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:36:09+08:00
- 采样时长：1m0.0210792s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 176091 | 176091 | 0 | 0 | 2934.47 | 10.611 | 12.744 | 13.900 | 16.756 | 41.380 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 162 | 0.001 |
| redis | 528435 | 3.001 |
| relation | 176091 | 1.000 |

- Cold compute：176091（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 176091 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 176091 | 0.203 |
| counter | 240 | 0.995 |
| hydrate | 176091 | 0.357 |
| inbox | 176091 | 0.223 |
| merge_dedup | 176091 | 0.015 |
| relation | 176091 | 9.619 |
| route | 176091 | 0.004 |
| total | 176091 | 10.448 |

## Redis 本轮边界增量

- Commands：1473706；input：446864785 bytes；output：2076272618 bytes
- Hits/Misses：11451838/166；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26043；safety epoch：417 -> 417

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.139 |
| client:loadtest | cpu_percent_total | 66.227 |
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
| docker:zg-canal | cpu_percent | 2.780 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.820 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.320 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 158.380 |
| docker:zg-kafka | memory_percent | 8.440 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 54.620 |
| docker:zg-zk | memory_percent | 1.290 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 14597553.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 30.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 12.372 |
| process:counter | cpu_seconds_total | 132.000 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43806720.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.547 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 36962304.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 334.929 |
| process:knowpost | cpu_seconds_total | 4949.453 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 88604672.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 271.134 |
| process:relation | cpu_seconds_total | 4642.453 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 88117248.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.875 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37015552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.953 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35164160.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 118646283.000 |
| redis | connected_clients | 241.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 417.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698728.000 |
| redis | keyspace_hits | 684407283.000 |
| redis | keyspace_misses | 97706.000 |
| redis | net_input_bytes | 27785228911.000 |
| redis | net_output_bytes | 125219759979.000 |
| redis | ops_per_sec | 26043.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22608.000 |
| redis | used_memory_bytes | 107086728.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
