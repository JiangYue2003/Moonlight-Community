# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-distributed-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:08:25+08:00
- 采样时长：1m0.0349587s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 77309 | 77309 | 0 | 0 | 1287.87 | 19.753 | 51.159 | 56.138 | 70.367 | 130.690 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 91152 | 1.179 |
| counter | 8990 | 0.116 |
| mysql | 5 | 0.000 |
| redis | 89751 | 1.161 |
| relation | 8990 | 0.116 |

- Cold compute：16868（0.218 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 35985 |
| l1_stale | 774 |
| l2_fresh | 32913 |
| l2_stale | 7637 |
| miss | 0 |

- L1+L2 Fresh ratio：89.12%
- Refresh max：queue=18 active=32 pending=50

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 16868 | 19.386 |
| counter | 8990 | 21.196 |
| hydrate | 16868 | 18.990 |
| inbox | 16868 | 19.385 |
| merge_dedup | 16868 | 0.004 |
| relation | 8990 | 23.593 |
| route | 16868 | 24.021 |
| total | 77309 | 18.922 |

## Redis 本轮边界增量

- Commands：272922；input：47300407 bytes；output：99370376 bytes
- Hits/Misses：441944/10389；run hit rate：97.70%
- Evicted/Rejected：0/0；ops/s max：10672；safety epoch：556 -> 556

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.827 |
| client:loadtest | cpu_percent_total | 109.233 |
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
| docker:zg-canal | cpu_percent | 2.250 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.800 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.050 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 201.490 |
| docker:zg-kafka | memory_percent | 8.760 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 49.200 |
| docker:zg-zk | memory_percent | 1.720 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23460992.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 60.794 |
| process:counter | cpu_seconds_total | 323.188 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54272000.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 476.083 |
| process:gateway | cpu_seconds_total | 1763.578 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 54169600.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 297.100 |
| process:knowpost | cpu_seconds_total | 3489.828 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 101462016.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 94.654 |
| process:relation | cpu_seconds_total | 384.516 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 57008128.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.557 |
| process:search | cpu_seconds_total | 0.812 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36392960.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 87.768 |
| process:user-storage | cpu_seconds_total | 422.172 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 52539392.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 212277975.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 556.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 707155.000 |
| redis | keyspace_hits | 1012472722.000 |
| redis | keyspace_misses | 6212295.000 |
| redis | net_input_bytes | 43810726414.000 |
| redis | net_output_bytes | 190845092989.000 |
| redis | ops_per_sec | 10672.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49744.000 |
| redis | used_memory_bytes | 126339184.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
