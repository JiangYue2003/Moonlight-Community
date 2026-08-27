# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-treatment-formal-nolog-20260816a`
- 开始时间：2026-08-16T23:52:48+08:00
- 采样时长：1m0.1625534s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 5897721 | 5897721 | 0 | 0 | 98294.95 | 0.528 | 1.078 | 1.173 | 1.683 | 22.417 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 146651 | 0.025 |
| counter | 10959 | 0.002 |
| mysql | 3 | 0.000 |
| redis | 145375 | 0.025 |
| relation | 10959 | 0.002 |

- Cold compute：22250（0.004 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 5817820 |
| l1_stale | 0 |
| l2_fresh | 79901 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=2 pending=2

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22250 | 0.170 |
| counter | 10959 | 0.507 |
| hydrate | 22250 | 0.190 |
| inbox | 22250 | 0.168 |
| merge_dedup | 22250 | 0.005 |
| relation | 10959 | 1.149 |
| route | 22250 | 0.828 |
| total | 5897721 | 0.008 |

## Redis 本轮边界增量

- Commands：459002；input：68770534 bytes；output：167412717 bytes
- Hits/Misses：627214/12825；run hit rate：98.00%
- Evicted/Rejected：0/0；ops/s max：12223；safety epoch：547 -> 547

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 24.247 |
| client:loadtest | cpu_percent_total | 387.959 |
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
| docker:zg-canal | cpu_percent | 2.780 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.520 |
| docker:zg-es | memory_percent | 12.030 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.520 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 247.410 |
| docker:zg-kafka | memory_percent | 8.850 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 75.080 |
| docker:zg-zk | memory_percent | 1.700 |
| docker:zg-zk | pids | 107.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23370981.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 46.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 23.197 |
| process:counter | cpu_seconds_total | 109.781 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 53899264.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 195.719 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 45604864.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 610.421 |
| process:knowpost | cpu_seconds_total | 2242.188 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 106737664.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 35.572 |
| process:relation | cpu_seconds_total | 97.516 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 57937920.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.453 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36274176.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.554 |
| process:user-storage | cpu_seconds_total | 59.672 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 41361408.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 209318863.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 547.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 701836.000 |
| redis | keyspace_hits | 1008599394.000 |
| redis | keyspace_misses | 6097731.000 |
| redis | net_input_bytes | 43356647809.000 |
| redis | net_output_bytes | 189990523284.000 |
| redis | ops_per_sec | 12223.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48807.000 |
| redis | used_memory_bytes | 123501224.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
