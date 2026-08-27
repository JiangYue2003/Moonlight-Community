# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:26:05+08:00
- 采样时长：1m0.0285471s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 177610 | 177610 | 0 | 0 | 2959.40 | 21.351 | 24.540 | 26.105 | 30.593 | 48.736 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 107 | 0.001 |
| redis | 532937 | 3.001 |
| relation | 177610 | 1.000 |

- Cold compute：177610（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 177610 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 177610 | 0.186 |
| counter | 240 | 0.718 |
| hydrate | 177610 | 0.287 |
| inbox | 177610 | 0.203 |
| merge_dedup | 177610 | 0.012 |
| relation | 177610 | 20.450 |
| route | 177610 | 0.004 |
| total | 177610 | 21.164 |

## Redis 本轮边界增量

- Commands：1486167；input：338628301 bytes；output：1546969749 bytes
- Hits/Misses：8353633/127；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26211；safety epoch：409 -> 409

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.679 |
| client:loadtest | cpu_percent_total | 74.860 |
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
| docker:zg-canal | cpu_percent | 2.790 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.620 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 6.340 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 186.990 |
| docker:zg-kafka | memory_percent | 8.420 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 55.880 |
| docker:zg-zk | memory_percent | 1.030 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 12968075.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 69.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.805 |
| process:counter | cpu_seconds_total | 94.219 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44732416.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 0.391 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 38469632.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 290.205 |
| process:knowpost | cpu_seconds_total | 3428.297 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 95072256.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 284.277 |
| process:relation | cpu_seconds_total | 3241.609 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 90144768.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.641 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37138432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.769 |
| process:user-storage | cpu_seconds_total | 0.672 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35287040.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 104978677.000 |
| redis | connected_clients | 237.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 409.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698754.000 |
| redis | keyspace_hits | 600587313.000 |
| redis | keyspace_misses | 96467.000 |
| redis | net_input_bytes | 24423796918.000 |
| redis | net_output_bytes | 109799549392.000 |
| redis | ops_per_sec | 26211.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22004.000 |
| redis | used_memory_bytes | 106632424.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
