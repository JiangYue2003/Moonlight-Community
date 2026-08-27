# Feed 压测报告：hybrid / rpc / page50-c32

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:21:47+08:00
- 采样时长：10.0258513s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 5772 | 5772 | 0 | 0 | 575.74 | 54.264 | 67.998 | 72.409 | 82.763 | 103.148 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.007 |
| mysql | 1295 | 0.224 |
| redis | 12839 | 2.224 |
| relation | 40 | 0.007 |

- Cold compute：5772（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 5772 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 5772 | 12.313 |
| counter | 40 | 5.024 |
| hydrate | 5772 | 41.921 |
| inbox | 5772 | 12.309 |
| merge_dedup | 5772 | 0.171 |
| relation | 40 | 5.608 |
| route | 5772 | 0.529 |
| total | 5772 | 55.152 |

## Redis 本轮边界增量

- Commands：44957；input：209329077 bytes；output：1326172700 bytes
- Hits/Misses：5921168/2166；run hit rate：99.96%
- Evicted/Rejected：0/0；ops/s max：4786；safety epoch：3602 -> 3602

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.633 |
| client:loadtest | cpu_percent_total | 10.130 |
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
| docker:zg-canal | cpu_percent | 1.860 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.240 |
| docker:zg-es | memory_percent | 11.990 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.200 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 40.190 |
| docker:zg-kafka | memory_percent | 7.150 |
| docker:zg-kafka | pids | 112.000 |
| docker:zg-zk | cpu_percent | 0.110 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 88431.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.098 |
| process:counter | cpu_seconds_total | 37.750 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 48439296.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 37515264.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 247.873 |
| process:knowpost | cpu_seconds_total | 671.859 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 101134336.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.793 |
| process:relation | cpu_seconds_total | 17.406 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 54636544.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.359 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42926080.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.344 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40947712.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 12910525.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3602.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597837.000 |
| redis | keyspace_hits | 68161249.000 |
| redis | keyspace_misses | 52088.000 |
| redis | net_input_bytes | 3244116097.000 |
| redis | net_output_bytes | 19190181214.000 |
| redis | ops_per_sec | 4786.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21696.000 |
| redis | used_memory_bytes | 106071728.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
