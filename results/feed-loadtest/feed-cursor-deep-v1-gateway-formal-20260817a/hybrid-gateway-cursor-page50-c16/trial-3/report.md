# Feed 压测报告：hybrid / gateway / cursor-page50-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:21:49+08:00
- 采样时长：1m0.0143363s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 109551 | 109551 | 0 | 0 | 1825.68 | 8.582 | 12.040 | 13.724 | 17.387 | 31.654 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：1.310126s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 125 | 0.001 |
| redis | 109676 | 1.001 |
| relation | 240 | 0.002 |

- Cold compute：109551（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2300571 | 21.000 |
| merge_candidates | 2410122 | 22.000 |
| redis_commands | 1424163 | 13.000 |
| redis_members | 2629224 | 24.000 |
| redis_roundtrips | 219102 | 2.000 |
| tie_members | 219102 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 109551 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 3.165 |
| cursor_decode | 109551 | 0.011 |
| cursor_seek | 109551 | 4.702 |
| hydrate | 109551 | 2.553 |
| merge_dedup | 109551 | 0.005 |
| relation | 240 | 4.077 |
| route | 109551 | 0.055 |
| total | 109551 | 7.333 |

## Redis 本轮边界增量

- Commands：1568097；input：240217498 bytes；output：492567898 bytes
- Hits/Misses：3731894/367；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：29507；safety epoch：3701 -> 3701

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.315 |
| client:loadtest | cpu_percent_total | 69.046 |
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
| docker:zg-canal | memory_percent | 4.230 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.650 |
| docker:zg-es | memory_percent | 12.350 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.970 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 155.100 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 45.820 |
| docker:zg-zk | memory_percent | 1.230 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 324642.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 12.453 |
| process:counter | cpu_seconds_total | 172.375 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47226880.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 258.538 |
| process:gateway | cpu_seconds_total | 3828.484 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51482624.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 205.060 |
| process:knowpost | cpu_seconds_total | 10822.984 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70750208.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.335 |
| process:relation | cpu_seconds_total | 29.047 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48205824.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.771 |
| process:search | cpu_seconds_total | 3.938 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38166528.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 101.352 |
| process:user-storage | cpu_seconds_total | 1370.609 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 61935616.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 188127714.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3701.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596574.000 |
| redis | keyspace_hits | 1563493075.000 |
| redis | keyspace_misses | 398850.000 |
| redis | net_input_bytes | 66358425901.000 |
| redis | net_output_bytes | 405883457380.000 |
| redis | ops_per_sec | 29507.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28949.000 |
| redis | used_memory_bytes | 101911376.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
