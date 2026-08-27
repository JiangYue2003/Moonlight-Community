# Feed 压测报告：hybrid / gateway / page1-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:36:29+08:00
- 采样时长：1m0.0116s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 113824 | 113824 | 0 | 0 | 1896.99 | 7.990 | 11.685 | 14.052 | 17.777 | 34.547 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 275 | 0.002 |
| redis | 227923 | 2.002 |
| relation | 240 | 0.002 |

- Cold compute：113824（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 4552960 | 40.000 |
| merge_candidates | 9219744 | 81.000 |
| redis_commands | 682944 | 6.000 |
| redis_members | 28000704 | 246.000 |
| redis_roundtrips | 113824 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 113824 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 113824 | 3.746 |
| counter | 240 | 3.925 |
| hydrate | 113824 | 3.386 |
| inbox | 113824 | 3.745 |
| merge_dedup | 113824 | 0.011 |
| relation | 240 | 4.924 |
| route | 113824 | 0.077 |
| total | 113824 | 7.241 |

## Redis 本轮边界增量

- Commands：823001；input：212645179 bytes；output：1927298967 bytes
- Hits/Misses：5243112/299；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：16750；safety epoch：3665 -> 3665

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.819 |
| client:loadtest | cpu_percent_total | 61.108 |
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
| docker:zg-canal | cpu_percent | 2.110 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.500 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.190 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 140.310 |
| docker:zg-kafka | memory_percent | 7.200 |
| docker:zg-kafka | pids | 113.000 |
| docker:zg-zk | cpu_percent | 46.780 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 102.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 219421.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.057 |
| process:counter | cpu_seconds_total | 99.953 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47108096.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 238.900 |
| process:gateway | cpu_seconds_total | 676.312 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50933760.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 181.962 |
| process:knowpost | cpu_seconds_total | 6923.875 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70787072.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.777 |
| process:relation | cpu_seconds_total | 13.531 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49143808.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 2.844 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38125568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 92.367 |
| process:user-storage | cpu_seconds_total | 241.078 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 58466304.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 149647858.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3665.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597177.000 |
| redis | keyspace_hits | 1055433961.000 |
| redis | keyspace_misses | 247741.000 |
| redis | net_input_bytes | 46267534060.000 |
| redis | net_output_bytes | 273812880285.000 |
| redis | ops_per_sec | 16750.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26228.000 |
| redis | used_memory_bytes | 102264424.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
