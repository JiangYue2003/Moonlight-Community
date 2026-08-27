# Feed 压测报告：hybrid / gateway / page1-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:33:58+08:00
- 采样时长：1m0.0130407s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 128608 | 128608 | 0 | 0 | 2143.32 | 7.076 | 10.419 | 12.470 | 16.067 | 31.598 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 159 | 0.001 |
| redis | 257375 | 2.001 |
| relation | 240 | 0.002 |

- Cold compute：128608（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 5144320 | 40.000 |
| merge_candidates | 10417248 | 81.000 |
| redis_commands | 771648 | 6.000 |
| redis_members | 31637568 | 246.000 |
| redis_roundtrips | 128608 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 128608 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 128608 | 3.325 |
| counter | 240 | 3.198 |
| hydrate | 128608 | 2.933 |
| inbox | 128608 | 3.324 |
| merge_dedup | 128608 | 0.011 |
| relation | 240 | 3.875 |
| route | 128608 | 0.060 |
| total | 128608 | 6.352 |

## Redis 本轮边界增量

- Commands：929608；input：240252603 bytes；output：2177633337 bytes
- Hits/Misses：5923324/159；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：19658；safety epoch：3663 -> 3663

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.894 |
| client:loadtest | cpu_percent_total | 62.304 |
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
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.250 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.810 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 23.000 |
| docker:zg-kafka | cpu_percent | 163.280 |
| docker:zg-kafka | memory_percent | 7.550 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 47.980 |
| docker:zg-zk | memory_percent | 0.990 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 217946.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.059 |
| process:counter | cpu_seconds_total | 93.547 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46645248.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 240.628 |
| process:gateway | cpu_seconds_total | 404.078 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50708480.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 169.893 |
| process:knowpost | cpu_seconds_total | 6719.531 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70799360.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.378 |
| process:relation | cpu_seconds_total | 12.594 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48848896.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.098 |
| process:search | cpu_seconds_total | 2.781 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38125568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 91.212 |
| process:user-storage | cpu_seconds_total | 148.969 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55599104.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 147950797.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3663.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597568.000 |
| redis | keyspace_hits | 1044679208.000 |
| redis | keyspace_misses | 246937.000 |
| redis | net_input_bytes | 45830829722.000 |
| redis | net_output_bytes | 269860106324.000 |
| redis | ops_per_sec | 19658.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26077.000 |
| redis | used_memory_bytes | 102467136.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
