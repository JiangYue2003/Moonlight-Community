# Feed 压测报告：hybrid / rpc / hot-read-c128

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:07:36+08:00
- 采样时长：1m0.0397728s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 483477 | 483477 | 0 | 0 | 8056.44 | 15.632 | 19.894 | 21.099 | 23.732 | 53.183 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 166 | 0.000 |
| counter | 9 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 121 | 0.000 |
| relation | 9 | 0.000 |

- Cold compute：18（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 483365 |
| l1_stale | 0 |
| l2_fresh | 112 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=1 pending=1

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 18 | 0.119 |
| counter | 9 | 1.884 |
| hydrate | 18 | 0.230 |
| inbox | 18 | 0.119 |
| merge_dedup | 18 | 0.057 |
| relation | 9 | 2.107 |
| route | 18 | 1.996 |
| total | 483477 | 0.005 |

## Redis 本轮边界增量

- Commands：53183；input：3802464 bytes；output：1429089 bytes
- Hits/Misses：1328/9；run hit rate：99.33%
- Evicted/Rejected：0/0；ops/s max：1102；safety epoch：471 -> 471

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.898 |
| client:loadtest | cpu_percent_total | 126.375 |
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
| docker:zg-canal | cpu_percent | 2.270 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.450 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.630 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 153.330 |
| docker:zg-kafka | memory_percent | 8.570 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 49.490 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20180380.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.065 |
| process:counter | cpu_seconds_total | 34.125 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 40271872.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.297 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37691392.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 236.808 |
| process:knowpost | cpu_seconds_total | 1247.766 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 106795008.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.960 |
| process:relation | cpu_seconds_total | 0.531 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 39112704.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.322 |
| process:search | cpu_seconds_total | 0.406 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37097472.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.297 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34512896.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 177738648.000 |
| redis | connected_clients | 16.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 471.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676902.000 |
| redis | keyspace_hits | 977434576.000 |
| redis | keyspace_misses | 5364838.000 |
| redis | net_input_bytes | 40560859264.000 |
| redis | net_output_bytes | 180192214718.000 |
| redis | ops_per_sec | 1102.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 38895.000 |
| redis | used_memory_bytes | 95628120.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
