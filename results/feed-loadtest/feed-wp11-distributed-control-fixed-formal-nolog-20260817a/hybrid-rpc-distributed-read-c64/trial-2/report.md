# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-control-fixed-formal-nolog-20260817a`
- 开始时间：2026-08-17T00:57:55+08:00
- 采样时长：1m0.0335243s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 413201 | 413201 | 0 | 0 | 6886.04 | 8.924 | 12.560 | 13.794 | 16.693 | 34.419 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 14400 | 0.035 |
| mysql | 201 | 0.000 |
| redis | 1239804 | 3.000 |
| relation | 413201 | 1.000 |

- Cold compute：413201（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 413201 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 413201 | 1.769 |
| counter | 14400 | 2.609 |
| hydrate | 413201 | 2.115 |
| inbox | 413201 | 2.188 |
| merge_dedup | 413201 | 0.004 |
| relation | 413201 | 2.811 |
| route | 413201 | 0.095 |
| total | 413201 | 8.999 |

## Redis 本轮边界增量

- Commands：3256155；input：377178009 bytes；output：1133462661 bytes
- Hits/Misses：7665020/447701；run hit rate：94.48%
- Evicted/Rejected：0/0；ops/s max：60204；safety epoch：584 -> 584

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.358 |
| client:loadtest | cpu_percent_total | 85.733 |
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
| docker:zg-canal | cpu_percent | 0.290 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.650 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 7.250 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 256.700 |
| docker:zg-kafka | memory_percent | 8.810 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 65.060 |
| docker:zg-zk | memory_percent | 1.680 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 26392851.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 33.000 |
| mysql | threads_running | 14.000 |
| process:counter | cpu_percent | 27.856 |
| process:counter | cpu_seconds_total | 32.609 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 53956608.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 2.320 |
| process:gateway | cpu_seconds_total | 52.672 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 51564544.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 381.935 |
| process:knowpost | cpu_seconds_total | 553.484 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 75329536.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 308.710 |
| process:relation | cpu_seconds_total | 456.938 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 59154432.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.188 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35864576.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 21.109 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 46714880.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 264025699.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 584.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 739463.000 |
| redis | keyspace_hits | 1085423938.000 |
| redis | keyspace_misses | 12202029.000 |
| redis | net_input_bytes | 50502598391.000 |
| redis | net_output_bytes | 202256559415.000 |
| redis | ops_per_sec | 60204.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52714.000 |
| redis | used_memory_bytes | 155661832.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
