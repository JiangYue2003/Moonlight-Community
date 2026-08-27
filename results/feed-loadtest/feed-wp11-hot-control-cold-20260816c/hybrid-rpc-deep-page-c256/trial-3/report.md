# Feed 压测报告：hybrid / rpc / deep-page-c256

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:48:44+08:00
- 采样时长：1m0.1202989s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 170039 | 170039 | 0 | 0 | 2829.02 | 76.131 | 151.555 | 184.337 | 255.066 | 561.815 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 148 | 0.001 |
| redis | 510265 | 3.001 |
| relation | 170039 | 1.000 |

- Cold compute：170039（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 170039 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 170039 | 0.197 |
| counter | 240 | 0.776 |
| hydrate | 170039 | 0.349 |
| inbox | 170039 | 0.216 |
| merge_dedup | 170039 | 0.015 |
| relation | 170039 | 89.116 |
| route | 170039 | 0.004 |
| total | 170039 | 89.924 |

## Redis 本轮边界增量

- Commands：1425229；input：431651866 bytes；output：2004794907 bytes
- Hits/Misses：10888432/170193；run hit rate：98.46%
- Evicted/Rejected：0/0；ops/s max：25388；safety epoch：427 -> 427

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.280 |
| client:loadtest | cpu_percent_total | 68.482 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.600 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 4.930 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 180.220 |
| docker:zg-kafka | memory_percent | 8.490 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 56.130 |
| docker:zg-zk | memory_percent | 1.050 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 16619053.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 71.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.821 |
| process:counter | cpu_seconds_total | 179.594 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43945984.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.734 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37048320.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 304.088 |
| process:knowpost | cpu_seconds_total | 6885.328 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 122855424.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 282.730 |
| process:relation | cpu_seconds_total | 6378.734 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 100642816.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.109 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37072896.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.550 |
| process:user-storage | cpu_seconds_total | 1.266 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35274752.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 135600137.000 |
| redis | connected_clients | 304.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 427.000 |
| redis | hit_rate | 0.998 |
| redis | keys | 696500.000 |
| redis | keyspace_hits | 813901435.000 |
| redis | keyspace_misses | 1805639.000 |
| redis | net_input_bytes | 32909566672.000 |
| redis | net_output_bytes | 149006030497.000 |
| redis | ops_per_sec | 25388.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23363.000 |
| redis | used_memory_bytes | 107542832.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
