# Feed 压测报告：hybrid / rpc / sequential-page-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:47:13+08:00
- 采样时长：1m0.298593s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 59200 | 59200 | 0 | 0 | 981.85 | 16.255 | 25.724 | 28.641 | 34.647 | 58.569 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1184/59200
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 222 | 0.004 |
| mysql | 2475 | 0.042 |
| redis | 120875 | 2.042 |
| relation | 222 | 0.004 |

- Cold compute：59200（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 31376000 | 530.000 |
| merge_candidates | 54441504 | 919.620 |
| redis_commands | 355200 | 6.000 |
| redis_members | 60316512 | 1018.860 |
| redis_roundtrips | 59200 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 59200 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 59200 | 5.280 |
| counter | 222 | 5.348 |
| hydrate | 59200 | 10.371 |
| inbox | 59200 | 5.277 |
| merge_dedup | 59200 | 0.104 |
| relation | 222 | 5.362 |
| route | 59200 | 0.122 |
| total | 59200 | 16.005 |

## Redis 本轮边界增量

- Commands：438625；input：1127081590 bytes；output：7673288693 bytes
- Hits/Misses：31734642/3180；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：7902；safety epoch：3717 -> 3717

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.443 |
| client:loadtest | cpu_percent_total | 23.088 |
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
| docker:zg-canal | cpu_percent | 2.220 |
| docker:zg-canal | memory_percent | 4.250 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.690 |
| docker:zg-es | memory_percent | 12.350 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.270 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 143.610 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 44.730 |
| docker:zg-zk | memory_percent | 1.060 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 354989.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 5.426 |
| process:counter | cpu_seconds_total | 221.875 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46039040.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 5120.078 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45793280.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 253.649 |
| process:knowpost | cpu_seconds_total | 12910.094 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 81719296.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.323 |
| process:relation | cpu_seconds_total | 36.594 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48586752.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.771 |
| process:search | cpu_seconds_total | 4.469 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38039552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 1810.828 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 48340992.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 236015044.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3717.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596570.000 |
| redis | keyspace_hits | 1711318696.000 |
| redis | keyspace_misses | 451333.000 |
| redis | net_input_bytes | 74845715763.000 |
| redis | net_output_bytes | 434846947838.000 |
| redis | ops_per_sec | 7902.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 30472.000 |
| redis | used_memory_bytes | 102488392.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
