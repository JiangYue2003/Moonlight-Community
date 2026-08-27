# Feed 压测报告：hybrid / rpc / distributed-read-c256

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:25:07+08:00
- 采样时长：1m0.0623965s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 488162 | 488162 | 0 | 0 | 8131.52 | 30.705 | 39.077 | 41.463 | 47.056 | 98.814 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2454 | 0.005 |
| counter | 180 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 2414 | 0.005 |
| relation | 180 | 0.000 |

- Cold compute：364（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 486800 |
| l1_stale | 0 |
| l2_fresh | 1362 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 364 | 0.161 |
| counter | 180 | 0.651 |
| hydrate | 364 | 0.207 |
| inbox | 364 | 0.161 |
| merge_dedup | 364 | 0.014 |
| relation | 180 | 1.450 |
| route | 364 | 1.059 |
| total | 488162 | 0.005 |

## Redis 本轮边界增量

- Commands：62833；input：5581267 bytes；output：6991264 bytes
- Hits/Misses：23895/180；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1322；safety epoch：485 -> 485

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.134 |
| client:loadtest | cpu_percent_total | 130.151 |
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
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.320 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.870 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 181.880 |
| docker:zg-kafka | memory_percent | 8.330 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 47.970 |
| docker:zg-zk | memory_percent | 1.480 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20184317.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 8.516 |
| process:counter | cpu_seconds_total | 82.734 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 44507136.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.516 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37593088.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 234.554 |
| process:knowpost | cpu_seconds_total | 3351.812 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 115146752.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.890 |
| process:relation | cpu_seconds_total | 3.531 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 47087616.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.859 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37044224.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.672 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34951168.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178793942.000 |
| redis | connected_clients | 55.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 485.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677159.000 |
| redis | keyspace_hits | 977737462.000 |
| redis | keyspace_misses | 5370628.000 |
| redis | net_input_bytes | 40650519165.000 |
| redis | net_output_bytes | 180284161962.000 |
| redis | ops_per_sec | 1322.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 39946.000 |
| redis | used_memory_bytes | 98171456.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
