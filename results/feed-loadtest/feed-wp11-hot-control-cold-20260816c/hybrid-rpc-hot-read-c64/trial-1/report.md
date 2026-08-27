# Feed 压测报告：hybrid / rpc / hot-read-c64

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:08:28+08:00
- 采样时长：1m0.0299977s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 182856 | 182856 | 0 | 0 | 3046.77 | 20.878 | 24.440 | 25.814 | 29.418 | 50.538 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 12 | 0.000 |
| mysql | 89 | 0.000 |
| redis | 548657 | 3.000 |
| relation | 182856 | 1.000 |

- Cold compute：182856（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 182856 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 182856 | 0.204 |
| counter | 12 | 1.860 |
| hydrate | 182856 | 0.301 |
| inbox | 182856 | 0.226 |
| merge_dedup | 182856 | 0.011 |
| relation | 182856 | 19.786 |
| route | 182856 | 0.003 |
| total | 182856 | 20.553 |

## Redis 本轮边界增量

- Commands：1522398；input：348332509 bytes；output：1592471249 bytes
- Hits/Misses：8594532/90；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26898；safety epoch：395 -> 395

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.482 |
| client:loadtest | cpu_percent_total | 71.709 |
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
| docker:zg-canal | cpu_percent | 3.180 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.250 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.630 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 181.940 |
| docker:zg-kafka | memory_percent | 8.420 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 61.300 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 9982597.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 70.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 8.505 |
| process:counter | cpu_seconds_total | 26.703 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 40898560.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.188 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37916672.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 306.825 |
| process:knowpost | cpu_seconds_total | 781.016 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 87388160.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 298.344 |
| process:relation | cpu_seconds_total | 733.578 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 74448896.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.250 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37068800.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 0.188 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35192832.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 80043663.000 |
| redis | connected_clients | 90.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 395.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698881.000 |
| redis | keyspace_hits | 460448666.000 |
| redis | keyspace_misses | 94542.000 |
| redis | net_input_bytes | 18738934512.000 |
| redis | net_output_bytes | 83838317245.000 |
| redis | ops_per_sec | 26898.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 20947.000 |
| redis | used_memory_bytes | 103975872.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
