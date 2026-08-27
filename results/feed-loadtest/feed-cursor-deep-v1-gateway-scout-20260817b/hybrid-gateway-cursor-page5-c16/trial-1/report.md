# Feed 压测报告：hybrid / gateway / cursor-page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:35:03+08:00
- 采样时长：10.008651s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 11262 | 11262 | 0 | 0 | 1125.35 | 13.629 | 19.681 | 23.879 | 30.205 | 40.743 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：181.261ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.004 |
| mysql | 0 | 0.000 |
| redis | 11262 | 1.000 |
| relation | 40 | 0.004 |

- Cold compute：11262（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 236502 | 21.000 |
| merge_candidates | 1689300 | 150.000 |
| redis_commands | 191454 | 17.000 |
| redis_members | 2657832 | 236.000 |
| redis_roundtrips | 22524 | 2.000 |
| tie_members | 1419012 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 11262 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 4.156 |
| cursor_decode | 11262 | 0.011 |
| cursor_seek | 11262 | 8.513 |
| hydrate | 11262 | 4.245 |
| merge_dedup | 11262 | 0.009 |
| relation | 40 | 5.676 |
| route | 11262 | 0.136 |
| total | 11262 | 12.915 |

## Redis 本轮边界增量

- Commands：206402；input：29194729 bytes；output：151425860 bytes
- Hits/Misses：429207/12；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：22226；safety epoch：3611 -> 3611

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.117 |
| client:loadtest | cpu_percent_total | 33.877 |
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
| docker:zg-canal | cpu_percent | 1.890 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.850 |
| docker:zg-es | memory_percent | 12.070 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.310 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 19.220 |
| docker:zg-kafka | memory_percent | 7.030 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.180 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 93739.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 24.716 |
| process:counter | cpu_seconds_total | 6.359 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 44453888.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 189.488 |
| process:gateway | cpu_seconds_total | 104.844 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51331072.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 128.573 |
| process:knowpost | cpu_seconds_total | 106.812 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 66191360.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.547 |
| process:relation | cpu_seconds_total | 0.531 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 45973504.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.337 |
| process:search | cpu_seconds_total | 0.328 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 41299968.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 56.476 |
| process:user-storage | cpu_seconds_total | 38.469 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54161408.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18460938.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3611.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597344.000 |
| redis | keyspace_hits | 94166833.000 |
| redis | keyspace_misses | 64331.000 |
| redis | net_input_bytes | 4518289403.000 |
| redis | net_output_bytes | 25211692601.000 |
| redis | ops_per_sec | 22226.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22493.000 |
| redis | used_memory_bytes | 103233320.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
