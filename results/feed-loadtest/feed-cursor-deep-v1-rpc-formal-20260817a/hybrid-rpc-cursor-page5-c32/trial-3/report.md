# Feed 压测报告：hybrid / rpc / cursor-page5-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:01:43+08:00
- 采样时长：1m0.0170387s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 166335 | 166335 | 0 | 0 | 2771.99 | 10.909 | 14.845 | 17.843 | 22.870 | 47.740 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：68.6396ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 266 | 0.002 |
| redis | 166601 | 1.002 |
| relation | 240 | 0.001 |

- Cold compute：166335（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 3493035 | 21.000 |
| merge_candidates | 24950250 | 150.000 |
| redis_commands | 2827695 | 17.000 |
| redis_members | 39255060 | 236.000 |
| redis_roundtrips | 332670 | 2.000 |
| tie_members | 20958210 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 166335 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 4.503 |
| cursor_decode | 166335 | 0.010 |
| cursor_seek | 166335 | 7.553 |
| hydrate | 166335 | 3.702 |
| merge_dedup | 166335 | 0.008 |
| relation | 240 | 5.046 |
| route | 166335 | 0.059 |
| total | 166335 | 11.333 |

## Redis 本轮边界增量

- Commands：3018859；input：429490899 bytes；output：2235621689 bytes
- Hits/Misses：6327969/266；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：54656；safety epoch：3638 -> 3638

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.926 |
| client:loadtest | cpu_percent_total | 46.810 |
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
| docker:zg-canal | cpu_percent | 0.870 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.610 |
| docker:zg-es | memory_percent | 12.220 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.650 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 133.600 |
| docker:zg-kafka | memory_percent | 7.540 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 42.800 |
| docker:zg-zk | memory_percent | 1.190 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 116744.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.652 |
| process:counter | cpu_seconds_total | 49.047 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46194688.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.234 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44687360.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 167.387 |
| process:knowpost | cpu_seconds_total | 2772.141 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70242304.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.326 |
| process:relation | cpu_seconds_total | 5.750 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47607808.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.788 |
| process:search | cpu_seconds_total | 1.906 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38346752.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 96.344 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44068864.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 64555313.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3638.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597220.000 |
| redis | keyspace_hits | 389861582.000 |
| redis | keyspace_misses | 102152.000 |
| redis | net_input_bytes | 17277620827.000 |
| redis | net_output_bytes | 129291319462.000 |
| redis | ops_per_sec | 54656.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 24142.000 |
| redis | used_memory_bytes | 106007672.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
