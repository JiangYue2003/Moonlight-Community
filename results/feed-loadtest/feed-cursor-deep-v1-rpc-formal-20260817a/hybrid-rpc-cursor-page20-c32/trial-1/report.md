# Feed 压测报告：hybrid / rpc / cursor-page20-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:14:16+08:00
- 采样时长：1m0.0229831s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 288681 | 288681 | 0 | 0 | 4811.06 | 6.288 | 8.779 | 10.495 | 13.297 | 26.076 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：331.8132ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 240 | 0.001 |
| redis | 288921 | 1.001 |
| relation | 240 | 0.001 |

- Cold compute：288681（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6062301 | 21.000 |
| merge_candidates | 18475584 | 64.000 |
| redis_commands | 4041534 | 14.000 |
| redis_members | 19341627 | 67.000 |
| redis_roundtrips | 577362 | 2.000 |
| tie_members | 866043 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 288681 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 2.956 |
| cursor_decode | 288681 | 0.010 |
| cursor_seek | 288681 | 4.274 |
| hydrate | 288681 | 2.122 |
| merge_dedup | 288681 | 0.007 |
| relation | 240 | 3.200 |
| route | 288681 | 0.023 |
| total | 288681 | 6.439 |

## Redis 本轮边界增量

- Commands：4366622；input：658505038 bytes；output：1825841912 bytes
- Hits/Misses：10111070/294；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：77171；safety epoch：3648 -> 3648

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.694 |
| client:loadtest | cpu_percent_total | 75.101 |
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
| docker:zg-canal | cpu_percent | 1.990 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.440 |
| docker:zg-es | memory_percent | 12.220 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.910 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 154.440 |
| docker:zg-kafka | memory_percent | 7.110 |
| docker:zg-kafka | pids | 98.000 |
| docker:zg-zk | cpu_percent | 43.870 |
| docker:zg-zk | memory_percent | 1.100 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 146619.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.194 |
| process:counter | cpu_seconds_total | 64.547 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46088192.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.344 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44064768.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 240.982 |
| process:knowpost | cpu_seconds_total | 4211.344 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71372800.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.829 |
| process:relation | cpu_seconds_total | 8.328 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47636480.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.203 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38113280.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.550 |
| process:user-storage | cpu_seconds_total | 96.906 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43790336.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 88698680.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3648.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597198.000 |
| redis | keyspace_hits | 649106251.000 |
| redis | keyspace_misses | 139531.000 |
| redis | net_input_bytes | 27918096378.000 |
| redis | net_output_bytes | 191361817891.000 |
| redis | ops_per_sec | 77171.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 24895.000 |
| redis | used_memory_bytes | 104284800.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
