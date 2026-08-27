# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T20:04:41+08:00
- 采样时长：1m0.0195498s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 150049 | 150049 | 0 | 0 | 2500.42 | 12.480 | 15.062 | 16.200 | 19.384 | 36.937 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 98 | 0.001 |
| redis | 450245 | 3.001 |
| relation | 150049 | 1.000 |

- Cold compute：150049（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 150049 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 150049 | 0.214 |
| counter | 240 | 0.819 |
| hydrate | 150049 | 0.321 |
| inbox | 150049 | 0.235 |
| merge_dedup | 150049 | 0.013 |
| relation | 150049 | 10.317 |
| route | 150049 | 0.004 |
| total | 150049 | 11.130 |

## Redis 本轮边界增量

- Commands：1258998；input：286282349 bytes；output：1306858778 bytes
- Hits/Misses：6908246/150147；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22746；safety epoch：439 -> 439

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.596 |
| client:loadtest | cpu_percent_total | 57.533 |
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
| docker:zg-canal | cpu_percent | 5.270 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.600 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 7.850 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 280.150 |
| docker:zg-kafka | memory_percent | 8.500 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 76.920 |
| docker:zg-zk | memory_percent | 1.400 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 18572436.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 86.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.038 |
| process:counter | cpu_seconds_total | 670.219 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44236800.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 253.691 |
| process:gateway | cpu_seconds_total | 1665.234 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 53743616.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 256.400 |
| process:knowpost | cpu_seconds_total | 8974.047 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 99540992.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 240.581 |
| process:relation | cpu_seconds_total | 8130.125 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 87334912.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 5.812 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37281792.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 122.843 |
| process:user-storage | cpu_seconds_total | 873.109 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 56692736.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 161545444.000 |
| redis | connected_clients | 197.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 439.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 676488.000 |
| redis | keyspace_hits | 903564486.000 |
| redis | keyspace_misses | 3757354.000 |
| redis | net_input_bytes | 37306975791.000 |
| redis | net_output_bytes | 166165002908.000 |
| redis | ops_per_sec | 22746.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35120.000 |
| redis | used_memory_bytes | 99990384.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
