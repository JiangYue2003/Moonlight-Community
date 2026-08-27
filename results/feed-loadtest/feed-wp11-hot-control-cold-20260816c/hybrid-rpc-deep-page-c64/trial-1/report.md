# Feed 压测报告：hybrid / rpc / deep-page-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:38:40+08:00
- 采样时长：1m0.0269768s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 173418 | 173418 | 0 | 0 | 2889.63 | 21.488 | 25.291 | 28.388 | 37.624 | 103.865 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 169 | 0.001 |
| redis | 520423 | 3.001 |
| relation | 173418 | 1.000 |

- Cold compute：173418（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 173418 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 173418 | 0.197 |
| counter | 240 | 1.026 |
| hydrate | 173418 | 0.350 |
| inbox | 173418 | 0.216 |
| merge_dedup | 173418 | 0.015 |
| relation | 173418 | 20.868 |
| route | 173418 | 0.004 |
| total | 173418 | 21.679 |

## Redis 本轮边界增量

- Commands：1452779；input：440181724 bytes；output：2044695378 bytes
- Hits/Misses：11188711/89549；run hit rate：99.21%
- Evicted/Rejected：0/0；ops/s max：25597；safety epoch：419 -> 419

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.223 |
| client:loadtest | cpu_percent_total | 67.574 |
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
| docker:zg-canal | cpu_percent | 2.520 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.420 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.440 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 169.840 |
| docker:zg-kafka | memory_percent | 8.440 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 54.150 |
| docker:zg-zk | memory_percent | 1.040 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 14999197.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 68.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.036 |
| process:counter | cpu_seconds_total | 141.219 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43757568.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.562 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 36986880.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 316.944 |
| process:knowpost | cpu_seconds_total | 5338.719 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 95899648.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 273.923 |
| process:relation | cpu_seconds_total | 4985.516 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 90361856.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.544 |
| process:search | cpu_seconds_total | 0.922 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37007360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 1.016 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35188736.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 122016080.000 |
| redis | connected_clients | 241.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 419.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698725.000 |
| redis | keyspace_hits | 710387951.000 |
| redis | keyspace_misses | 187480.000 |
| redis | net_input_bytes | 28803544455.000 |
| redis | net_output_bytes | 129946515194.000 |
| redis | ops_per_sec | 25597.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22759.000 |
| redis | used_memory_bytes | 106920096.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
