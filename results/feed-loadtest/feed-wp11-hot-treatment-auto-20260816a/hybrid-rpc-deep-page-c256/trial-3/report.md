# Feed 压测报告：hybrid / rpc / deep-page-c256

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:42:41+08:00
- 采样时长：1m0.0660281s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 183403 | 183403 | 0 | 0 | 3054.11 | 70.746 | 140.994 | 171.132 | 242.700 | 569.075 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 0 | 0.000 |
| counter | 220 | 0.001 |
| mysql | 183403 | 1.000 |
| redis | 183403 | 1.000 |
| relation | 220 | 0.001 |

- Cold compute：183403（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 183403 |
| l1_fresh | 0 |
| l1_stale | 0 |
| l2_fresh | 0 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 183403 | 0.170 |
| counter | 220 | 0.719 |
| hydrate | 183403 | 63.462 |
| inbox | 183403 | 0.168 |
| merge_dedup | 183403 | 0.014 |
| relation | 220 | 1.509 |
| route | 183403 | 0.007 |
| total | 183403 | 63.671 |

## Redis 本轮边界增量

- Commands：1159998；input：83520536 bytes；output：483644817 bytes
- Hits/Misses：1107188/220；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：21128；safety epoch：499 -> 499

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.123 |
| client:loadtest | cpu_percent_total | 65.969 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.750 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.960 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 173.930 |
| docker:zg-kafka | memory_percent | 8.390 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 58.870 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22729842.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 87.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.054 |
| process:counter | cpu_seconds_total | 139.969 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 44453888.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.772 |
| process:gateway | cpu_seconds_total | 0.672 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 36818944.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 474.477 |
| process:knowpost | cpu_seconds_total | 7234.547 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 134729728.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.094 |
| process:relation | cpu_seconds_total | 8.516 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 49717248.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 1.266 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36990976.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.797 |
| process:user-storage | cpu_seconds_total | 1.109 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34873344.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 195086319.000 |
| redis | connected_clients | 205.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 499.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677199.000 |
| redis | keyspace_hits | 993138182.000 |
| redis | keyspace_misses | 5378385.000 |
| redis | net_input_bytes | 41827300046.000 |
| redis | net_output_bytes | 186998663912.000 |
| redis | ops_per_sec | 21128.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 41000.000 |
| redis | used_memory_bytes | 101732248.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
