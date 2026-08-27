# Feed 压测报告：hybrid / gateway / cursor-page20-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:07:52+08:00
- 采样时长：1m0.0199039s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 97072 | 97072 | 0 | 0 | 1617.52 | 19.113 | 26.935 | 30.225 | 39.664 | 84.654 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：399.0998ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 208 | 0.002 |
| redis | 97280 | 1.002 |
| relation | 240 | 0.002 |

- Cold compute：97072（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2038512 | 21.000 |
| merge_candidates | 6212608 | 64.000 |
| redis_commands | 1359008 | 14.000 |
| redis_members | 6503824 | 67.000 |
| redis_roundtrips | 194144 | 2.000 |
| tie_members | 291216 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 97072 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 7.091 |
| cursor_decode | 97072 | 0.011 |
| cursor_seek | 97072 | 12.063 |
| hydrate | 97072 | 6.190 |
| merge_dedup | 97072 | 0.007 |
| relation | 240 | 7.841 |
| route | 97072 | 0.155 |
| total | 97072 | 18.431 |

## Redis 本轮边界增量

- Commands：1474218；input：221674025 bytes；output：614146079 bytes
- Hits/Misses：3404547/448；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：28117；safety epoch：3690 -> 3690

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.391 |
| client:loadtest | cpu_percent_total | 54.253 |
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
| docker:zg-canal | cpu_percent | 2.140 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.450 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 157.000 |
| docker:zg-etcd | cpu_percent | 5.740 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 147.720 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 43.110 |
| docker:zg-zk | memory_percent | 1.200 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 264886.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 4.986 |
| process:counter | cpu_seconds_total | 145.578 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46661632.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 219.278 |
| process:gateway | cpu_seconds_total | 2912.906 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52699136.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 166.408 |
| process:knowpost | cpu_seconds_total | 9528.344 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71581696.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.644 |
| process:relation | cpu_seconds_total | 23.953 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48324608.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 3.594 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38158336.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 78.989 |
| process:user-storage | cpu_seconds_total | 1060.672 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 62308352.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 177051526.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3690.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596059.000 |
| redis | keyspace_hits | 1359386680.000 |
| redis | keyspace_misses | 297027.000 |
| redis | net_input_bytes | 58469145867.000 |
| redis | net_output_bytes | 361822042877.000 |
| redis | ops_per_sec | 28117.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28111.000 |
| redis | used_memory_bytes | 103731648.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
