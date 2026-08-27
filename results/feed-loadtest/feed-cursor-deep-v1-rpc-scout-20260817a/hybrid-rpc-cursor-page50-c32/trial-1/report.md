# Feed 压测报告：hybrid / rpc / cursor-page50-c32

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:22:46+08:00
- 采样时长：10.0071138s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 72841 | 72841 | 0 | 0 | 7282.40 | 4.184 | 5.985 | 6.994 | 8.791 | 16.042 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：720.2073ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.001 |
| mysql | 50 | 0.001 |
| redis | 72891 | 1.001 |
| relation | 40 | 0.001 |

- Cold compute：72841（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 72841 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 2.056 |
| cursor_decode | 72841 | 0.009 |
| cursor_seek | 72841 | 2.696 |
| hydrate | 72841 | 1.366 |
| merge_dedup | 72841 | 0.005 |
| relation | 40 | 2.612 |
| route | 72841 | 0.012 |
| total | 72841 | 4.094 |

## Redis 本轮边界增量

- Commands：1025798；input：158677015 bytes；output：327116046 bytes
- Hits/Misses：2477782/79；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：106277；safety epoch：3605 -> 3605

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 9.964 |
| client:loadtest | cpu_percent_total | 159.418 |
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
| docker:zg-canal | cpu_percent | 1.910 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.600 |
| docker:zg-es | memory_percent | 12.000 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.110 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 146.570 |
| docker:zg-kafka | memory_percent | 7.510 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 0.170 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 92491.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.097 |
| process:counter | cpu_seconds_total | 39.266 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 47116288.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 38232064.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 286.834 |
| process:knowpost | cpu_seconds_total | 772.719 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 80986112.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 17.656 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 55226368.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.308 |
| process:search | cpu_seconds_total | 2.469 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42962944.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.344 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 41086976.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 15747803.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3605.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597873.000 |
| redis | keyspace_hits | 82897011.000 |
| redis | keyspace_misses | 61620.000 |
| redis | net_input_bytes | 3958567572.000 |
| redis | net_output_bytes | 21886828333.000 |
| redis | ops_per_sec | 106277.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21755.000 |
| redis | used_memory_bytes | 104623888.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
