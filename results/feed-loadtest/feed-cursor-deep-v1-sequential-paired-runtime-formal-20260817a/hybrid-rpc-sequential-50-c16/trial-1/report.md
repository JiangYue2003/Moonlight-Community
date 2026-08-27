# Feed 压测报告：hybrid / rpc / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 开始时间：2026-08-17T19:54:26+08:00
- 采样时长：1m0.0788074s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 296350 | 296350 | 0 | 0 | 4934.19 | 3.102 | 4.782 | 5.472 | 7.076 | 14.629 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：5927/296350
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 1002 | 0.003 |
| redis | 303279 | 1.023 |
| relation | 240 | 0.001 |

- Cold compute：296350（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6335963 | 21.380 |
| merge_candidates | 17206081 | 58.060 |
| redis_commands | 4178535 | 14.100 |
| redis_members | 21621696 | 72.960 |
| redis_roundtrips | 586773 | 1.980 |
| tie_members | 5215760 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 296350 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 5927 | 1.180 |
| counter | 240 | 1.453 |
| cursor_decode | 290423 | 0.010 |
| cursor_seek | 290423 | 2.040 |
| hydrate | 296350 | 0.985 |
| inbox | 5927 | 1.179 |
| merge_dedup | 296350 | 0.006 |
| relation | 240 | 2.103 |
| route | 296350 | 0.010 |
| total | 296350 | 3.038 |

## Redis 本轮边界增量

- Commands：4511948；input：681337854 bytes；output：1963797463 bytes
- Hits/Misses：10520575/1348；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：83267；safety epoch：3767 -> 3767

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.262 |
| client:loadtest | cpu_percent_total | 68.192 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.330 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.410 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 110.960 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 49.780 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 548118.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.975 |
| process:counter | cpu_seconds_total | 321.641 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46616576.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 6559.031 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45187072.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 247.033 |
| process:knowpost | cpu_seconds_total | 18977.203 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71512064.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.326 |
| process:relation | cpu_seconds_total | 51.953 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49131520.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.551 |
| process:search | cpu_seconds_total | 6.234 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38051840.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.550 |
| process:user-storage | cpu_seconds_total | 2342.562 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 47407104.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 307172016.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3767.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596567.000 |
| redis | keyspace_hits | 2863091408.000 |
| redis | keyspace_misses | 747472.000 |
| redis | net_input_bytes | 119705449468.000 |
| redis | net_output_bytes | 702615574223.000 |
| redis | ops_per_sec | 83267.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34505.000 |
| redis | used_memory_bytes | 102214568.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
