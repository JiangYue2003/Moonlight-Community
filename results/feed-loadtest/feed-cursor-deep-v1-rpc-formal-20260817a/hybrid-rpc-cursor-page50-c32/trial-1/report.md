# Feed 压测报告：hybrid / rpc / cursor-page50-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:29:27+08:00
- 采样时长：1m0.0320252s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 463824 | 463824 | 0 | 0 | 7730.01 | 3.808 | 5.684 | 6.509 | 8.243 | 14.000 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：657.9572ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 286 | 0.001 |
| redis | 464110 | 1.001 |
| relation | 240 | 0.001 |

- Cold compute：463824（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9740304 | 21.000 |
| merge_candidates | 10204128 | 22.000 |
| redis_commands | 6029712 | 13.000 |
| redis_members | 11131776 | 24.000 |
| redis_roundtrips | 927648 | 2.000 |
| tie_members | 927648 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 463824 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 1.533 |
| cursor_decode | 463824 | 0.010 |
| cursor_seek | 463824 | 2.560 |
| hydrate | 463824 | 1.285 |
| merge_dedup | 463824 | 0.005 |
| relation | 240 | 2.516 |
| route | 463824 | 0.011 |
| total | 463824 | 3.876 |

## Redis 本轮边界增量

- Commands：6530214；input：1010240409 bytes；output：2082917601 bytes
- Hits/Misses：15777261/286；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：117228；safety epoch：3660 -> 3660

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.734 |
| client:loadtest | cpu_percent_total | 139.743 |
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
| docker:zg-canal | cpu_percent | 0.170 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.710 |
| docker:zg-es | memory_percent | 12.330 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.130 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 133.920 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 50.670 |
| docker:zg-zk | memory_percent | 1.230 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 215633.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 16.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.418 |
| process:counter | cpu_seconds_total | 85.359 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45846528.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.406 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43900928.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 343.301 |
| process:knowpost | cpu_seconds_total | 6193.453 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 73584640.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.548 |
| process:relation | cpu_seconds_total | 11.656 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48312320.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.548 |
| process:search | cpu_seconds_total | 2.625 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38129664.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.095 |
| process:user-storage | cpu_seconds_total | 98.547 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43986944.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 131537906.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3660.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597779.000 |
| redis | keyspace_hits | 1000485500.000 |
| redis | keyspace_misses | 243713.000 |
| redis | net_input_bytes | 43172998012.000 |
| redis | net_output_bytes | 262263248515.000 |
| redis | ops_per_sec | 117228.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 25806.000 |
| redis | used_memory_bytes | 103367760.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
