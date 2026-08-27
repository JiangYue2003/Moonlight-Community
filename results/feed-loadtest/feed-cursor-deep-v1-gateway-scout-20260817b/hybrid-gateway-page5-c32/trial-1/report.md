# Feed 压测报告：hybrid / gateway / page5-c32

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:34:45+08:00
- 采样时长：10.0103219s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 12464 | 12464 | 0 | 0 | 1245.24 | 27.714 | 35.360 | 39.594 | 59.508 | 89.390 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.003 |
| mysql | 197 | 0.016 |
| redis | 25125 | 2.016 |
| relation | 40 | 0.003 |

- Cold compute：12464（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1495680 | 120.000 |
| merge_candidates | 3003824 | 241.000 |
| redis_commands | 74784 | 6.000 |
| redis_members | 7740144 | 621.000 |
| redis_roundtrips | 12464 | 1.000 |

| page cache source | requests |
|---|---:|
| bypass | 12464 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 12464 | 12.486 |
| counter | 40 | 4.827 |
| hydrate | 12464 | 11.699 |
| inbox | 12464 | 12.484 |
| merge_dedup | 12464 | 0.025 |
| relation | 40 | 7.477 |
| route | 12464 | 0.360 |
| total | 12464 | 24.614 |

## Redis 本轮边界增量

- Commands：89630；input：58196578 bytes；output：570929700 bytes
- Hits/Misses：1571477/248；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：10873；safety epoch：3610 -> 3610

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.893 |
| client:loadtest | cpu_percent_total | 30.281 |
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
| docker:zg-canal | cpu_percent | 1.800 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.760 |
| docker:zg-es | memory_percent | 12.080 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.000 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 22.080 |
| docker:zg-kafka | memory_percent | 7.030 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 0.200 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 93633.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.095 |
| process:counter | cpu_seconds_total | 5.469 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 44261376.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 158.403 |
| process:gateway | cpu_seconds_total | 87.516 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52506624.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 164.960 |
| process:knowpost | cpu_seconds_total | 91.297 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 66510848.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.188 |
| process:relation | cpu_seconds_total | 0.453 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 45625344.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.234 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 41107456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 49.160 |
| process:user-storage | cpu_seconds_total | 32.188 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 50970624.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 18178711.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3610.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597337.000 |
| redis | keyspace_hits | 93587970.000 |
| redis | keyspace_misses | 64298.000 |
| redis | net_input_bytes | 4478698773.000 |
| redis | net_output_bytes | 25007588894.000 |
| redis | ops_per_sec | 10873.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22474.000 |
| redis | used_memory_bytes | 102844568.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
