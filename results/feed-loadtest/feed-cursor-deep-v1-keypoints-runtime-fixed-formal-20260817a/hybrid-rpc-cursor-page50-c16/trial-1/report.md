# Feed 压测报告：hybrid / rpc / cursor-page50-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:29:44+08:00
- 采样时长：1m0.0264744s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 454871 | 454871 | 0 | 0 | 7580.98 | 2.084 | 3.125 | 3.359 | 4.350 | 11.763 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：782.2623ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 135 | 0.000 |
| redis | 455006 | 1.000 |
| relation | 240 | 0.001 |

- Cold compute：454871（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9552291 | 21.000 |
| merge_candidates | 10007162 | 22.000 |
| redis_commands | 5913323 | 13.000 |
| redis_members | 10916904 | 24.000 |
| redis_roundtrips | 909742 | 2.000 |
| tie_members | 909742 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 454871 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 0.965 |
| cursor_decode | 454871 | 0.010 |
| cursor_seek | 454871 | 1.213 |
| hydrate | 454871 | 0.630 |
| merge_dedup | 454871 | 0.004 |
| relation | 240 | 1.540 |
| route | 454871 | 0.006 |
| total | 454871 | 1.869 |

## Redis 本轮边界增量

- Commands：6404354；input：990727601 bytes；output：2042740038 bytes
- Hits/Misses：15472752/393；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：120914；safety epoch：3749 -> 3749

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.836 |
| client:loadtest | cpu_percent_total | 141.370 |
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
| docker:zg-canal | cpu_percent | 2.490 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.670 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.310 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 133.710 |
| docker:zg-kafka | memory_percent | 7.620 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 50.000 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 496125.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.974 |
| process:counter | cpu_seconds_total | 273.031 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46276608.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.773 |
| process:gateway | cpu_seconds_total | 5478.500 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43823104.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 333.566 |
| process:knowpost | cpu_seconds_total | 16486.875 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72159232.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.096 |
| process:relation | cpu_seconds_total | 45.266 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49356800.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 5.500 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38039552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.549 |
| process:user-storage | cpu_seconds_total | 1948.094 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 46084096.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 271036892.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3749.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596565.000 |
| redis | keyspace_hits | 2474392074.000 |
| redis | keyspace_misses | 656121.000 |
| redis | net_input_bytes | 103695735459.000 |
| redis | net_output_bytes | 616499615047.000 |
| redis | ops_per_sec | 120914.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 33023.000 |
| redis | used_memory_bytes | 101841832.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
