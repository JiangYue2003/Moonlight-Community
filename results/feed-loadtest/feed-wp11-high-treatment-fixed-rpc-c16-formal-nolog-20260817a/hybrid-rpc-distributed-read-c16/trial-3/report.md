# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-fixed-rpc-c16-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:37:44+08:00
- 采样时长：1m0.0312193s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 566215 | 566215 | 0 | 0 | 9436.74 | 1.648 | 2.974 | 3.732 | 5.783 | 16.317 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1027974 | 1.816 |
| counter | 147699 | 0.261 |
| mysql | 22 | 0.000 |
| redis | 1090322 | 1.926 |
| relation | 147699 | 0.261 |

- Cold compute：227580（0.402 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 157840 |
| l1_stale | 474 |
| l2_fresh | 248706 |
| l2_stale | 138275 |
| miss | 20920 |

- L1+L2 Fresh ratio：71.80%
- Refresh max：queue=0 active=29 pending=29

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 227580 | 0.977 |
| counter | 147699 | 1.387 |
| hydrate | 227580 | 0.968 |
| inbox | 227580 | 0.976 |
| merge_dedup | 227580 | 0.002 |
| relation | 147699 | 1.580 |
| route | 227580 | 1.939 |
| total | 566215 | 1.472 |

## Redis 本轮边界增量

- Commands：3488503；input：479043136 bytes；output：763213667 bytes
- Hits/Misses：4582447/376943；run hit rate：92.40%
- Evicted/Rejected：0/0；ops/s max：62637；safety epoch：616 -> 616

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.667 |
| client:loadtest | cpu_percent_total | 122.671 |
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
| docker:zg-canal | cpu_percent | 2.860 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.190 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.480 |
| docker:zg-etcd | memory_percent | 0.310 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 213.110 |
| docker:zg-kafka | memory_percent | 8.840 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 54.470 |
| docker:zg-zk | memory_percent | 1.810 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 30119083.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 105.458 |
| process:counter | cpu_seconds_total | 380.250 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 54886400.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 618.344 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 45223936.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 442.480 |
| process:knowpost | cpu_seconds_total | 2922.344 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 237051904.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 175.235 |
| process:relation | cpu_seconds_total | 576.172 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 57090048.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 0.500 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35782656.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 152.141 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 41082880.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 314857675.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 616.000 |
| redis | hit_rate | 0.984 |
| redis | keys | 811489.000 |
| redis | keyspace_hits | 1161749062.000 |
| redis | keyspace_misses | 19243666.000 |
| redis | net_input_bytes | 56145014883.000 |
| redis | net_output_bytes | 213375765422.000 |
| redis | ops_per_sec | 62637.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 55103.000 |
| redis | used_memory_bytes | 228782688.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
