# Feed 压测报告：hybrid / gateway / cursor-page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:51:31+08:00
- 采样时长：1m0.0098535s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 78500 | 78500 | 0 | 0 | 1308.25 | 11.522 | 17.270 | 20.254 | 25.700 | 51.968 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：117.0913ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 150 | 0.002 |
| redis | 78650 | 1.002 |
| relation | 240 | 0.003 |

- Cold compute：78500（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1648500 | 21.000 |
| merge_candidates | 11775000 | 150.000 |
| redis_commands | 1334500 | 17.000 |
| redis_members | 18526000 | 236.000 |
| redis_roundtrips | 157000 | 2.000 |
| tie_members | 9891000 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 78500 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 4.101 |
| cursor_decode | 78500 | 0.011 |
| cursor_seek | 78500 | 7.333 |
| hydrate | 78500 | 3.654 |
| merge_dedup | 78500 | 0.009 |
| relation | 240 | 5.016 |
| route | 78500 | 0.104 |
| total | 78500 | 11.114 |

## Redis 本轮边界增量

- Commands：1437971；input：203483157 bytes；output：1055405805 bytes
- Hits/Misses：2990336/150；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：27858；safety epoch：3677 -> 3677

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.381 |
| client:loadtest | cpu_percent_total | 38.093 |
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
| docker:zg-canal | cpu_percent | 1.670 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.530 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.130 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 133.550 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 46.450 |
| docker:zg-zk | memory_percent | 1.170 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 234163.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.958 |
| process:counter | cpu_seconds_total | 121.812 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45813760.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 164.002 |
| process:gateway | cpu_seconds_total | 1826.703 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51081216.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 138.052 |
| process:knowpost | cpu_seconds_total | 8098.703 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70234112.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.871 |
| process:relation | cpu_seconds_total | 18.781 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48406528.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 3.266 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38158336.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 61.597 |
| process:user-storage | cpu_seconds_total | 664.250 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 59273216.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 161490980.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3677.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596899.000 |
| redis | keyspace_hits | 1158923247.000 |
| redis | keyspace_misses | 258837.000 |
| redis | net_input_bytes | 50469547232.000 |
| redis | net_output_bytes | 311383200609.000 |
| redis | ops_per_sec | 27858.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27131.000 |
| redis | used_memory_bytes | 103752352.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
