# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-treatment-fixed-rpc-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:22:50+08:00
- 采样时长：1m0.1703169s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 6124737 | 6124737 | 0 | 0 | 102075.49 | 0.529 | 1.075 | 1.162 | 1.689 | 13.576 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 146791 | 0.024 |
| counter | 10815 | 0.002 |
| mysql | 4 | 0.000 |
| redis | 145470 | 0.024 |
| relation | 10815 | 0.002 |

- Cold compute：22273（0.004 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 6044765 |
| l2_fresh | 79972 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=2 pending=2

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22273 | 0.174 |
| counter | 10815 | 0.515 |
| hydrate | 22273 | 0.201 |
| inbox | 22273 | 0.172 |
| merge_dedup | 22273 | 0.006 |
| relation | 10815 | 1.128 |
| route | 22273 | 0.809 |
| total | 6124737 | 0.008 |

## Redis 本轮边界增量

- Commands：460290；input：68954305 bytes；output：167502061 bytes
- Hits/Misses：626692/12670；run hit rate：98.02%
- Evicted/Rejected：0/0；ops/s max：13377；safety epoch：604 -> 604

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 24.715 |
| client:loadtest | cpu_percent_total | 395.440 |
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
| docker:zg-canal | cpu_percent | 3.170 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.630 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.260 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 267.110 |
| docker:zg-kafka | memory_percent | 8.840 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 70.420 |
| docker:zg-zk | memory_percent | 1.710 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29174805.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 39.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 29.342 |
| process:counter | cpu_seconds_total | 12.672 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 52994048.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.047 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 37240832.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 647.930 |
| process:knowpost | cpu_seconds_total | 359.109 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 102940672.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 37.864 |
| process:relation | cpu_seconds_total | 9.453 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 54542336.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.047 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35528704.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.589 |
| process:user-storage | cpu_seconds_total | 0.188 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 34377728.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 292317971.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 604.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 708276.000 |
| redis | keyspace_hits | 1132320908.000 |
| redis | keyspace_misses | 16905901.000 |
| redis | net_input_bytes | 53052403040.000 |
| redis | net_output_bytes | 208305692123.000 |
| redis | ops_per_sec | 13377.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54209.000 |
| redis | used_memory_bytes | 124078168.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
