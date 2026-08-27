# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:22:37+08:00
- 采样时长：1m0.0398399s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 485139 | 485139 | 0 | 0 | 8084.14 | 15.612 | 19.693 | 20.876 | 23.220 | 50.145 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2447 | 0.005 |
| counter | 183 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 2412 | 0.005 |
| relation | 183 | 0.000 |

- Cold compute：367（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 483793 |
| l1_stale | 0 |
| l2_fresh | 1346 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 367 | 0.161 |
| counter | 183 | 0.642 |
| hydrate | 367 | 0.194 |
| inbox | 367 | 0.161 |
| merge_dedup | 367 | 0.009 |
| relation | 183 | 1.436 |
| route | 367 | 1.056 |
| total | 485139 | 0.005 |

## Redis 本轮边界增量

- Commands：62924；input：5596126 bytes；output：6998979 bytes
- Hits/Misses：24097/183；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1382；safety epoch：483 -> 483

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.999 |
| client:loadtest | cpu_percent_total | 127.988 |
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
| docker:zg-canal | cpu_percent | 0.150 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.410 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.950 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 157.350 |
| docker:zg-kafka | memory_percent | 8.610 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 17.500 |
| docker:zg-zk | memory_percent | 1.390 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20183593.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.746 |
| process:counter | cpu_seconds_total | 75.250 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43888640.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.516 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37842944.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 235.480 |
| process:knowpost | cpu_seconds_total | 3051.625 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 144637952.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 3.016 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 46882816.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.781 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37027840.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.778 |
| process:user-storage | cpu_seconds_total | 0.672 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34902016.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178636279.000 |
| redis | connected_clients | 55.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 483.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677134.000 |
| redis | keyspace_hits | 977678120.000 |
| redis | keyspace_misses | 5369558.000 |
| redis | net_input_bytes | 40636419386.000 |
| redis | net_output_bytes | 180267168907.000 |
| redis | ops_per_sec | 1382.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 39796.000 |
| redis | used_memory_bytes | 98011384.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
