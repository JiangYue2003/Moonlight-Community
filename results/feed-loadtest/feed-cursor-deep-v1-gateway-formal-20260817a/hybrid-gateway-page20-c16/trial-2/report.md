# Feed 压测报告：hybrid / gateway / page20-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:57:46+08:00
- 采样时长：1m0.015719s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 62266 | 62266 | 0 | 0 | 1037.57 | 13.887 | 22.941 | 25.289 | 30.828 | 63.658 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.004 |
| mysql | 1946 | 0.031 |
| redis | 126478 | 2.031 |
| relation | 240 | 0.004 |

- Cold compute：62266（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 26151720 | 420.000 |
| merge_candidates | 52365706 | 841.000 |
| redis_commands | 373596 | 6.000 |
| redis_members | 57346986 | 921.000 |
| redis_roundtrips | 62266 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 62266 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 62266 | 7.141 |
| counter | 240 | 7.343 |
| hydrate | 62266 | 7.069 |
| inbox | 62266 | 7.139 |
| merge_dedup | 62266 | 0.084 |
| relation | 240 | 7.620 |
| route | 62266 | 0.230 |
| total | 62266 | 14.639 |

## Redis 本轮边界增量

- Commands：457013；input：945331673 bytes；output：6700168994 bytes
- Hits/Misses：26530140/2649；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：8163；safety epoch：3682 -> 3682

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.357 |
| client:loadtest | cpu_percent_total | 21.713 |
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
| docker:zg-canal | cpu_percent | 2.190 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.500 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.520 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 154.630 |
| docker:zg-kafka | memory_percent | 7.240 |
| docker:zg-kafka | pids | 111.000 |
| docker:zg-zk | cpu_percent | 43.030 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 241584.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.646 |
| process:counter | cpu_seconds_total | 129.141 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46542848.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 96.857 |
| process:gateway | cpu_seconds_total | 2203.766 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51105792.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 238.597 |
| process:knowpost | cpu_seconds_total | 8598.844 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 74866688.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.550 |
| process:relation | cpu_seconds_total | 21.109 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49332224.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 3.406 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38162432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 43.332 |
| process:user-storage | cpu_seconds_total | 803.312 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54542336.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 167510547.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3682.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597195.000 |
| redis | keyspace_hits | 1227339715.000 |
| redis | keyspace_misses | 270308.000 |
| redis | net_input_bytes | 53246664515.000 |
| redis | net_output_bytes | 329711971872.000 |
| redis | ops_per_sec | 8163.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27505.000 |
| redis | used_memory_bytes | 102739728.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
