# Feed 压测报告：hybrid / gateway / page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:44:00+08:00
- 采样时长：1m0.0141036s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 94122 | 94122 | 0 | 0 | 1568.51 | 10.071 | 14.183 | 17.146 | 24.079 | 46.040 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 693 | 0.007 |
| redis | 188937 | 2.007 |
| relation | 240 | 0.003 |

- Cold compute：94122（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 11294640 | 120.000 |
| merge_candidates | 22683402 | 241.000 |
| redis_commands | 564732 | 6.000 |
| redis_members | 58449762 | 621.000 |
| redis_roundtrips | 94122 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 94122 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 94122 | 4.838 |
| counter | 240 | 4.218 |
| hydrate | 94122 | 4.338 |
| inbox | 94122 | 4.837 |
| merge_dedup | 94122 | 0.026 |
| relation | 240 | 4.913 |
| route | 94122 | 0.109 |
| total | 94122 | 9.354 |

## Redis 本轮边界增量

- Commands：683062；input：439776870 bytes；output：4311647218 bytes
- Hits/Misses：11866116/747；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：15296；safety epoch：3671 -> 3671

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.187 |
| client:loadtest | cpu_percent_total | 34.992 |
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
| docker:zg-canal | cpu_percent | 2.210 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.450 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.140 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 19.000 |
| docker:zg-kafka | cpu_percent | 106.460 |
| docker:zg-kafka | memory_percent | 7.330 |
| docker:zg-kafka | pids | 115.000 |
| docker:zg-zk | cpu_percent | 47.520 |
| docker:zg-zk | memory_percent | 1.150 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 226172.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.188 |
| process:counter | cpu_seconds_total | 110.531 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46964736.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 138.254 |
| process:gateway | cpu_seconds_total | 1332.906 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50610176.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 224.629 |
| process:knowpost | cpu_seconds_total | 7558.312 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71946240.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 10.074 |
| process:relation | cpu_seconds_total | 16.344 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48631808.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 3.062 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38137856.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 70.387 |
| process:user-storage | cpu_seconds_total | 481.016 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 57724928.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 154631485.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3671.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596953.000 |
| redis | keyspace_hits | 1112668609.000 |
| redis | keyspace_misses | 252763.000 |
| redis | net_input_bytes | 48447914535.000 |
| redis | net_output_bytes | 294680959200.000 |
| redis | ops_per_sec | 15296.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26679.000 |
| redis | used_memory_bytes | 102431280.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
