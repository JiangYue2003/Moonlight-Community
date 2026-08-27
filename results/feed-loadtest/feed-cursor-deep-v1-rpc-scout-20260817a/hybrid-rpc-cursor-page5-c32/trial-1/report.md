# Feed 压测报告：hybrid / rpc / cursor-page5-c32

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:18:56+08:00
- 采样时长：10.0106088s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 28109 | 28109 | 0 | 0 | 2808.68 | 10.785 | 14.628 | 17.194 | 22.521 | 35.144 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：75.786ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.001 |
| mysql | 63 | 0.002 |
| redis | 28172 | 1.002 |
| relation | 40 | 0.001 |

- Cold compute：28109（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 28109 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 3.092 |
| cursor_decode | 28109 | 0.010 |
| cursor_seek | 28109 | 7.467 |
| hydrate | 28109 | 3.642 |
| merge_dedup | 28109 | 0.008 |
| relation | 40 | 3.261 |
| route | 28109 | 0.051 |
| total | 28109 | 11.178 |

## Redis 本轮边界增量

- Commands：510365；input：72613772 bytes；output：377792593 bytes
- Hits/Misses：1069277/135；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：52560；safety epoch：3593 -> 3593

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.731 |
| client:loadtest | cpu_percent_total | 43.704 |
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
| docker:zg-canal | cpu_percent | 0.130 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.080 |
| docker:zg-es | memory_percent | 11.960 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.040 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 10.550 |
| docker:zg-kafka | memory_percent | 7.020 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 39.540 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 84145.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 25.338 |
| process:counter | cpu_seconds_total | 34.094 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 47431680.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 38326272.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 151.826 |
| process:knowpost | cpu_seconds_total | 438.312 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 77971456.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.829 |
| process:relation | cpu_seconds_total | 16.812 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 52658176.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.219 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42930176.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.266 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40906752.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 8868190.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3593.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597237.000 |
| redis | keyspace_hits | 24155016.000 |
| redis | keyspace_misses | 31484.000 |
| redis | net_input_bytes | 1434526502.000 |
| redis | net_output_bytes | 8850656558.000 |
| redis | ops_per_sec | 52560.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21525.000 |
| redis | used_memory_bytes | 105707648.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
