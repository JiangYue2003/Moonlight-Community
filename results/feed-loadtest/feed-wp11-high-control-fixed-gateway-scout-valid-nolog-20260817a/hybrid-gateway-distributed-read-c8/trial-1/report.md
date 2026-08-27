# Feed 压测报告：hybrid / gateway / distributed-read-c8

- Run ID：`feed-wp11-high-control-fixed-gateway-scout-valid-nolog-20260817a`
- 开始时间：2026-08-17T01:10:10+08:00
- 采样时长：10.0039183s
- 并发：8
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 30059 | 30059 | 0 | 0 | 3005.20 | 2.642 | 3.517 | 3.808 | 4.543 | 7.194 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 19059 | 0.634 |
| mysql | 0 | 0.000 |
| redis | 90177 | 3.000 |
| relation | 30059 | 1.000 |

- Cold compute：30059（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 30059 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 30059 | 0.207 |
| counter | 19059 | 0.465 |
| hydrate | 30059 | 0.219 |
| inbox | 30059 | 0.198 |
| merge_dedup | 30059 | 0.002 |
| relation | 30059 | 0.927 |
| route | 30059 | 0.306 |
| total | 30059 | 1.876 |

## Redis 本轮边界增量

- Commands：335615；input：26525004 bytes；output：57276202 bytes
- Hits/Misses：501070/57561；run hit rate：89.70%
- Evicted/Rejected：0/0；ops/s max：37702；safety epoch：595 -> 595

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.270 |
| client:loadtest | cpu_percent_total | 52.323 |
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
| docker:zg-canal | cpu_percent | 2.680 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.250 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.160 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 132.780 |
| docker:zg-kafka | memory_percent | 8.730 |
| docker:zg-kafka | pids | 97.000 |
| docker:zg-zk | cpu_percent | 0.180 |
| docker:zg-zk | memory_percent | 1.550 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 27621206.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 72.021 |
| process:counter | cpu_seconds_total | 179.812 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 53968896.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 186.855 |
| process:gateway | cpu_seconds_total | 448.531 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 50401280.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 280.705 |
| process:knowpost | cpu_seconds_total | 1507.031 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 80248832.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 205.945 |
| process:relation | cpu_seconds_total | 1200.969 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 57475072.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.609 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35987456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 92.930 |
| process:user-storage | cpu_seconds_total | 165.703 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 48885760.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 275484252.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 595.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 716619.000 |
| redis | keyspace_hits | 1107045261.000 |
| redis | keyspace_misses | 13939087.000 |
| redis | net_input_bytes | 51629775039.000 |
| redis | net_output_bytes | 205213759885.000 |
| redis | ops_per_sec | 37702.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53399.000 |
| redis | used_memory_bytes | 129313320.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
