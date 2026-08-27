# Feed 压测报告：hybrid / rpc / deep-page-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:39:55+08:00
- 采样时长：1m0.0264143s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 169988 | 169988 | 0 | 0 | 2832.45 | 21.327 | 25.722 | 27.965 | 84.272 | 118.788 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 163 | 0.001 |
| redis | 510127 | 3.001 |
| relation | 169988 | 1.000 |

- Cold compute：169988（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 169988 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 169988 | 0.200 |
| counter | 240 | 0.867 |
| hydrate | 169988 | 0.354 |
| inbox | 169988 | 0.222 |
| merge_dedup | 169988 | 0.015 |
| relation | 169988 | 21.274 |
| route | 169988 | 0.004 |
| total | 169988 | 22.097 |

## Redis 本轮边界增量

- Commands：1425181；input：431551392 bytes；output：2004196595 bytes
- Hits/Misses：10885147/170163；run hit rate：98.46%
- Evicted/Rejected：0/0；ops/s max：25566；safety epoch：420 -> 420

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.056 |
| client:loadtest | cpu_percent_total | 64.893 |
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
| docker:zg-canal | cpu_percent | 0.170 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.770 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.390 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 190.090 |
| docker:zg-kafka | memory_percent | 8.440 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 47.000 |
| docker:zg-zk | memory_percent | 1.310 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 15200255.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 71.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.825 |
| process:counter | cpu_seconds_total | 146.031 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43896832.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.562 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 36990976.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 299.992 |
| process:knowpost | cpu_seconds_total | 5530.312 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 96018432.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 282.103 |
| process:relation | cpu_seconds_total | 5156.391 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 90935296.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.969 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37036032.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.765 |
| process:user-storage | cpu_seconds_total | 1.031 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35201024.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 123702893.000 |
| redis | connected_clients | 241.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 420.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698621.000 |
| redis | keyspace_hits | 723236567.000 |
| redis | keyspace_misses | 388353.000 |
| redis | net_input_bytes | 29313269539.000 |
| redis | net_output_bytes | 132312324184.000 |
| redis | ops_per_sec | 25566.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22834.000 |
| redis | used_memory_bytes | 107679920.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
