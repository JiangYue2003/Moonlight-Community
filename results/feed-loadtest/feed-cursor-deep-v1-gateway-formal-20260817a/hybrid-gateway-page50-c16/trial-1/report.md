# Feed 压测报告：hybrid / gateway / page50-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:11:42+08:00
- 采样时长：1m0.0245323s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 29695 | 29695 | 0 | 0 | 494.73 | 31.674 | 41.166 | 44.319 | 51.409 | 77.326 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.008 |
| mysql | 3471 | 0.117 |
| redis | 62861 | 2.117 |
| relation | 240 | 0.008 |

- Cold compute：29695（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 30288900 | 1020.000 |
| merge_candidates | 44512805 | 1499.000 |
| redis_commands | 178170 | 6.000 |
| redis_members | 44542500 | 1500.000 |
| redis_roundtrips | 29695 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 29695 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 29695 | 6.943 |
| counter | 240 | 5.783 |
| hydrate | 29695 | 23.640 |
| inbox | 29695 | 6.939 |
| merge_dedup | 29695 | 0.168 |
| relation | 240 | 6.985 |
| route | 29695 | 0.387 |
| total | 29695 | 31.327 |

## Redis 本轮边界增量

- Commands：232269；input：1075907423 bytes；output：6823857848 bytes
- Hits/Misses：30469556/4966；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：4437；safety epoch：3693 -> 3693

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.786 |
| client:loadtest | cpu_percent_total | 12.573 |
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
| docker:zg-canal | cpu_percent | 0.180 |
| docker:zg-canal | memory_percent | 4.220 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.080 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.950 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 112.830 |
| docker:zg-kafka | memory_percent | 7.360 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 44.780 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 271165.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.194 |
| process:counter | cpu_seconds_total | 151.719 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46198784.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 60.326 |
| process:gateway | cpu_seconds_total | 3207.688 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52129792.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 218.975 |
| process:knowpost | cpu_seconds_total | 9857.844 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 81825792.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.414 |
| process:relation | cpu_seconds_total | 25.734 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48672768.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.538 |
| process:search | cpu_seconds_total | 3.703 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38162432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 31.855 |
| process:user-storage | cpu_seconds_total | 1160.109 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 53133312.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 180967076.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3693.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596613.000 |
| redis | keyspace_hits | 1402622089.000 |
| redis | keyspace_misses | 313206.000 |
| redis | net_input_bytes | 60248815411.000 |
| redis | net_output_bytes | 371145233701.000 |
| redis | ops_per_sec | 4437.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28341.000 |
| redis | used_memory_bytes | 103059080.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
