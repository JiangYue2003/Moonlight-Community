# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-distributed-treatment-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:41:53+08:00
- 采样时长：10.0166652s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 689599 | 689599 | 0 | 0 | 68958.39 | 0.519 | 0.632 | 1.043 | 1.318 | 13.538 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 24648 | 0.036 |
| counter | 1250 | 0.002 |
| mysql | 0 | 0.000 |
| redis | 24478 | 0.035 |
| relation | 1250 | 0.002 |

- Cold compute：3638（0.005 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 675865 |
| l1_stale | 0 |
| l2_fresh | 13734 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=1 pending=1

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 3638 | 0.154 |
| counter | 1250 | 0.454 |
| hydrate | 3638 | 0.186 |
| inbox | 3638 | 0.152 |
| merge_dedup | 3638 | 0.005 |
| relation | 1250 | 0.959 |
| route | 3638 | 0.494 |
| total | 689599 | 0.010 |

## Redis 本轮边界增量

- Commands：71600；input：11154073 bytes；output：28206615 bytes
- Hits/Misses：100109/1551；run hit rate：98.47%
- Evicted/Rejected：0/0；ops/s max：11589；safety epoch：535 -> 535

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 22.424 |
| client:loadtest | cpu_percent_total | 358.777 |
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
| docker:zg-canal | cpu_percent | 0.140 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.430 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.700 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 184.500 |
| docker:zg-kafka | memory_percent | 8.260 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 63.190 |
| docker:zg-zk | memory_percent | 1.420 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23266037.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 41.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 37.882 |
| process:counter | cpu_seconds_total | 15.594 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 47915008.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 35.234 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 50708480.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 417.357 |
| process:knowpost | cpu_seconds_total | 112.844 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 91942912.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 26.286 |
| process:relation | cpu_seconds_total | 8.812 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 53870592.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36175872.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 11.125 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 47276032.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 205373118.000 |
| redis | connected_clients | 162.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 535.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 689255.000 |
| redis | keyspace_hits | 1003478366.000 |
| redis | keyspace_misses | 5962527.000 |
| redis | net_input_bytes | 42775669427.000 |
| redis | net_output_bytes | 188693685487.000 |
| redis | ops_per_sec | 11589.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48102.000 |
| redis | used_memory_bytes | 108579280.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
