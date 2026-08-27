# Feed 压测报告：hybrid / gateway / page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-page5-recheck-c`
- 开始时间：2026-08-17T19:06:48+08:00
- 采样时长：10.0067024s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 14291 | 14291 | 0 | 0 | 1428.29 | 11.280 | 15.388 | 18.509 | 26.361 | 39.240 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.003 |
| mysql | 117 | 0.008 |
| redis | 28699 | 2.008 |
| relation | 40 | 0.003 |

- Cold compute：14291（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1714920 | 120.000 |
| merge_candidates | 3444131 | 241.000 |
| redis_commands | 85746 | 6.000 |
| redis_members | 8874711 | 621.000 |
| redis_roundtrips | 14291 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 14291 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 14291 | 5.286 |
| counter | 40 | 4.067 |
| hydrate | 14291 | 4.838 |
| inbox | 14291 | 5.285 |
| merge_dedup | 14291 | 0.026 |
| relation | 40 | 4.592 |
| route | 14291 | 0.111 |
| total | 14291 | 10.302 |

## Redis 本轮边界增量

- Commands：103727；input：66772918 bytes；output：654657412 bytes
- Hits/Misses：1801745/171；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：11775；safety epoch：3735 -> 3735

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.991 |
| client:loadtest | cpu_percent_total | 31.854 |
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
| docker:zg-canal | memory_percent | 4.270 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.200 |
| docker:zg-es | memory_percent | 12.350 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 0.890 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 152.900 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 0.120 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 454939.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.732 |
| process:counter | cpu_seconds_total | 241.422 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46374912.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 123.930 |
| process:gateway | cpu_seconds_total | 5462.281 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50712576.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 167.305 |
| process:knowpost | cpu_seconds_total | 14520.172 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72560640.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.155 |
| process:relation | cpu_seconds_total | 41.422 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48467968.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 4.812 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38027264.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 60.335 |
| process:user-storage | cpu_seconds_total | 1941.094 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 52117504.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 242632121.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3735.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 595999.000 |
| redis | keyspace_hits | 2078104464.000 |
| redis | keyspace_misses | 577624.000 |
| redis | net_input_bytes | 87983801219.000 |
| redis | net_output_bytes | 524789869617.000 |
| redis | ops_per_sec | 11775.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31597.000 |
| redis | used_memory_bytes | 101562160.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
