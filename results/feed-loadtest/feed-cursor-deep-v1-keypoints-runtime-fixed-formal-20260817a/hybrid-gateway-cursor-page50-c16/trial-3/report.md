# Feed 压测报告：hybrid / gateway / cursor-page50-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:47:32+08:00
- 采样时长：1m0.0158714s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 130839 | 130839 | 0 | 0 | 2180.45 | 7.070 | 9.947 | 11.369 | 14.233 | 30.630 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：987.8418ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 122 | 0.001 |
| redis | 130961 | 1.001 |
| relation | 240 | 0.002 |

- Cold compute：130839（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2747619 | 21.000 |
| merge_candidates | 2878458 | 22.000 |
| redis_commands | 1700907 | 13.000 |
| redis_members | 3140136 | 24.000 |
| redis_roundtrips | 261678 | 2.000 |
| tie_members | 261678 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 130839 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 2.625 |
| cursor_decode | 130839 | 0.011 |
| cursor_seek | 130839 | 3.879 |
| hydrate | 130839 | 2.111 |
| merge_dedup | 130839 | 0.005 |
| relation | 240 | 3.684 |
| route | 130839 | 0.039 |
| total | 130839 | 6.052 |

## Redis 本轮边界增量

- Commands：1867698；input：286586751 bytes；output：588161896 bytes
- Hits/Misses：4455695/362；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：34504；safety epoch：3763 -> 3763

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.551 |
| client:loadtest | cpu_percent_total | 72.819 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.490 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.860 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 164.280 |
| docker:zg-kafka | memory_percent | 7.630 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 51.130 |
| docker:zg-zk | memory_percent | 1.370 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 531348.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 10.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.837 |
| process:counter | cpu_seconds_total | 311.969 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46505984.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 270.358 |
| process:gateway | cpu_seconds_total | 6558.953 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51687424.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 224.998 |
| process:knowpost | cpu_seconds_total | 18384.016 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71479296.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.098 |
| process:relation | cpu_seconds_total | 50.438 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49254400.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 6.062 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38051840.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 96.744 |
| process:user-storage | cpu_seconds_total | 2342.016 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 64856064.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 300181349.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3763.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596557.000 |
| redis | keyspace_hits | 2736803482.000 |
| redis | keyspace_misses | 718817.000 |
| redis | net_input_bytes | 114847576919.000 |
| redis | net_output_bytes | 672763273687.000 |
| redis | ops_per_sec | 34504.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34091.000 |
| redis | used_memory_bytes | 101796672.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
