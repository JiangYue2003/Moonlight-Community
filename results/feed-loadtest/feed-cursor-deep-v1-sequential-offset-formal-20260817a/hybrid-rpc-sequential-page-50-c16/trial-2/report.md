# Feed 压测报告：hybrid / rpc / sequential-page-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:48:28+08:00
- 采样时长：1m0.2092142s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 59200 | 59200 | 0 | 0 | 983.31 | 16.177 | 25.913 | 28.836 | 34.638 | 53.194 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1184/59200
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 223 | 0.004 |
| mysql | 3307 | 0.056 |
| redis | 121707 | 2.056 |
| relation | 223 | 0.004 |

- Cold compute：59200（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 31376000 | 530.000 |
| merge_candidates | 54441504 | 919.620 |
| redis_commands | 355200 | 6.000 |
| redis_members | 60316512 | 1018.860 |
| redis_roundtrips | 59200 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 59200 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 59200 | 5.247 |
| counter | 223 | 4.973 |
| hydrate | 59200 | 10.381 |
| inbox | 59200 | 5.244 |
| merge_dedup | 59200 | 0.102 |
| relation | 223 | 5.184 |
| route | 59200 | 0.119 |
| total | 59200 | 15.975 |

## Redis 本轮边界增量

- Commands：440259；input：1127435275 bytes；output：7673061671 bytes
- Hits/Misses：31733167/4680；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：8206；safety epoch：3718 -> 3718

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.364 |
| client:loadtest | cpu_percent_total | 21.825 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 4.250 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.160 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.300 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 166.060 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 46.370 |
| docker:zg-zk | memory_percent | 1.060 |
| docker:zg-zk | pids | 83.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 360798.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.652 |
| process:counter | cpu_seconds_total | 223.375 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47087616.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 5120.078 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45772800.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 250.226 |
| process:knowpost | cpu_seconds_total | 13053.203 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 81219584.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.645 |
| process:relation | cpu_seconds_total | 37.031 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49106944.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 4.484 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38043648.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.548 |
| process:user-storage | cpu_seconds_total | 1810.922 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 48398336.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 236542593.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3718.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596569.000 |
| redis | keyspace_hits | 1748946297.000 |
| redis | keyspace_misses | 458805.000 |
| redis | net_input_bytes | 76183336506.000 |
| redis | net_output_bytes | 443945512476.000 |
| redis | ops_per_sec | 8206.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 30548.000 |
| redis | used_memory_bytes | 102392496.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
