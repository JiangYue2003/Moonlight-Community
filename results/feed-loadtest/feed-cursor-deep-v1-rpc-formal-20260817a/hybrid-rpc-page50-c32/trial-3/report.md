# Feed 压测报告：hybrid / rpc / page50-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:24:19+08:00
- 采样时长：1m0.0276783s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 35474 | 35474 | 0 | 0 | 590.99 | 52.834 | 65.639 | 70.333 | 80.028 | 105.416 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.007 |
| mysql | 11399 | 0.321 |
| redis | 82347 | 2.321 |
| relation | 240 | 0.007 |

- Cold compute：35474（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 36183480 | 1020.000 |
| merge_candidates | 53175526 | 1499.000 |
| redis_commands | 212844 | 6.000 |
| redis_members | 53211000 | 1500.000 |
| redis_roundtrips | 35474 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 35474 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 35474 | 11.860 |
| counter | 240 | 7.668 |
| hydrate | 35474 | 40.928 |
| inbox | 35474 | 11.856 |
| merge_dedup | 35474 | 0.181 |
| relation | 240 | 9.871 |
| route | 35474 | 0.583 |
| total | 35474 | 53.776 |

## Redis 本轮边界增量

- Commands：279272；input：1287162076 bytes；output：8150042177 bytes
- Hits/Misses：36387558/16209；run hit rate：99.96%
- Evicted/Rejected：0/0；ops/s max：5116；safety epoch：3656 -> 3656

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.797 |
| client:loadtest | cpu_percent_total | 12.755 |
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
| docker:zg-canal | cpu_percent | 2.220 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.490 |
| docker:zg-es | memory_percent | 12.240 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.280 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 178.650 |
| docker:zg-kafka | memory_percent | 7.130 |
| docker:zg-kafka | pids | 98.000 |
| docker:zg-zk | cpu_percent | 15.820 |
| docker:zg-zk | memory_percent | 1.100 |
| docker:zg-zk | pids | 92.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 212646.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 33.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.648 |
| process:counter | cpu_seconds_total | 76.656 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46592000.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.406 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43884544.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 266.253 |
| process:knowpost | cpu_seconds_total | 5372.656 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 100737024.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 10.531 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49971200.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 2.469 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38117376.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.323 |
| process:user-storage | cpu_seconds_total | 97.906 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43945984.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 100970792.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3656.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597796.000 |
| redis | keyspace_hits | 926688510.000 |
| redis | keyspace_misses | 239782.000 |
| redis | net_input_bytes | 38445830662.000 |
| redis | net_output_bytes | 252505087852.000 |
| redis | ops_per_sec | 5116.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 25498.000 |
| redis | used_memory_bytes | 105527136.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
