# Feed 压测报告：hybrid / gateway / hot-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T17:01:18+08:00
- 采样时长：1m0.0913651s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 140981 | 140981 | 0 | 0 | 2346.52 | 94.997 | 175.218 | 210.460 | 294.660 | 696.864 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 164 | 0.001 |
| redis | 423107 | 3.001 |
| relation | 140981 | 1.000 |

- Cold compute：140981（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 140981 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 140981 | 1.398 |
| counter | 12 | 4.318 |
| hydrate | 140981 | 1.564 |
| inbox | 140981 | 1.552 |
| merge_dedup | 140981 | 0.013 |
| relation | 140981 | 102.139 |
| route | 140981 | 0.012 |
| total | 140981 | 106.702 |

## Redis 本轮边界增量

- Commands：1183240；input：269242414 bytes；output：1227846644 bytes
- Hits/Misses：6485352/141145；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：21914；safety epoch：437 -> 437

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.673 |
| client:loadtest | cpu_percent_total | 58.765 |
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
| docker:zg-canal | cpu_percent | 2.960 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.440 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 8.090 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 240.000 |
| docker:zg-kafka | memory_percent | 8.450 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 76.110 |
| docker:zg-zk | memory_percent | 1.300 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 18317917.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 90.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 13.140 |
| process:counter | cpu_seconds_total | 225.922 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 42594304.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 258.917 |
| process:gateway | cpu_seconds_total | 1454.141 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 70426624.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 264.328 |
| process:knowpost | cpu_seconds_total | 8527.688 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 122101760.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 259.887 |
| process:relation | cpu_seconds_total | 7902.484 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 100179968.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.625 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37076992.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 135.909 |
| process:user-storage | cpu_seconds_total | 758.469 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 57663488.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 149876775.000 |
| redis | connected_clients | 347.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 437.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 684766.000 |
| redis | keyspace_hits | 891885153.000 |
| redis | keyspace_misses | 3502519.000 |
| redis | net_input_bytes | 36150454471.000 |
| redis | net_output_bytes | 163771187579.000 |
| redis | ops_per_sec | 21914.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 24117.000 |
| redis | used_memory_bytes | 114215064.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
