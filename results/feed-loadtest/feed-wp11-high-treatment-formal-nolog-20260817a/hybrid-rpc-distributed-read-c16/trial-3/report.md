# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-formal-nolog-20260817a`
- 开始时间：2026-08-17T00:35:14+08:00
- 采样时长：1m0.03237s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 549819 | 549819 | 0 | 0 | 9163.24 | 1.814 | 3.128 | 3.751 | 5.885 | 36.172 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1009783 | 1.837 |
| counter | 146677 | 0.267 |
| mysql | 19 | 0.000 |
| redis | 1072563 | 1.951 |
| relation | 146677 | 0.267 |

- Cold compute：224469（0.408 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 149935 |
| l1_stale | 421 |
| l2_fresh | 240651 |
| l2_stale | 137754 |
| miss | 21058 |

- L1+L2 Fresh ratio：71.04%
- Refresh max：queue=0 active=25 pending=25

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 224469 | 1.002 |
| counter | 146677 | 1.409 |
| hydrate | 224469 | 0.988 |
| inbox | 224469 | 1.001 |
| merge_dedup | 224469 | 0.002 |
| relation | 146677 | 1.598 |
| route | 224469 | 1.979 |
| total | 549819 | 1.524 |

## Redis 本轮边界增量

- Commands：3440355；input：472239870 bytes；output：750061303 bytes
- Hits/Misses：4518626/373238；run hit rate：92.37%
- Evicted/Rejected：0/0；ops/s max：62674；safety epoch：567 -> 567

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.389 |
| client:loadtest | cpu_percent_total | 118.217 |
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
| docker:zg-canal | cpu_percent | 2.660 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.000 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.220 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 236.110 |
| docker:zg-kafka | memory_percent | 8.820 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 65.670 |
| docker:zg-zk | memory_percent | 1.500 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 24116383.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 116.802 |
| process:counter | cpu_seconds_total | 611.281 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54419456.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 1840.922 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 50163712.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 434.594 |
| process:knowpost | cpu_seconds_total | 4622.656 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 236969984.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 157.359 |
| process:relation | cpu_seconds_total | 741.531 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62496768.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.548 |
| process:search | cpu_seconds_total | 1.469 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36184064.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.547 |
| process:user-storage | cpu_seconds_total | 456.812 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 43601920.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 228117868.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 567.000 |
| redis | hit_rate | 0.993 |
| redis | keys | 803016.000 |
| redis | keyspace_hits | 1030911847.000 |
| redis | keyspace_misses | 7924854.000 |
| redis | net_input_bytes | 45850778749.000 |
| redis | net_output_bytes | 193777077840.000 |
| redis | ops_per_sec | 62674.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51353.000 |
| redis | used_memory_bytes | 220388024.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
