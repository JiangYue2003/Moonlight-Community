# Feed 压测报告：hybrid / rpc / page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:19:38+08:00
- 采样时长：1m0.0151368s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 70742 | 70742 | 0 | 0 | 1178.84 | 11.894 | 20.784 | 22.672 | 26.804 | 51.789 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 1601 | 0.023 |
| redis | 143085 | 2.023 |
| relation | 240 | 0.003 |

- Cold compute：70742（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 29711640 | 420.000 |
| merge_candidates | 59494022 | 841.000 |
| redis_commands | 424452 | 6.000 |
| redis_members | 65153382 | 921.000 |
| redis_roundtrips | 70742 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 70742 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 70742 | 6.378 |
| counter | 240 | 5.570 |
| hydrate | 70742 | 6.582 |
| inbox | 70742 | 6.376 |
| merge_dedup | 70742 | 0.081 |
| relation | 240 | 5.896 |
| route | 70742 | 0.173 |
| total | 70742 | 13.325 |

## Redis 本轮边界增量

- Commands：517198；input：1073862004 bytes；output：7612229432 bytes
- Hits/Misses：30140666/2901；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：9214；safety epoch：3741 -> 3741

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.726 |
| client:loadtest | cpu_percent_total | 27.623 |
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
| docker:zg-canal | cpu_percent | 0.140 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.270 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.810 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 183.170 |
| docker:zg-kafka | memory_percent | 7.160 |
| docker:zg-kafka | pids | 98.000 |
| docker:zg-zk | cpu_percent | 47.120 |
| docker:zg-zk | memory_percent | 1.310 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 469316.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 25.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.873 |
| process:counter | cpu_seconds_total | 258.391 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47038464.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 5478.438 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 43773952.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 250.429 |
| process:knowpost | cpu_seconds_total | 15293.328 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 80076800.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.869 |
| process:relation | cpu_seconds_total | 43.094 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49405952.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 5.188 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38027264.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.003 |
| process:user-storage | cpu_seconds_total | 1947.531 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 45998080.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 246297366.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3741.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 595980.000 |
| redis | keyspace_hits | 2256290827.000 |
| redis | keyspace_misses | 607348.000 |
| redis | net_input_bytes | 94377356852.000 |
| redis | net_output_bytes | 569878671082.000 |
| redis | ops_per_sec | 9214.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 32417.000 |
| redis | used_memory_bytes | 101946624.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
