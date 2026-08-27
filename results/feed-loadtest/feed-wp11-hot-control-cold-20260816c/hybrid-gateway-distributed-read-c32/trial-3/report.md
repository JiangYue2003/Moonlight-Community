# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T20:07:12+08:00
- 采样时长：1m0.0184711s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 152548 | 152548 | 0 | 0 | 2542.15 | 12.353 | 14.771 | 15.794 | 18.454 | 30.940 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 112 | 0.001 |
| redis | 457756 | 3.001 |
| relation | 152548 | 1.000 |

- Cold compute：152548（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 152548 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 152548 | 0.218 |
| counter | 240 | 0.923 |
| hydrate | 152548 | 0.327 |
| inbox | 152548 | 0.238 |
| merge_dedup | 152548 | 0.013 |
| relation | 152548 | 10.136 |
| route | 152548 | 0.005 |
| total | 152548 | 10.960 |

## Redis 本轮边界增量

- Commands：1278964；input：290985938 bytes；output：1328598902 bytes
- Hits/Misses：7023185/152660；run hit rate：97.87%
- Evicted/Rejected：0/0；ops/s max：23045；safety epoch：441 -> 441

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.720 |
| client:loadtest | cpu_percent_total | 59.513 |
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
| docker:zg-canal | cpu_percent | 3.140 |
| docker:zg-canal | memory_percent | 7.850 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 3.670 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 7.500 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 314.320 |
| docker:zg-kafka | memory_percent | 8.520 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 76.830 |
| docker:zg-zk | memory_percent | 1.340 |
| docker:zg-zk | pids | 103.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 18932705.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 29.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 11.590 |
| process:counter | cpu_seconds_total | 679.438 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 44838912.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 241.399 |
| process:gateway | cpu_seconds_total | 1961.969 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 53981184.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 274.740 |
| process:knowpost | cpu_seconds_total | 9308.562 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 99205120.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 248.155 |
| process:relation | cpu_seconds_total | 8445.531 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 86781952.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 5.859 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37277696.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 144.515 |
| process:user-storage | cpu_seconds_total | 1035.031 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 57815040.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 164569966.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 441.000 |
| redis | hit_rate | 0.996 |
| redis | keys | 676488.000 |
| redis | keyspace_hits | 920127817.000 |
| redis | keyspace_misses | 4117407.000 |
| redis | net_input_bytes | 37993824578.000 |
| redis | net_output_bytes | 169298507451.000 |
| redis | ops_per_sec | 23045.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35271.000 |
| redis | used_memory_bytes | 99860984.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
