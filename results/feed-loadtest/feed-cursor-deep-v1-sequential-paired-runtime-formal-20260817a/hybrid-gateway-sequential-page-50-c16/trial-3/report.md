# Feed 压测报告：hybrid / gateway / sequential-page-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 开始时间：2026-08-17T20:00:45+08:00
- 采样时长：1m0.4318391s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 46350 | 46350 | 0 | 0 | 767.01 | 19.416 | 33.358 | 37.925 | 47.448 | 78.739 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：927/46350
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 221 | 0.005 |
| mysql | 4872 | 0.105 |
| redis | 97572 | 2.105 |
| relation | 221 | 0.005 |

- Cold compute：46350（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 24565500 | 530.000 |
| merge_candidates | 42624387 | 919.620 |
| redis_commands | 278100 | 6.000 |
| redis_members | 47224161 | 1018.860 |
| redis_roundtrips | 46350 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 46350 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 46350 | 6.483 |
| counter | 221 | 6.101 |
| hydrate | 46350 | 12.966 |
| inbox | 46350 | 6.481 |
| merge_dedup | 46350 | 0.103 |
| relation | 221 | 6.427 |
| route | 46350 | 0.188 |
| total | 46350 | 19.869 |

## Redis 本轮边界增量

- Commands：347988；input：883210616 bytes；output：6007289041 bytes
- Hits/Misses：24844511/5705；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：7264；safety epoch：3772 -> 3772

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.256 |
| client:loadtest | cpu_percent_total | 20.090 |
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
| docker:zg-canal | cpu_percent | 2.320 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.030 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.160 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 139.450 |
| docker:zg-kafka | memory_percent | 7.320 |
| docker:zg-kafka | pids | 115.000 |
| docker:zg-zk | cpu_percent | 46.660 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 570591.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 4.649 |
| process:counter | cpu_seconds_total | 330.719 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45285376.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 77.493 |
| process:gateway | cpu_seconds_total | 6689.469 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51462144.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 220.973 |
| process:knowpost | cpu_seconds_total | 19666.875 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 82968576.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.199 |
| process:relation | cpu_seconds_total | 53.250 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48689152.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 6.344 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38064128.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 34.081 |
| process:user-storage | cpu_seconds_total | 2394.766 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55500800.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 319305700.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3772.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2984607970.000 |
| redis | keyspace_misses | 772011.000 |
| redis | net_input_bytes | 124759875543.000 |
| redis | net_output_bytes | 730613698251.000 |
| redis | ops_per_sec | 7264.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34885.000 |
| redis | used_memory_bytes | 102836032.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
