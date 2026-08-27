# Feed 压测报告：hybrid / gateway / cursor-page5-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:52:46+08:00
- 采样时长：1m0.0217788s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 74369 | 74369 | 0 | 0 | 1239.14 | 25.190 | 35.046 | 39.575 | 52.900 | 84.170 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：91.9529ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 161 | 0.002 |
| redis | 74530 | 1.002 |
| relation | 240 | 0.003 |

- Cold compute：74369（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1561749 | 21.000 |
| merge_candidates | 11155350 | 150.000 |
| redis_commands | 1264273 | 17.000 |
| redis_members | 17551084 | 236.000 |
| redis_roundtrips | 148738 | 2.000 |
| tie_members | 9370494 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 74369 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 8.520 |
| cursor_decode | 74369 | 0.011 |
| cursor_seek | 74369 | 16.215 |
| hydrate | 74369 | 8.159 |
| merge_dedup | 74369 | 0.009 |
| relation | 240 | 9.838 |
| route | 74369 | 0.257 |
| total | 74369 | 24.653 |

## Redis 本轮边界增量

- Commands：1354103；input：192185426 bytes；output：999707929 bytes
- Hits/Misses：2833320/168；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：24342；safety epoch：3678 -> 3678

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.455 |
| client:loadtest | cpu_percent_total | 39.283 |
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
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.120 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.130 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 143.660 |
| docker:zg-kafka | memory_percent | 7.570 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 46.250 |
| docker:zg-zk | memory_percent | 1.000 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 234746.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.422 |
| process:counter | cpu_seconds_total | 123.219 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45817856.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 153.245 |
| process:gateway | cpu_seconds_total | 1916.672 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52011008.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 133.934 |
| process:knowpost | cpu_seconds_total | 8178.469 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71716864.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.425 |
| process:relation | cpu_seconds_total | 19.391 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48164864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 3.266 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38137856.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 61.191 |
| process:user-storage | cpu_seconds_total | 695.859 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55099392.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 163102986.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3678.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596892.000 |
| redis | keyspace_hits | 1162291482.000 |
| redis | keyspace_misses | 259108.000 |
| redis | net_input_bytes | 50698169870.000 |
| redis | net_output_bytes | 312571682520.000 |
| redis | ops_per_sec | 24342.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27205.000 |
| redis | used_memory_bytes | 106050280.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
