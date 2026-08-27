# Feed 压测报告：hybrid / rpc / page1-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:40:26+08:00
- 采样时长：1m0.0221935s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 306187 | 306187 | 0 | 0 | 5102.96 | 2.863 | 4.696 | 5.211 | 6.304 | 13.443 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 178 | 0.001 |
| redis | 612552 | 2.001 |
| relation | 240 | 0.001 |

- Cold compute：306187（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 12247480 | 40.000 |
| merge_candidates | 24801147 | 81.000 |
| redis_commands | 1837122 | 6.000 |
| redis_members | 75322002 | 246.000 |
| redis_roundtrips | 306187 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 306187 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 306187 | 1.543 |
| counter | 240 | 1.536 |
| hydrate | 306187 | 1.345 |
| inbox | 306187 | 1.542 |
| merge_dedup | 306187 | 0.010 |
| relation | 240 | 2.144 |
| route | 306187 | 0.017 |
| total | 306187 | 2.935 |

## Redis 本轮边界增量

- Commands：2180121；input：570050776 bytes；output：5183653158 bytes
- Hits/Misses：14091954/268；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：39150；safety epoch：3621 -> 3621

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.574 |
| client:loadtest | cpu_percent_total | 73.176 |
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
| docker:zg-canal | cpu_percent | 2.460 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 2.590 |
| docker:zg-es | memory_percent | 12.120 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.930 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 153.500 |
| docker:zg-kafka | memory_percent | 7.520 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 46.930 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 97660.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 8.518 |
| process:counter | cpu_seconds_total | 16.906 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 44982272.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.062 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 47456256.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 281.665 |
| process:knowpost | cpu_seconds_total | 467.969 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 68648960.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.775 |
| process:relation | cpu_seconds_total | 1.750 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47632384.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.550 |
| process:search | cpu_seconds_total | 0.766 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42217472.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 94.406 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 45600768.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 23311577.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3621.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597985.000 |
| redis | keyspace_hits | 135741008.000 |
| redis | keyspace_misses | 86202.000 |
| redis | net_input_bytes | 6218394164.000 |
| redis | net_output_bytes | 37007719036.000 |
| redis | ops_per_sec | 39150.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22865.000 |
| redis | used_memory_bytes | 102259056.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
