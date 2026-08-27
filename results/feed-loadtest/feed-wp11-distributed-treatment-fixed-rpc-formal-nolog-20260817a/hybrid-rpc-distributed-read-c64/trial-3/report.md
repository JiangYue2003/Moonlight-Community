# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-treatment-fixed-rpc-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:25:14+08:00
- 采样时长：1m0.1547888s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 5999125 | 5999125 | 0 | 0 | 99984.97 | 0.529 | 1.080 | 1.174 | 1.705 | 23.901 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 146861 | 0.024 |
| counter | 10890 | 0.002 |
| mysql | 3 | 0.000 |
| redis | 145555 | 0.024 |
| relation | 10890 | 0.002 |

- Cold compute：22313（0.004 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 5919203 |
| l1_stale | 0 |
| l2_fresh | 79922 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=1 pending=1

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22313 | 0.169 |
| counter | 10890 | 0.521 |
| hydrate | 22313 | 0.204 |
| inbox | 22313 | 0.167 |
| merge_dedup | 22313 | 0.005 |
| relation | 10890 | 1.167 |
| route | 22313 | 0.836 |
| total | 5999125 | 0.008 |

## Redis 本轮边界增量

- Commands：461016；input：69054039 bytes；output：167590790 bytes
- Hits/Misses：627832/12750；run hit rate：98.01%
- Evicted/Rejected：0/0；ops/s max：13478；safety epoch：606 -> 606

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 24.704 |
| client:loadtest | cpu_percent_total | 395.256 |
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
| docker:zg-canal | cpu_percent | 0.200 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.100 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.810 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 283.730 |
| docker:zg-kafka | memory_percent | 8.810 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 69.820 |
| docker:zg-zk | memory_percent | 1.730 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29201478.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 38.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 27.032 |
| process:counter | cpu_seconds_total | 31.938 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 53473280.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.770 |
| process:gateway | cpu_seconds_total | 0.062 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 37318656.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 578.689 |
| process:knowpost | cpu_seconds_total | 1051.438 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 96931840.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 37.883 |
| process:relation | cpu_seconds_total | 28.625 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 56705024.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35565568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.250 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 34627584.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 293382354.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 606.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 709167.000 |
| redis | keyspace_hits | 1133762595.000 |
| redis | keyspace_misses | 16939380.000 |
| redis | net_input_bytes | 53212156654.000 |
| redis | net_output_bytes | 208680404230.000 |
| redis | ops_per_sec | 13478.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54353.000 |
| redis | used_memory_bytes | 126076424.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
