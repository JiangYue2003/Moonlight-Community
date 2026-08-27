# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:31:07+08:00
- 采样时长：1m0.0773567s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 160722 | 160722 | 0 | 0 | 2675.85 | 80.795 | 158.992 | 194.721 | 284.030 | 804.310 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 110 | 0.001 |
| redis | 482276 | 3.001 |
| relation | 160722 | 1.000 |

- Cold compute：160722（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 160722 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 160722 | 0.168 |
| counter | 240 | 0.790 |
| hydrate | 160722 | 0.266 |
| inbox | 160722 | 0.180 |
| merge_dedup | 160722 | 0.012 |
| relation | 160722 | 94.408 |
| route | 160722 | 0.004 |
| total | 160722 | 95.060 |

## Redis 本轮边界增量

- Commands：1351118；input：306846850 bytes；output：1400013484 bytes
- Hits/Misses：7559910/114；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：25623；safety epoch：413 -> 413

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.098 |
| client:loadtest | cpu_percent_total | 65.567 |
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
| docker:zg-es | cpu_percent | 0.560 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.140 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 196.160 |
| docker:zg-kafka | memory_percent | 8.440 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 91.480 |
| docker:zg-zk | memory_percent | 1.340 |
| docker:zg-zk | pids | 102.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 13785750.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 74.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 11.657 |
| process:counter | cpu_seconds_total | 112.938 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44146688.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.778 |
| process:gateway | cpu_seconds_total | 0.500 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37957632.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 296.864 |
| process:knowpost | cpu_seconds_total | 4180.391 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 113131520.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 282.457 |
| process:relation | cpu_seconds_total | 3947.219 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 105013248.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.750 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37195776.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.545 |
| process:user-storage | cpu_seconds_total | 0.812 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35569664.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 111836000.000 |
| redis | connected_clients | 237.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 413.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698734.000 |
| redis | keyspace_hits | 638978745.000 |
| redis | keyspace_misses | 97012.000 |
| redis | net_input_bytes | 25981849812.000 |
| redis | net_output_bytes | 116909289264.000 |
| redis | ops_per_sec | 25623.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22307.000 |
| redis | used_memory_bytes | 106588560.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
