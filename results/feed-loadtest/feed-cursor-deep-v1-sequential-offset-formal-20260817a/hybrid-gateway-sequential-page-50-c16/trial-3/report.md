# Feed 压测报告：hybrid / gateway / sequential-page-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:57:24+08:00
- 采样时长：1m0.8027544s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 47200 | 47200 | 0 | 0 | 776.33 | 19.282 | 33.021 | 37.328 | 46.550 | 85.712 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：944/47200
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 225 | 0.005 |
| mysql | 5198 | 0.110 |
| redis | 99598 | 2.110 |
| relation | 225 | 0.005 |

- Cold compute：47200（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 25016000 | 530.000 |
| merge_candidates | 43406064 | 919.620 |
| redis_commands | 283200 | 6.000 |
| redis_members | 48090192 | 1018.860 |
| redis_roundtrips | 47200 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 47200 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 47200 | 6.477 |
| counter | 225 | 5.662 |
| hydrate | 47200 | 12.770 |
| inbox | 47200 | 6.474 |
| merge_dedup | 47200 | 0.101 |
| relation | 225 | 6.125 |
| route | 47200 | 0.178 |
| total | 47200 | 19.654 |

## Redis 本轮边界增量

- Commands：354727；input：899492613 bytes；output：6117398879 bytes
- Hits/Misses：25299754/6184；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：7067；safety epoch：3725 -> 3725

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.304 |
| client:loadtest | cpu_percent_total | 20.867 |
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
| docker:zg-canal | memory_percent | 4.260 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.450 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.830 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 142.970 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 45.760 |
| docker:zg-zk | memory_percent | 1.070 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 419588.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 9.895 |
| process:counter | cpu_seconds_total | 231.453 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46182400.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 86.441 |
| process:gateway | cpu_seconds_total | 5243.844 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51240960.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 223.133 |
| process:knowpost | cpu_seconds_total | 14062.953 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 80314368.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 39.312 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48435200.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 4.609 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 36.355 |
| process:user-storage | cpu_seconds_total | 1861.719 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 52604928.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 239992314.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3725.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 1995065444.000 |
| redis | keyspace_misses | 528238.000 |
| redis | net_input_bytes | 84936087318.000 |
| redis | net_output_bytes | 503458583542.000 |
| redis | ops_per_sec | 7067.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31084.000 |
| redis | used_memory_bytes | 102534120.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
