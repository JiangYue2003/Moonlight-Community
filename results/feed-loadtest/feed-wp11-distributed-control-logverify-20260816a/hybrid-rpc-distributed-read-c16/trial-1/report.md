# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-distributed-control-logverify-20260816a`
- 开始时间：2026-08-16T23:29:58+08:00
- 采样时长：5.0026152s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 31567 | 31567 | 0 | 0 | 6312.06 | 2.369 | 3.442 | 3.829 | 4.775 | 7.605 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 1200 | 0.038 |
| mysql | 0 | 0.000 |
| redis | 94701 | 3.000 |
| relation | 31567 | 1.000 |

- Cold compute：31567（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 31567 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 31567 | 0.359 |
| counter | 1200 | 0.685 |
| hydrate | 31567 | 0.412 |
| inbox | 31567 | 0.376 |
| merge_dedup | 31567 | 0.004 |
| relation | 31567 | 1.085 |
| route | 31567 | 0.029 |
| total | 31567 | 2.284 |

## Redis 本轮边界增量

- Commands：251251；input：28920184 bytes；output：86405425 bytes
- Hits/Misses：585000/34211；run hit rate：94.48%
- Evicted/Rejected：0/0；ops/s max：51954；safety epoch：523 -> 523

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.305 |
| client:loadtest | cpu_percent_total | 100.885 |
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
| docker:zg-canal | cpu_percent | 0.200 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.670 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.130 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 70.950 |
| docker:zg-kafka | memory_percent | 8.230 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 53.830 |
| docker:zg-zk | memory_percent | 1.400 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 22779747.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 34.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 26.975 |
| process:counter | cpu_seconds_total | 5.453 |
| process:counter | pid | 8320.000 |
| process:counter | process_start_ms | 1786894081426.000 |
| process:counter | rss_bytes | 45142016.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.000 |
| process:gateway | pid | 9852.000 |
| process:gateway | process_start_ms | 1786894099200.000 |
| process:gateway | rss_bytes | 36798464.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 360.993 |
| process:knowpost | cpu_seconds_total | 29.297 |
| process:knowpost | pid | 7360.000 |
| process:knowpost | process_start_ms | 1786894091315.000 |
| process:knowpost | rss_bytes | 61464576.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 277.661 |
| process:relation | cpu_seconds_total | 18.484 |
| process:relation | pid | 26448.000 |
| process:relation | process_start_ms | 1786894085233.000 |
| process:relation | rss_bytes | 49541120.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.109 |
| process:search | pid | 19932.000 |
| process:search | process_start_ms | 1786894095566.000 |
| process:search | rss_bytes | 35344384.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 0.062 |
| process:user-storage | pid | 7096.000 |
| process:user-storage | process_start_ms | 1786894077784.000 |
| process:user-storage | rss_bytes | 34263040.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 200944218.000 |
| redis | connected_clients | 50.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 523.000 |
| redis | hit_rate | 0.995 |
| redis | keys | 677094.000 |
| redis | keyspace_hits | 994287451.000 |
| redis | keyspace_misses | 5433911.000 |
| redis | net_input_bytes | 42271286317.000 |
| redis | net_output_bytes | 187300721392.000 |
| redis | ops_per_sec | 51954.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 47382.000 |
| redis | used_memory_bytes | 97297176.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
