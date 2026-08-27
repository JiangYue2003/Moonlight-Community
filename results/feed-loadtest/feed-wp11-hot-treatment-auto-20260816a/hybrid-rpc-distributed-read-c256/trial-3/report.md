# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:27:37+08:00
- 采样时长：1m0.0624283s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 473927 | 473927 | 0 | 0 | 7894.27 | 31.861 | 38.962 | 41.151 | 46.043 | 231.844 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2442 | 0.005 |
| counter | 181 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 2401 | 0.005 |
| relation | 181 | 0.000 |

- Cold compute：365（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 472580 |
| l1_stale | 0 |
| l2_fresh | 1347 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=1 pending=1

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 365 | 0.144 |
| counter | 181 | 0.689 |
| hydrate | 365 | 0.213 |
| inbox | 365 | 0.141 |
| merge_dedup | 365 | 0.017 |
| relation | 181 | 1.454 |
| route | 365 | 1.077 |
| total | 473927 | 0.006 |

## Redis 本轮边界增量

- Commands：62851；input：5583389 bytes；output：6968301 bytes
- Hits/Misses：23950/181；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1327；safety epoch：487 -> 487

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.551 |
| client:loadtest | cpu_percent_total | 120.812 |
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
| docker:zg-canal | cpu_percent | 2.070 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.430 |
| docker:zg-es | memory_percent | 11.950 |
| docker:zg-es | pids | 157.000 |
| docker:zg-etcd | cpu_percent | 4.320 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 174.420 |
| docker:zg-kafka | memory_percent | 8.610 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 44.390 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20185026.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 8.166 |
| process:counter | cpu_seconds_total | 90.531 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 42463232.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.547 |
| process:gateway | cpu_seconds_total | 0.562 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37306368.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 247.934 |
| process:knowpost | cpu_seconds_total | 3646.953 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 133935104.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.549 |
| process:relation | cpu_seconds_total | 3.953 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 48525312.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.948 |
| process:search | cpu_seconds_total | 0.922 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37097472.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.719 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34787328.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178951385.000 |
| redis | connected_clients | 55.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 487.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677174.000 |
| redis | keyspace_hits | 977796492.000 |
| redis | keyspace_misses | 5371665.000 |
| redis | net_input_bytes | 40664588357.000 |
| redis | net_output_bytes | 180301110420.000 |
| redis | ops_per_sec | 1327.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 40096.000 |
| redis | used_memory_bytes | 98122568.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
