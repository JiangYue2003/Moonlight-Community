# Feed 压测报告：hybrid / gateway / distributed-read-c8

- Run ID：`feed-wp11-high-treatment-fixed-gateway-c8-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:40:38+08:00
- 采样时长：1m0.0127191s
- 并发：8
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 81314 | 81314 | 0 | 0 | 1355.07 | 4.123 | 15.541 | 19.290 | 25.322 | 42.638 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 167986 | 2.066 |
| counter | 59722 | 0.734 |
| mysql | 25 | 0.000 |
| redis | 275591 | 3.389 |
| relation | 59722 | 0.734 |

- Cold compute：66305（0.815 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 4623 |
| l1_stale | 16 |
| l2_fresh | 13775 |
| l2_stale | 27016 |
| miss | 35884 |

- L1+L2 Fresh ratio：22.63%
- Refresh max：queue=0 active=9 pending=9

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 66305 | 1.075 |
| counter | 59722 | 1.480 |
| hydrate | 66305 | 1.007 |
| inbox | 66305 | 1.074 |
| merge_dedup | 66305 | 0.002 |
| relation | 59722 | 2.235 |
| route | 66305 | 3.367 |
| total | 81314 | 4.603 |

## Redis 本轮边界增量

- Commands：1085111；input：141603806 bytes；output：164946004 bytes
- Hits/Misses：1307327/156416；run hit rate：89.31%
- Evicted/Rejected：0/0；ops/s max：28132；safety epoch：618 -> 618

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.286 |
| client:loadtest | cpu_percent_total | 36.581 |
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
| docker:zg-canal | cpu_percent | 2.170 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.060 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.250 |
| docker:zg-etcd | memory_percent | 0.310 |
| docker:zg-etcd | pids | 26.000 |
| docker:zg-kafka | cpu_percent | 175.530 |
| docker:zg-kafka | memory_percent | 8.850 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 53.550 |
| docker:zg-zk | memory_percent | 1.590 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 30281204.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 97.410 |
| process:counter | cpu_seconds_total | 485.641 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 55197696.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 194.501 |
| process:gateway | cpu_seconds_total | 835.703 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 51154944.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 271.505 |
| process:knowpost | cpu_seconds_total | 3259.938 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 163061760.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 139.117 |
| process:relation | cpu_seconds_total | 729.797 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 56627200.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.594 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35622912.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 83.633 |
| process:user-storage | cpu_seconds_total | 244.484 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 51097600.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 317901220.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 618.000 |
| redis | hit_rate | 0.984 |
| redis | keys | 799809.000 |
| redis | keyspace_hits | 1165385823.000 |
| redis | keyspace_misses | 19663531.000 |
| redis | net_input_bytes | 56540921033.000 |
| redis | net_output_bytes | 213850218185.000 |
| redis | ops_per_sec | 28132.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 55277.000 |
| redis | used_memory_bytes | 215369944.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
