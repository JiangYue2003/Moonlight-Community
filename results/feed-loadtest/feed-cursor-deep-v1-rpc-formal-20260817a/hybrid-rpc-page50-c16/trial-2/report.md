# Feed 压测报告：hybrid / rpc / page50-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:19:18+08:00
- 采样时长：1m0.0186355s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 34876 | 34876 | 0 | 0 | 581.12 | 26.948 | 33.568 | 35.945 | 41.641 | 59.120 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.007 |
| mysql | 4397 | 0.126 |
| redis | 74149 | 2.126 |
| relation | 240 | 0.007 |

- Cold compute：34876（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 35573520 | 1020.000 |
| merge_candidates | 52279124 | 1499.000 |
| redis_commands | 209256 | 6.000 |
| redis_members | 52314000 | 1500.000 |
| redis_roundtrips | 34876 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 34876 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 34876 | 5.976 |
| counter | 240 | 5.300 |
| hydrate | 34876 | 20.601 |
| inbox | 34876 | 5.973 |
| merge_dedup | 34876 | 0.171 |
| relation | 240 | 5.934 |
| route | 34876 | 0.288 |
| total | 34876 | 27.223 |

## Redis 本轮边界增量

- Commands：272742；input：1263858967 bytes；output：8014213431 bytes
- Hits/Misses：35783508/6730；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：5004；safety epoch：3652 -> 3652

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.823 |
| client:loadtest | cpu_percent_total | 13.173 |
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
| docker:zg-es | cpu_percent | 0.900 |
| docker:zg-es | memory_percent | 12.220 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.460 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 142.970 |
| docker:zg-kafka | memory_percent | 7.580 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 46.090 |
| docker:zg-zk | memory_percent | 1.180 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 160625.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 26.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.648 |
| process:counter | cpu_seconds_total | 71.969 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46993408.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.782 |
| process:gateway | cpu_seconds_total | 258.375 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43880448.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 254.077 |
| process:knowpost | cpu_seconds_total | 4785.578 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 81186816.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.097 |
| process:relation | cpu_seconds_total | 9.469 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48730112.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.312 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38105088.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.005 |
| process:user-storage | cpu_seconds_total | 97.469 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43859968.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 99652622.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3652.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597802.000 |
| redis | keyspace_hits | 757622385.000 |
| redis | keyspace_misses | 166805.000 |
| redis | net_input_bytes | 32464350530.000 |
| redis | net_output_bytes | 214637648985.000 |
| redis | ops_per_sec | 5004.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 25197.000 |
| redis | used_memory_bytes | 103609472.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
