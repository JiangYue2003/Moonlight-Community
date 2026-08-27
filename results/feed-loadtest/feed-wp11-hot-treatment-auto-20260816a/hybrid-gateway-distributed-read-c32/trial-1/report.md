# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:59:01+08:00
- 采样时长：1m0.0252138s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 311688 | 311688 | 0 | 0 | 5194.44 | 5.995 | 7.879 | 9.013 | 11.710 | 25.350 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 2437 | 0.008 |
| counter | 179 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 2412 | 0.008 |
| relation | 179 | 0.001 |

- Cold compute：365（0.001 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 310346 |
| l1_stale | 0 |
| l2_fresh | 1342 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 365 | 0.221 |
| counter | 179 | 0.786 |
| hydrate | 365 | 0.274 |
| inbox | 365 | 0.221 |
| merge_dedup | 365 | 0.012 |
| relation | 179 | 1.851 |
| route | 365 | 1.307 |
| total | 311688 | 0.008 |

## Redis 本轮边界增量

- Commands：62870；input：5586830 bytes；output：6989126 bytes
- Hits/Misses：23911/179；run hit rate：99.26%
- Evicted/Rejected：0/0；ops/s max：1428；safety epoch：509 -> 509

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.820 |
| client:loadtest | cpu_percent_total | 109.121 |
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
| docker:zg-canal | cpu_percent | 2.950 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.890 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.240 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 218.830 |
| docker:zg-kafka | memory_percent | 8.650 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 66.860 |
| docker:zg-zk | memory_percent | 1.410 |
| docker:zg-zk | pids | 103.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22731279.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 28.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 11.611 |
| process:counter | cpu_seconds_total | 195.672 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 44150784.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 422.033 |
| process:gateway | cpu_seconds_total | 2684.250 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 54095872.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 217.442 |
| process:knowpost | cpu_seconds_total | 8480.078 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 113856512.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.545 |
| process:relation | cpu_seconds_total | 9.375 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 48222208.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.782 |
| process:search | cpu_seconds_total | 1.750 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36999168.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 218.328 |
| process:user-storage | cpu_seconds_total | 1343.156 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 77025280.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 195968134.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 509.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677291.000 |
| redis | keyspace_hits | 993183028.000 |
| redis | keyspace_misses | 5380036.000 |
| redis | net_input_bytes | 41891747795.000 |
| redis | net_output_bytes | 187027023909.000 |
| redis | ops_per_sec | 1428.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 41980.000 |
| redis | used_memory_bytes | 100947168.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
