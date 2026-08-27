# Feed 压测报告：hybrid / rpc / same-second-cursor-c1

- Run ID：`feed-cursor-deep-v1-smoke`
- 开始时间：2026-08-17T16:13:45+08:00
- 采样时长：7.3343ms
- 并发：1
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 5 | 5 | 0 | 0 | 681.73 | 1.035 | 3.197 | 3.197 | 3.197 | 3.197 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：2
- 准备请求：20；准备耗时：32.118ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 1 | 0.200 |
| mysql | 0 | 0.000 |
| redis | 5 | 1.000 |
| relation | 1 | 0.200 |

- Cold compute：5（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 5 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 1 | 0.526 |
| cursor_decode | 5 | 0.000 |
| cursor_seek | 5 | 0.519 |
| hydrate | 5 | 0.206 |
| merge_dedup | 5 | 0.000 |
| relation | 1 | 1.071 |
| route | 5 | 0.429 |
| total | 5 | 1.155 |

## Redis 本轮边界增量

- Commands：125；input：14079 bytes；output：71636 bytes
- Hits/Misses：222/0；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：984；safety epoch：3584 -> 3584

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
| docker:zg-canal | cpu_percent | 0.090 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.980 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 0.330 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 4.470 |
| docker:zg-kafka | memory_percent | 7.020 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.100 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 82514.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 0.000 |
| process:counter | cpu_seconds_total | 27.125 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 41635840.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 38420480.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 79.976 |
| process:knowpost | cpu_seconds_total | 220.375 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 68845568.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 16.375 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 48447488.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.062 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 43012096.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.172 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40427520.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5233824.000 |
| redis | connected_clients | 96.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3584.000 |
| redis | hit_rate | 0.699 |
| redis | keys | 597499.000 |
| redis | keyspace_hits | 65232.000 |
| redis | keyspace_misses | 28072.000 |
| redis | net_input_bytes | 412195490.000 |
| redis | net_output_bytes | 103999435.000 |
| redis | ops_per_sec | 984.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21204.000 |
| redis | used_memory_bytes | 100550480.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
