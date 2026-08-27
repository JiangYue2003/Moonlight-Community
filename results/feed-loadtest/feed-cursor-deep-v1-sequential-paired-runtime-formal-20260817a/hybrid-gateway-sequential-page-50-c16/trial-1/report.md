# Feed 压测报告：hybrid / gateway / sequential-page-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 开始时间：2026-08-17T19:58:12+08:00
- 采样时长：1m0.5200694s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 56000 | 56000 | 0 | 0 | 925.37 | 16.842 | 27.042 | 30.230 | 36.684 | 61.041 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1120/56000
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 225 | 0.004 |
| mysql | 4593 | 0.082 |
| redis | 116593 | 2.082 |
| relation | 225 | 0.004 |

- Cold compute：56000（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 29680000 | 530.000 |
| merge_candidates | 51498720 | 919.620 |
| redis_commands | 336000 | 6.000 |
| redis_members | 57056160 | 1018.860 |
| redis_roundtrips | 56000 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 56000 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 56000 | 5.581 |
| counter | 225 | 5.236 |
| hydrate | 56000 | 10.498 |
| inbox | 56000 | 5.578 |
| merge_dedup | 56000 | 0.105 |
| relation | 225 | 5.760 |
| route | 56000 | 0.131 |
| total | 56000 | 16.444 |

## Redis 本轮边界增量

- Commands：418611；input：1066759636 bytes；output：7258214002 bytes
- Hits/Misses：30017429/5255；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：7385；safety epoch：3770 -> 3770

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.373 |
| client:loadtest | cpu_percent_total | 21.971 |
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
| docker:zg-canal | cpu_percent | 2.010 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.500 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.910 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 93.220 |
| docker:zg-kafka | memory_percent | 7.480 |
| docker:zg-kafka | pids | 119.000 |
| docker:zg-zk | cpu_percent | 35.930 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 557900.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 36.783 |
| process:counter | cpu_seconds_total | 327.531 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45645824.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 85.232 |
| process:gateway | cpu_seconds_total | 6602.562 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51183616.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 245.623 |
| process:knowpost | cpu_seconds_total | 19411.672 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 79974400.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.548 |
| process:relation | cpu_seconds_total | 52.469 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49987584.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.550 |
| process:search | cpu_seconds_total | 6.328 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38064128.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 40.295 |
| process:user-storage | cpu_seconds_total | 2360.781 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54480896.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 318446271.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3770.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2923797039.000 |
| redis | keyspace_misses | 758020.000 |
| redis | net_input_bytes | 122597491903.000 |
| redis | net_output_bytes | 715909653687.000 |
| redis | ops_per_sec | 7385.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34731.000 |
| redis | used_memory_bytes | 102636848.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
