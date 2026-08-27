# Feed 压测报告：hybrid / rpc / sequential-page-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:51:01+08:00
- 采样时长：1m1.0317634s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 60750 | 60750 | 0 | 0 | 995.44 | 32.842 | 50.654 | 56.482 | 68.448 | 118.172 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1215/60750
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 235 | 0.004 |
| mysql | 8051 | 0.133 |
| redis | 129551 | 2.133 |
| relation | 235 | 0.004 |

- Cold compute：60750（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 32197500 | 530.000 |
| merge_candidates | 55866915 | 919.620 |
| redis_commands | 364500 | 6.000 |
| redis_members | 61895745 | 1018.860 |
| redis_roundtrips | 60750 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 60750 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 60750 | 10.103 |
| counter | 235 | 8.574 |
| hydrate | 60750 | 21.077 |
| inbox | 60750 | 10.100 |
| merge_dedup | 60750 | 0.105 |
| relation | 235 | 9.581 |
| route | 60750 | 0.310 |
| total | 60750 | 31.739 |

## Redis 本轮边界增量

- Commands：449391；input：1157622083 bytes；output：7872993847 bytes
- Hits/Misses：32559062/10115；run hit rate：99.97%
- Evicted/Rejected：0/0；ops/s max：8320；safety epoch：3720 -> 3720

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.323 |
| client:loadtest | cpu_percent_total | 21.172 |
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
| docker:zg-canal | cpu_percent | 0.150 |
| docker:zg-canal | memory_percent | 4.250 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.740 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.670 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 21.000 |
| docker:zg-kafka | cpu_percent | 115.800 |
| docker:zg-kafka | memory_percent | 7.470 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 41.510 |
| docker:zg-zk | memory_percent | 1.060 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 377895.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.876 |
| process:counter | cpu_seconds_total | 225.609 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46358528.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 5120.078 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44810240.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 294.876 |
| process:knowpost | cpu_seconds_total | 13359.422 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 94003200.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.095 |
| process:relation | cpu_seconds_total | 37.734 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49057792.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 4.531 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38039552.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.549 |
| process:user-storage | cpu_seconds_total | 1811.109 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 48279552.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 237603265.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3720.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 1824893110.000 |
| redis | keyspace_misses | 479721.000 |
| redis | net_input_bytes | 78883960373.000 |
| redis | net_output_bytes | 462310043721.000 |
| redis | ops_per_sec | 8320.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 30701.000 |
| redis | used_memory_bytes | 104413632.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
