# Feed 压测报告：hybrid / gateway / hot-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:51:14+08:00
- 采样时长：1m0.0208943s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 140963 | 140963 | 0 | 0 | 2348.97 | 12.411 | 15.368 | 17.029 | 76.127 | 120.951 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 106 | 0.001 |
| redis | 422995 | 3.001 |
| relation | 140963 | 1.000 |

- Cold compute：140963（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 140963 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 140963 | 0.216 |
| counter | 12 | 1.665 |
| hydrate | 140963 | 0.331 |
| inbox | 140963 | 0.237 |
| merge_dedup | 140963 | 0.013 |
| relation | 140963 | 10.964 |
| route | 140963 | 0.003 |
| total | 140963 | 11.789 |

## Redis 本轮边界增量

- Commands：1185860；input：269396588 bytes；output：1227744984 bytes
- Hits/Misses：6484566/141085；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：21290；safety epoch：429 -> 429

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.485 |
| client:loadtest | cpu_percent_total | 55.762 |
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
| docker:zg-canal | cpu_percent | 2.810 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.480 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 7.500 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 199.320 |
| docker:zg-kafka | memory_percent | 8.390 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 75.290 |
| docker:zg-zk | memory_percent | 1.200 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 16953863.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 31.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 11.603 |
| process:counter | cpu_seconds_total | 189.438 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43970560.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 225.868 |
| process:gateway | cpu_seconds_total | 280.438 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 51777536.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 244.782 |
| process:knowpost | cpu_seconds_total | 7207.234 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 93679616.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 241.255 |
| process:relation | cpu_seconds_total | 6672.891 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 85983232.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.322 |
| process:search | cpu_seconds_total | 1.156 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37068800.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 123.505 |
| process:user-storage | cpu_seconds_total | 151.422 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 55320576.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 138420075.000 |
| redis | connected_clients | 314.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 429.000 |
| redis | hit_rate | 0.998 |
| redis | keys | 694520.000 |
| redis | keyspace_hits | 829275333.000 |
| redis | keyspace_misses | 2140148.000 |
| redis | net_input_bytes | 33548861530.000 |
| redis | net_output_bytes | 151916981256.000 |
| redis | ops_per_sec | 21290.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23513.000 |
| redis | used_memory_bytes | 107213848.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
