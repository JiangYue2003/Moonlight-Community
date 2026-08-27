# Feed 压测报告：hybrid / rpc / hot-read-c256

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:10:06+08:00
- 采样时长：1m0.0499849s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 481796 | 481796 | 0 | 0 | 8027.02 | 31.472 | 39.660 | 41.566 | 46.702 | 98.594 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 165 | 0.000 |
| counter | 10 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 121 | 0.000 |
| relation | 10 | 0.000 |

- Cold compute：19（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 481688 |
| l1_stale | 0 |
| l2_fresh | 108 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 19 | 0.055 |
| counter | 10 | 1.892 |
| hydrate | 19 | 0.218 |
| inbox | 19 | 0.055 |
| merge_dedup | 19 | 0.028 |
| relation | 10 | 1.883 |
| route | 19 | 1.987 |
| total | 481796 | 0.005 |

## Redis 本轮边界增量

- Commands：53214；input：3807415 bytes；output：1432860 bytes
- Hits/Misses：1396/10；run hit rate：99.29%
- Evicted/Rejected：0/0；ops/s max：1099；safety epoch：473 -> 473

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.996 |
| client:loadtest | cpu_percent_total | 127.940 |
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
| docker:zg-canal | cpu_percent | 1.720 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.260 |
| docker:zg-es | memory_percent | 11.940 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.600 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 139.630 |
| docker:zg-kafka | memory_percent | 8.280 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 49.050 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20180550.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 8.528 |
| process:counter | cpu_seconds_total | 40.984 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 40316928.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.866 |
| process:gateway | cpu_seconds_total | 0.406 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37666816.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 236.246 |
| process:knowpost | cpu_seconds_total | 1549.547 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 119783424.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.775 |
| process:relation | cpu_seconds_total | 0.656 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 39219200.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 0.453 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 37097472.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.549 |
| process:user-storage | cpu_seconds_total | 0.375 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34557952.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 177871718.000 |
| redis | connected_clients | 16.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 473.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676931.000 |
| redis | keyspace_hits | 977437924.000 |
| redis | keyspace_misses | 5364972.000 |
| redis | net_input_bytes | 40570403674.000 |
| redis | net_output_bytes | 180195736398.000 |
| redis | ops_per_sec | 1099.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 39045.000 |
| redis | used_memory_bytes | 95632432.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
