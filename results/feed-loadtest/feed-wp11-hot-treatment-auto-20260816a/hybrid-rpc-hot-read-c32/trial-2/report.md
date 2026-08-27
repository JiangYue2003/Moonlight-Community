# Feed 压测报告：hybrid / rpc / hot-read-c32

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:00:05+08:00
- 采样时长：1m0.0305447s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 484725 | 484725 | 0 | 0 | 8078.18 | 3.785 | 5.008 | 5.388 | 6.371 | 20.070 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 173 | 0.000 |
| counter | 9 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 119 | 0.000 |
| relation | 9 | 0.000 |

- Cold compute：18（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 484606 |
| l2_fresh | 119 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 18 | 0.145 |
| counter | 9 | 1.915 |
| hydrate | 18 | 0.118 |
| inbox | 18 | 0.116 |
| merge_dedup | 18 | 0.000 |
| relation | 9 | 2.130 |
| route | 18 | 2.022 |
| total | 484725 | 0.005 |

## Redis 本轮边界增量

- Commands：53121；input：3796970 bytes；output：1422973 bytes
- Hits/Misses：1326/9；run hit rate：99.33%
- Evicted/Rejected：0/0；ops/s max：1091；safety epoch：465 -> 465

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.901 |
| client:loadtest | cpu_percent_total | 126.420 |
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
| docker:zg-canal | cpu_percent | 0.180 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.320 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.260 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 179.200 |
| docker:zg-kafka | memory_percent | 8.580 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 49.670 |
| docker:zg-zk | memory_percent | 1.230 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20179871.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.747 |
| process:counter | cpu_seconds_total | 13.453 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 39718912.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 2.337 |
| process:gateway | cpu_seconds_total | 0.125 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37277696.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 246.852 |
| process:knowpost | cpu_seconds_total | 339.969 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 115589120.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.774 |
| process:relation | cpu_seconds_total | 0.234 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 39325696.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.203 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37412864.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.895 |
| process:user-storage | cpu_seconds_total | 0.172 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34787328.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 177338834.000 |
| redis | connected_clients | 16.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 465.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676784.000 |
| redis | keyspace_hits | 977424818.000 |
| redis | keyspace_misses | 5364426.000 |
| redis | net_input_bytes | 40532203465.000 |
| redis | net_output_bytes | 180181679111.000 |
| redis | ops_per_sec | 1091.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 38444.000 |
| redis | used_memory_bytes | 95590888.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
