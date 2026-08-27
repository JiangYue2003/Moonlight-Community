# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-control-fixed-rpc-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:12:13+08:00
- 采样时长：1m0.0251058s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 326083 | 326083 | 0 | 0 | 5434.52 | 5.310 | 8.639 | 10.361 | 14.306 | 33.956 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 140153 | 0.430 |
| mysql | 67 | 0.000 |
| redis | 978316 | 3.000 |
| relation | 326083 | 1.000 |

- Cold compute：326083（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 326083 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 326083 | 0.960 |
| counter | 140153 | 1.363 |
| hydrate | 326083 | 1.040 |
| inbox | 326083 | 1.069 |
| merge_dedup | 326083 | 0.002 |
| relation | 326083 | 1.935 |
| route | 326083 | 0.594 |
| total | 326083 | 5.619 |

## Redis 本轮边界增量

- Commands：3266290；input：275037989 bytes；output：611478379 bytes
- Hits/Misses：5114172/624606；run hit rate：89.12%
- Evicted/Rejected：0/0；ops/s max：67955；safety epoch：598 -> 598

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.393 |
| client:loadtest | cpu_percent_total | 86.292 |
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
| docker:zg-canal | cpu_percent | 3.780 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.660 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 7.180 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 221.360 |
| docker:zg-kafka | memory_percent | 8.810 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 68.190 |
| docker:zg-zk | memory_percent | 1.710 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 28021003.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 50.000 |
| mysql | threads_running | 11.000 |
| process:counter | cpu_percent | 119.261 |
| process:counter | cpu_seconds_total | 254.484 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54939648.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.046 |
| process:gateway | cpu_seconds_total | 494.594 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 52477952.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 371.908 |
| process:knowpost | cpu_seconds_total | 1793.781 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 86355968.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 294.348 |
| process:relation | cpu_seconds_total | 1403.734 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 59953152.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.641 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 36040704.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.557 |
| process:user-storage | cpu_seconds_total | 183.906 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 50245632.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 279676166.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 598.000 |
| redis | hit_rate | 0.987 |
| redis | keys | 714327.000 |
| redis | keyspace_hits | 1113398096.000 |
| redis | keyspace_misses | 14704577.000 |
| redis | net_input_bytes | 51976348379.000 |
| redis | net_output_bytes | 205967052295.000 |
| redis | ops_per_sec | 67955.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53572.000 |
| redis | used_memory_bytes | 128772448.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
