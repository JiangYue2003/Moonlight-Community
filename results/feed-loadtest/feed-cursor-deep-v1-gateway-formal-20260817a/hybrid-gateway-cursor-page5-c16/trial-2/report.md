# Feed 压测报告：hybrid / gateway / cursor-page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:50:16+08:00
- 采样时长：1m0.0099072s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 68316 | 68316 | 0 | 0 | 1138.50 | 13.432 | 19.736 | 23.305 | 29.816 | 61.449 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：156.9776ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.004 |
| mysql | 126 | 0.002 |
| redis | 68442 | 1.002 |
| relation | 240 | 0.004 |

- Cold compute：68316（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1434636 | 21.000 |
| merge_candidates | 10247400 | 150.000 |
| redis_commands | 1161372 | 17.000 |
| redis_members | 16122576 | 236.000 |
| redis_roundtrips | 136632 | 2.000 |
| tie_members | 8607816 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 68316 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 4.669 |
| cursor_decode | 68316 | 0.010 |
| cursor_seek | 68316 | 8.467 |
| hydrate | 68316 | 4.251 |
| merge_dedup | 68316 | 0.009 |
| relation | 240 | 5.367 |
| route | 68316 | 0.129 |
| total | 68316 | 12.868 |

## Redis 本轮边界增量

- Commands：1252427；input：177119722 bytes；output：918522243 bytes
- Hits/Misses：2603354/135；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：23550；safety epoch：3676 -> 3676

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.301 |
| client:loadtest | cpu_percent_total | 36.817 |
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
| docker:zg-canal | cpu_percent | 2.450 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.690 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.590 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 117.700 |
| docker:zg-kafka | memory_percent | 7.490 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 45.520 |
| docker:zg-zk | memory_percent | 1.010 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 233609.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.970 |
| process:counter | cpu_seconds_total | 119.266 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45801472.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 143.234 |
| process:gateway | cpu_seconds_total | 1737.594 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51159040.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 129.191 |
| process:knowpost | cpu_seconds_total | 8016.062 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70320128.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.099 |
| process:relation | cpu_seconds_total | 18.297 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48459776.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 3.203 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38158336.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 62.839 |
| process:user-storage | cpu_seconds_total | 633.656 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55250944.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 159811114.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3676.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596907.000 |
| redis | keyspace_hits | 1155437677.000 |
| redis | keyspace_misses | 258601.000 |
| redis | net_input_bytes | 50232137809.000 |
| redis | net_output_bytes | 310153081558.000 |
| redis | ops_per_sec | 23550.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27055.000 |
| redis | used_memory_bytes | 103652192.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
