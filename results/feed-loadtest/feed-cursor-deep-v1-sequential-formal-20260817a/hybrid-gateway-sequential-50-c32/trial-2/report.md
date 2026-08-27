# Feed 压测报告：hybrid / gateway / sequential-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-formal-20260817a`
- 开始时间：2026-08-17T18:40:17+08:00
- 采样时长：1m0.4511004s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 82050 | 82050 | 0 | 0 | 1357.44 | 22.678 | 33.660 | 38.084 | 49.551 | 98.023 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1641/82050
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 237 | 0.003 |
| mysql | 1566 | 0.019 |
| redis | 85257 | 1.039 |
| relation | 237 | 0.003 |

- Cold compute：82050（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1754229 | 21.380 |
| merge_candidates | 4763823 | 58.060 |
| redis_commands | 1156905 | 14.100 |
| redis_members | 5986368 | 72.960 |
| redis_roundtrips | 162459 | 1.980 |
| tie_members | 1444080 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 82050 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1641 | 7.311 |
| counter | 237 | 8.652 |
| cursor_decode | 80409 | 0.011 |
| cursor_seek | 80409 | 14.393 |
| hydrate | 82050 | 7.407 |
| inbox | 1641 | 7.309 |
| merge_dedup | 82050 | 0.007 |
| relation | 237 | 8.958 |
| route | 82050 | 0.195 |
| total | 82050 | 21.875 |

## Redis 本轮边界增量

- Commands：1257134；input：189267612 bytes；output：543678781 bytes
- Hits/Misses：2916186/2219；run hit rate：99.92%
- Evicted/Rejected：0/0；ops/s max：25133；safety epoch：3715 -> 3715

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.119 |
| client:loadtest | cpu_percent_total | 49.911 |
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
| docker:zg-canal | cpu_percent | 1.540 |
| docker:zg-canal | memory_percent | 4.240 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.320 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.830 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 144.160 |
| docker:zg-kafka | memory_percent | 7.610 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 47.540 |
| docker:zg-zk | memory_percent | 1.050 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 349288.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 5.578 |
| process:counter | cpu_seconds_total | 211.594 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46600192.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 208.201 |
| process:gateway | cpu_seconds_total | 4986.312 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52486144.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 181.679 |
| process:knowpost | cpu_seconds_total | 12641.891 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72097792.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.873 |
| process:relation | cpu_seconds_total | 35.406 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49184768.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 4.266 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38039552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 77.460 |
| process:user-storage | cpu_seconds_total | 1765.172 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55554048.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 233659335.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3715.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596551.000 |
| redis | keyspace_hits | 1669656176.000 |
| redis | keyspace_misses | 436514.000 |
| redis | net_input_bytes | 73239037186.000 |
| redis | net_output_bytes | 424987476800.000 |
| redis | ops_per_sec | 25133.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 30057.000 |
| redis | used_memory_bytes | 103990808.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
