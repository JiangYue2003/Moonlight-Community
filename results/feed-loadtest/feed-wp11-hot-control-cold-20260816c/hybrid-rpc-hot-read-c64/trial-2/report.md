# Feed 压测报告：hybrid / rpc / hot-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:09:44+08:00
- 采样时长：1m0.0257928s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 183261 | 183261 | 0 | 0 | 3053.70 | 20.745 | 24.655 | 26.184 | 29.971 | 57.719 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 95 | 0.001 |
| redis | 549878 | 3.001 |
| relation | 183261 | 1.000 |

- Cold compute：183261（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 183261 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 183261 | 0.207 |
| counter | 12 | 1.755 |
| hydrate | 183261 | 0.304 |
| inbox | 183261 | 0.234 |
| merge_dedup | 183261 | 0.011 |
| relation | 183261 | 19.723 |
| route | 183261 | 0.003 |
| total | 183261 | 20.503 |

## Redis 本轮边界增量

- Commands：1525908；input：349114581 bytes；output：1596005433 bytes
- Hits/Misses：8613560/97；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27456；safety epoch：396 -> 396

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.684 |
| client:loadtest | cpu_percent_total | 74.942 |
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
| docker:zg-canal | cpu_percent | 0.190 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.600 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.870 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 240.000 |
| docker:zg-kafka | memory_percent | 8.400 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 66.990 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 10199922.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 69.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 12.360 |
| process:counter | cpu_seconds_total | 31.484 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 41201664.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.188 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37081088.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 307.398 |
| process:knowpost | cpu_seconds_total | 968.047 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 93540352.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 275.442 |
| process:relation | cpu_seconds_total | 912.094 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 74588160.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.250 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37085184.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.188 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35192832.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 81854455.000 |
| redis | connected_clients | 135.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 396.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698864.000 |
| redis | keyspace_hits | 470648451.000 |
| redis | keyspace_misses | 94663.000 |
| redis | net_input_bytes | 19152613604.000 |
| redis | net_output_bytes | 85728298352.000 |
| redis | ops_per_sec | 27456.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21023.000 |
| redis | used_memory_bytes | 107015776.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
