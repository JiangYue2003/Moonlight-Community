# Feed 压测报告：hybrid / rpc / page5-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:50:27+08:00
- 采样时长：1m0.0153589s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 139477 | 139477 | 0 | 0 | 2324.45 | 6.161 | 10.288 | 11.336 | 13.744 | 24.220 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 688 | 0.005 |
| redis | 279642 | 2.005 |
| relation | 240 | 0.002 |

- Cold compute：139477（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 16737240 | 120.000 |
| merge_candidates | 33613957 | 241.000 |
| redis_commands | 836862 | 6.000 |
| redis_members | 86615217 | 621.000 |
| redis_roundtrips | 139477 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 139477 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 139477 | 3.354 |
| counter | 240 | 3.053 |
| hydrate | 139477 | 3.145 |
| inbox | 139477 | 3.353 |
| merge_dedup | 139477 | 0.031 |
| relation | 240 | 3.519 |
| route | 139477 | 0.053 |
| total | 139477 | 6.632 |

## Redis 本轮边界增量

- Commands：1008056；input：651529066 bytes；output：6389217184 bytes
- Hits/Misses：17580788/919；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：17836；safety epoch：3629 -> 3629

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.766 |
| client:loadtest | cpu_percent_total | 44.260 |
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
| docker:zg-canal | cpu_percent | 1.910 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 80.000 |
| docker:zg-es | cpu_percent | 3.660 |
| docker:zg-es | memory_percent | 12.200 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 5.130 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 204.430 |
| docker:zg-kafka | memory_percent | 7.520 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 46.920 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 106603.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.188 |
| process:counter | cpu_seconds_total | 33.594 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45436928.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 258.141 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45285376.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 258.524 |
| process:knowpost | cpu_seconds_total | 1721.078 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72159232.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.870 |
| process:relation | cpu_seconds_total | 3.156 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48291840.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 6.973 |
| process:search | cpu_seconds_total | 1.641 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42328064.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.999 |
| process:user-storage | cpu_seconds_total | 95.312 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44183552.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 39961109.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3629.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597874.000 |
| redis | keyspace_hits | 282566963.000 |
| redis | keyspace_misses | 94650.000 |
| redis | net_input_bytes | 11947381183.000 |
| redis | net_output_bytes | 90742347054.000 |
| redis | ops_per_sec | 17836.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23466.000 |
| redis | used_memory_bytes | 102642088.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
