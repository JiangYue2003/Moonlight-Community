# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:13:52+08:00
- 采样时长：1m0.0282776s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 482844 | 482844 | 0 | 0 | 8047.05 | 3.816 | 4.984 | 5.365 | 6.282 | 16.015 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2451 | 0.005 |
| counter | 180 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 2416 | 0.005 |
| relation | 180 | 0.000 |

- Cold compute：367（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 481494 |
| l1_stale | 0 |
| l2_fresh | 1350 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=1 pending=1

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 367 | 0.149 |
| counter | 180 | 0.694 |
| hydrate | 367 | 0.216 |
| inbox | 367 | 0.147 |
| merge_dedup | 367 | 0.012 |
| relation | 180 | 1.474 |
| route | 367 | 1.082 |
| total | 482844 | 0.006 |

## Redis 本轮边界增量

- Commands：62791；input：5589074 bytes；output：7002420 bytes
- Hits/Misses：24027/180；run hit rate：99.26%
- Evicted/Rejected：0/0；ops/s max：1427；safety epoch：476 -> 476

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.614 |
| client:loadtest | cpu_percent_total | 121.818 |
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
| docker:zg-canal | cpu_percent | 2.010 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 80.000 |
| docker:zg-es | cpu_percent | 2.300 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.540 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 161.950 |
| docker:zg-kafka | memory_percent | 8.360 |
| docker:zg-kafka | pids | 119.000 |
| docker:zg-zk | cpu_percent | 0.170 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20181081.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.746 |
| process:counter | cpu_seconds_total | 51.109 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43220992.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.438 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37167104.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 248.529 |
| process:knowpost | cpu_seconds_total | 1991.844 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 110882816.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.557 |
| process:relation | cpu_seconds_total | 0.906 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 46116864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.531 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37007360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.484 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34516992.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178084713.000 |
| redis | connected_clients | 55.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 476.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677002.000 |
| redis | keyspace_hits | 977471290.000 |
| redis | keyspace_misses | 5365674.000 |
| redis | net_input_bytes | 40587088068.000 |
| redis | net_output_bytes | 180207792347.000 |
| redis | ops_per_sec | 1427.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 39271.000 |
| redis | used_memory_bytes | 98226632.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
