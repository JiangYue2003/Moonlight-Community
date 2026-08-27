# Feed 压测报告：hybrid / gateway / page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:42:45+08:00
- 采样时长：1m0.0117584s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 95405 | 95405 | 0 | 0 | 1589.95 | 9.854 | 13.722 | 17.294 | 22.442 | 42.449 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 662 | 0.007 |
| redis | 191472 | 2.007 |
| relation | 240 | 0.003 |

- Cold compute：95405（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 11448600 | 120.000 |
| merge_candidates | 22992605 | 241.000 |
| redis_commands | 572430 | 6.000 |
| redis_members | 59246505 | 621.000 |
| redis_roundtrips | 95405 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 95405 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 95405 | 4.779 |
| counter | 240 | 4.497 |
| hydrate | 95405 | 4.280 |
| inbox | 95405 | 4.778 |
| merge_dedup | 95405 | 0.025 |
| relation | 240 | 5.128 |
| route | 95405 | 0.104 |
| total | 95405 | 9.231 |

## Redis 本轮边界增量

- Commands：691970；input：445736814 bytes；output：4370421486 bytes
- Hits/Misses：12027829/694；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：13883；safety epoch：3670 -> 3670

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.190 |
| client:loadtest | cpu_percent_total | 35.045 |
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
| docker:zg-canal | cpu_percent | 0.280 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.350 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 154.000 |
| docker:zg-etcd | cpu_percent | 5.040 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 122.840 |
| docker:zg-kafka | memory_percent | 7.500 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 51.000 |
| docker:zg-zk | memory_percent | 1.150 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 224833.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 5.426 |
| process:counter | cpu_seconds_total | 108.719 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45985792.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 149.412 |
| process:gateway | cpu_seconds_total | 1253.438 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51806208.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 189.760 |
| process:knowpost | cpu_seconds_total | 7451.562 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71000064.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.024 |
| process:relation | cpu_seconds_total | 15.906 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48205824.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 3.016 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38133760.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 61.190 |
| process:user-storage | cpu_seconds_total | 446.969 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54087680.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 153835035.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3670.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596969.000 |
| redis | keyspace_hits | 1098896139.000 |
| redis | keyspace_misses | 251742.000 |
| redis | net_input_bytes | 47937200932.000 |
| redis | net_output_bytes | 289676547813.000 |
| redis | ops_per_sec | 13883.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26604.000 |
| redis | used_memory_bytes | 102412216.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
