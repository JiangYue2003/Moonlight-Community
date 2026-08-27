# Feed 压测报告：hybrid / gateway / page20-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:01:32+08:00
- 采样时长：1m0.0233236s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 54524 | 54524 | 0 | 0 | 908.44 | 31.649 | 50.147 | 58.437 | 83.734 | 170.163 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.004 |
| mysql | 4420 | 0.081 |
| redis | 113468 | 2.081 |
| relation | 240 | 0.004 |

- Cold compute：54524（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 22900080 | 420.000 |
| merge_candidates | 45854684 | 841.000 |
| redis_commands | 327144 | 6.000 |
| redis_members | 50216604 | 921.000 |
| redis_roundtrips | 54524 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 54524 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 54524 | 16.600 |
| counter | 240 | 14.616 |
| hydrate | 54524 | 16.795 |
| inbox | 54524 | 16.598 |
| merge_dedup | 54524 | 0.080 |
| relation | 240 | 14.929 |
| route | 54524 | 0.727 |
| total | 54524 | 34.315 |

## Redis 本轮边界增量

- Commands：398565；input：828121543 bytes；output：5866552837 bytes
- Hits/Misses：23229145/5520；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：8199；safety epoch：3685 -> 3685

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.272 |
| client:loadtest | cpu_percent_total | 20.357 |
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
| docker:zg-canal | cpu_percent | 2.290 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.480 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.430 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 158.740 |
| docker:zg-kafka | memory_percent | 7.580 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 35.850 |
| docker:zg-zk | memory_percent | 1.210 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 256745.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 2.323 |
| process:counter | cpu_seconds_total | 131.719 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47386624.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 83.528 |
| process:gateway | cpu_seconds_total | 2347.031 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51884032.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 232.127 |
| process:knowpost | cpu_seconds_total | 8993.406 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 79282176.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.550 |
| process:relation | cpu_seconds_total | 21.969 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49594368.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 3.469 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38162432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 44.122 |
| process:user-storage | cpu_seconds_total | 863.875 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 52981760.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 169031534.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3685.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596519.000 |
| redis | keyspace_hits | 1315279482.000 |
| redis | keyspace_misses | 287320.000 |
| redis | net_input_bytes | 56381988433.000 |
| redis | net_output_bytes | 351921658982.000 |
| redis | ops_per_sec | 8199.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27731.000 |
| redis | used_memory_bytes | 104059928.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
