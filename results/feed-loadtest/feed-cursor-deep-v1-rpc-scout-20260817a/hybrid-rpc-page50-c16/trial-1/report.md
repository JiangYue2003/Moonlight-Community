# Feed 压测报告：hybrid / rpc / page50-c16

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:21:28+08:00
- 采样时长：10.0119408s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 5827 | 5827 | 0 | 0 | 582.04 | 26.968 | 33.565 | 35.626 | 40.521 | 52.619 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.007 |
| mysql | 266 | 0.046 |
| redis | 11920 | 2.046 |
| relation | 40 | 0.007 |

- Cold compute：5827（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 5827 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 5827 | 6.016 |
| counter | 40 | 5.833 |
| hydrate | 5827 | 20.547 |
| inbox | 5827 | 6.012 |
| merge_dedup | 5827 | 0.167 |
| relation | 40 | 6.101 |
| route | 5827 | 0.278 |
| total | 5827 | 27.175 |

## Redis 本轮边界增量

- Commands：44718；input：210991147 bytes；output：1339110529 bytes
- Hits/Misses：5979370/400；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：4722；safety epoch：3601 -> 3601

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.005 |
| client:loadtest | cpu_percent_total | 16.075 |
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
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.860 |
| docker:zg-es | memory_percent | 11.970 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.690 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 128.240 |
| docker:zg-kafka | memory_percent | 7.020 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 44.790 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 86808.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 28.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.677 |
| process:counter | cpu_seconds_total | 37.281 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 47669248.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 37494784.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 245.488 |
| process:knowpost | cpu_seconds_total | 644.547 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 87814144.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.779 |
| process:relation | cpu_seconds_total | 17.375 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 53964800.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.359 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42962944.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.344 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40947712.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12846715.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3601.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597824.000 |
| redis | keyspace_hits | 60281224.000 |
| redis | keyspace_misses | 49537.000 |
| redis | net_input_bytes | 2965331129.000 |
| redis | net_output_bytes | 17425253894.000 |
| redis | ops_per_sec | 4722.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21677.000 |
| redis | used_memory_bytes | 103377992.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
