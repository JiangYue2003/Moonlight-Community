# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-high-treatment-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T00:48:07+08:00
- 采样时长：1m0.0245228s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 77930 | 77930 | 0 | 0 | 1298.42 | 8.345 | 30.138 | 35.606 | 45.208 | 94.459 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 158275 | 2.031 |
| counter | 58188 | 0.747 |
| mysql | 18 | 0.000 |
| redis | 266745 | 3.423 |
| relation | 58188 | 0.747 |

- Cold compute：64312（0.825 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 4070 |
| l1_stale | 26 |
| l2_fresh | 12669 |
| l2_stale | 24967 |
| miss | 36198 |

- L1+L2 Fresh ratio：21.48%
- Refresh max：queue=0 active=13 pending=13

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 64312 | 2.543 |
| counter | 58188 | 3.079 |
| hydrate | 64312 | 2.415 |
| inbox | 64312 | 2.542 |
| merge_dedup | 64312 | 0.002 |
| relation | 58188 | 3.886 |
| route | 64312 | 6.324 |
| total | 77930 | 10.775 |

## Redis 本轮边界增量

- Commands：1014929；input：134660335 bytes；output：157709035 bytes
- Hits/Misses：1267224/153257；run hit rate：89.21%
- Evicted/Rejected：0/0；ops/s max：24556；safety epoch：577 -> 577

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.605 |
| client:loadtest | cpu_percent_total | 41.676 |
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
| docker:zg-canal | cpu_percent | 2.720 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.430 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.930 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 165.020 |
| docker:zg-kafka | memory_percent | 8.820 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 52.200 |
| docker:zg-zk | memory_percent | 1.670 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 25297383.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 95.248 |
| process:counter | cpu_seconds_total | 1047.984 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54857728.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 188.543 |
| process:gateway | cpu_seconds_total | 2177.172 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 51240960.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 301.982 |
| process:knowpost | cpu_seconds_total | 6685.094 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 162615296.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 142.484 |
| process:relation | cpu_seconds_total | 1436.672 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 62263296.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.688 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36208640.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 87.473 |
| process:user-storage | cpu_seconds_total | 589.406 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 57085952.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 254556261.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 577.000 |
| redis | hit_rate | 0.990 |
| redis | keys | 799011.000 |
| redis | keyspace_hits | 1064930649.000 |
| redis | keyspace_misses | 10931480.000 |
| redis | net_input_bytes | 49409553391.000 |
| redis | net_output_bytes | 199258303334.000 |
| redis | ops_per_sec | 24556.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52126.000 |
| redis | used_memory_bytes | 217145864.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
