# Feed 压测报告：hybrid / rpc / page1-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:45:27+08:00
- 采样时长：1m0.0239785s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 316123 | 316123 | 0 | 0 | 5268.35 | 5.581 | 8.710 | 10.172 | 12.273 | 23.822 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 471 | 0.001 |
| redis | 632717 | 2.001 |
| relation | 240 | 0.001 |

- Cold compute：316123（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 12644920 | 40.000 |
| merge_candidates | 25605963 | 81.000 |
| redis_commands | 1896738 | 6.000 |
| redis_members | 77766258 | 246.000 |
| redis_roundtrips | 316123 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 316123 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 316123 | 3.024 |
| counter | 240 | 3.183 |
| hydrate | 316123 | 2.772 |
| inbox | 316123 | 3.023 |
| merge_dedup | 316123 | 0.010 |
| relation | 240 | 3.323 |
| route | 316123 | 0.032 |
| total | 316123 | 5.860 |

## Redis 本轮边界增量

- Commands：2244563；input：588165149 bytes；output：5351680523 bytes
- Hits/Misses：14548660/606；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：40445；safety epoch：3625 -> 3625

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.471 |
| client:loadtest | cpu_percent_total | 71.534 |
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
| docker:zg-canal | cpu_percent | 0.370 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.260 |
| docker:zg-es | memory_percent | 12.180 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.560 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 150.440 |
| docker:zg-kafka | memory_percent | 7.090 |
| docker:zg-kafka | pids | 98.000 |
| docker:zg-zk | cpu_percent | 48.190 |
| docker:zg-zk | memory_percent | 1.100 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 101518.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 31.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 11.610 |
| process:counter | cpu_seconds_total | 25.500 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45531136.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 258.094 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45101056.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 298.019 |
| process:knowpost | cpu_seconds_total | 1102.281 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71147520.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.321 |
| process:relation | cpu_seconds_total | 2.344 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48369664.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.567 |
| process:search | cpu_seconds_total | 1.188 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42360832.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.532 |
| process:user-storage | cpu_seconds_total | 94.953 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44175360.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 33700933.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3625.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597732.000 |
| redis | keyspace_hits | 202918302.000 |
| redis | keyspace_misses | 89076.000 |
| redis | net_input_bytes | 8935980455.000 |
| redis | net_output_bytes | 61719012483.000 |
| redis | ops_per_sec | 40445.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23166.000 |
| redis | used_memory_bytes | 103226760.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
