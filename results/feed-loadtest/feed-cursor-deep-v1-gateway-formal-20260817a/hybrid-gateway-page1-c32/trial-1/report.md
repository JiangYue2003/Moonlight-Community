# Feed 压测报告：hybrid / gateway / page1-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:37:44+08:00
- 采样时长：1m0.0192037s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 118068 | 118068 | 0 | 0 | 1967.44 | 15.754 | 21.265 | 24.136 | 34.673 | 70.748 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 496 | 0.004 |
| redis | 236632 | 2.004 |
| relation | 240 | 0.002 |

- Cold compute：118068（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 4722720 | 40.000 |
| merge_candidates | 9563508 | 81.000 |
| redis_commands | 708408 | 6.000 |
| redis_members | 29044728 | 246.000 |
| redis_roundtrips | 118068 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 118068 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 118068 | 7.559 |
| counter | 240 | 7.040 |
| hydrate | 118068 | 7.276 |
| inbox | 118068 | 7.558 |
| merge_dedup | 118068 | 0.011 |
| relation | 240 | 8.030 |
| route | 118068 | 0.164 |
| total | 118068 | 15.033 |

## Redis 本轮边界增量

- Commands：843176；input：219877349 bytes；output：1998910149 bytes
- Hits/Misses：5438074/529；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：17483；safety epoch：3666 -> 3666

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.614 |
| client:loadtest | cpu_percent_total | 57.820 |
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
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.640 |
| docker:zg-es | memory_percent | 12.330 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.530 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 152.240 |
| docker:zg-kafka | memory_percent | 7.550 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 3.370 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 85.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 220400.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 29.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.198 |
| process:counter | cpu_seconds_total | 101.812 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45367296.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 255.616 |
| process:gateway | cpu_seconds_total | 816.656 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52690944.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 201.296 |
| process:knowpost | cpu_seconds_total | 7036.125 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71856128.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.419 |
| process:relation | cpu_seconds_total | 14.000 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49213440.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.549 |
| process:search | cpu_seconds_total | 2.922 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38109184.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 85.155 |
| process:user-storage | cpu_seconds_total | 289.641 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 57139200.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 150652057.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3666.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597071.000 |
| redis | keyspace_hits | 1061889559.000 |
| redis | keyspace_misses | 248370.000 |
| redis | net_input_bytes | 46528793077.000 |
| redis | net_output_bytes | 276185915494.000 |
| redis | ops_per_sec | 17483.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26303.000 |
| redis | used_memory_bytes | 103268824.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
