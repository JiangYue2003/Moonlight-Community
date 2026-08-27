# Feed 压测报告：hybrid / rpc / sequential-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-formal-20260817a`
- 开始时间：2026-08-17T18:32:43+08:00
- 采样时长：1m0.1270864s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 295850 | 295850 | 0 | 0 | 4922.10 | 6.119 | 8.778 | 10.280 | 13.438 | 36.572 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：5917/295850
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 241 | 0.001 |
| mysql | 1159 | 0.004 |
| redis | 302926 | 1.024 |
| relation | 241 | 0.001 |

- Cold compute：295850（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6325273 | 21.380 |
| merge_candidates | 17177051 | 58.060 |
| redis_commands | 4171485 | 14.100 |
| redis_members | 21585216 | 72.960 |
| redis_roundtrips | 585783 | 1.980 |
| tie_members | 5206960 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 295850 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 5917 | 2.207 |
| counter | 241 | 3.012 |
| cursor_decode | 289933 | 0.010 |
| cursor_seek | 289933 | 4.216 |
| hydrate | 295850 | 2.076 |
| inbox | 5917 | 2.205 |
| merge_dedup | 295850 | 0.007 |
| relation | 241 | 3.215 |
| route | 295850 | 0.020 |
| total | 295850 | 6.293 |

## Redis 本轮边界增量

- Commands：4504908；input：680288660 bytes；output：1960426857 bytes
- Hits/Misses：10502523/1741；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：80286；safety epoch：3709 -> 3709

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.210 |
| client:loadtest | cpu_percent_total | 67.357 |
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
| docker:zg-canal | memory_percent | 4.240 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.610 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.840 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 149.190 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 42.730 |
| docker:zg-zk | memory_percent | 1.260 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 336968.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.744 |
| process:counter | cpu_seconds_total | 194.047 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46313472.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.783 |
| process:gateway | cpu_seconds_total | 4321.938 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 46129152.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 256.135 |
| process:knowpost | cpu_seconds_total | 11954.250 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71462912.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.097 |
| process:relation | cpu_seconds_total | 32.531 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49233920.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.109 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38035456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.323 |
| process:user-storage | cpu_seconds_total | 1533.719 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44883968.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 219523680.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3709.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596552.000 |
| redis | keyspace_hits | 1636912684.000 |
| redis | keyspace_misses | 423019.000 |
| redis | net_input_bytes | 71112336192.000 |
| redis | net_output_bytes | 418878023652.000 |
| redis | ops_per_sec | 80286.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 29602.000 |
| redis | used_memory_bytes | 103313424.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
