# Feed 压测报告：hybrid / rpc / hot-read-c256

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:11:22+08:00
- 采样时长：1m0.0571746s
- 并发：256
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 480779 | 480779 | 0 | 0 | 8009.18 | 32.037 | 39.088 | 41.203 | 46.754 | 118.546 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 172 | 0.000 |
| counter | 9 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 120 | 0.000 |
| relation | 9 | 0.000 |

- Cold compute：18（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 480661 |
| l1_stale | 0 |
| l2_fresh | 118 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 18 | 0.177 |
| counter | 9 | 1.716 |
| hydrate | 18 | 0.207 |
| inbox | 18 | 0.177 |
| merge_dedup | 18 | 0.000 |
| relation | 9 | 1.991 |
| route | 18 | 1.854 |
| total | 480779 | 0.005 |

## Redis 本轮边界增量

- Commands：53545；input：3828845 bytes；output：1435061 bytes
- Hits/Misses：1327/9；run hit rate：99.33%
- Evicted/Rejected：0/0；ops/s max：1108；safety epoch：474 -> 474

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.168 |
| client:loadtest | cpu_percent_total | 130.683 |
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
| docker:zg-canal | cpu_percent | 2.070 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.580 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.490 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 134.270 |
| docker:zg-kafka | memory_percent | 8.220 |
| docker:zg-kafka | pids | 112.000 |
| docker:zg-zk | cpu_percent | 46.710 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 20180634.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.731 |
| process:counter | cpu_seconds_total | 44.531 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 40321024.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.406 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 37933056.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 235.577 |
| process:knowpost | cpu_seconds_total | 1698.953 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 118104064.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.773 |
| process:relation | cpu_seconds_total | 0.672 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 39231488.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.975 |
| process:search | cpu_seconds_total | 0.469 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36966400.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.320 |
| process:user-storage | cpu_seconds_total | 0.453 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34611200.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 177938573.000 |
| redis | connected_clients | 16.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 474.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 676943.000 |
| redis | keyspace_hits | 977439486.000 |
| redis | keyspace_misses | 5365035.000 |
| redis | net_input_bytes | 40575191903.000 |
| redis | net_output_bytes | 180197490805.000 |
| redis | ops_per_sec | 1108.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 39121.000 |
| redis | used_memory_bytes | 95642160.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
