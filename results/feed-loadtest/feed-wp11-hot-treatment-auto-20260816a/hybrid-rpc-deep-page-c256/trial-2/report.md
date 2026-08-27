# Feed 压测报告：hybrid / rpc / deep-page-c256

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:41:26+08:00
- 采样时长：1m0.0679379s
- 并发：256
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 183299 | 183299 | 0 | 0 | 3052.24 | 70.767 | 141.141 | 171.798 | 243.365 | 627.474 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 0 | 0.000 |
| counter | 220 | 0.001 |
| mysql | 183299 | 1.000 |
| redis | 183299 | 1.000 |
| relation | 220 | 0.001 |

- Cold compute：183299（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 183299 |
| l1_fresh | 0 |
| l1_stale | 0 |
| l2_fresh | 0 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 183299 | 0.166 |
| counter | 220 | 0.717 |
| hydrate | 183299 | 63.499 |
| inbox | 183299 | 0.165 |
| merge_dedup | 183299 | 0.014 |
| relation | 220 | 1.392 |
| route | 183299 | 0.007 |
| total | 183299 | 63.704 |

## Redis 本轮边界增量

- Commands：1159374；input：83476262 bytes；output：483372148 bytes
- Hits/Misses：1106564/220；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：21355；safety epoch：498 -> 498

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.066 |
| client:loadtest | cpu_percent_total | 65.057 |
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
| docker:zg-canal | cpu_percent | 2.800 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 81.000 |
| docker:zg-es | cpu_percent | 4.140 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.490 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 198.830 |
| docker:zg-kafka | memory_percent | 8.630 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 68.060 |
| docker:zg-zk | memory_percent | 1.430 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22513574.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 87.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 13.147 |
| process:counter | cpu_seconds_total | 135.594 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 44097536.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.817 |
| process:gateway | cpu_seconds_total | 0.656 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 36814848.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 477.340 |
| process:knowpost | cpu_seconds_total | 6934.859 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 131489792.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.094 |
| process:relation | cpu_seconds_total | 8.078 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 48488448.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.250 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36990976.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.526 |
| process:user-storage | cpu_seconds_total | 1.062 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 34893824.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 193716999.000 |
| redis | connected_clients | 204.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 498.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677188.000 |
| redis | keyspace_hits | 991834531.000 |
| redis | keyspace_misses | 5377885.000 |
| redis | net_input_bytes | 41728628118.000 |
| redis | net_output_bytes | 186429524576.000 |
| redis | ops_per_sec | 21355.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 40926.000 |
| redis | used_memory_bytes | 101652928.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
