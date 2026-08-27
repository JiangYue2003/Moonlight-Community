# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:01:37+08:00
- 采样时长：1m0.0170345s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 77139 | 77139 | 0 | 0 | 1285.42 | 9.647 | 25.577 | 29.405 | 37.677 | 69.235 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 91164 | 1.182 |
| counter | 8969 | 0.116 |
| mysql | 7 | 0.000 |
| redis | 90638 | 1.175 |
| relation | 8969 | 0.116 |

- Cold compute：16941（0.220 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 36498 |
| l1_stale | 300 |
| l2_fresh | 32871 |
| l2_stale | 7470 |
| miss | 0 |

- L1+L2 Fresh ratio：89.93%
- Refresh max：queue=0 active=27 pending=27

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 16941 | 7.611 |
| counter | 8969 | 9.526 |
| hydrate | 16941 | 7.601 |
| inbox | 16941 | 7.609 |
| merge_dedup | 16941 | 0.003 |
| relation | 8969 | 12.239 |
| route | 16941 | 11.552 |
| total | 77139 | 7.242 |

## Redis 本轮边界增量

- Commands：280548；input：47956210 bytes；output：100541460 bytes
- Hits/Misses：443406/10385；run hit rate：97.71%
- Evicted/Rejected：0/0；ops/s max：13985；safety epoch：551 -> 551

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.569 |
| client:loadtest | cpu_percent_total | 105.100 |
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
| docker:zg-es | cpu_percent | 2.730 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.070 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 176.950 |
| docker:zg-kafka | memory_percent | 8.740 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 55.550 |
| docker:zg-zk | memory_percent | 1.450 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23402853.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 48.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 58.809 |
| process:counter | cpu_seconds_total | 192.500 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 53866496.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 430.979 |
| process:gateway | cpu_seconds_total | 667.625 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 52850688.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 280.213 |
| process:knowpost | cpu_seconds_total | 2635.953 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 99688448.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 105.274 |
| process:relation | cpu_seconds_total | 187.672 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 58040320.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 4.147 |
| process:search | cpu_seconds_total | 0.609 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36343808.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 94.398 |
| process:user-storage | cpu_seconds_total | 180.141 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 52609024.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 210518360.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 551.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 704366.000 |
| redis | keyspace_hits | 1009835075.000 |
| redis | keyspace_misses | 6139303.000 |
| redis | net_input_bytes | 43519465676.000 |
| redis | net_output_bytes | 190262319478.000 |
| redis | ops_per_sec | 13985.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49336.000 |
| redis | used_memory_bytes | 122663408.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
