# Feed 压测报告：hybrid / rpc / cursor-page20-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:15:31+08:00
- 采样时长：1m0.0242507s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 287373 | 287373 | 0 | 0 | 4789.16 | 6.320 | 8.812 | 10.445 | 13.359 | 30.994 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：357.855ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 255 | 0.001 |
| redis | 287628 | 1.001 |
| relation | 240 | 0.001 |

- Cold compute：287373（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6034833 | 21.000 |
| merge_candidates | 18391872 | 64.000 |
| redis_commands | 4023222 | 14.000 |
| redis_members | 19253991 | 67.000 |
| redis_roundtrips | 574746 | 2.000 |
| tie_members | 862119 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 287373 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 2.801 |
| cursor_decode | 287373 | 0.010 |
| cursor_seek | 287373 | 4.293 |
| hydrate | 287373 | 2.132 |
| merge_dedup | 287373 | 0.007 |
| relation | 240 | 3.098 |
| route | 287373 | 0.024 |
| total | 287373 | 6.469 |

## Redis 本轮边界增量

- Commands：4347057；input：655529458 bytes；output：1817580518 bytes
- Hits/Misses：10065330/255；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：79083；safety epoch：3649 -> 3649

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.834 |
| client:loadtest | cpu_percent_total | 77.339 |
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
| docker:zg-canal | cpu_percent | 0.180 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.710 |
| docker:zg-es | memory_percent | 12.220 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.430 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 128.070 |
| docker:zg-kafka | memory_percent | 7.260 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 44.770 |
| docker:zg-zk | memory_percent | 1.150 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 147390.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.983 |
| process:counter | cpu_seconds_total | 66.562 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46104576.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 258.359 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44044288.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 250.798 |
| process:knowpost | cpu_seconds_total | 4356.062 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70688768.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.549 |
| process:relation | cpu_seconds_total | 8.453 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48181248.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.387 |
| process:search | cpu_seconds_total | 2.266 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38113280.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.549 |
| process:user-storage | cpu_seconds_total | 97.094 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43827200.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 93841599.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3649.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597196.000 |
| redis | keyspace_hits | 661006464.000 |
| redis | keyspace_misses | 140163.000 |
| redis | net_input_bytes | 28693381580.000 |
| redis | net_output_bytes | 193512964507.000 |
| redis | ops_per_sec | 79083.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 24970.000 |
| redis | used_memory_bytes | 104161648.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
