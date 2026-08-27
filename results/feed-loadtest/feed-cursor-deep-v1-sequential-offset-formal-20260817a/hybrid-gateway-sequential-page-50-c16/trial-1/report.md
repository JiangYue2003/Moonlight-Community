# Feed 压测报告：hybrid / gateway / sequential-page-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:54:52+08:00
- 采样时长：1m0.5334257s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 52750 | 52750 | 0 | 0 | 871.46 | 17.223 | 29.120 | 33.196 | 42.315 | 86.904 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1055/52750
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 223 | 0.004 |
| mysql | 4621 | 0.088 |
| redis | 110121 | 2.088 |
| relation | 223 | 0.004 |

- Cold compute：52750（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 27957500 | 530.000 |
| merge_candidates | 48509955 | 919.620 |
| redis_commands | 316500 | 6.000 |
| redis_members | 53744865 | 1018.860 |
| redis_roundtrips | 52750 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 52750 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 52750 | 5.862 |
| counter | 223 | 5.541 |
| hydrate | 52750 | 11.247 |
| inbox | 52750 | 5.860 |
| merge_dedup | 52750 | 0.104 |
| relation | 223 | 5.877 |
| route | 52750 | 0.146 |
| total | 52750 | 17.489 |

## Redis 本轮边界增量

- Commands：394931；input：1004941862 bytes；output：6836930629 bytes
- Hits/Misses：28275296/5351；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：7343；safety epoch：3723 -> 3723

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.226 |
| client:loadtest | cpu_percent_total | 19.617 |
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
| docker:zg-canal | cpu_percent | 1.970 |
| docker:zg-canal | memory_percent | 4.260 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.180 |
| docker:zg-es | memory_percent | 12.350 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.770 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 112.240 |
| docker:zg-kafka | memory_percent | 7.320 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 47.240 |
| docker:zg-zk | memory_percent | 1.070 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 407031.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 3.873 |
| process:counter | cpu_seconds_total | 229.000 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46436352.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 75.837 |
| process:gateway | cpu_seconds_total | 5159.500 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50802688.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 276.326 |
| process:knowpost | cpu_seconds_total | 13814.812 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 81448960.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 38.688 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49020928.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.562 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 35.605 |
| process:user-storage | cpu_seconds_total | 1828.906 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 53227520.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 239156072.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3723.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 1935889271.000 |
| redis | keyspace_misses | 514356.000 |
| redis | net_input_bytes | 82831815795.000 |
| redis | net_output_bytes | 489149882058.000 |
| redis | ops_per_sec | 7343.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 30931.000 |
| redis | used_memory_bytes | 102478488.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
