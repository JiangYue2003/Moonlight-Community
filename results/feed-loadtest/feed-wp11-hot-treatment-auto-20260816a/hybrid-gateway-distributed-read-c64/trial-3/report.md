# Feed 压测报告：hybrid / gateway / distributed-read-c64

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T22:10:19+08:00
- 采样时长：1m0.0311151s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 299555 | 299555 | 0 | 0 | 4991.78 | 12.460 | 16.914 | 20.933 | 28.468 | 101.456 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2439 | 0.008 |
| counter | 180 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 2404 | 0.008 |
| relation | 180 | 0.001 |

- Cold compute：362（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 298202 |
| l1_stale | 0 |
| l2_fresh | 1353 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 362 | 0.250 |
| counter | 180 | 0.892 |
| hydrate | 362 | 0.350 |
| inbox | 362 | 0.246 |
| merge_dedup | 362 | 0.018 |
| relation | 180 | 2.017 |
| route | 362 | 1.458 |
| total | 299555 | 0.009 |

## Redis 本轮边界增量

- Commands：62910；input：5577973 bytes；output：6966914 bytes
- Hits/Misses：23799/180；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1363；safety epoch：515 -> 515

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.371 |
| client:loadtest | cpu_percent_total | 117.934 |
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
| docker:zg-canal | cpu_percent | 0.180 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.340 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.230 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 281.490 |
| docker:zg-kafka | memory_percent | 8.650 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 72.280 |
| docker:zg-zk | memory_percent | 1.560 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22733242.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 12.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 12.393 |
| process:counter | cpu_seconds_total | 233.781 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 45592576.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 515.034 |
| process:gateway | cpu_seconds_total | 4198.719 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 56180736.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 210.895 |
| process:knowpost | cpu_seconds_total | 9209.734 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 111988736.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.090 |
| process:relation | cpu_seconds_total | 11.312 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 48386048.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.546 |
| process:search | cpu_seconds_total | 2.141 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36990976.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 212.432 |
| process:user-storage | cpu_seconds_total | 2104.906 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 73588736.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 196635304.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 515.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677411.000 |
| redis | keyspace_hits | 993346846.000 |
| redis | keyspace_misses | 5382817.000 |
| redis | net_input_bytes | 41946931865.000 |
| redis | net_output_bytes | 187078302680.000 |
| redis | ops_per_sec | 1363.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 42658.000 |
| redis | used_memory_bytes | 100502592.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
