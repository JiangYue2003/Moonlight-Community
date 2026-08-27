# Feed 压测报告：hybrid / gateway / page1-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:40:14+08:00
- 采样时长：1m0.0129881s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 106016 | 106016 | 0 | 0 | 1766.77 | 17.625 | 22.525 | 26.413 | 37.505 | 74.063 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 627 | 0.006 |
| redis | 212659 | 2.006 |
| relation | 240 | 0.002 |

- Cold compute：106016（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 4240640 | 40.000 |
| merge_candidates | 8587296 | 81.000 |
| redis_commands | 636096 | 6.000 |
| redis_members | 26079936 | 246.000 |
| redis_roundtrips | 106016 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 106016 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 106016 | 8.440 |
| counter | 240 | 7.935 |
| hydrate | 106016 | 8.125 |
| inbox | 106016 | 8.438 |
| merge_dedup | 106016 | 0.011 |
| relation | 240 | 9.372 |
| route | 106016 | 0.204 |
| total | 106016 | 16.803 |

## Redis 本轮边界增量

- Commands：757875；input：197483208 bytes；output：1794870203 bytes
- Hits/Misses：4883571/630；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：14289；safety epoch：3668 -> 3668

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.394 |
| client:loadtest | cpu_percent_total | 54.311 |
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
| docker:zg-canal | cpu_percent | 2.130 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.590 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.460 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 150.030 |
| docker:zg-kafka | memory_percent | 7.570 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 43.710 |
| docker:zg-zk | memory_percent | 1.160 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 222475.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.646 |
| process:counter | cpu_seconds_total | 105.219 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45379584.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 232.821 |
| process:gateway | cpu_seconds_total | 1090.016 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51806208.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 181.050 |
| process:knowpost | cpu_seconds_total | 7244.703 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70914048.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.097 |
| process:relation | cpu_seconds_total | 15.312 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48750592.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 2.984 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38129664.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 88.289 |
| process:user-storage | cpu_seconds_total | 381.359 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55738368.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 152310046.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3668.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596962.000 |
| redis | keyspace_hits | 1072517616.000 |
| redis | keyspace_misses | 249673.000 |
| redis | net_input_bytes | 46959090515.000 |
| redis | net_output_bytes | 280091799257.000 |
| redis | ops_per_sec | 14289.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26453.000 |
| redis | used_memory_bytes | 103197152.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
