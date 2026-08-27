# Feed 压测报告：hybrid / rpc / hot-read-c32

- Run ID：`feed-wp11-hot-control-validation-20260816b`
- 开始时间：2026-08-16T16:02:40+08:00
- 采样时长：5.0088814s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 15540 | 15540 | 0 | 0 | 3103.14 | 10.182 | 11.978 | 12.782 | 15.043 | 18.867 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 1 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 46620 | 3.000 |
| relation | 15540 | 1.000 |

- Cold compute：15540（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 15540 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 15540 | 0.193 |
| counter | 1 | 1.553 |
| hydrate | 15540 | 0.284 |
| inbox | 15540 | 0.211 |
| merge_dedup | 15540 | 0.011 |
| relation | 15540 | 9.130 |
| route | 15540 | 0.003 |
| total | 15540 | 9.852 |

## Redis 本轮边界增量

- Commands：129282；input：29594778 bytes；output：134513378 bytes
- Hits/Misses：730413/0；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：27478；safety epoch：389 -> 389

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.933 |
| client:loadtest | cpu_percent_total | 78.922 |
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
| docker:zg-canal | cpu_percent | 0.120 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.470 |
| docker:zg-es | memory_percent | 11.830 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 3.870 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 39.560 |
| docker:zg-kafka | memory_percent | 8.140 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 0.120 |
| docker:zg-zk | memory_percent | 1.010 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3100.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3100.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 9122923.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 36.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.806 |
| process:counter | cpu_seconds_total | 4.062 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 38559744.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.156 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 36364288.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 279.406 |
| process:knowpost | cpu_seconds_total | 21.094 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 66797568.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 264.508 |
| process:relation | cpu_seconds_total | 15.859 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 53985280.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.094 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 36577280.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.062 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 34324480.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 72761597.000 |
| redis | connected_clients | 76.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 389.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 698928.000 |
| redis | keyspace_hits | 420095592.000 |
| redis | keyspace_misses | 93212.000 |
| redis | net_input_bytes | 17093924049.000 |
| redis | net_output_bytes | 76359136120.000 |
| redis | ops_per_sec | 27478.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 20544.000 |
| redis | used_memory_bytes | 103238232.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
