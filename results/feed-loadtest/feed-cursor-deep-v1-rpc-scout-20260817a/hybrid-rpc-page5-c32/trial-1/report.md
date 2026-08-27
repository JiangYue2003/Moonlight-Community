# Feed 压测报告：hybrid / rpc / page5-c32

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:17:59+08:00
- 采样时长：10.0096827s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 23529 | 23529 | 0 | 0 | 2351.00 | 12.402 | 17.610 | 22.997 | 27.631 | 41.851 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 206 | 0.009 |
| redis | 47264 | 2.009 |
| relation | 40 | 0.002 |

- Cold compute：23529（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 23529 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 23529 | 6.752 |
| counter | 40 | 3.855 |
| hydrate | 23529 | 6.475 |
| inbox | 23529 | 6.751 |
| merge_dedup | 23529 | 0.024 |
| relation | 40 | 3.017 |
| route | 23529 | 0.086 |
| total | 23529 | 13.379 |

## Redis 本轮边界增量

- Commands：167967；input：109785795 bytes；output：1077764021 bytes
- Hits/Misses：2965624/296；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：17206；safety epoch：3590 -> 3590

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.946 |
| client:loadtest | cpu_percent_total | 47.142 |
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
| docker:zg-canal | cpu_percent | 1.670 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.010 |
| docker:zg-es | memory_percent | 11.960 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.970 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 16.460 |
| docker:zg-kafka | memory_percent | 7.030 |
| docker:zg-kafka | pids | 100.000 |
| docker:zg-zk | cpu_percent | 35.220 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 83691.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 30.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 1.550 |
| process:counter | cpu_seconds_total | 32.719 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 46084096.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 38285312.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 213.605 |
| process:knowpost | cpu_seconds_total | 374.453 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 77901824.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.775 |
| process:relation | cpu_seconds_total | 16.719 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 54837248.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 2.219 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42905600.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.234 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40722432.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7302749.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3590.000 |
| redis | hit_rate | 0.998 |
| redis | keys | 597314.000 |
| redis | keyspace_hits | 17519817.000 |
| redis | keyspace_misses | 31130.000 |
| redis | net_input_bytes | 1101746669.000 |
| redis | net_output_bytes | 6467793164.000 |
| redis | ops_per_sec | 17206.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21468.000 |
| redis | used_memory_bytes | 103755432.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
