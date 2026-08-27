# Feed 压测报告：hybrid / rpc / hot-read-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:13:30+08:00
- 采样时长：1m0.0397561s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 184580 | 184580 | 0 | 0 | 3074.93 | 41.391 | 48.234 | 51.132 | 58.212 | 94.879 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 89 | 0.000 |
| redis | 553829 | 3.000 |
| relation | 184580 | 1.000 |

- Cold compute：184580（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 184580 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 184580 | 0.200 |
| counter | 12 | 1.780 |
| hydrate | 184580 | 0.292 |
| inbox | 184580 | 0.219 |
| merge_dedup | 184580 | 0.011 |
| relation | 184580 | 40.427 |
| route | 184580 | 0.003 |
| total | 184580 | 41.173 |

## Redis 本轮边界增量

- Commands：1536189；input：351575833 bytes；output：1607472738 bytes
- Hits/Misses：8675561/89；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27405；safety epoch：399 -> 399

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.546 |
| client:loadtest | cpu_percent_total | 72.738 |
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
| docker:zg-canal | cpu_percent | 3.180 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.470 |
| docker:zg-es | memory_percent | 11.840 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 6.450 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 209.050 |
| docker:zg-kafka | memory_percent | 8.370 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 60.420 |
| docker:zg-zk | memory_percent | 1.180 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 10852639.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 73.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 11.277 |
| process:counter | cpu_seconds_total | 45.562 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40734720.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.219 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37285888.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 317.036 |
| process:knowpost | cpu_seconds_total | 1531.516 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 97611776.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 277.998 |
| process:relation | cpu_seconds_total | 1447.922 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 85082112.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.772 |
| process:search | cpu_seconds_total | 0.328 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37138432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.025 |
| process:user-storage | cpu_seconds_total | 0.281 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35332096.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 87293091.000 |
| redis | connected_clients | 177.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 399.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698832.000 |
| redis | keyspace_hits | 501278559.000 |
| redis | keyspace_misses | 95060.000 |
| redis | net_input_bytes | 20394959711.000 |
| redis | net_output_bytes | 91403937184.000 |
| redis | ops_per_sec | 27405.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21249.000 |
| redis | used_memory_bytes | 105596680.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
