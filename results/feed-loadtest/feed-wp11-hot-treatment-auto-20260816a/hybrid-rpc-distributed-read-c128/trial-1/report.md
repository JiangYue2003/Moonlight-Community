# Feed 压测报告：hybrid / rpc / distributed-read-c128

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:21:23+08:00
- 采样时长：1m0.0419571s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 484272 | 484272 | 0 | 0 | 8069.51 | 15.732 | 19.602 | 20.722 | 22.950 | 49.324 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2456 | 0.005 |
| counter | 183 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 2415 | 0.005 |
| relation | 183 | 0.000 |

- Cold compute：366（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 482914 |
| l1_stale | 0 |
| l2_fresh | 1358 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 366 | 0.159 |
| counter | 183 | 0.653 |
| hydrate | 366 | 0.200 |
| inbox | 366 | 0.157 |
| merge_dedup | 366 | 0.014 |
| relation | 183 | 1.423 |
| route | 366 | 1.060 |
| total | 484272 | 0.005 |

## Redis 本轮边界增量

- Commands：62922；input：5589301 bytes；output：6998733 bytes
- Hits/Misses：24057/183；run hit rate：99.25%
- Evicted/Rejected：0/0；ops/s max：1297；safety epoch：482 -> 482

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.903 |
| client:loadtest | cpu_percent_total | 126.448 |
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
| docker:zg-canal | cpu_percent | 1.790 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.650 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.620 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 167.670 |
| docker:zg-kafka | memory_percent | 8.590 |
| docker:zg-kafka | pids | 122.000 |
| docker:zg-zk | cpu_percent | 44.950 |
| docker:zg-zk | memory_percent | 1.250 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20183239.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 25.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.744 |
| process:counter | cpu_seconds_total | 72.062 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43769856.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 0.516 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37814272.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 230.850 |
| process:knowpost | cpu_seconds_total | 2905.453 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 104644608.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.325 |
| process:relation | cpu_seconds_total | 2.766 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 47468544.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.781 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37015552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.802 |
| process:user-storage | cpu_seconds_total | 0.656 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34889728.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 178557432.000 |
| redis | connected_clients | 55.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 482.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677120.000 |
| redis | keyspace_hits | 977648520.000 |
| redis | keyspace_misses | 5368971.000 |
| redis | net_input_bytes | 40629354612.000 |
| redis | net_output_bytes | 180258680836.000 |
| redis | ops_per_sec | 1297.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 39722.000 |
| redis | used_memory_bytes | 97949200.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
