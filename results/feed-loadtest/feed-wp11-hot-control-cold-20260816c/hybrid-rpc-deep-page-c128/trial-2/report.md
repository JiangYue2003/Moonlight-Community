# Feed 压测报告：hybrid / rpc / deep-page-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:43:42+08:00
- 采样时长：1m0.0388825s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 174147 | 174147 | 0 | 0 | 2901.13 | 43.374 | 50.614 | 54.026 | 64.796 | 106.980 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 169 | 0.001 |
| redis | 522610 | 3.001 |
| relation | 174147 | 1.000 |

- Cold compute：174147（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 174147 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 174147 | 0.205 |
| counter | 240 | 0.656 |
| hydrate | 174147 | 0.357 |
| inbox | 174147 | 0.223 |
| merge_dedup | 174147 | 0.015 |
| relation | 174147 | 42.823 |
| route | 174147 | 0.004 |
| total | 174147 | 43.655 |

## Redis 本轮边界增量

- Commands：1458286；input：441997450 bytes；output：2053188920 bytes
- Hits/Misses：11151287/174358；run hit rate：98.46%
- Evicted/Rejected：0/0；ops/s max：25700；safety epoch：423 -> 423

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.153 |
| client:loadtest | cpu_percent_total | 66.441 |
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
| docker:zg-canal | cpu_percent | 1.050 |
| docker:zg-canal | memory_percent | 7.860 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 0.710 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 5.620 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 187.470 |
| docker:zg-kafka | memory_percent | 8.450 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 58.400 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 15811871.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 74.000 |
| mysql | threads_running | 7.000 |
| process:counter | cpu_percent | 19.337 |
| process:counter | cpu_seconds_total | 160.047 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43651072.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 2.319 |
| process:gateway | cpu_seconds_total | 0.641 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37015552.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 312.111 |
| process:knowpost | cpu_seconds_total | 6120.031 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 101371904.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 284.520 |
| process:relation | cpu_seconds_total | 5679.844 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 94736384.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.062 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37056512.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.188 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35217408.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 128830998.000 |
| redis | connected_clients | 241.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 423.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 698020.000 |
| redis | keyspace_hits | 762322213.000 |
| redis | keyspace_misses | 999383.000 |
| redis | net_input_bytes | 30863629254.000 |
| redis | net_output_bytes | 139509073470.000 |
| redis | ops_per_sec | 25700.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23061.000 |
| redis | used_memory_bytes | 106564008.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
