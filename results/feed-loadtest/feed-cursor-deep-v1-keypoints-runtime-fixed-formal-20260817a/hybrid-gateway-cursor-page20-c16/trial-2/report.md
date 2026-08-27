# Feed 压测报告：hybrid / gateway / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:38:37+08:00
- 采样时长：1m0.0097608s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 73986 | 73986 | 0 | 0 | 1233.01 | 12.714 | 19.222 | 22.187 | 28.588 | 56.590 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：395.6889ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 107 | 0.001 |
| redis | 74093 | 1.001 |
| relation | 240 | 0.003 |

- Cold compute：73986（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1553706 | 21.000 |
| merge_candidates | 4735104 | 64.000 |
| redis_commands | 1035804 | 14.000 |
| redis_members | 4957062 | 67.000 |
| redis_roundtrips | 147972 | 2.000 |
| tie_members | 221958 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 73986 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 4.660 |
| cursor_decode | 73986 | 0.012 |
| cursor_seek | 73986 | 7.369 |
| hydrate | 73986 | 3.902 |
| merge_dedup | 73986 | 0.007 |
| relation | 240 | 5.577 |
| route | 73986 | 0.119 |
| total | 73986 | 11.413 |

## Redis 本轮边界增量

- Commands：1133888；input：169606101 bytes；output：468329697 bytes
- Hits/Misses：2596638/356；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：23033；safety epoch：3756 -> 3756

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.985 |
| client:loadtest | cpu_percent_total | 47.753 |
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
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.410 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.380 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 270.260 |
| docker:zg-kafka | memory_percent | 8.130 |
| docker:zg-kafka | pids | 147.000 |
| docker:zg-zk | cpu_percent | 44.940 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 508695.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 13.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.065 |
| process:counter | cpu_seconds_total | 288.594 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46571520.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 213.286 |
| process:gateway | cpu_seconds_total | 5863.750 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51613696.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 161.141 |
| process:knowpost | cpu_seconds_total | 17520.266 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71659520.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.423 |
| process:relation | cpu_seconds_total | 47.469 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48615424.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.548 |
| process:search | cpu_seconds_total | 5.750 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 70.568 |
| process:user-storage | cpu_seconds_total | 2095.594 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55451648.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 291313208.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3756.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596106.000 |
| redis | keyspace_hits | 2614191907.000 |
| redis | keyspace_misses | 674745.000 |
| redis | net_input_bytes | 109955089356.000 |
| redis | net_output_bytes | 646819967107.000 |
| redis | ops_per_sec | 23033.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 33556.000 |
| redis | used_memory_bytes | 101932000.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
