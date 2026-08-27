# Feed 压测报告：hybrid / rpc / page50-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:23:04+08:00
- 采样时长：1m0.0273161s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 35208 | 35208 | 0 | 0 | 586.55 | 53.045 | 66.803 | 72.171 | 82.549 | 113.055 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.007 |
| mysql | 11637 | 0.331 |
| redis | 82053 | 2.331 |
| relation | 240 | 0.007 |

- Cold compute：35208（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 35912160 | 1020.000 |
| merge_candidates | 52776792 | 1499.000 |
| redis_commands | 211248 | 6.000 |
| redis_members | 52812000 | 1500.000 |
| redis_roundtrips | 35208 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 35208 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 35208 | 11.805 |
| counter | 240 | 7.935 |
| hydrate | 35208 | 41.359 |
| inbox | 35208 | 11.801 |
| merge_dedup | 35208 | 0.180 |
| relation | 240 | 10.036 |
| route | 35208 | 0.619 |
| total | 35208 | 54.188 |

## Redis 本轮边界增量

- Commands：276576；input：1277373740 bytes；output：8089018482 bytes
- Hits/Misses：36115342/15508；run hit rate：99.96%
- Evicted/Rejected：0/0；ops/s max：5024；safety epoch：3655 -> 3655

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.711 |
| client:loadtest | cpu_percent_total | 11.375 |
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
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.670 |
| docker:zg-es | memory_percent | 12.240 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.620 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 159.260 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 43.480 |
| docker:zg-zk | memory_percent | 1.170 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 198027.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 35.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.615 |
| process:counter | cpu_seconds_total | 75.453 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47190016.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.406 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43884544.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 281.432 |
| process:knowpost | cpu_seconds_total | 5223.344 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 101154816.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.099 |
| process:relation | cpu_seconds_total | 10.188 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49082368.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 2.438 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38117376.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.546 |
| process:user-storage | cpu_seconds_total | 97.797 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43843584.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 100637390.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3655.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597797.000 |
| redis | keyspace_hits | 883943042.000 |
| redis | keyspace_misses | 219029.000 |
| redis | net_input_bytes | 36933051938.000 |
| redis | net_output_bytes | 242930815044.000 |
| redis | ops_per_sec | 5024.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 25423.000 |
| redis | used_memory_bytes | 105448928.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
