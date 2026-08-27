# Feed 压测报告：hybrid / rpc / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-formal-20260817a`
- 开始时间：2026-08-17T18:27:43+08:00
- 采样时长：1m0.0897744s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 260600 | 260600 | 0 | 0 | 4338.04 | 3.299 | 5.471 | 6.421 | 8.736 | 24.558 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：5212/260600
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 543 | 0.002 |
| redis | 266355 | 1.022 |
| relation | 240 | 0.001 |

- Cold compute：260600（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 5571628 | 21.380 |
| merge_candidates | 15130436 | 58.060 |
| redis_commands | 3674460 | 14.100 |
| redis_members | 19013376 | 72.960 |
| redis_roundtrips | 515988 | 1.980 |
| tie_members | 4586560 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 260600 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 5212 | 1.289 |
| counter | 240 | 1.601 |
| cursor_decode | 255388 | 0.010 |
| cursor_seek | 255388 | 2.331 |
| hydrate | 260600 | 1.132 |
| inbox | 5212 | 1.288 |
| merge_dedup | 260600 | 0.006 |
| relation | 240 | 2.360 |
| route | 260600 | 0.013 |
| total | 260600 | 3.475 |

## Redis 本轮边界增量

- Commands：3971497；input：599314496 bytes；output：1727062310 bytes
- Hits/Misses：9252721/787；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：79553；safety epoch：3705 -> 3705

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.844 |
| client:loadtest | cpu_percent_total | 61.497 |
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
| docker:zg-canal | cpu_percent | 1.610 |
| docker:zg-canal | memory_percent | 4.230 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.930 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.670 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 147.250 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 33.370 |
| docker:zg-zk | memory_percent | 1.030 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 328457.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.427 |
| process:counter | cpu_seconds_total | 185.016 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46444544.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4321.891 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51281920.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 254.786 |
| process:knowpost | cpu_seconds_total | 11363.625 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70909952.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 8.516 |
| process:relation | cpu_seconds_total | 31.219 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49688576.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.551 |
| process:search | cpu_seconds_total | 4.078 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38035456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.549 |
| process:user-storage | cpu_seconds_total | 1533.500 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 59932672.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 198300439.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3705.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596555.000 |
| redis | keyspace_hits | 1587458346.000 |
| redis | keyspace_misses | 413001.000 |
| redis | net_input_bytes | 67907898927.000 |
| redis | net_output_bytes | 409646577719.000 |
| redis | ops_per_sec | 79553.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 29302.000 |
| redis | used_memory_bytes | 102086000.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
