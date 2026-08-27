# Feed 压测报告：hybrid / gateway / page20-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:56:32+08:00
- 采样时长：1m0.0139418s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 55328 | 55328 | 0 | 0 | 921.98 | 16.068 | 25.304 | 29.970 | 38.443 | 69.022 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.004 |
| mysql | 1580 | 0.029 |
| redis | 112236 | 2.029 |
| relation | 240 | 0.004 |

- Cold compute：55328（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 23237760 | 420.000 |
| merge_candidates | 46530848 | 841.000 |
| redis_commands | 331968 | 6.000 |
| redis_members | 50957088 | 921.000 |
| redis_roundtrips | 55328 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 55328 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 55328 | 7.939 |
| counter | 240 | 6.799 |
| hydrate | 55328 | 8.096 |
| inbox | 55328 | 7.936 |
| merge_dedup | 55328 | 0.082 |
| relation | 240 | 7.508 |
| route | 55328 | 0.277 |
| total | 55328 | 16.505 |

## Redis 本轮边界增量

- Commands：405767；input：839842223 bytes；output：5953720612 bytes
- Hits/Misses：23575505/1691；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：8227；safety epoch：3681 -> 3681

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.411 |
| client:loadtest | cpu_percent_total | 22.573 |
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
| docker:zg-canal | cpu_percent | 2.050 |
| docker:zg-canal | memory_percent | 4.230 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 1.130 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.950 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 156.090 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 44.050 |
| docker:zg-zk | memory_percent | 1.010 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 238118.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.874 |
| process:counter | cpu_seconds_total | 127.844 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46366720.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 110.792 |
| process:gateway | cpu_seconds_total | 2155.656 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51388416.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 220.516 |
| process:knowpost | cpu_seconds_total | 8463.328 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 74604544.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.838 |
| process:relation | cpu_seconds_total | 20.859 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48775168.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.778 |
| process:search | cpu_seconds_total | 3.391 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38162432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 41.845 |
| process:user-storage | cpu_seconds_total | 782.406 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54784000.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 166976075.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3681.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597196.000 |
| redis | keyspace_hits | 1196574107.000 |
| redis | keyspace_misses | 266155.000 |
| redis | net_input_bytes | 52149888079.000 |
| redis | net_output_bytes | 321941987615.000 |
| redis | ops_per_sec | 8227.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27431.000 |
| redis | used_memory_bytes | 102897136.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
