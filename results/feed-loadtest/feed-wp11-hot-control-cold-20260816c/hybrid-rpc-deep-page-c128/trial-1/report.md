# Feed 压测报告：hybrid / rpc / deep-page-c128

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:42:27+08:00
- 采样时长：1m0.0446795s
- 并发：128
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 170714 | 170714 | 0 | 0 | 2843.69 | 43.280 | 51.036 | 57.450 | 108.314 | 153.384 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 165 | 0.001 |
| redis | 512307 | 3.001 |
| relation | 170714 | 1.000 |

- Cold compute：170714（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 170714 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 170714 | 0.198 |
| counter | 240 | 0.732 |
| hydrate | 170714 | 0.351 |
| inbox | 170714 | 0.217 |
| merge_dedup | 170714 | 0.014 |
| relation | 170714 | 43.709 |
| route | 170714 | 0.004 |
| total | 170714 | 44.520 |

## Redis 本轮边界增量

- Commands：1430848；input：433365871 bytes；output：2012749996 bytes
- Hits/Misses：10931619/170881；run hit rate：98.46%
- Evicted/Rejected：0/0；ops/s max：26126；safety epoch：422 -> 422

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.219 |
| client:loadtest | cpu_percent_total | 67.502 |
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
| docker:zg-canal | cpu_percent | 2.660 |
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.670 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 150.000 |
| docker:zg-etcd | cpu_percent | 6.110 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 205.380 |
| docker:zg-kafka | memory_percent | 8.450 |
| docker:zg-kafka | pids | 120.000 |
| docker:zg-zk | cpu_percent | 59.450 |
| docker:zg-zk | memory_percent | 1.050 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 15606738.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 74.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 10.381 |
| process:counter | cpu_seconds_total | 155.016 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43397120.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.772 |
| process:gateway | cpu_seconds_total | 0.578 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 36974592.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 320.062 |
| process:knowpost | cpu_seconds_total | 5920.391 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 101675008.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 273.203 |
| process:relation | cpu_seconds_total | 5504.297 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 97546240.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 3.114 |
| process:search | cpu_seconds_total | 1.062 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37056512.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 1.141 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35188736.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 127112310.000 |
| redis | connected_clients | 241.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 422.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 698220.000 |
| redis | keyspace_hits | 749212619.000 |
| redis | keyspace_misses | 794419.000 |
| redis | net_input_bytes | 30343706320.000 |
| redis | net_output_bytes | 137095233936.000 |
| redis | ops_per_sec | 26126.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22986.000 |
| redis | used_memory_bytes | 106557192.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
