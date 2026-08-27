# Feed 压测报告：hybrid / rpc / deep-page-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:44:57+08:00
- 采样时长：1m0.0373704s
- 并发：128
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 167606 | 167606 | 0 | 0 | 2792.23 | 43.463 | 51.215 | 57.650 | 119.365 | 162.990 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 163 | 0.001 |
| redis | 502981 | 3.001 |
| relation | 167606 | 1.000 |

- Cold compute：167606（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 167606 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 167606 | 0.221 |
| counter | 240 | 0.802 |
| hydrate | 167606 | 0.378 |
| inbox | 167606 | 0.271 |
| merge_dedup | 167606 | 0.014 |
| relation | 167606 | 44.406 |
| route | 167606 | 0.004 |
| total | 167606 | 45.322 |

## Redis 本轮边界增量

- Commands：1405975；input：425554595 bytes；output：1976144387 bytes
- Hits/Misses：10732711/167769；run hit rate：98.46%
- Evicted/Rejected：0/0；ops/s max：25248；safety epoch：424 -> 424

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.218 |
| client:loadtest | cpu_percent_total | 67.484 |
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
| docker:zg-canal | cpu_percent | 2.410 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.950 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.340 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 179.270 |
| docker:zg-kafka | memory_percent | 8.460 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 52.980 |
| docker:zg-zk | memory_percent | 1.050 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 16009466.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 74.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 10.763 |
| process:counter | cpu_seconds_total | 164.594 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43778048.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.555 |
| process:gateway | cpu_seconds_total | 0.672 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37040128.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 296.682 |
| process:knowpost | cpu_seconds_total | 6307.609 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 108986368.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 280.570 |
| process:relation | cpu_seconds_total | 5849.031 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 93700096.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.062 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37052416.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 1.203 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35217408.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 130489251.000 |
| redis | connected_clients | 304.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 424.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 697802.000 |
| redis | keyspace_hits | 774947430.000 |
| redis | keyspace_misses | 1196734.000 |
| redis | net_input_bytes | 31364524720.000 |
| redis | net_output_bytes | 141833735148.000 |
| redis | ops_per_sec | 25248.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23136.000 |
| redis | used_memory_bytes | 107771488.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
