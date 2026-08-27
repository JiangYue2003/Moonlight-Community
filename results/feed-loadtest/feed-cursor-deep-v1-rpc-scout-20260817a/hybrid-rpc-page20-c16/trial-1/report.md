# Feed 压测报告：hybrid / rpc / page20-c16

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:19:34+08:00
- 采样时长：10.0118808s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 11811 | 11811 | 0 | 0 | 1179.76 | 11.927 | 20.535 | 22.576 | 26.855 | 38.609 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.003 |
| mysql | 29 | 0.002 |
| redis | 23651 | 2.002 |
| relation | 40 | 0.003 |

- Cold compute：11811（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 11811 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 11811 | 6.430 |
| counter | 40 | 5.448 |
| hydrate | 11811 | 6.522 |
| inbox | 11811 | 6.428 |
| merge_dedup | 11811 | 0.080 |
| relation | 40 | 6.524 |
| route | 11811 | 0.165 |
| total | 11811 | 13.306 |

## Redis 本轮边界增量

- Commands：85906；input：179198785 bytes；output：1270991507 bytes
- Hits/Misses：5032705/44；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：8982；safety epoch：3595 -> 3595

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.951 |
| client:loadtest | cpu_percent_total | 31.213 |
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
| docker:zg-canal | cpu_percent | 1.780 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.330 |
| docker:zg-es | memory_percent | 11.960 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.840 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 137.920 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 0.170 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 84390.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 25.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.564 |
| process:counter | cpu_seconds_total | 34.797 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 47153152.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 37470208.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 217.948 |
| process:knowpost | cpu_seconds_total | 484.625 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 80252928.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.795 |
| process:relation | cpu_seconds_total | 16.875 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 52883456.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.281 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42901504.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.281 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40910848.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 9669633.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3595.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597400.000 |
| redis | keyspace_hits | 32204343.000 |
| redis | keyspace_misses | 36524.000 |
| redis | net_input_bytes | 1768610345.000 |
| redis | net_output_bytes | 11027063186.000 |
| redis | ops_per_sec | 8982.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21563.000 |
| redis | used_memory_bytes | 102741216.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
