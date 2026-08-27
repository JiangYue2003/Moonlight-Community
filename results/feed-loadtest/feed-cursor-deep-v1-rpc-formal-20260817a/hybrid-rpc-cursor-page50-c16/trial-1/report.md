# Feed 压测报告：hybrid / rpc / cursor-page50-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:25:35+08:00
- 采样时长：1m0.02688s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 455330 | 455330 | 0 | 0 | 7588.64 | 2.087 | 3.133 | 3.391 | 4.365 | 12.828 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：748.089ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 189 | 0.000 |
| redis | 455519 | 1.000 |
| relation | 240 | 0.001 |

- Cold compute：455330（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9561930 | 21.000 |
| merge_candidates | 10017260 | 22.000 |
| redis_commands | 5919290 | 13.000 |
| redis_members | 10927920 | 24.000 |
| redis_roundtrips | 910660 | 2.000 |
| tie_members | 910660 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 455330 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 0.962 |
| cursor_decode | 455330 | 0.009 |
| cursor_seek | 455330 | 1.211 |
| hydrate | 455330 | 0.629 |
| merge_dedup | 455330 | 0.004 |
| relation | 240 | 1.623 |
| route | 455330 | 0.007 |
| total | 455330 | 1.866 |

## Redis 本轮边界增量

- Commands：6411229；input：991766853 bytes；output：2044799026 bytes
- Hits/Misses：15488536/215；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：116993；safety epoch：3657 -> 3657

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.681 |
| client:loadtest | cpu_percent_total | 138.896 |
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
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.190 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 157.000 |
| docker:zg-etcd | cpu_percent | 4.680 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 109.690 |
| docker:zg-kafka | memory_percent | 7.460 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 44.340 |
| docker:zg-zk | memory_percent | 0.990 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 213340.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.191 |
| process:counter | cpu_seconds_total | 78.844 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46600192.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.406 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43864064.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 325.190 |
| process:knowpost | cpu_seconds_total | 5577.859 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71073792.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.081 |
| process:relation | cpu_seconds_total | 10.797 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47230976.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.275 |
| process:search | cpu_seconds_total | 2.516 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38117376.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.548 |
| process:user-storage | cpu_seconds_total | 98.031 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44077056.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 108530974.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3657.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597795.000 |
| redis | keyspace_hits | 944940445.000 |
| redis | keyspace_misses | 240274.000 |
| redis | net_input_bytes | 39614892031.000 |
| redis | net_output_bytes | 254918594608.000 |
| redis | ops_per_sec | 116993.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 25575.000 |
| redis | used_memory_bytes | 102576216.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
