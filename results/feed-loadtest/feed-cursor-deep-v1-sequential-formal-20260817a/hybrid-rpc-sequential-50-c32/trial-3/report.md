# Feed 压测报告：hybrid / rpc / sequential-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-formal-20260817a`
- 开始时间：2026-08-17T18:33:58+08:00
- 采样时长：1m0.1918278s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 293750 | 293750 | 0 | 0 | 4881.87 | 6.162 | 8.894 | 10.418 | 13.507 | 26.857 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：5875/293750
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 241 | 0.001 |
| mysql | 1371 | 0.005 |
| redis | 300996 | 1.025 |
| relation | 241 | 0.001 |

- Cold compute：293750（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6280375 | 21.380 |
| merge_candidates | 17055125 | 58.060 |
| redis_commands | 4141875 | 14.100 |
| redis_members | 21432000 | 72.960 |
| redis_roundtrips | 581625 | 1.980 |
| tie_members | 5170000 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 293750 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 5875 | 2.224 |
| counter | 241 | 2.981 |
| cursor_decode | 287875 | 0.010 |
| cursor_seek | 287875 | 4.250 |
| hydrate | 293750 | 2.092 |
| inbox | 5875 | 2.224 |
| merge_dedup | 293750 | 0.006 |
| relation | 241 | 3.197 |
| route | 293750 | 0.021 |
| total | 293750 | 6.343 |

## Redis 本轮边界增量

- Commands：4473159；input：675483594 bytes；output：1946505997 bytes
- Hits/Misses：10427971/1792；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：80889；safety epoch：3710 -> 3710

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.366 |
| client:loadtest | cpu_percent_total | 69.855 |
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
| docker:zg-canal | cpu_percent | 2.240 |
| docker:zg-canal | memory_percent | 4.260 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 0.500 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.670 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 141.120 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 46.400 |
| docker:zg-zk | memory_percent | 1.040 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 339200.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 28.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.419 |
| process:counter | cpu_seconds_total | 196.125 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46489600.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4321.953 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44879872.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 282.641 |
| process:knowpost | cpu_seconds_total | 12105.422 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72036352.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.550 |
| process:relation | cpu_seconds_total | 32.703 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49393664.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 4.141 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38043648.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.550 |
| process:user-storage | cpu_seconds_total | 1533.844 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44261376.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 224791910.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3710.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596552.000 |
| redis | keyspace_hits | 1649188279.000 |
| redis | keyspace_misses | 425444.000 |
| redis | net_input_bytes | 71907725652.000 |
| redis | net_output_bytes | 421169421823.000 |
| redis | ops_per_sec | 80889.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 29677.000 |
| redis | used_memory_bytes | 103272304.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
