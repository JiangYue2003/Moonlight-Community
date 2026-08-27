# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-treatment-scout-nolog-20260816a`
- 开始时间：2026-08-16T23:42:11+08:00
- 采样时长：10.0271508s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 912752 | 912752 | 0 | 0 | 91270.56 | 0.528 | 1.099 | 1.256 | 1.994 | 11.386 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 24726 | 0.027 |
| counter | 1304 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 24471 | 0.027 |
| relation | 1304 | 0.001 |

- Cold compute：3618（0.004 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 898880 |
| l1_stale | 0 |
| l2_fresh | 13872 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=1 pending=1

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 3618 | 0.163 |
| counter | 1304 | 0.543 |
| hydrate | 3618 | 0.187 |
| inbox | 3618 | 0.162 |
| merge_dedup | 3618 | 0.006 |
| relation | 1304 | 1.125 |
| route | 3618 | 0.610 |
| total | 912752 | 0.009 |

## Redis 本轮边界增量

- Commands：71958；input：11124509 bytes；output：28226551 bytes
- Hits/Misses：100134/1605；run hit rate：98.42%
- Evicted/Rejected：0/0；ops/s max：13364；safety epoch：536 -> 536

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 24.173 |
| client:loadtest | cpu_percent_total | 386.762 |
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
| docker:zg-canal | cpu_percent | 3.300 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.710 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.290 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 140.220 |
| docker:zg-kafka | memory_percent | 8.280 |
| docker:zg-kafka | pids | 101.000 |
| docker:zg-zk | cpu_percent | 0.200 |
| docker:zg-zk | memory_percent | 1.420 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23269664.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 41.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 23.183 |
| process:counter | cpu_seconds_total | 17.797 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 50475008.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 35.250 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 47845376.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 521.397 |
| process:knowpost | cpu_seconds_total | 170.875 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 97943552.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 30.910 |
| process:relation | cpu_seconds_total | 10.734 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 55345152.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36208640.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 11.125 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 47288320.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 205496580.000 |
| redis | connected_clients | 200.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 536.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 690415.000 |
| redis | keyspace_hits | 1003643459.000 |
| redis | keyspace_misses | 5967977.000 |
| redis | net_input_bytes | 42794195728.000 |
| redis | net_output_bytes | 188733443625.000 |
| redis | ops_per_sec | 13364.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48120.000 |
| redis | used_memory_bytes | 110928096.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
