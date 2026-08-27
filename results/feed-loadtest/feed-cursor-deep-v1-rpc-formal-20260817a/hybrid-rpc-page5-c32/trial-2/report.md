# Feed 压测报告：hybrid / rpc / page5-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:52:57+08:00
- 采样时长：1m0.0170495s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 142963 | 142963 | 0 | 0 | 2382.44 | 12.145 | 19.871 | 22.454 | 27.247 | 50.609 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 1392 | 0.010 |
| redis | 287318 | 2.010 |
| relation | 240 | 0.002 |

- Cold compute：142963（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 17155560 | 120.000 |
| merge_candidates | 34454083 | 241.000 |
| redis_commands | 857778 | 6.000 |
| redis_members | 88780023 | 621.000 |
| redis_roundtrips | 142963 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 142963 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 142963 | 6.661 |
| counter | 240 | 6.153 |
| hydrate | 142963 | 6.311 |
| inbox | 142963 | 6.660 |
| merge_dedup | 142963 | 0.031 |
| relation | 240 | 6.206 |
| route | 142963 | 0.118 |
| total | 142963 | 13.172 |

## Redis 本轮边界增量

- Commands：1020235；input：666961376 bytes；output：6548543524 bytes
- Hits/Misses：18019369/1460；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：18255；safety epoch：3631 -> 3631

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.773 |
| client:loadtest | cpu_percent_total | 44.362 |
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
| docker:zg-canal | cpu_percent | 0.170 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.550 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.560 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 135.880 |
| docker:zg-kafka | memory_percent | 7.540 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 41.180 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 110853.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 29.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 8.526 |
| process:counter | cpu_seconds_total | 36.422 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46346240.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 258.156 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45371392.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 262.573 |
| process:knowpost | cpu_seconds_total | 2025.625 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 74518528.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 3.828 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48992256.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.593 |
| process:search | cpu_seconds_total | 1.734 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 41799680.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.549 |
| process:user-storage | cpu_seconds_total | 95.562 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44023808.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 42359955.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3631.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597906.000 |
| redis | keyspace_hits | 324803313.000 |
| redis | keyspace_misses | 98394.000 |
| redis | net_input_bytes | 13511322757.000 |
| redis | net_output_bytes | 106091938269.000 |
| redis | ops_per_sec | 18255.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23616.000 |
| redis | used_memory_bytes | 103355832.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
