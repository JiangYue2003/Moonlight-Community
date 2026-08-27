# Feed 压测报告：hybrid / rpc / sequential-page-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:52:18+08:00
- 采样时长：1m0.86288s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 60800 | 60800 | 0 | 0 | 999.06 | 32.972 | 50.585 | 56.022 | 67.603 | 117.644 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1216/60800
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 233 | 0.004 |
| mysql | 8478 | 0.139 |
| redis | 130078 | 2.139 |
| relation | 233 | 0.004 |

- Cold compute：60800（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 32224000 | 530.000 |
| merge_candidates | 55912896 | 919.620 |
| redis_commands | 364800 | 6.000 |
| redis_members | 61946688 | 1018.860 |
| redis_roundtrips | 60800 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 60800 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 60800 | 9.974 |
| counter | 233 | 8.534 |
| hydrate | 60800 | 21.059 |
| inbox | 60800 | 9.971 |
| merge_dedup | 60800 | 0.106 |
| relation | 233 | 9.527 |
| route | 60800 | 0.314 |
| total | 60800 | 31.598 |

## Redis 本轮边界增量

- Commands：449969；input：1158652698 bytes；output：7879411516 bytes
- Hits/Misses：32585426/10498；run hit rate：99.97%
- Evicted/Rejected：0/0；ops/s max：8297；safety epoch：3721 -> 3721

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.272 |
| client:loadtest | cpu_percent_total | 20.358 |
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
| docker:zg-es | cpu_percent | 2.680 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.660 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 163.760 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 45.070 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 389192.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 26.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.642 |
| process:counter | cpu_seconds_total | 226.766 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46366720.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.547 |
| process:gateway | cpu_seconds_total | 5120.125 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44810240.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 283.701 |
| process:knowpost | cpu_seconds_total | 13516.609 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 93777920.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.326 |
| process:relation | cpu_seconds_total | 38.016 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49700864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.547 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.534 |
| process:user-storage | cpu_seconds_total | 1811.234 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44171264.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 238141023.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3721.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 1863480066.000 |
| redis | keyspace_misses | 493337.000 |
| redis | net_input_bytes | 80256578435.000 |
| redis | net_output_bytes | 471640808994.000 |
| redis | ops_per_sec | 8297.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 30777.000 |
| redis | used_memory_bytes | 104426752.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
