# Feed 压测报告：hybrid / gateway / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-formal-20260817a`
- 开始时间：2026-08-17T18:35:14+08:00
- 采样时长：1m0.1777462s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 113450 | 113450 | 0 | 0 | 1885.54 | 7.969 | 12.834 | 14.753 | 19.894 | 40.926 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：2269/113450
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 237 | 0.002 |
| mysql | 1143 | 0.010 |
| redis | 116862 | 1.030 |
| relation | 237 | 0.002 |

- Cold compute：113450（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2425561 | 21.380 |
| merge_candidates | 6586907 | 58.060 |
| redis_commands | 1599645 | 14.100 |
| redis_members | 8277312 | 72.960 |
| redis_roundtrips | 224631 | 1.980 |
| tie_members | 1996720 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 113450 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2269 | 2.721 |
| counter | 237 | 3.260 |
| cursor_decode | 111181 | 0.011 |
| cursor_seek | 111181 | 4.746 |
| hydrate | 113450 | 2.514 |
| inbox | 2269 | 2.718 |
| merge_dedup | 113450 | 0.007 |
| relation | 237 | 4.118 |
| route | 113450 | 0.046 |
| total | 113450 | 7.289 |

## Redis 本轮边界增量

- Commands：1746578；input：262172359 bytes；output：752098078 bytes
- Hits/Misses：4030845/1566；run hit rate：99.96%
- Evicted/Rejected：0/0；ops/s max：39737；safety epoch：3711 -> 3711

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.343 |
| client:loadtest | cpu_percent_total | 53.487 |
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
| docker:zg-canal | memory_percent | 4.240 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.830 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.090 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 25.000 |
| docker:zg-kafka | cpu_percent | 148.740 |
| docker:zg-kafka | memory_percent | 7.150 |
| docker:zg-kafka | pids | 98.000 |
| docker:zg-zk | cpu_percent | 52.060 |
| docker:zg-zk | memory_percent | 1.040 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 341088.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 13.169 |
| process:counter | cpu_seconds_total | 199.844 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45674496.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 217.997 |
| process:gateway | cpu_seconds_total | 4455.000 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51380224.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 185.820 |
| process:knowpost | cpu_seconds_total | 12216.797 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70729728.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.871 |
| process:relation | cpu_seconds_total | 33.203 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48766976.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 4.188 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38043648.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 79.018 |
| process:user-storage | cpu_seconds_total | 1581.641 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 57536512.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 227006806.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3711.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596552.000 |
| redis | keyspace_hits | 1654298556.000 |
| redis | keyspace_misses | 427460.000 |
| redis | net_input_bytes | 72240204475.000 |
| redis | net_output_bytes | 422123027617.000 |
| redis | ops_per_sec | 39737.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 29753.000 |
| redis | used_memory_bytes | 102310536.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
