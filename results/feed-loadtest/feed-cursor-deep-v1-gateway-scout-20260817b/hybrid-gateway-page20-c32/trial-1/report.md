# Feed 压测报告：hybrid / gateway / page20-c32

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:36:00+08:00
- 采样时长：10.0288821s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 8921 | 8921 | 0 | 0 | 889.58 | 35.026 | 49.311 | 56.044 | 80.402 | 132.177 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.004 |
| mysql | 141 | 0.016 |
| redis | 17983 | 2.016 |
| relation | 40 | 0.004 |

- Cold compute：8921（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3746820 | 420.000 |
| merge_candidates | 7502561 | 841.000 |
| redis_commands | 53526 | 6.000 |
| redis_members | 8216241 | 921.000 |
| redis_roundtrips | 8921 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 8921 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 8921 | 16.953 |
| counter | 40 | 8.205 |
| hydrate | 8921 | 17.345 |
| inbox | 8921 | 16.951 |
| merge_dedup | 8921 | 0.076 |
| relation | 40 | 7.637 |
| route | 8921 | 0.555 |
| total | 8921 | 35.044 |

## Redis 本轮边界增量

- Commands：64653；input：135346607 bytes；output：959975663 bytes
- Hits/Misses：3801425/179；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：7162；safety epoch：3614 -> 3614

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.383 |
| client:loadtest | cpu_percent_total | 22.124 |
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
| docker:zg-canal | cpu_percent | 0.130 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.020 |
| docker:zg-es | memory_percent | 12.080 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 3.940 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 142.740 |
| docker:zg-kafka | memory_percent | 7.520 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 6.390 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 94542.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 29.943 |
| process:counter | cpu_seconds_total | 7.797 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45834240.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 86.215 |
| process:gateway | cpu_seconds_total | 141.141 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52301824.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 201.710 |
| process:knowpost | cpu_seconds_total | 163.969 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71852032.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 0.906 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 46174208.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.375 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 41725952.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 39.041 |
| process:user-storage | cpu_seconds_total | 52.516 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 50241536.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18915110.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3614.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597663.000 |
| redis | keyspace_hits | 104320385.000 |
| redis | keyspace_misses | 70322.000 |
| redis | net_input_bytes | 4900501713.000 |
| redis | net_output_bytes | 27833613829.000 |
| redis | ops_per_sec | 7162.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22549.000 |
| redis | used_memory_bytes | 103802280.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
