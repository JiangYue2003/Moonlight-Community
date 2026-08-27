# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:33:38+08:00
- 采样时长：1m0.1469937s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 167880 | 167880 | 0 | 0 | 2791.79 | 76.070 | 157.031 | 190.457 | 266.790 | 686.097 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 96 | 0.001 |
| redis | 503736 | 3.001 |
| relation | 167880 | 1.000 |

- Cold compute：167880（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 167880 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 167880 | 0.168 |
| counter | 240 | 0.767 |
| hydrate | 167880 | 0.266 |
| inbox | 167880 | 0.183 |
| merge_dedup | 167880 | 0.012 |
| relation | 167880 | 90.444 |
| route | 167880 | 0.004 |
| total | 167880 | 91.098 |

## Redis 本轮边界增量

- Commands：1408869；input：320347145 bytes；output：1462312186 bytes
- Hits/Misses：7896354/97；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：25352；safety epoch：415 -> 415

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.449 |
| client:loadtest | cpu_percent_total | 71.180 |
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
| docker:zg-es | cpu_percent | 3.600 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.770 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 160.740 |
| docker:zg-kafka | memory_percent | 8.410 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 48.640 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 14190754.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 73.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 12.370 |
| process:counter | cpu_seconds_total | 121.812 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43503616.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.516 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 36950016.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 289.949 |
| process:knowpost | cpu_seconds_total | 4550.172 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 108556288.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 278.594 |
| process:relation | cpu_seconds_total | 4300.453 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 100937728.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.770 |
| process:search | cpu_seconds_total | 0.812 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37015552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.558 |
| process:user-storage | cpu_seconds_total | 0.891 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35094528.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 115234401.000 |
| redis | connected_clients | 241.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 415.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698723.000 |
| redis | keyspace_hits | 657995640.000 |
| redis | keyspace_misses | 97268.000 |
| redis | net_input_bytes | 26753696193.000 |
| redis | net_output_bytes | 120430989471.000 |
| redis | ops_per_sec | 25352.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22457.000 |
| redis | used_memory_bytes | 106751072.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
