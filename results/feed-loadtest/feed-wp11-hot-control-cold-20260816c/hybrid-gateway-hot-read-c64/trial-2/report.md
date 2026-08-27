# Feed 压测报告：hybrid / gateway / hot-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:55:01+08:00
- 采样时长：1m0.0294389s
- 并发：64
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 146111 | 146111 | 0 | 0 | 2434.44 | 25.243 | 31.103 | 34.719 | 41.207 | 132.962 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 123 | 0.001 |
| redis | 438456 | 3.001 |
| relation | 146111 | 1.000 |

- Cold compute：146111（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 146111 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 146111 | 0.260 |
| counter | 12 | 1.865 |
| hydrate | 146111 | 0.377 |
| inbox | 146111 | 0.290 |
| merge_dedup | 146111 | 0.013 |
| relation | 146111 | 23.531 |
| route | 146111 | 0.003 |
| total | 146111 | 24.499 |

## Redis 本轮边界增量

- Commands：1225626；input：278985252 bytes；output：1272509699 bytes
- Hits/Misses：6721372/146234；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22233；safety epoch：432 -> 432

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.698 |
| client:loadtest | cpu_percent_total | 59.164 |
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
| docker:zg-canal | cpu_percent | 2.940 |
| docker:zg-canal | memory_percent | 7.860 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 0.870 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 7.890 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 25.000 |
| docker:zg-kafka | cpu_percent | 222.130 |
| docker:zg-kafka | memory_percent | 8.460 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 63.420 |
| docker:zg-zk | memory_percent | 1.270 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 17472556.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 48.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 9.994 |
| process:counter | cpu_seconds_total | 202.906 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 42770432.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 242.006 |
| process:gateway | cpu_seconds_total | 715.297 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 54226944.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 269.083 |
| process:knowpost | cpu_seconds_total | 7697.328 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 106311680.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 242.948 |
| process:relation | cpu_seconds_total | 7131.703 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 87351296.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.556 |
| process:search | cpu_seconds_total | 1.422 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37068800.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 126.809 |
| process:user-storage | cpu_seconds_total | 379.922 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 56418304.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 142779085.000 |
| redis | connected_clients | 314.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 432.000 |
| redis | hit_rate | 0.997 |
| redis | keys | 690287.000 |
| redis | keyspace_hits | 853093631.000 |
| redis | keyspace_misses | 2658342.000 |
| redis | net_input_bytes | 34538601103.000 |
| redis | net_output_bytes | 156426618918.000 |
| redis | ops_per_sec | 22233.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23740.000 |
| redis | used_memory_bytes | 106564792.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
