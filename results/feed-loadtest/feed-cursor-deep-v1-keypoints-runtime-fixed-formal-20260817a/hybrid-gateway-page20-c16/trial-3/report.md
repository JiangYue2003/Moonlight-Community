# Feed 压测报告：hybrid / gateway / page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:36:04+08:00
- 采样时长：1m0.0138527s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 61862 | 61862 | 0 | 0 | 1030.87 | 14.063 | 22.652 | 25.788 | 32.813 | 59.012 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.004 |
| mysql | 2220 | 0.036 |
| redis | 125944 | 2.036 |
| relation | 240 | 0.004 |

- Cold compute：61862（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 25982040 | 420.000 |
| merge_candidates | 52025942 | 841.000 |
| redis_commands | 371172 | 6.000 |
| redis_members | 56974902 | 921.000 |
| redis_roundtrips | 61862 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 61862 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 61862 | 7.176 |
| counter | 240 | 5.992 |
| hydrate | 61862 | 7.139 |
| inbox | 61862 | 7.174 |
| merge_dedup | 61862 | 0.085 |
| relation | 240 | 6.683 |
| route | 61862 | 0.225 |
| total | 61862 | 14.742 |

## Redis 本轮边界增量

- Commands：454268；input：939255548 bytes；output：6656654633 bytes
- Hits/Misses：26357575/3109；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：8278；safety epoch：3754 -> 3754

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.386 |
| client:loadtest | cpu_percent_total | 22.182 |
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
| docker:zg-canal | cpu_percent | 0.180 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.600 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.560 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 149.830 |
| docker:zg-kafka | memory_percent | 7.630 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 48.480 |
| docker:zg-zk | memory_percent | 1.310 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 507534.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 3.891 |
| process:counter | cpu_seconds_total | 281.641 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47480832.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 98.609 |
| process:gateway | cpu_seconds_total | 5621.688 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51654656.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 238.226 |
| process:knowpost | cpu_seconds_total | 17319.094 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 75350016.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.549 |
| process:relation | cpu_seconds_total | 46.391 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48812032.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 5.625 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38043648.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 45.651 |
| process:user-storage | cpu_seconds_total | 2011.609 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 59285504.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 287984732.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3754.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596243.000 |
| redis | keyspace_hits | 2606584335.000 |
| redis | keyspace_misses | 673363.000 |
| redis | net_input_bytes | 109457439785.000 |
| redis | net_output_bytes | 645442796480.000 |
| redis | ops_per_sec | 8278.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 33403.000 |
| redis | used_memory_bytes | 102129872.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
