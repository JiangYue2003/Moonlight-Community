# Feed 压测报告：hybrid / rpc / deep-page-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:41:10+08:00
- 采样时长：1m0.0282157s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 173375 | 173375 | 0 | 0 | 2888.83 | 21.710 | 25.585 | 27.454 | 33.073 | 64.428 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 158 | 0.001 |
| redis | 520283 | 3.001 |
| relation | 173375 | 1.000 |

- Cold compute：173375（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 173375 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 173375 | 0.211 |
| counter | 240 | 0.725 |
| hydrate | 173375 | 0.365 |
| inbox | 173375 | 0.230 |
| merge_dedup | 173375 | 0.015 |
| relation | 173375 | 20.840 |
| route | 173375 | 0.004 |
| total | 173375 | 21.693 |

## Redis 本轮边界增量

- Commands：1452194；input：440055154 bytes；output：2044103509 bytes
- Hits/Misses：11101930/173535；run hit rate：98.46%
- Evicted/Rejected：0/0；ops/s max：26130；safety epoch：421 -> 421

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.083 |
| client:loadtest | cpu_percent_total | 65.334 |
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
| docker:zg-canal | cpu_percent | 2.440 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.030 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.800 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 179.130 |
| docker:zg-kafka | memory_percent | 8.440 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 60.690 |
| docker:zg-zk | memory_percent | 1.040 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 15404110.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 13.153 |
| process:counter | cpu_seconds_total | 150.438 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43900928.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.562 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 36990976.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 316.458 |
| process:knowpost | cpu_seconds_total | 5725.875 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 95891456.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 295.152 |
| process:relation | cpu_seconds_total | 5330.078 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 90488832.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.969 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37056512.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 1.125 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35217408.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 125411652.000 |
| redis | connected_clients | 241.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 421.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 698460.000 |
| redis | keyspace_hits | 736265182.000 |
| redis | keyspace_misses | 592016.000 |
| redis | net_input_bytes | 29830014422.000 |
| redis | net_output_bytes | 134711246120.000 |
| redis | ops_per_sec | 26130.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22909.000 |
| redis | used_memory_bytes | 106636216.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
