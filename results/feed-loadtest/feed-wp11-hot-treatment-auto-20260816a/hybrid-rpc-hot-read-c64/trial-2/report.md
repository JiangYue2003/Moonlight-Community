# Feed 压测报告：hybrid / rpc / hot-read-c64

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:03:51+08:00
- 采样时长：1m0.0351327s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 488973 | 488973 | 0 | 0 | 8148.58 | 7.812 | 9.947 | 10.584 | 12.181 | 26.861 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 163 | 0.000 |
| counter | 10 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 120 | 0.000 |
| relation | 10 | 0.000 |

- Cold compute：19（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 488867 |
| l1_stale | 0 |
| l2_fresh | 106 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 19 | 0.167 |
| counter | 10 | 1.700 |
| hydrate | 19 | 0.143 |
| inbox | 19 | 0.167 |
| merge_dedup | 19 | 0.000 |
| relation | 10 | 2.047 |
| route | 19 | 1.972 |
| total | 488973 | 0.005 |

## Redis 本轮边界增量

- Commands：53161；input：3802895 bytes；output：1429175 bytes
- Hits/Misses：1395/10；run hit rate：99.29%
- Evicted/Rejected：0/0；ops/s max：1095；safety epoch：468 -> 468

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.831 |
| client:loadtest | cpu_percent_total | 125.291 |
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
| docker:zg-canal | cpu_percent | 2.080 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.430 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.540 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 173.040 |
| docker:zg-kafka | memory_percent | 8.570 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 45.940 |
| docker:zg-zk | memory_percent | 1.230 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20180124.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.972 |
| process:counter | cpu_seconds_total | 23.484 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 39985152.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.172 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37027840.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 235.108 |
| process:knowpost | cpu_seconds_total | 794.609 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 110878720.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.775 |
| process:relation | cpu_seconds_total | 0.406 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 39030784.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.281 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37007360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 0.219 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34713600.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 177539083.000 |
| redis | connected_clients | 16.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 468.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676850.000 |
| redis | keyspace_hits | 977429680.000 |
| redis | keyspace_misses | 5364625.000 |
| redis | net_input_bytes | 40546552366.000 |
| redis | net_output_bytes | 180186947651.000 |
| redis | ops_per_sec | 1095.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 38670.000 |
| redis | used_memory_bytes | 95608584.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
