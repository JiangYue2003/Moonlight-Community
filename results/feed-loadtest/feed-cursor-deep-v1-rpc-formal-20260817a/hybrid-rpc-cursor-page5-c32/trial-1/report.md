# Feed 压测报告：hybrid / rpc / cursor-page5-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:59:12+08:00
- 采样时长：1m0.0156542s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 164558 | 164558 | 0 | 0 | 2742.44 | 11.031 | 15.095 | 17.925 | 23.137 | 40.124 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：81.3766ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 219 | 0.001 |
| redis | 164777 | 1.001 |
| relation | 240 | 0.001 |

- Cold compute：164558（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3455718 | 21.000 |
| merge_candidates | 24683700 | 150.000 |
| redis_commands | 2797486 | 17.000 |
| redis_members | 38835688 | 236.000 |
| redis_roundtrips | 329116 | 2.000 |
| tie_members | 20734308 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 164558 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 3.923 |
| cursor_decode | 164558 | 0.010 |
| cursor_seek | 164558 | 7.629 |
| hydrate | 164558 | 3.747 |
| merge_dedup | 164558 | 0.009 |
| relation | 240 | 4.404 |
| route | 164558 | 0.058 |
| total | 164558 | 11.454 |

## Redis 本轮边界增量

- Commands：2986620；input：424893578 bytes；output：2211747327 bytes
- Hits/Misses：6260489/220；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：53819；safety epoch：3636 -> 3636

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.768 |
| client:loadtest | cpu_percent_total | 44.285 |
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
| docker:zg-canal | cpu_percent | 2.110 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.180 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.730 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 144.530 |
| docker:zg-kafka | memory_percent | 7.550 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 26.710 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 115412.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 25.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.199 |
| process:counter | cpu_seconds_total | 46.047 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46505984.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 258.219 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45256704.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 174.195 |
| process:knowpost | cpu_seconds_total | 2571.547 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71192576.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.549 |
| process:relation | cpu_seconds_total | 5.297 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48783360.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.875 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38346752.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.550 |
| process:user-storage | cpu_seconds_total | 96.156 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44072960.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 57432255.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3636.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597390.000 |
| redis | keyspace_hits | 374942679.000 |
| redis | keyspace_misses | 101469.000 |
| redis | net_input_bytes | 16264701248.000 |
| redis | net_output_bytes | 124020653948.000 |
| redis | ops_per_sec | 53819.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23992.000 |
| redis | used_memory_bytes | 103524184.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
