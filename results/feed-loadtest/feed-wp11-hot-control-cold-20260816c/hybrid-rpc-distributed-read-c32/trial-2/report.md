# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:21:04+08:00
- 采样时长：1m0.019824s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 177843 | 177843 | 0 | 0 | 2963.71 | 10.627 | 12.443 | 13.240 | 15.375 | 24.585 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 103 | 0.001 |
| redis | 533632 | 3.001 |
| relation | 177843 | 1.000 |

- Cold compute：177843（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 177843 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 177843 | 0.172 |
| counter | 240 | 0.689 |
| hydrate | 177843 | 0.273 |
| inbox | 177843 | 0.187 |
| merge_dedup | 177843 | 0.012 |
| relation | 177843 | 9.664 |
| route | 177843 | 0.003 |
| total | 177843 | 10.335 |

## Redis 本轮边界增量

- Commands：1488013；input：339060642 bytes；output：1548997969 bytes
- Hits/Misses：8364602/109；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26178；safety epoch：405 -> 405

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.611 |
| client:loadtest | cpu_percent_total | 73.778 |
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
| docker:zg-canal | cpu_percent | 0.180 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.990 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.960 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 149.670 |
| docker:zg-kafka | memory_percent | 8.390 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 52.250 |
| docker:zg-zk | memory_percent | 1.030 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 12129602.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 39.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.822 |
| process:counter | cpu_seconds_total | 74.828 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43630592.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 0.328 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 38256640.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 308.140 |
| process:knowpost | cpu_seconds_total | 2666.109 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 86523904.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 269.482 |
| process:relation | cpu_seconds_total | 2524.641 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 88297472.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.547 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37101568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 0.516 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35266560.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 97956305.000 |
| redis | connected_clients | 237.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 405.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698781.000 |
| redis | keyspace_hits | 561214273.000 |
| redis | keyspace_misses | 95892.000 |
| redis | net_input_bytes | 22826492011.000 |
| redis | net_output_bytes | 102507894292.000 |
| redis | ops_per_sec | 26178.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21703.000 |
| redis | used_memory_bytes | 106597888.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
