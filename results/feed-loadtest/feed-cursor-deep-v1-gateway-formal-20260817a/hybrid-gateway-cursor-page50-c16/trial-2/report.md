# Feed 压测报告：hybrid / gateway / cursor-page50-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:20:32+08:00
- 采样时长：1m0.0159333s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 121019 | 121019 | 0 | 0 | 2016.76 | 7.571 | 11.121 | 12.620 | 16.009 | 29.058 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：1.0824031s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 126 | 0.001 |
| redis | 121145 | 1.001 |
| relation | 240 | 0.002 |

- Cold compute：121019（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2541399 | 21.000 |
| merge_candidates | 2662418 | 22.000 |
| redis_commands | 1573247 | 13.000 |
| redis_members | 2904456 | 24.000 |
| redis_roundtrips | 242038 | 2.000 |
| tie_members | 242038 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 121019 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 2.945 |
| cursor_decode | 121019 | 0.011 |
| cursor_seek | 121019 | 4.237 |
| hydrate | 121019 | 2.313 |
| merge_dedup | 121019 | 0.005 |
| relation | 240 | 3.868 |
| route | 121019 | 0.046 |
| total | 121019 | 6.619 |

## Redis 本轮边界增量

- Commands：1729650；input：265208172 bytes；output：544067781 bytes
- Hits/Misses：4121809/366；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：33670；safety epoch：3700 -> 3700

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.828 |
| client:loadtest | cpu_percent_total | 77.245 |
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
| docker:zg-canal | cpu_percent | 0.190 |
| docker:zg-canal | memory_percent | 4.220 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.740 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.930 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 133.740 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 54.470 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 323954.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 17.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 10.580 |
| process:counter | cpu_seconds_total | 167.812 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46460928.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 266.408 |
| process:gateway | cpu_seconds_total | 3668.688 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51634176.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 211.600 |
| process:knowpost | cpu_seconds_total | 10695.078 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70639616.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.261 |
| process:relation | cpu_seconds_total | 28.641 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47521792.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.763 |
| process:search | cpu_seconds_total | 3.922 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38166528.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 93.528 |
| process:user-storage | cpu_seconds_total | 1317.266 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 58425344.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 186248577.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3700.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596577.000 |
| redis | keyspace_hits | 1559035634.000 |
| redis | keyspace_misses | 397464.000 |
| redis | net_input_bytes | 66070936996.000 |
| redis | net_output_bytes | 405291240926.000 |
| redis | ops_per_sec | 33670.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28871.000 |
| redis | used_memory_bytes | 101788008.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
