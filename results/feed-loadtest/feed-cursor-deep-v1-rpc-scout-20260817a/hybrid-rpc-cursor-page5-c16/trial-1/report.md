# Feed 压测报告：hybrid / rpc / cursor-page5-c16

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:18:37+08:00
- 采样时长：10.0063621s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 27301 | 27301 | 0 | 0 | 2728.95 | 5.444 | 8.281 | 9.708 | 11.839 | 22.647 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：95.0338ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.001 |
| mysql | 40 | 0.001 |
| redis | 27341 | 1.001 |
| relation | 40 | 0.001 |

- Cold compute：27301（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 27301 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 2.268 |
| cursor_decode | 27301 | 0.011 |
| cursor_seek | 27301 | 3.809 |
| hydrate | 27301 | 1.802 |
| merge_dedup | 27301 | 0.007 |
| relation | 40 | 2.535 |
| route | 27301 | 0.026 |
| total | 27301 | 5.656 |

## Redis 本轮边界增量

- Commands：497532；input：70640931 bytes；output：366982581 bytes
- Hits/Misses：1038654/52；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：52860；safety epoch：3592 -> 3592

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.879 |
| client:loadtest | cpu_percent_total | 46.064 |
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
| docker:zg-canal | cpu_percent | 1.820 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.730 |
| docker:zg-es | memory_percent | 11.960 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.890 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 40.480 |
| docker:zg-kafka | memory_percent | 7.040 |
| docker:zg-kafka | pids | 99.000 |
| docker:zg-zk | cpu_percent | 7.370 |
| docker:zg-zk | memory_percent | 1.050 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 83950.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.419 |
| process:counter | cpu_seconds_total | 33.672 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 47271936.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 38326272.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 150.186 |
| process:knowpost | cpu_seconds_total | 419.750 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 75923456.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 16.734 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 55214080.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.219 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42917888.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.266 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40906752.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8191610.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3592.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597265.000 |
| redis | keyspace_hits | 22742756.000 |
| redis | keyspace_misses | 31229.000 |
| redis | net_input_bytes | 1338451885.000 |
| redis | net_output_bytes | 8351653296.000 |
| redis | ops_per_sec | 52860.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21506.000 |
| redis | used_memory_bytes | 103325800.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
