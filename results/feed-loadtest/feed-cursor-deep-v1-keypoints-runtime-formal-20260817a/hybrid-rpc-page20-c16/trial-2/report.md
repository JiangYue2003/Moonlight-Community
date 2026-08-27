# Feed 压测报告：hybrid / rpc / page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-formal-20260817a`
- 开始时间：2026-08-17T19:13:09+08:00
- 采样时长：1m0.0150644s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 70035 | 70035 | 0 | 0 | 1167.06 | 11.970 | 20.952 | 22.945 | 27.780 | 46.988 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Redis 本轮边界增量

- Commands：511763；input：1063081408 bytes；output：7536177416 bytes
- Hits/Misses：29839676/2708；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：9174；safety epoch：3738 -> 3738

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.632 |
| client:loadtest | cpu_percent_total | 26.113 |
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
| docker:zg-canal | cpu_percent | 0.140 |
| docker:zg-canal | memory_percent | 4.270 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.720 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.460 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 131.480 |
| docker:zg-kafka | memory_percent | 7.610 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 44.430 |
| docker:zg-zk | memory_percent | 1.330 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 460608.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.692 |
| process:counter | cpu_seconds_total | 250.719 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46981120.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 5478.406 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44998656.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 288.150 |
| process:knowpost | cpu_seconds_total | 14841.281 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 75538432.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.096 |
| process:relation | cpu_seconds_total | 42.234 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48738304.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.556 |
| process:search | cpu_seconds_total | 5.047 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38027264.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.003 |
| process:user-storage | cpu_seconds_total | 1947.125 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 46403584.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 244371944.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3738.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596016.000 |
| redis | keyspace_hits | 2149568351.000 |
| redis | keyspace_misses | 590470.000 |
| redis | net_input_bytes | 90567059202.000 |
| redis | net_output_bytes | 542922818940.000 |
| redis | ops_per_sec | 9174.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 32028.000 |
| redis | used_memory_bytes | 102062184.000 |

## 缺失指标

- feed_metrics

## 说明

- SLA values are reference lines, not pass/fail gates.
