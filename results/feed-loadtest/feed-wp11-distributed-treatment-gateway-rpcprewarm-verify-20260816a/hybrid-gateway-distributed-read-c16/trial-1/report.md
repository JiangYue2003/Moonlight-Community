# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-gateway-rpcprewarm-verify-20260816a`
- 开始时间：2026-08-17T00:00:34+08:00
- 采样时长：5.0095618s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 16448 | 16448 | 0 | 0 | 3283.99 | 4.135 | 8.570 | 10.367 | 14.373 | 28.923 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 9085 | 0.552 |
| counter | 120 | 0.007 |
| mysql | 0 | 0.000 |
| redis | 9046 | 0.550 |
| relation | 120 | 0.007 |

- Cold compute：1294（0.079 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 11237 |
| l1_stale | 8 |
| l2_fresh | 5052 |
| l2_stale | 151 |
| miss | 0 |

- L1+L2 Fresh ratio：99.03%
- Refresh max：queue=0 active=14 pending=14

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1294 | 3.241 |
| counter | 120 | 3.963 |
| hydrate | 1294 | 3.246 |
| inbox | 1294 | 3.239 |
| merge_dedup | 1294 | 0.004 |
| relation | 120 | 5.787 |
| route | 1294 | 0.911 |
| total | 16448 | 1.373 |

## Redis 本轮边界增量

- Commands：23367；input：3912009 bytes；output：10428187 bytes
- Hits/Misses：34386/227；run hit rate：99.34%
- Evicted/Rejected：0/0；ops/s max：13921；safety epoch：550 -> 550

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 12.671 |
| client:loadtest | cpu_percent_total | 202.737 |
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
| docker:zg-canal | cpu_percent | 0.130 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.460 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 3.900 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 59.080 |
| docker:zg-kafka | memory_percent | 8.270 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 40.010 |
| docker:zg-zk | memory_percent | 1.700 |
| docker:zg-zk | pids | 107.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23389957.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 43.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 42.894 |
| process:counter | cpu_seconds_total | 161.719 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 52563968.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 464.234 |
| process:gateway | cpu_seconds_total | 460.500 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 52146176.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 257.361 |
| process:knowpost | cpu_seconds_total | 2467.219 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 103829504.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 50.692 |
| process:relation | cpu_seconds_total | 146.531 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 57577472.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.578 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36360192.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 138.697 |
| process:user-storage | cpu_seconds_total | 128.938 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 49995776.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 210121577.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 550.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 702555.000 |
| redis | keyspace_hits | 1009305316.000 |
| redis | keyspace_misses | 6123310.000 |
| redis | net_input_bytes | 43457808737.000 |
| redis | net_output_bytes | 190146054942.000 |
| redis | ops_per_sec | 13921.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49218.000 |
| redis | used_memory_bytes | 120888320.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
