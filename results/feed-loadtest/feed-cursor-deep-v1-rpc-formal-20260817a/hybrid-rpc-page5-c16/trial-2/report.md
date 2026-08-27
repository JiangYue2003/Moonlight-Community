# Feed 压测报告：hybrid / rpc / page5-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:49:12+08:00
- 采样时长：1m0.0158212s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 140297 | 140297 | 0 | 0 | 2338.11 | 6.103 | 10.204 | 11.252 | 13.790 | 22.563 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 608 | 0.004 |
| redis | 281202 | 2.004 |
| relation | 240 | 0.002 |

- Cold compute：140297（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 16835640 | 120.000 |
| merge_candidates | 33811577 | 241.000 |
| redis_commands | 841782 | 6.000 |
| redis_members | 87124437 | 621.000 |
| redis_roundtrips | 140297 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 140297 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 140297 | 3.337 |
| counter | 240 | 3.181 |
| hydrate | 140297 | 3.122 |
| inbox | 140297 | 3.335 |
| merge_dedup | 140297 | 0.031 |
| relation | 240 | 3.418 |
| route | 140297 | 0.053 |
| total | 140297 | 6.591 |

## Redis 本轮边界增量

- Commands：1014359；input：655395237 bytes；output：6426778734 bytes
- Hits/Misses：17684061/970；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：17904；safety epoch：3628 -> 3628

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.869 |
| client:loadtest | cpu_percent_total | 45.899 |
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
| docker:zg-canal | cpu_percent | 1.810 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.970 |
| docker:zg-es | memory_percent | 12.200 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.310 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 154.390 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 29.250 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 105168.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.902 |
| process:counter | cpu_seconds_total | 31.266 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45568000.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.125 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45113344.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 266.414 |
| process:knowpost | cpu_seconds_total | 1573.453 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72540160.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.099 |
| process:relation | cpu_seconds_total | 2.922 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48197632.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.469 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42565632.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 95.203 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44331008.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 38770955.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3628.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597845.000 |
| redis | keyspace_hits | 261874432.000 |
| redis | keyspace_misses | 93377.000 |
| redis | net_input_bytes | 11180229773.000 |
| redis | net_output_bytes | 83222135222.000 |
| redis | ops_per_sec | 17904.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23391.000 |
| redis | used_memory_bytes | 102475944.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
