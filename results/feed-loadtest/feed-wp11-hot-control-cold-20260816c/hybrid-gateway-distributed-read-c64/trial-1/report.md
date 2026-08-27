# Feed 压测报告：hybrid / gateway / distributed-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T20:08:27+08:00
- 采样时长：1m0.0277848s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 151240 | 151240 | 0 | 0 | 2519.94 | 24.965 | 29.418 | 31.328 | 35.994 | 55.526 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 127 | 0.001 |
| redis | 453847 | 3.001 |
| relation | 151240 | 1.000 |

- Cold compute：151240（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 151240 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 151240 | 0.226 |
| counter | 240 | 0.897 |
| hydrate | 151240 | 0.337 |
| inbox | 151240 | 0.248 |
| merge_dedup | 151240 | 0.013 |
| relation | 151240 | 22.890 |
| route | 151240 | 0.005 |
| total | 151240 | 23.743 |

## Redis 本轮边界增量

- Commands：1268547；input：288529684 bytes；output：1317216127 bytes
- Hits/Misses：6963003/151367；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22730；safety epoch：442 -> 442

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.639 |
| client:loadtest | cpu_percent_total | 58.228 |
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
| docker:zg-canal | cpu_percent | 0.190 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 6.090 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 8.830 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 319.760 |
| docker:zg-kafka | memory_percent | 9.000 |
| docker:zg-kafka | pids | 151.000 |
| docker:zg-zk | cpu_percent | 81.890 |
| docker:zg-zk | memory_percent | 1.360 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 19112073.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 48.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 11.478 |
| process:counter | cpu_seconds_total | 684.109 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44322816.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 255.944 |
| process:gateway | cpu_seconds_total | 2110.281 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 56352768.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 267.430 |
| process:knowpost | cpu_seconds_total | 9479.281 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 104837120.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 248.960 |
| process:relation | cpu_seconds_total | 8603.359 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 87429120.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 5.875 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37277696.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 132.241 |
| process:user-storage | cpu_seconds_total | 1113.906 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 57012224.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 166074943.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 442.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 676533.000 |
| redis | keyspace_hits | 928368019.000 |
| redis | keyspace_misses | 4296523.000 |
| redis | net_input_bytes | 38335530576.000 |
| redis | net_output_bytes | 170857391937.000 |
| redis | ops_per_sec | 22730.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35347.000 |
| redis | used_memory_bytes | 99865816.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
