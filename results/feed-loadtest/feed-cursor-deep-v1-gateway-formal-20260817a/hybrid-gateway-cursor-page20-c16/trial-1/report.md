# Feed 压测报告：hybrid / gateway / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:04:03+08:00
- 采样时长：1m0.0135284s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 113472 | 113472 | 0 | 0 | 1891.04 | 8.062 | 11.784 | 13.626 | 17.300 | 32.714 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：399.7245ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 104 | 0.001 |
| redis | 113576 | 1.001 |
| relation | 240 | 0.002 |

- Cold compute：113472（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2382912 | 21.000 |
| merge_candidates | 7262208 | 64.000 |
| redis_commands | 1588608 | 14.000 |
| redis_members | 7602624 | 67.000 |
| redis_roundtrips | 226944 | 2.000 |
| tie_members | 340416 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 113472 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 2.949 |
| cursor_decode | 113472 | 0.012 |
| cursor_seek | 113472 | 4.738 |
| hydrate | 113472 | 2.516 |
| merge_dedup | 113472 | 0.007 |
| relation | 240 | 3.977 |
| route | 113472 | 0.053 |
| total | 113472 | 7.330 |

## Redis 本轮边界增量

- Commands：1735340；input：260002486 bytes；output：718141873 bytes
- Hits/Misses：3978698/344；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：33358；safety epoch：3687 -> 3687

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.798 |
| client:loadtest | cpu_percent_total | 60.768 |
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
| docker:zg-canal | cpu_percent | 2.450 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.630 |
| docker:zg-es | memory_percent | 12.350 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.990 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 150.780 |
| docker:zg-kafka | memory_percent | 7.680 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 48.550 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 263010.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 8.514 |
| process:counter | cpu_seconds_total | 136.359 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46743552.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 220.666 |
| process:gateway | cpu_seconds_total | 2522.781 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50950144.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 176.725 |
| process:knowpost | cpu_seconds_total | 9223.812 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70414336.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.416 |
| process:relation | cpu_seconds_total | 22.672 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48431104.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.781 |
| process:search | cpu_seconds_total | 3.547 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38174720.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 86.572 |
| process:user-storage | cpu_seconds_total | 930.094 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 56233984.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 171597941.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3687.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596257.000 |
| redis | keyspace_hits | 1346888634.000 |
| redis | keyspace_misses | 294522.000 |
| redis | net_input_bytes | 57652357291.000 |
| redis | net_output_bytes | 359560090203.000 |
| redis | ops_per_sec | 33358.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27882.000 |
| redis | used_memory_bytes | 102411640.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
