# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:38:07+08:00
- 采样时长：1m0.0351418s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 562758 | 562758 | 0 | 0 | 9378.92 | 1.682 | 3.002 | 3.728 | 5.727 | 35.710 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1021923 | 1.816 |
| counter | 147358 | 0.262 |
| mysql | 29 | 0.000 |
| redis | 1084483 | 1.927 |
| relation | 147358 | 0.262 |

- Cold compute：226289（0.402 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 156373 |
| l1_stale | 459 |
| l2_fresh | 247325 |
| l2_stale | 137614 |
| miss | 20987 |

- L1+L2 Fresh ratio：71.74%
- Refresh max：queue=0 active=27 pending=27

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 226289 | 0.988 |
| counter | 147358 | 1.396 |
| hydrate | 226289 | 0.978 |
| inbox | 226289 | 0.987 |
| merge_dedup | 226289 | 0.002 |
| relation | 147358 | 1.576 |
| route | 226289 | 1.950 |
| total | 562758 | 1.487 |

## Redis 本轮边界增量

- Commands：3471606；input：476375566 bytes；output：759040248 bytes
- Hits/Misses：4558992/375533；run hit rate：92.39%
- Evicted/Rejected：0/0；ops/s max：64310；safety epoch：569 -> 569

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.748 |
| client:loadtest | cpu_percent_total | 123.964 |
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
| docker:zg-canal | cpu_percent | 3.110 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.250 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.760 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 214.150 |
| docker:zg-kafka | memory_percent | 8.780 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 51.450 |
| docker:zg-zk | memory_percent | 1.640 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 24338004.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 34.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 125.555 |
| process:counter | cpu_seconds_total | 681.203 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54431744.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1840.953 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 45789184.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 455.004 |
| process:knowpost | cpu_seconds_total | 4971.156 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 238194688.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 185.949 |
| process:relation | cpu_seconds_total | 847.031 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62636032.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.516 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36200448.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.547 |
| process:user-storage | cpu_seconds_total | 456.938 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 43601920.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 233284788.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 569.000 |
| redis | hit_rate | 0.992 |
| redis | keys | 799874.000 |
| redis | keyspace_hits | 1037478610.000 |
| redis | keyspace_misses | 8502021.000 |
| redis | net_input_bytes | 46545132682.000 |
| redis | net_output_bytes | 194845970562.000 |
| redis | ops_per_sec | 64310.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51526.000 |
| redis | used_memory_bytes | 216852176.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
