# Feed 压测报告：hybrid / gateway / page50-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:12:57+08:00
- 采样时长：1m0.0143065s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 30121 | 30121 | 0 | 0 | 501.91 | 30.822 | 41.351 | 44.861 | 52.091 | 75.279 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.008 |
| mysql | 4289 | 0.142 |
| redis | 64531 | 2.142 |
| relation | 240 | 0.008 |

- Cold compute：30121（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 30723420 | 1020.000 |
| merge_candidates | 45151379 | 1499.000 |
| redis_commands | 180726 | 6.000 |
| redis_members | 45181500 | 1500.000 |
| redis_roundtrips | 30121 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 30121 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 30121 | 6.807 |
| counter | 240 | 5.672 |
| hydrate | 30121 | 23.356 |
| inbox | 30121 | 6.804 |
| merge_dedup | 30121 | 0.174 |
| relation | 240 | 6.050 |
| route | 30121 | 0.376 |
| total | 30121 | 30.906 |

## Redis 本轮边界增量

- Commands：237371；input：1091712903 bytes；output：6921521599 bytes
- Hits/Misses：30905050/6551；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：4582；safety epoch：3694 -> 3694

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.875 |
| client:loadtest | cpu_percent_total | 14.007 |
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
| docker:zg-canal | cpu_percent | 0.980 |
| docker:zg-canal | memory_percent | 4.230 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 0.620 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.400 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 131.360 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 0.290 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 278178.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.194 |
| process:counter | cpu_seconds_total | 153.453 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47161344.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 62.708 |
| process:gateway | cpu_seconds_total | 3234.531 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51417088.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 249.536 |
| process:knowpost | cpu_seconds_total | 9987.031 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 84668416.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.118 |
| process:relation | cpu_seconds_total | 26.062 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49152000.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 3.703 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38162432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 27.952 |
| process:user-storage | cpu_seconds_total | 1171.453 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 51785728.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 181250890.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3694.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596604.000 |
| redis | keyspace_hits | 1438792002.000 |
| redis | keyspace_misses | 323633.000 |
| redis | net_input_bytes | 61527471091.000 |
| redis | net_output_bytes | 379246132197.000 |
| redis | ops_per_sec | 4582.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28416.000 |
| redis | used_memory_bytes | 103082768.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
