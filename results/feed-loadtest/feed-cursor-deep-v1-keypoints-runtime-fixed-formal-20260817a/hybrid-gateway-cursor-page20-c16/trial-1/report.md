# Feed 压测报告：hybrid / gateway / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:37:20+08:00
- 采样时长：1m0.0133885s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 101899 | 101899 | 0 | 0 | 1698.16 | 9.016 | 13.693 | 15.719 | 20.340 | 35.884 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：434.4858ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 138 | 0.001 |
| redis | 102037 | 1.001 |
| relation | 240 | 0.002 |

- Cold compute：101899（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2139879 | 21.000 |
| merge_candidates | 6521536 | 64.000 |
| redis_commands | 1426586 | 14.000 |
| redis_members | 6827233 | 67.000 |
| redis_roundtrips | 203798 | 2.000 |
| tie_members | 305697 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 101899 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 3.258 |
| cursor_decode | 101899 | 0.011 |
| cursor_seek | 101899 | 5.333 |
| hydrate | 101899 | 2.825 |
| merge_dedup | 101899 | 0.007 |
| relation | 240 | 4.043 |
| route | 101899 | 0.063 |
| total | 101899 | 8.243 |

## Redis 本轮边界增量

- Commands：1559602；input：233551176 bytes；output：644929466 bytes
- Hits/Misses：3573605/378；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：31845；safety epoch：3755 -> 3755

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.246 |
| client:loadtest | cpu_percent_total | 51.942 |
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
| docker:zg-es | cpu_percent | 0.470 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.770 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 154.600 |
| docker:zg-kafka | memory_percent | 7.620 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 45.440 |
| docker:zg-zk | memory_percent | 1.340 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 508133.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.071 |
| process:counter | cpu_seconds_total | 284.875 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47489024.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 221.387 |
| process:gateway | cpu_seconds_total | 5741.984 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51503104.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 163.999 |
| process:knowpost | cpu_seconds_total | 17423.281 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71450624.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.065 |
| process:relation | cpu_seconds_total | 46.797 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48476160.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.548 |
| process:search | cpu_seconds_total | 5.688 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 79.015 |
| process:user-storage | cpu_seconds_total | 2055.484 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 57716736.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 289946528.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3755.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596155.000 |
| redis | keyspace_hits | 2611073597.000 |
| redis | keyspace_misses | 673967.000 |
| redis | net_input_bytes | 109751030540.000 |
| redis | net_output_bytes | 646255269262.000 |
| redis | ops_per_sec | 31845.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 33479.000 |
| redis | used_memory_bytes | 102105224.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
