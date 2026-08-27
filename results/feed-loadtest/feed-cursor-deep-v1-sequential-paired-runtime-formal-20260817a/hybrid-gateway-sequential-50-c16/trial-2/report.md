# Feed 压测报告：hybrid / gateway / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 开始时间：2026-08-17T20:03:18+08:00
- 采样时长：1m0.2745344s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 108300 | 108300 | 0 | 0 | 1797.02 | 8.549 | 12.370 | 14.378 | 18.533 | 33.040 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：2166/108300
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 237 | 0.002 |
| mysql | 1126 | 0.010 |
| redis | 111592 | 1.030 |
| relation | 237 | 0.002 |

- Cold compute：108300（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2315454 | 21.380 |
| merge_candidates | 6287898 | 58.060 |
| redis_commands | 1527030 | 14.100 |
| redis_members | 7901568 | 72.960 |
| redis_roundtrips | 214434 | 1.980 |
| tie_members | 1906080 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 108300 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2166 | 2.738 |
| counter | 237 | 3.173 |
| cursor_decode | 106134 | 0.011 |
| cursor_seek | 106134 | 4.977 |
| hydrate | 108300 | 2.612 |
| inbox | 2166 | 2.736 |
| merge_dedup | 108300 | 0.006 |
| relation | 237 | 4.070 |
| route | 108300 | 0.046 |
| total | 108300 | 7.613 |

## Redis 本轮边界增量

- Commands：1668736；input：250355790 bytes；output：717996610 bytes
- Hits/Misses：3848202/1476；run hit rate：99.96%
- Evicted/Rejected：0/0；ops/s max：31175；safety epoch：3774 -> 3774

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.574 |
| client:loadtest | cpu_percent_total | 57.186 |
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
| docker:zg-canal | cpu_percent | 2.250 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.730 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.960 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 176.510 |
| docker:zg-kafka | memory_percent | 7.650 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 28.300 |
| docker:zg-zk | memory_percent | 1.180 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 574141.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.064 |
| process:counter | cpu_seconds_total | 337.656 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47378432.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 227.593 |
| process:gateway | cpu_seconds_total | 6980.344 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51544064.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 183.468 |
| process:knowpost | cpu_seconds_total | 19889.594 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72683520.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.775 |
| process:relation | cpu_seconds_total | 54.109 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48709632.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 6.359 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38064128.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 92.235 |
| process:user-storage | cpu_seconds_total | 2492.453 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 57479168.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 323758897.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3774.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596562.000 |
| redis | keyspace_hits | 2994875024.000 |
| redis | keyspace_misses | 775673.000 |
| redis | net_input_bytes | 125428018627.000 |
| redis | net_output_bytes | 732529658766.000 |
| redis | ops_per_sec | 31175.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35037.000 |
| redis | used_memory_bytes | 102376744.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
