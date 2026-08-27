# Feed 压测报告：hybrid / rpc / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:13:00+08:00
- 采样时长：1m0.0220436s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 288949 | 288949 | 0 | 0 | 4815.56 | 3.156 | 4.765 | 5.406 | 6.800 | 14.080 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：279.8451ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 171 | 0.001 |
| redis | 289120 | 1.001 |
| relation | 240 | 0.001 |

- Cold compute：288949（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6067929 | 21.000 |
| merge_candidates | 18492736 | 64.000 |
| redis_commands | 4045286 | 14.000 |
| redis_members | 19359583 | 67.000 |
| redis_roundtrips | 577898 | 2.000 |
| tie_members | 866847 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 288949 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 1.207 |
| cursor_decode | 288949 | 0.010 |
| cursor_seek | 288949 | 2.069 |
| hydrate | 288949 | 1.006 |
| merge_dedup | 288949 | 0.006 |
| relation | 240 | 1.931 |
| route | 288949 | 0.011 |
| total | 288949 | 3.105 |

## Redis 本轮边界增量

- Commands：4370660；input：659098054 bytes；output：1827555605 bytes
- Hits/Misses：10120560/185；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：79969；safety epoch：3647 -> 3647

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.752 |
| client:loadtest | cpu_percent_total | 76.040 |
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
| docker:zg-canal | cpu_percent | 2.120 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.440 |
| docker:zg-es | memory_percent | 12.220 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.490 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 145.290 |
| docker:zg-kafka | memory_percent | 7.120 |
| docker:zg-kafka | pids | 99.000 |
| docker:zg-zk | cpu_percent | 45.640 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 145873.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.970 |
| process:counter | cpu_seconds_total | 62.484 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45961216.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.344 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44060672.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 244.619 |
| process:knowpost | cpu_seconds_total | 4069.109 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 69836800.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.548 |
| process:relation | cpu_seconds_total | 8.047 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47497216.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 2.203 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38113280.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 96.859 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43859968.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 83528581.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3647.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597207.000 |
| redis | keyspace_hits | 637142790.000 |
| redis | keyspace_misses | 138866.000 |
| redis | net_input_bytes | 27138692011.000 |
| redis | net_output_bytes | 189199245629.000 |
| redis | ops_per_sec | 79969.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 24819.000 |
| redis | used_memory_bytes | 102736536.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
