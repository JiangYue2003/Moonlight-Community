# Feed 压测报告：hybrid / rpc / cursor-page5-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:55:28+08:00
- 采样时长：1m0.0135514s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 157492 | 157492 | 0 | 0 | 2624.75 | 5.657 | 8.706 | 10.056 | 12.576 | 23.213 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：61.8811ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 142 | 0.001 |
| redis | 157634 | 1.001 |
| relation | 240 | 0.002 |

- Cold compute：157492（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3307332 | 21.000 |
| merge_candidates | 23623800 | 150.000 |
| redis_commands | 2677364 | 17.000 |
| redis_members | 37168112 | 236.000 |
| redis_roundtrips | 314984 | 2.000 |
| tie_members | 19843992 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 157492 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 2.374 |
| cursor_decode | 157492 | 0.010 |
| cursor_seek | 157492 | 3.962 |
| hydrate | 157492 | 1.872 |
| merge_dedup | 157492 | 0.008 |
| relation | 240 | 2.752 |
| route | 157492 | 0.029 |
| total | 157492 | 5.883 |

## Redis 本轮边界增量

- Commands：2871369；input：407544870 bytes；output：2117047162 bytes
- Hits/Misses：5992073/153；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：52435；safety epoch：3633 -> 3633

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.742 |
| client:loadtest | cpu_percent_total | 43.870 |
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
| docker:zg-canal | cpu_percent | 2.210 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.740 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.030 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 127.740 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 45.590 |
| docker:zg-zk | memory_percent | 1.150 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 113628.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.972 |
| process:counter | cpu_seconds_total | 40.078 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46714880.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.967 |
| process:gateway | cpu_seconds_total | 258.188 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44965888.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 169.756 |
| process:knowpost | cpu_seconds_total | 2274.906 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 69509120.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.323 |
| process:relation | cpu_seconds_total | 4.609 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47837184.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 1.812 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 40022016.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.548 |
| process:user-storage | cpu_seconds_total | 95.797 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44179456.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 46954900.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3633.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597683.000 |
| redis | keyspace_hits | 353052607.000 |
| redis | keyspace_misses | 100631.000 |
| redis | net_input_bytes | 14776695128.000 |
| redis | net_output_bytes | 116286771945.000 |
| redis | ops_per_sec | 52435.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23767.000 |
| redis | used_memory_bytes | 103754312.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
