# Feed 压测报告：hybrid / gateway / hot-read-c256

- Run ID：`feed-wp11-hot-control-diagnostic-c256-20260816d`
- 开始时间：2026-08-16T20:00:57+08:00
- 采样时长：10.0768129s
- 并发：256
- 重复轮次：1 / 1
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 24711 | 24660 | 51 | 0 | 2447.73 | 89.061 | 168.709 | 203.389 | 287.904 | 662.926 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 4 | 0.000 |
| mysql | 8 | 0.000 |
| redis | 73988 | 3.000 |
| relation | 24660 | 1.000 |

- Cold compute：24660（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 24660 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 24660 | 0.251 |
| counter | 4 | 7.904 |
| hydrate | 24660 | 0.376 |
| inbox | 24660 | 0.292 |
| merge_dedup | 24660 | 0.014 |
| relation | 24660 | 100.990 |
| route | 24660 | 0.010 |
| total | 24660 | 101.958 |

## Redis 本轮边界增量

- Commands：206622；input：47114449 bytes；output：214733996 bytes
- Hits/Misses：1134156/24980；run hit rate：97.84%
- Evicted/Rejected：0/0；ops/s max：21584；safety epoch：438 -> 438

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.303 |
| client:loadtest | cpu_percent_total | 68.846 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.060 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 6.430 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 195.550 |
| docker:zg-kafka | memory_percent | 8.350 |
| docker:zg-kafka | pids | 119.000 |
| docker:zg-zk | cpu_percent | 80.900 |
| docker:zg-zk | memory_percent | 1.210 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 18367140.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 72.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 8.516 |
| process:counter | cpu_seconds_total | 656.719 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 42164224.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 224.339 |
| process:gateway | cpu_seconds_total | 1497.312 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 70029312.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 263.943 |
| process:knowpost | cpu_seconds_total | 8776.250 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 113557504.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 227.674 |
| process:relation | cpu_seconds_total | 7952.984 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 92872704.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 5.547 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37273600.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 127.638 |
| process:user-storage | cpu_seconds_total | 783.484 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 49451008.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 159659910.000 |
| redis | connected_clients | 347.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 438.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 676397.000 |
| redis | keyspace_hits | 894138552.000 |
| redis | keyspace_misses | 3551771.000 |
| redis | net_input_bytes | 36904345180.000 |
| redis | net_output_bytes | 164378432824.000 |
| redis | ops_per_sec | 21584.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34846.000 |
| redis | used_memory_bytes | 102347272.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
