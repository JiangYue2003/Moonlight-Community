# Feed 压测报告：hybrid / gateway / page1-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:35:14+08:00
- 采样时长：1m0.0118923s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 85267 | 85267 | 0 | 0 | 1420.98 | 11.069 | 15.293 | 19.432 | 24.354 | 47.808 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 178 | 0.002 |
| redis | 170712 | 2.002 |
| relation | 240 | 0.003 |

- Cold compute：85267（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3410680 | 40.000 |
| merge_candidates | 6906627 | 81.000 |
| redis_commands | 511602 | 6.000 |
| redis_members | 20975682 | 246.000 |
| redis_roundtrips | 85267 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 85267 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 85267 | 4.973 |
| counter | 240 | 5.213 |
| hydrate | 85267 | 4.655 |
| inbox | 85267 | 4.972 |
| merge_dedup | 85267 | 0.011 |
| relation | 240 | 6.575 |
| route | 85267 | 0.126 |
| total | 85267 | 9.789 |

## Redis 本轮边界增量

- Commands：618146；input：159340853 bytes；output：1443827901 bytes
- Hits/Misses：3929525/238；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：11928；safety epoch：3664 -> 3664

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.287 |
| client:loadtest | cpu_percent_total | 52.594 |
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
| docker:zg-canal | cpu_percent | 0.370 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.500 |
| docker:zg-es | memory_percent | 12.330 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.960 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 154.500 |
| docker:zg-kafka | memory_percent | 7.330 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 47.620 |
| docker:zg-zk | memory_percent | 0.990 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 218643.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 7.747 |
| process:counter | cpu_seconds_total | 96.547 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46854144.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 229.429 |
| process:gateway | cpu_seconds_total | 539.312 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50868224.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 154.932 |
| process:knowpost | cpu_seconds_total | 6816.453 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70725632.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 6.197 |
| process:relation | cpu_seconds_total | 13.031 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49422336.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 2.828 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38125568.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 85.171 |
| process:user-storage | cpu_seconds_total | 193.766 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54923264.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 148696070.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3664.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597343.000 |
| redis | keyspace_hits | 1049392533.000 |
| redis | keyspace_misses | 247322.000 |
| redis | net_input_bytes | 46022259198.000 |
| redis | net_output_bytes | 271592094837.000 |
| redis | ops_per_sec | 11928.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26153.000 |
| redis | used_memory_bytes | 102559304.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
