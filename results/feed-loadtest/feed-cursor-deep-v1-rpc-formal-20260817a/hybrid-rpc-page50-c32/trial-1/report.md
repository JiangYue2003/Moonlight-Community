# Feed 压测报告：hybrid / rpc / page50-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:21:49+08:00
- 采样时长：1m0.0240782s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 35025 | 35025 | 0 | 0 | 583.54 | 53.478 | 67.036 | 71.412 | 81.794 | 107.950 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.007 |
| mysql | 10493 | 0.300 |
| redis | 80543 | 2.300 |
| relation | 240 | 0.007 |

- Cold compute：35025（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 35725500 | 1020.000 |
| merge_candidates | 52502475 | 1499.000 |
| redis_commands | 210150 | 6.000 |
| redis_members | 52537500 | 1500.000 |
| redis_roundtrips | 35025 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 35025 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 35025 | 11.964 |
| counter | 240 | 8.218 |
| hydrate | 35025 | 41.460 |
| inbox | 35025 | 11.960 |
| merge_dedup | 35025 | 0.180 |
| relation | 240 | 10.384 |
| route | 35025 | 0.635 |
| total | 35025 | 54.463 |

## Redis 本轮边界增量

- Commands：275016；input：1270704862 bytes；output：8046993356 bytes
- Hits/Misses：35927784/15309；run hit rate：99.96%
- Evicted/Rejected：0/0；ops/s max：4895；safety epoch：3654 -> 3654

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.768 |
| client:loadtest | cpu_percent_total | 12.287 |
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
| docker:zg-es | cpu_percent | 1.460 |
| docker:zg-es | memory_percent | 12.240 |
| docker:zg-es | pids | 154.000 |
| docker:zg-etcd | cpu_percent | 4.360 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 133.800 |
| docker:zg-kafka | memory_percent | 7.560 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 44.650 |
| docker:zg-zk | memory_percent | 0.980 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 182908.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 35.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.105 |
| process:counter | cpu_seconds_total | 74.328 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46493696.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.406 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43884544.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 275.969 |
| process:knowpost | cpu_seconds_total | 5076.719 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 103215104.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.325 |
| process:relation | cpu_seconds_total | 9.828 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48472064.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 2.391 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38109184.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.550 |
| process:user-storage | cpu_seconds_total | 97.688 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43941888.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 100306777.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3654.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597798.000 |
| redis | keyspace_hits | 841520233.000 |
| redis | keyspace_misses | 198808.000 |
| redis | net_input_bytes | 35431788791.000 |
| redis | net_output_bytes | 233428850520.000 |
| redis | ops_per_sec | 4895.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 25348.000 |
| redis | used_memory_bytes | 105717352.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
