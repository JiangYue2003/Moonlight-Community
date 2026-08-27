# Feed 压测报告：hybrid / gateway / cursor-page50-c32

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:37:59+08:00
- 采样时长：10.0118539s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 18000 | 18000 | 0 | 0 | 1798.07 | 17.302 | 24.623 | 27.291 | 35.963 | 59.180 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：1.0706499s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 124 | 0.007 |
| redis | 18124 | 1.007 |
| relation | 40 | 0.002 |

- Cold compute：18000（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 378000 | 21.000 |
| merge_candidates | 396000 | 22.000 |
| redis_commands | 234000 | 13.000 |
| redis_members | 432000 | 24.000 |
| redis_roundtrips | 36000 | 2.000 |
| tie_members | 36000 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 18000 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 6.452 |
| cursor_decode | 18000 | 0.012 |
| cursor_seek | 18000 | 10.581 |
| hydrate | 18000 | 5.451 |
| merge_dedup | 18000 | 0.005 |
| relation | 40 | 8.287 |
| route | 18000 | 0.127 |
| total | 18000 | 16.183 |

## Redis 本轮边界增量

- Commands：255430；input：39333299 bytes；output：80873966 bytes
- Hits/Misses：613126/136；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：27526；safety epoch：3620 -> 3620

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.536 |
| client:loadtest | cpu_percent_total | 72.570 |
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
| docker:zg-canal | cpu_percent | 1.830 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.680 |
| docker:zg-es | memory_percent | 12.080 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.360 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 71.620 |
| docker:zg-kafka | memory_percent | 7.270 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 45.970 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 96979.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 13.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.970 |
| process:counter | cpu_seconds_total | 11.891 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45436928.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 266.625 |
| process:gateway | cpu_seconds_total | 258.016 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 53149696.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 206.029 |
| process:knowpost | cpu_seconds_total | 300.766 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 68378624.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.872 |
| process:relation | cpu_seconds_total | 1.594 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 45125632.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.778 |
| process:search | cpu_seconds_total | 0.547 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42201088.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 98.349 |
| process:user-storage | cpu_seconds_total | 94.047 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 53473280.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 20682144.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3620.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 598278.000 |
| redis | keyspace_hits | 119144421.000 |
| redis | keyspace_misses | 85570.000 |
| redis | net_input_bytes | 5542555177.000 |
| redis | net_output_bytes | 30901541459.000 |
| redis | ops_per_sec | 27526.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22668.000 |
| redis | used_memory_bytes | 103264880.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
