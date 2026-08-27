# Feed 压测报告：hybrid / rpc / hot-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:16:01+08:00
- 采样时长：1m0.0776897s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 182201 | 182201 | 0 | 0 | 3033.47 | 71.733 | 139.190 | 168.118 | 236.857 | 512.440 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 104 | 0.001 |
| redis | 546707 | 3.001 |
| relation | 182201 | 1.000 |

- Cold compute：182201（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 182201 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 182201 | 0.205 |
| counter | 12 | 1.631 |
| hydrate | 182201 | 0.298 |
| inbox | 182201 | 0.227 |
| merge_dedup | 182201 | 0.011 |
| relation | 182201 | 83.129 |
| route | 182201 | 0.003 |
| total | 182201 | 83.893 |

## Redis 本轮边界增量

- Commands：1517172；input：347102397 bytes；output：1586768860 bytes
- Hits/Misses：8563733/104；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26916；safety epoch：401 -> 401

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.693 |
| client:loadtest | cpu_percent_total | 75.085 |
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
| docker:zg-canal | cpu_percent | 2.560 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.400 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.770 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 192.610 |
| docker:zg-kafka | memory_percent | 8.400 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 58.820 |
| docker:zg-zk | memory_percent | 1.030 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 11283568.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 73.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 11.505 |
| process:counter | cpu_seconds_total | 55.250 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40820736.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.767 |
| process:gateway | cpu_seconds_total | 0.281 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 38121472.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 292.172 |
| process:knowpost | cpu_seconds_total | 1905.547 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 115605504.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 292.222 |
| process:relation | cpu_seconds_total | 1805.594 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 95145984.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.069 |
| process:search | cpu_seconds_total | 0.469 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37146624.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.359 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35467264.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 90884550.000 |
| redis | connected_clients | 189.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 401.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698811.000 |
| redis | keyspace_hits | 521498365.000 |
| redis | keyspace_misses | 95339.000 |
| redis | net_input_bytes | 21215160729.000 |
| redis | net_output_bytes | 95150612060.000 |
| redis | ops_per_sec | 26916.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21400.000 |
| redis | used_memory_bytes | 105762040.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
