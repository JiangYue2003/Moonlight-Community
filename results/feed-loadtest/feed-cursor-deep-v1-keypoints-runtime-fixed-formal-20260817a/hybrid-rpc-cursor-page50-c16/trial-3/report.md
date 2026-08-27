# Feed 压测报告：hybrid / rpc / cursor-page50-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:32:17+08:00
- 采样时长：1m0.0267715s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 464187 | 464187 | 0 | 0 | 7736.26 | 2.076 | 2.889 | 3.269 | 4.252 | 10.613 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：729.411ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 151 | 0.000 |
| redis | 464338 | 1.000 |
| relation | 240 | 0.001 |

- Cold compute：464187（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9747927 | 21.000 |
| merge_candidates | 10212114 | 22.000 |
| redis_commands | 6034431 | 13.000 |
| redis_members | 11140488 | 24.000 |
| redis_roundtrips | 928374 | 2.000 |
| tie_members | 928374 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 464187 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 1.027 |
| cursor_decode | 464187 | 0.009 |
| cursor_seek | 464187 | 1.189 |
| hydrate | 464187 | 0.619 |
| merge_dedup | 464187 | 0.004 |
| relation | 240 | 1.682 |
| route | 464187 | 0.007 |
| total | 464187 | 1.835 |

## Redis 本轮边界增量

- Commands：6534779；input：1010974088 bytes；output：2084557655 bytes
- Hits/Misses：15789481/407；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：118001；safety epoch：3751 -> 3751

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.437 |
| client:loadtest | cpu_percent_total | 134.992 |
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
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.600 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.020 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 170.600 |
| docker:zg-kafka | memory_percent | 7.620 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 49.350 |
| docker:zg-zk | memory_percent | 1.390 |
| docker:zg-zk | pids | 107.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 497595.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 16.854 |
| process:counter | cpu_seconds_total | 277.328 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45568000.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 5478.531 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43823104.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 332.715 |
| process:knowpost | cpu_seconds_total | 16899.609 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 73007104.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.320 |
| process:relation | cpu_seconds_total | 45.672 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49266688.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 5.594 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.549 |
| process:user-storage | cpu_seconds_total | 1948.391 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 46039040.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 286330830.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3751.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596558.000 |
| redis | keyspace_hits | 2511317253.000 |
| redis | keyspace_misses | 658959.000 |
| redis | net_input_bytes | 106061056178.000 |
| redis | net_output_bytes | 621382267498.000 |
| redis | ops_per_sec | 118001.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 33176.000 |
| redis | used_memory_bytes | 101797024.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
