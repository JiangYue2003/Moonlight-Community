# Feed 压测报告：hybrid / gateway / hot-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:53:46+08:00
- 采样时长：1m0.0271017s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 143408 | 143408 | 0 | 0 | 2389.49 | 25.250 | 30.354 | 33.012 | 96.426 | 132.672 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 118 | 0.001 |
| redis | 430342 | 3.001 |
| relation | 143408 | 1.000 |

- Cold compute：143408（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 143408 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 143408 | 0.253 |
| counter | 12 | 1.922 |
| hydrate | 143408 | 0.371 |
| inbox | 143408 | 0.280 |
| merge_dedup | 143408 | 0.013 |
| relation | 143408 | 24.038 |
| route | 143408 | 0.003 |
| total | 143408 | 24.984 |

## Redis 本轮边界增量

- Commands：1204571；input：273937332 bytes；output：1249002767 bytes
- Hits/Misses：6597040/143526；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：22040；safety epoch：431 -> 431

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.595 |
| client:loadtest | cpu_percent_total | 57.526 |
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
| docker:zg-canal | cpu_percent | 0.210 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.220 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 7.160 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 261.010 |
| docker:zg-kafka | memory_percent | 8.460 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 76.960 |
| docker:zg-zk | memory_percent | 1.050 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 17300947.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 47.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 8.544 |
| process:counter | cpu_seconds_total | 198.281 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 42770432.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 232.718 |
| process:gateway | cpu_seconds_total | 569.672 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 54677504.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 273.684 |
| process:knowpost | cpu_seconds_total | 7531.375 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 101879808.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 242.242 |
| process:relation | cpu_seconds_total | 6978.938 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 87416832.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.281 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37068800.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 124.604 |
| process:user-storage | cpu_seconds_total | 303.078 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 61812736.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 141336700.000 |
| redis | connected_clients | 314.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 431.000 |
| redis | hit_rate | 0.997 |
| redis | keys | 691803.000 |
| redis | keyspace_hits | 845213005.000 |
| redis | keyspace_misses | 2486893.000 |
| redis | net_input_bytes | 34211120602.000 |
| redis | net_output_bytes | 154934539217.000 |
| redis | ops_per_sec | 22040.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23665.000 |
| redis | used_memory_bytes | 107067456.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
