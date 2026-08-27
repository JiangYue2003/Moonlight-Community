# Feed 压测报告：hybrid / gateway / sequential-page-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:56:08+08:00
- 采样时长：1m0.1793403s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 46400 | 46400 | 0 | 0 | 771.07 | 18.733 | 35.102 | 40.855 | 52.790 | 113.726 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：928/46400
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 219 | 0.005 |
| mysql | 4773 | 0.103 |
| redis | 97573 | 2.103 |
| relation | 219 | 0.005 |

- Cold compute：46400（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 24592000 | 530.000 |
| merge_candidates | 42670368 | 919.620 |
| redis_commands | 278400 | 6.000 |
| redis_members | 47275104 | 1018.860 |
| redis_roundtrips | 46400 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 46400 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 46400 | 6.512 |
| counter | 219 | 6.001 |
| hydrate | 46400 | 12.856 |
| inbox | 46400 | 6.509 |
| merge_dedup | 46400 | 0.100 |
| relation | 219 | 6.860 |
| route | 46400 | 0.189 |
| total | 46400 | 19.786 |

## Redis 本轮边界增量

- Commands：348481；input：884131089 bytes；output：6013818935 bytes
- Hits/Misses：24871501/5435；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：7179；safety epoch：3724 -> 3724

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.191 |
| client:loadtest | cpu_percent_total | 19.058 |
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
| docker:zg-canal | cpu_percent | 2.030 |
| docker:zg-canal | memory_percent | 4.260 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.060 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.860 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 144.840 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 45.280 |
| docker:zg-zk | memory_percent | 1.070 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 413079.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 15.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.652 |
| process:counter | cpu_seconds_total | 230.422 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46424064.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 77.436 |
| process:gateway | cpu_seconds_total | 5200.031 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51281920.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 215.494 |
| process:knowpost | cpu_seconds_total | 13939.750 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 80519168.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.321 |
| process:relation | cpu_seconds_total | 39.016 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48398336.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.578 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 37.164 |
| process:user-storage | cpu_seconds_total | 1845.312 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54005760.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 239573598.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3724.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 1965477716.000 |
| redis | keyspace_misses | 520883.000 |
| redis | net_input_bytes | 83883845393.000 |
| redis | net_output_bytes | 496304290960.000 |
| redis | ops_per_sec | 7179.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31007.000 |
| redis | used_memory_bytes | 102742360.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
