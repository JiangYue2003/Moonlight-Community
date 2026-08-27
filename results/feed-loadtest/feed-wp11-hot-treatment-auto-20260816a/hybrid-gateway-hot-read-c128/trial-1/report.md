# Feed 压测报告：hybrid / gateway / hot-read-c128

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:51:30+08:00
- 采样时长：1m0.0352715s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 291984 | 291984 | 0 | 0 | 4865.23 | 24.508 | 33.770 | 46.429 | 96.077 | 401.259 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 156 | 0.001 |
| counter | 8 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 118 | 0.000 |
| relation | 8 | 0.000 |

- Cold compute：17（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 291879 |
| l1_stale | 0 |
| l2_fresh | 105 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 17 | 0.280 |
| counter | 8 | 11.156 |
| hydrate | 17 | 0.341 |
| inbox | 17 | 0.280 |
| merge_dedup | 17 | 0.036 |
| relation | 8 | 2.612 |
| route | 17 | 6.479 |
| total | 291984 | 0.007 |

## Redis 本轮边界增量

- Commands：53269；input：3805067 bytes；output：1421307 bytes
- Hits/Misses：1257/8；run hit rate：99.37%
- Evicted/Rejected：0/0；ops/s max：1109；safety epoch：506 -> 506

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.791 |
| client:loadtest | cpu_percent_total | 108.660 |
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
| docker:zg-canal | cpu_percent | 2.700 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.100 |
| docker:zg-es | memory_percent | 11.940 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.950 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 204.360 |
| docker:zg-kafka | memory_percent | 8.620 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 65.450 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 84.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22730636.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 6.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 12.367 |
| process:counter | cpu_seconds_total | 172.781 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43024384.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 422.464 |
| process:gateway | cpu_seconds_total | 1856.438 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 60108800.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 210.984 |
| process:knowpost | cpu_seconds_total | 8079.547 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 113573888.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 8.984 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 44466176.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 1.516 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36990976.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 221.837 |
| process:user-storage | cpu_seconds_total | 921.453 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 66322432.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 195555713.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 506.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677265.000 |
| redis | keyspace_hits | 993149558.000 |
| redis | keyspace_misses | 5378849.000 |
| redis | net_input_bytes | 41860921567.000 |
| redis | net_output_bytes | 187011066599.000 |
| redis | ops_per_sec | 1109.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 41529.000 |
| redis | used_memory_bytes | 100087640.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
