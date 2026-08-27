# Feed 压测报告：hybrid / rpc / deep-page-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:34:53+08:00
- 采样时长：1m0.0220653s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 169782 | 169782 | 0 | 0 | 2829.27 | 10.787 | 12.774 | 13.689 | 23.246 | 100.472 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 133 | 0.001 |
| redis | 509479 | 3.001 |
| relation | 169782 | 1.000 |

- Cold compute：169782（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 169782 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 169782 | 0.207 |
| counter | 240 | 0.751 |
| hydrate | 169782 | 0.362 |
| inbox | 169782 | 0.230 |
| merge_dedup | 169782 | 0.015 |
| relation | 169782 | 9.974 |
| route | 169782 | 0.004 |
| total | 169782 | 10.820 |

## Redis 本轮边界增量

- Commands：1423535；input：431026854 bytes；output：2001945723 bytes
- Hits/Misses：11041787/133；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：25668；safety epoch：416 -> 416

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.079 |
| client:loadtest | cpu_percent_total | 65.262 |
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
| docker:zg-canal | cpu_percent | 2.740 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 81.000 |
| docker:zg-es | cpu_percent | 2.960 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 6.200 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 191.220 |
| docker:zg-kafka | memory_percent | 8.450 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 60.340 |
| docker:zg-zk | memory_percent | 1.040 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 14390419.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 52.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.839 |
| process:counter | cpu_seconds_total | 126.828 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44036096.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.018 |
| process:gateway | cpu_seconds_total | 0.547 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 36954112.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 306.901 |
| process:knowpost | cpu_seconds_total | 4746.891 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 85884928.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 269.000 |
| process:relation | cpu_seconds_total | 4469.312 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 87322624.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.875 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37007360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.938 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35069952.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 116910277.000 |
| redis | connected_clients | 241.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 416.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698731.000 |
| redis | keyspace_hits | 670959832.000 |
| redis | keyspace_misses | 97479.000 |
| redis | net_input_bytes | 27260096651.000 |
| redis | net_output_bytes | 122781566806.000 |
| redis | ops_per_sec | 25668.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22532.000 |
| redis | used_memory_bytes | 106922048.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
