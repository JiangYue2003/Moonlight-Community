# Feed 压测报告：hybrid / rpc / sequential-50-c1

- Run ID：`feed-cursor-deep-v1-smoke`
- 开始时间：2026-08-17T16:13:40+08:00
- 采样时长：38.9346ms
- 并发：1
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 50 | 50 | 0 | 0 | 1284.20 | 0.545 | 1.114 | 1.136 | 4.322 | 4.322 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1/50
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 1 | 0.020 |
| mysql | 0 | 0.000 |
| redis | 51 | 1.020 |
| relation | 1 | 0.020 |

- Cold compute：50（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 50 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1 | 0.000 |
| counter | 1 | 1.678 |
| cursor_decode | 49 | 0.011 |
| cursor_seek | 49 | 0.318 |
| hydrate | 50 | 0.183 |
| inbox | 1 | 0.000 |
| merge_dedup | 50 | 0.000 |
| relation | 1 | 1.587 |
| route | 50 | 0.065 |
| total | 50 | 0.582 |

## Redis 本轮边界增量

- Commands：787；input：115663 bytes；output：335640 bytes
- Hits/Misses：1802/0；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：691；safety epoch：3583 -> 3583

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.016 |
| client:loadtest | cpu_percent_total | 80.263 |
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
| docker:zg-es | cpu_percent | 1.200 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 0.410 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 21.480 |
| docker:zg-kafka | memory_percent | 7.010 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.100 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 82484.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 0.000 |
| process:counter | cpu_seconds_total | 26.953 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 41627648.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 38420480.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 66.682 |
| process:knowpost | cpu_seconds_total | 219.844 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 67403776.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 16.375 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 48439296.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.047 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 43012096.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.172 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40415232.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5230610.000 |
| redis | connected_clients | 96.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3583.000 |
| redis | hit_rate | 0.694 |
| redis | keys | 597488.000 |
| redis | keyspace_hits | 63568.000 |
| redis | keyspace_misses | 28060.000 |
| redis | net_input_bytes | 411951976.000 |
| redis | net_output_bytes | 103517587.000 |
| redis | ops_per_sec | 691.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21199.000 |
| redis | used_memory_bytes | 100462672.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
