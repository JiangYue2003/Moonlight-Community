# Feed 压测报告：hybrid / gateway / page50-c16

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:36:58+08:00
- 采样时长：10.0183599s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4389 | 4389 | 0 | 0 | 438.12 | 32.737 | 53.059 | 57.347 | 68.248 | 84.465 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.009 |
| mysql | 871 | 0.198 |
| redis | 9649 | 2.198 |
| relation | 40 | 0.009 |

- Cold compute：4389（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 4476780 | 1020.000 |
| merge_candidates | 6579111 | 1499.000 |
| redis_commands | 26334 | 6.000 |
| redis_members | 6583500 | 1500.000 |
| redis_roundtrips | 4389 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 4389 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 4389 | 7.686 |
| counter | 40 | 4.847 |
| hydrate | 4389 | 26.940 |
| inbox | 4389 | 7.682 |
| merge_dedup | 4389 | 0.171 |
| relation | 40 | 6.527 |
| route | 4389 | 0.434 |
| total | 4389 | 35.425 |

## Redis 本轮边界增量

- Commands：35355；input：159229549 bytes；output：1008460739 bytes
- Hits/Misses：4502826/1551；run hit rate：99.97%
- Evicted/Rejected：0/0；ops/s max：3624；safety epoch：3617 -> 3617

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.751 |
| client:loadtest | cpu_percent_total | 12.009 |
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
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.750 |
| docker:zg-es | memory_percent | 12.080 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.090 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 24.140 |
| docker:zg-kafka | memory_percent | 7.030 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.100 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 96188.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 28.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 3.098 |
| process:counter | cpu_seconds_total | 9.531 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45342720.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 42.600 |
| process:gateway | cpu_seconds_total | 192.469 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50954240.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 201.332 |
| process:knowpost | cpu_seconds_total | 227.578 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 77803520.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.593 |
| process:relation | cpu_seconds_total | 1.219 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47968256.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.438 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42000384.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 19.910 |
| process:user-storage | cpu_seconds_total | 71.656 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 53989376.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19789693.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3617.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 598242.000 |
| redis | keyspace_hits | 111948985.000 |
| redis | keyspace_misses | 84641.000 |
| redis | net_input_bytes | 5229603287.000 |
| redis | net_output_bytes | 29466961876.000 |
| redis | ops_per_sec | 3624.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22607.000 |
| redis | used_memory_bytes | 102778736.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
