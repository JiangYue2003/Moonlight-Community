# Feed 压测报告：hybrid / rpc / cursor-page50-c1

- Run ID：`feed-cursor-deep-v1-smoke-recheck`
- 开始时间：2026-08-17T16:14:57+08:00
- 采样时长：11.9847ms
- 并发：1
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 5 | 5 | 0 | 0 | 417.20 | 1.712 | 5.322 | 5.322 | 5.322 | 5.322 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：713.2322ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 5 | 1.000 |
| mysql | 1 | 0.200 |
| redis | 6 | 1.200 |
| relation | 5 | 1.000 |

- Cold compute：5（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 5 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 5 | 0.796 |
| cursor_decode | 5 | 0.000 |
| cursor_seek | 5 | 0.134 |
| hydrate | 5 | 0.309 |
| merge_dedup | 5 | 0.000 |
| relation | 5 | 0.738 |
| route | 5 | 1.850 |
| total | 5 | 2.293 |

## Redis 本轮边界增量

- Commands：210；input：15449 bytes；output：29474 bytes
- Hits/Misses：305/1；run hit rate：99.67%
- Evicted/Rejected：0/0；ops/s max：8718；safety epoch：3585 -> 3585

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.000 |
| client:loadtest | cpu_percent_total | 0.000 |
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
| docker:zg-canal | cpu_percent | 0.110 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.960 |
| docker:zg-es | memory_percent | 11.940 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 0.620 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 175.440 |
| docker:zg-kafka | memory_percent | 7.020 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.100 |
| docker:zg-zk | memory_percent | 0.890 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 82725.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 51.343 |
| process:counter | cpu_seconds_total | 28.594 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 41734144.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 38424576.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 0.000 |
| process:knowpost | cpu_seconds_total | 227.016 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 65789952.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 16.453 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 48713728.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.109 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42708992.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.172 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40480768.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5319891.000 |
| redis | connected_clients | 96.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3585.000 |
| redis | hit_rate | 0.864 |
| redis | keys | 597567.000 |
| redis | keyspace_hits | 184072.000 |
| redis | keyspace_misses | 28892.000 |
| redis | net_input_bytes | 422516319.000 |
| redis | net_output_bytes | 124160881.000 |
| redis | ops_per_sec | 8718.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21276.000 |
| redis | used_memory_bytes | 100509688.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
