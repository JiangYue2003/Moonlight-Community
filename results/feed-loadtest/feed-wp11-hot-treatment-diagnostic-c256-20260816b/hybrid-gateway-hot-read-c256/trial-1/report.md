# Feed 压测报告：hybrid / gateway / hot-read-c256

- Run ID：`feed-wp11-hot-treatment-diagnostic-c256-20260816b`
- 开始时间：2026-08-16T22:16:57+08:00
- 采样时长：10.0344067s
- 并发：256
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 53315 | 53315 | 0 | 0 | 5314.87 | 47.507 | 58.445 | 80.849 | 106.765 | 213.238 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 27 | 0.001 |
| counter | 1 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 20 | 0.000 |
| relation | 1 | 0.000 |

- Cold compute：3（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 53297 |
| l1_stale | 0 |
| l2_fresh | 18 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 3 | 0.000 |
| counter | 1 | 2.586 |
| hydrate | 3 | 0.535 |
| inbox | 3 | 0.000 |
| merge_dedup | 3 | 0.000 |
| relation | 1 | 2.207 |
| route | 3 | 1.598 |
| total | 53315 | 0.007 |

## Redis 本轮边界增量

- Commands：8885；input：635745 bytes；output：241927 bytes
- Hits/Misses：210/1；run hit rate：99.53%
- Evicted/Rejected：0/0；ops/s max：1091；safety epoch：519 -> 519

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.338 |
| client:loadtest | cpu_percent_total | 117.409 |
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
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.600 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.560 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 233.790 |
| docker:zg-kafka | memory_percent | 8.620 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 40.210 |
| docker:zg-zk | memory_percent | 1.300 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22734352.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 6.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 9.324 |
| process:counter | cpu_seconds_total | 254.234 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 44011520.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 388.313 |
| process:gateway | cpu_seconds_total | 5032.812 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 68235264.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 196.378 |
| process:knowpost | cpu_seconds_total | 9609.297 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 113152000.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 12.391 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 46678016.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 71.722 |
| process:search | cpu_seconds_total | 2.312 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36990976.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 208.601 |
| process:user-storage | cpu_seconds_total | 2520.266 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 55615488.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 196980112.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 519.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677471.000 |
| redis | keyspace_hits | 993436290.000 |
| redis | keyspace_misses | 5383946.000 |
| redis | net_input_bytes | 41975601666.000 |
| redis | net_output_bytes | 187105952293.000 |
| redis | ops_per_sec | 1091.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 43006.000 |
| redis | used_memory_bytes | 100188808.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
