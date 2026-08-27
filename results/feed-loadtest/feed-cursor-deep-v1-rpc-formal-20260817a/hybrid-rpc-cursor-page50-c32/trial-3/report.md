# Feed 压测报告：hybrid / rpc / cursor-page50-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:32:01+08:00
- 采样时长：1m0.030818s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 451906 | 451906 | 0 | 0 | 7531.44 | 3.930 | 5.828 | 6.678 | 8.533 | 19.719 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：796.3994ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 275 | 0.001 |
| redis | 452181 | 1.001 |
| relation | 240 | 0.001 |

- Cold compute：451906（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9490026 | 21.000 |
| merge_candidates | 9941932 | 22.000 |
| redis_commands | 5874778 | 13.000 |
| redis_members | 10845744 | 24.000 |
| redis_roundtrips | 903812 | 2.000 |
| tie_members | 903812 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 451906 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 1.655 |
| cursor_decode | 451906 | 0.010 |
| cursor_seek | 451906 | 2.638 |
| hydrate | 451906 | 1.324 |
| merge_dedup | 451906 | 0.005 |
| relation | 240 | 2.430 |
| route | 451906 | 0.011 |
| total | 451906 | 3.994 |

## Redis 本轮边界增量

- Commands：6363380；input：984344174 bytes；output：2029417357 bytes
- Hits/Misses：15372046/289；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：117062；safety epoch：3662 -> 3662

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.898 |
| client:loadtest | cpu_percent_total | 142.375 |
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
| docker:zg-canal | cpu_percent | 2.170 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.350 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.120 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 108.840 |
| docker:zg-kafka | memory_percent | 7.490 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 47.730 |
| docker:zg-zk | memory_percent | 1.230 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 217368.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.189 |
| process:counter | cpu_seconds_total | 89.578 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45690880.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 258.453 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43900928.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 369.523 |
| process:knowpost | cpu_seconds_total | 6608.859 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72237056.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.825 |
| process:relation | cpu_seconds_total | 12.312 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48582656.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 2.688 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38125568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 4.902 |
| process:user-storage | cpu_seconds_total | 98.781 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44044288.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 146760632.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3662.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597777.000 |
| redis | keyspace_hits | 1037235244.000 |
| redis | keyspace_misses | 246138.000 |
| redis | net_input_bytes | 45527193926.000 |
| redis | net_output_bytes | 267122675395.000 |
| redis | ops_per_sec | 117062.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 25960.000 |
| redis | used_memory_bytes | 103507384.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
