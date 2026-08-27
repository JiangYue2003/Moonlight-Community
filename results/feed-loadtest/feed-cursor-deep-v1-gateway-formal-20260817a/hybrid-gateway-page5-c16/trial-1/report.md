# Feed 压测报告：hybrid / gateway / page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:41:30+08:00
- 采样时长：1m0.0120604s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 86219 | 86219 | 0 | 0 | 1436.85 | 11.018 | 15.400 | 19.928 | 25.897 | 42.444 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 613 | 0.007 |
| redis | 173051 | 2.007 |
| relation | 240 | 0.003 |

- Cold compute：86219（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 10346280 | 120.000 |
| merge_candidates | 20778779 | 241.000 |
| redis_commands | 517314 | 6.000 |
| redis_members | 53541999 | 621.000 |
| redis_roundtrips | 86219 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 86219 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 86219 | 5.240 |
| counter | 240 | 4.332 |
| hydrate | 86219 | 4.812 |
| inbox | 86219 | 5.238 |
| merge_dedup | 86219 | 0.026 |
| relation | 240 | 5.018 |
| route | 86219 | 0.128 |
| total | 86219 | 10.249 |

## Redis 本轮边界增量

- Commands：625757；input：402819634 bytes；output：3949640845 bytes
- Hits/Misses：10870457/620；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：14386；safety epoch：3669 -> 3669

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.210 |
| client:loadtest | cpu_percent_total | 35.357 |
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
| docker:zg-canal | cpu_percent | 0.150 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.530 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.690 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 128.650 |
| docker:zg-kafka | memory_percent | 7.370 |
| docker:zg-kafka | pids | 118.000 |
| docker:zg-zk | cpu_percent | 28.340 |
| docker:zg-zk | memory_percent | 1.000 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 223514.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.421 |
| process:counter | cpu_seconds_total | 106.859 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45977600.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 134.700 |
| process:gateway | cpu_seconds_total | 1171.469 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51326976.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 184.311 |
| process:knowpost | cpu_seconds_total | 7344.938 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71630848.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.098 |
| process:relation | cpu_seconds_total | 15.719 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48582656.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 3.016 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38137856.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 64.269 |
| process:user-storage | cpu_seconds_total | 413.547 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 59670528.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 153030963.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3669.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596997.000 |
| redis | keyspace_hits | 1084976092.000 |
| redis | keyspace_misses | 250751.000 |
| redis | net_input_bytes | 47421080766.000 |
| redis | net_output_bytes | 284618506196.000 |
| redis | ops_per_sec | 14386.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26529.000 |
| redis | used_memory_bytes | 102476336.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
