# Feed 压测报告：hybrid / rpc / page1-c64

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:17:21+08:00
- 采样时长：10.0075608s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 53225 | 53225 | 0 | 0 | 5320.16 | 11.354 | 14.536 | 17.650 | 23.933 | 51.344 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 106450 | 2.000 |
| relation | 40 | 0.001 |

- Cold compute：53225（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 53225 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 53225 | 6.089 |
| counter | 40 | 3.406 |
| hydrate | 53225 | 5.614 |
| inbox | 53225 | 6.088 |
| merge_dedup | 53225 | 0.011 |
| relation | 40 | 3.683 |
| route | 53225 | 0.066 |
| total | 53225 | 11.801 |

## Redis 本轮边界增量

- Commands：375753；input：98862171 bytes；output：901024467 bytes
- Hits/Misses：2449599/15；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：40691；safety epoch：3588 -> 3588

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.743 |
| client:loadtest | cpu_percent_total | 75.880 |
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
| docker:zg-canal | cpu_percent | 0.140 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.360 |
| docker:zg-es | memory_percent | 11.950 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 3.810 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 28.330 |
| docker:zg-kafka | memory_percent | 7.020 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.110 |
| docker:zg-zk | memory_percent | 0.890 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 83053.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 1.571 |
| process:counter | cpu_seconds_total | 31.969 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 45768704.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 38264832.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 247.549 |
| process:knowpost | cpu_seconds_total | 326.016 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 82022400.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 16.656 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 53534720.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.188 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42860544.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.219 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40538112.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 6852400.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3588.000 |
| redis | hit_rate | 0.997 |
| redis | keys | 597343.000 |
| redis | keyspace_hits | 9758968.000 |
| redis | keyspace_misses | 29283.000 |
| redis | net_input_bytes | 813502721.000 |
| redis | net_output_bytes | 3647085697.000 |
| redis | ops_per_sec | 40691.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21430.000 |
| redis | used_memory_bytes | 104427040.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
