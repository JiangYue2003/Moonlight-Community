# Feed 压测报告：hybrid / gateway / hot-read-c128

- Run ID：`feed-wp11-hot-treatment-auto-20260816a`
- 开始时间：2026-08-16T21:52:45+08:00
- 采样时长：1m0.0391364s
- 并发：128
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 294296 | 294296 | 0 | 0 | 4903.42 | 24.792 | 35.775 | 44.709 | 65.750 | 295.149 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 158 | 0.001 |
| counter | 9 | 0.000 |
| mysql | 0 | 0.000 |
| redis | 122 | 0.000 |
| relation | 9 | 0.000 |

- Cold compute：19（0.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 294195 |
| l1_stale | 0 |
| l2_fresh | 101 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 19 | 0.353 |
| counter | 9 | 1.941 |
| hydrate | 19 | 0.164 |
| inbox | 19 | 0.315 |
| merge_dedup | 19 | 0.028 |
| relation | 9 | 2.639 |
| route | 19 | 2.170 |
| total | 294296 | 0.007 |

## Redis 本轮边界增量

- Commands：53309；input：3815437 bytes；output：1436777 bytes
- Hits/Misses：1372/9；run hit rate：99.35%
- Evicted/Rejected：0/0；ops/s max：1116；safety epoch：507 -> 507

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.936 |
| client:loadtest | cpu_percent_total | 110.969 |
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
| docker:zg-canal | cpu_percent | 2.470 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.390 |
| docker:zg-es | memory_percent | 11.930 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 7.030 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 318.330 |
| docker:zg-kafka | memory_percent | 8.640 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 57.300 |
| docker:zg-zk | memory_percent | 1.470 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3103.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3103.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22730721.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 5.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.039 |
| process:counter | cpu_seconds_total | 176.750 |
| process:counter | pid | 30092.000 |
| process:counter | process_start_ms | 1786884930660.000 |
| process:counter | rss_bytes | 43053056.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 446.871 |
| process:gateway | cpu_seconds_total | 2124.438 |
| process:gateway | pid | 28604.000 |
| process:gateway | process_start_ms | 1786884948386.000 |
| process:gateway | rss_bytes | 60268544.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 192.984 |
| process:knowpost | cpu_seconds_total | 8202.422 |
| process:knowpost | pid | 10768.000 |
| process:knowpost | process_start_ms | 1786884940053.000 |
| process:knowpost | rss_bytes | 130011136.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.774 |
| process:relation | cpu_seconds_total | 9.078 |
| process:relation | pid | 5972.000 |
| process:relation | process_start_ms | 1786884935386.000 |
| process:relation | rss_bytes | 44486656.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.547 |
| process:search | pid | 4808.000 |
| process:search | process_start_ms | 1786884944422.000 |
| process:search | rss_bytes | 36974592.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 214.174 |
| process:user-storage | cpu_seconds_total | 1057.250 |
| process:user-storage | pid | 16208.000 |
| process:user-storage | process_start_ms | 1786884926376.000 |
| process:user-storage | rss_bytes | 82968576.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 195623277.000 |
| redis | connected_clients | 215.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 507.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677257.000 |
| redis | keyspace_hits | 993151231.000 |
| redis | keyspace_misses | 5378918.000 |
| redis | net_input_bytes | 41865766090.000 |
| redis | net_output_bytes | 187012846691.000 |
| redis | ops_per_sec | 1116.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 41604.000 |
| redis | used_memory_bytes | 100031408.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
