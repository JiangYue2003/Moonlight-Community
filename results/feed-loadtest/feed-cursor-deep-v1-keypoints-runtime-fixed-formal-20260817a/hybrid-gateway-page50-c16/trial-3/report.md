# Feed 压测报告：hybrid / gateway / page50-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:43:40+08:00
- 采样时长：1m0.0279712s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 29036 | 29036 | 0 | 0 | 483.72 | 31.321 | 45.010 | 49.203 | 57.086 | 82.085 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.008 |
| mysql | 5639 | 0.194 |
| redis | 63711 | 2.194 |
| relation | 240 | 0.008 |

- Cold compute：29036（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 29616720 | 1020.000 |
| merge_candidates | 43524964 | 1499.000 |
| redis_commands | 174216 | 6.000 |
| redis_members | 43554000 | 1500.000 |
| redis_roundtrips | 29036 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 29036 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 29036 | 7.014 |
| counter | 240 | 5.368 |
| hydrate | 29036 | 24.278 |
| inbox | 29036 | 7.011 |
| merge_dedup | 29036 | 0.170 |
| relation | 240 | 5.819 |
| route | 29036 | 0.403 |
| total | 29036 | 32.058 |

## Redis 本轮边界增量

- Commands：230431；input：1052734845 bytes；output：6671972794 bytes
- Hits/Misses：29790543/7843；run hit rate：99.97%
- Evicted/Rejected：0/0；ops/s max：4287；safety epoch：3760 -> 3760

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.800 |
| client:loadtest | cpu_percent_total | 12.807 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.660 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.680 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 147.380 |
| docker:zg-kafka | memory_percent | 7.630 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 42.260 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 529271.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.877 |
| process:counter | cpu_seconds_total | 297.000 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46931968.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 57.339 |
| process:gateway | cpu_seconds_total | 6069.828 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50909184.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 238.575 |
| process:knowpost | cpu_seconds_total | 17984.484 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 84611072.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.873 |
| process:relation | cpu_seconds_total | 48.938 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49926144.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 5.922 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38055936.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 24.795 |
| process:user-storage | cpu_seconds_total | 2172.062 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 51499008.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 294038894.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3760.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2722200081.000 |
| redis | keyspace_misses | 715319.000 |
| redis | net_input_bytes | 113906688926.000 |
| redis | net_output_bytes | 670824087623.000 |
| redis | ops_per_sec | 4287.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 33859.000 |
| redis | used_memory_bytes | 103215168.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
