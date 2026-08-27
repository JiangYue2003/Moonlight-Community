# Feed 压测报告：hybrid / gateway / page20-c16

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:35:41+08:00
- 采样时长：10.0165373s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 7598 | 7598 | 0 | 0 | 758.55 | 20.697 | 29.756 | 38.478 | 48.810 | 77.862 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.005 |
| mysql | 232 | 0.031 |
| redis | 15428 | 2.031 |
| relation | 40 | 0.005 |

- Cold compute：7598（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3191160 | 420.000 |
| merge_candidates | 6389918 | 841.000 |
| redis_commands | 45588 | 6.000 |
| redis_members | 6997758 | 921.000 |
| redis_roundtrips | 7598 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 7598 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 7598 | 9.590 |
| counter | 40 | 5.622 |
| hydrate | 7598 | 9.991 |
| inbox | 7598 | 9.587 |
| merge_dedup | 7598 | 0.074 |
| relation | 40 | 5.861 |
| route | 7598 | 0.373 |
| total | 7598 | 20.128 |

## Redis 本轮边界增量

- Commands：56016；input：115355974 bytes；output：817607156 bytes
- Hits/Misses：3237732/279；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：6586；safety epoch：3613 -> 3613

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.180 |
| client:loadtest | cpu_percent_total | 18.875 |
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
| docker:zg-canal | cpu_percent | 2.040 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.100 |
| docker:zg-es | memory_percent | 12.080 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.310 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 12.780 |
| docker:zg-kafka | memory_percent | 7.090 |
| docker:zg-kafka | pids | 99.000 |
| docker:zg-zk | cpu_percent | 43.890 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 94315.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 29.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.646 |
| process:counter | cpu_seconds_total | 7.375 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45469696.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 86.718 |
| process:gateway | cpu_seconds_total | 132.391 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51470336.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 171.053 |
| process:knowpost | cpu_seconds_total | 142.906 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 69287936.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.870 |
| process:relation | cpu_seconds_total | 0.781 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 46219264.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.344 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 41656320.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 34.056 |
| process:user-storage | cpu_seconds_total | 48.797 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 51798016.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18827291.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3613.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597658.000 |
| redis | keyspace_hits | 99405230.000 |
| redis | keyspace_misses | 70131.000 |
| redis | net_input_bytes | 4725226171.000 |
| redis | net_output_bytes | 26592359824.000 |
| redis | ops_per_sec | 6586.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22530.000 |
| redis | used_memory_bytes | 102487800.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
