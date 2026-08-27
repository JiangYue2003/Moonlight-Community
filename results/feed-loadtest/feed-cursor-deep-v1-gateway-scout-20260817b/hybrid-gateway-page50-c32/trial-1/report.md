# Feed 压测报告：hybrid / gateway / page50-c32

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:37:17+08:00
- 采样时长：10.0386841s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 4001 | 4001 | 0 | 0 | 398.56 | 85.146 | 103.168 | 108.569 | 117.323 | 153.322 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.010 |
| mysql | 142 | 0.035 |
| redis | 8144 | 2.035 |
| relation | 40 | 0.010 |

- Cold compute：4001（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 4081020 | 1020.000 |
| merge_candidates | 5997499 | 1499.000 |
| redis_commands | 24006 | 6.000 |
| redis_members | 6001500 | 1500.000 |
| redis_roundtrips | 4001 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 4001 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 4001 | 17.628 |
| counter | 40 | 9.391 |
| hydrate | 4001 | 59.695 |
| inbox | 4001 | 17.624 |
| merge_dedup | 4001 | 0.164 |
| relation | 40 | 9.388 |
| route | 4001 | 1.143 |
| total | 4001 | 78.830 |

## Redis 本轮边界增量

- Commands：30360；input：144840897 bytes；output：919471281 bytes
- Hits/Misses：4105980/310；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：3576；safety epoch：3618 -> 3618

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.798 |
| client:loadtest | cpu_percent_total | 12.763 |
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
| docker:zg-canal | cpu_percent | 0.120 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.130 |
| docker:zg-es | memory_percent | 12.080 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.160 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 53.270 |
| docker:zg-kafka | memory_percent | 7.270 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 0.180 |
| docker:zg-zk | memory_percent | 0.930 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 96517.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 30.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 1.594 |
| process:counter | cpu_seconds_total | 9.844 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45387776.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 237.856 |
| process:gateway | cpu_seconds_total | 197.031 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52158464.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 215.175 |
| process:knowpost | cpu_seconds_total | 247.922 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 92090368.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.548 |
| process:relation | cpu_seconds_total | 1.328 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47190016.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.453 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42074112.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 14.712 |
| process:user-storage | cpu_seconds_total | 73.125 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 48820224.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19832750.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3618.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 598261.000 |
| redis | keyspace_hits | 117137783.000 |
| redis | keyspace_misses | 85171.000 |
| redis | net_input_bytes | 5412982210.000 |
| redis | net_output_bytes | 30628959555.000 |
| redis | ops_per_sec | 3576.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22626.000 |
| redis | used_memory_bytes | 105302472.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
