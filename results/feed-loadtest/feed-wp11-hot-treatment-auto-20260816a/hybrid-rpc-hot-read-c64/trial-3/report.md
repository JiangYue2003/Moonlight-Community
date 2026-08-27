# Feed 压测报告：hybrid / rpc / hot-read-c64

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:05:06+08:00
- 采样时长：1m0.0338282s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 490433 | 490433 | 0 | 0 | 8173.10 | 7.643 | 10.009 | 10.594 | 11.870 | 24.290 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 168 | 0.000 |
| counter | 10 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 122 | 0.000 |
| relation | 10 | 0.000 |

- Cold compute：19（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 490322 |
| l1_stale | 0 |
| l2_fresh | 111 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 19 | 0.061 |
| counter | 10 | 1.826 |
| hydrate | 19 | 0.147 |
| inbox | 19 | 0.061 |
| merge_dedup | 19 | 0.055 |
| relation | 10 | 1.949 |
| route | 19 | 1.987 |
| total | 490433 | 0.005 |

## Redis 本轮边界增量

- Commands：53215；input：3807527 bytes；output：1434828 bytes
- Hits/Misses：1397/10；run hit rate：99.29%
- Evicted/Rejected：0/0；ops/s max：1114；safety epoch：469 -> 469

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.850 |
| client:loadtest | cpu_percent_total | 125.606 |
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
| docker:zg-canal | cpu_percent | 2.120 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.580 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.510 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 110.800 |
| docker:zg-kafka | memory_percent | 8.460 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 48.340 |
| docker:zg-zk | memory_percent | 1.380 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20180211.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 9.618 |
| process:counter | cpu_seconds_total | 27.141 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 40120320.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 5.412 |
| process:gateway | cpu_seconds_total | 0.281 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37543936.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 235.438 |
| process:knowpost | cpu_seconds_total | 946.688 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 114921472.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.548 |
| process:relation | cpu_seconds_total | 0.438 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 39100416.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.328 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37064704.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.266 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34721792.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 177605611.000 |
| redis | connected_clients | 16.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 469.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676858.000 |
| redis | keyspace_hits | 977431346.000 |
| redis | keyspace_misses | 5364700.000 |
| redis | net_input_bytes | 40551324911.000 |
| redis | net_output_bytes | 180188706445.000 |
| redis | ops_per_sec | 1114.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 38745.000 |
| redis | used_memory_bytes | 95531704.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
