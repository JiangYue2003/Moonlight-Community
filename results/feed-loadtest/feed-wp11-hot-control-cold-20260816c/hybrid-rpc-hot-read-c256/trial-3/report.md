# Feed 压测报告：hybrid / rpc / hot-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:18:31+08:00
- 采样时长：1m0.089457s
- 并发：256
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 179733 | 179733 | 0 | 0 | 2991.85 | 72.760 | 141.096 | 170.445 | 240.071 | 559.201 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 121 | 0.001 |
| redis | 539320 | 3.001 |
| relation | 179733 | 1.000 |

- Cold compute：179733（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 179733 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 179733 | 0.203 |
| counter | 12 | 1.845 |
| hydrate | 179733 | 0.299 |
| inbox | 179733 | 0.222 |
| merge_dedup | 179733 | 0.011 |
| relation | 179733 | 84.294 |
| route | 179733 | 0.003 |
| total | 179733 | 85.054 |

## Redis 本轮边界增量

- Commands：1497445；input：342461700 bytes；output：1565289954 bytes
- Hits/Misses：8447720/121；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27160；safety epoch：403 -> 403

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.533 |
| client:loadtest | cpu_percent_total | 72.522 |
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
| docker:zg-canal | cpu_percent | 1.570 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.080 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.830 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 230.980 |
| docker:zg-kafka | memory_percent | 8.440 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 61.480 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 11711858.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 73.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 11.529 |
| process:counter | cpu_seconds_total | 65.312 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40882176.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.312 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 38293504.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 295.322 |
| process:knowpost | cpu_seconds_total | 2283.250 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 109244416.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 286.204 |
| process:relation | cpu_seconds_total | 2168.203 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 100896768.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.484 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37146624.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.323 |
| process:user-storage | cpu_seconds_total | 0.453 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35516416.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 94453916.000 |
| redis | connected_clients | 233.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 403.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698800.000 |
| redis | keyspace_hits | 541593758.000 |
| redis | keyspace_misses | 95621.000 |
| redis | net_input_bytes | 22030315059.000 |
| redis | net_output_bytes | 98874239693.000 |
| redis | ops_per_sec | 27160.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21550.000 |
| redis | used_memory_bytes | 106598536.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
