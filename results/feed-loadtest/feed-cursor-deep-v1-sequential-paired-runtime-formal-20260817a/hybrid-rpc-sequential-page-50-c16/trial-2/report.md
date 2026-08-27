# Feed 压测报告：hybrid / rpc / sequential-page-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 开始时间：2026-08-17T19:51:53+08:00
- 采样时长：1m0.684768s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 60000 | 60000 | 0 | 0 | 988.78 | 16.187 | 25.618 | 28.651 | 34.488 | 59.508 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1200/60000
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 227 | 0.004 |
| mysql | 3131 | 0.052 |
| redis | 123131 | 2.052 |
| relation | 227 | 0.004 |

- Cold compute：60000（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 31800000 | 530.000 |
| merge_candidates | 55177200 | 919.620 |
| redis_commands | 360000 | 6.000 |
| redis_members | 61131600 | 1018.860 |
| redis_roundtrips | 60000 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 60000 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 60000 | 5.228 |
| counter | 227 | 5.140 |
| hydrate | 60000 | 10.323 |
| inbox | 60000 | 5.225 |
| merge_dedup | 60000 | 0.101 |
| relation | 227 | 5.115 |
| route | 60000 | 0.118 |
| total | 60000 | 15.896 |

## Redis 本轮边界增量

- Commands：445692；input：1142597695 bytes；output：7776782558 bytes
- Hits/Misses：32162239/4520；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：8152；safety epoch：3765 -> 3765

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.479 |
| client:loadtest | cpu_percent_total | 23.662 |
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
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.600 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.450 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 111.150 |
| docker:zg-kafka | memory_percent | 7.540 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 38.740 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 540215.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.099 |
| process:counter | cpu_seconds_total | 317.484 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46071808.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.550 |
| process:gateway | cpu_seconds_total | 6559.016 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 47349760.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 260.874 |
| process:knowpost | cpu_seconds_total | 18681.562 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 80965632.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.323 |
| process:relation | cpu_seconds_total | 51.188 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49545216.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 6.125 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38031360.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 6.469 |
| process:user-storage | cpu_seconds_total | 2342.281 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 48631808.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 301298338.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3765.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2812509777.000 |
| redis | keyspace_misses | 738512.000 |
| redis | net_input_bytes | 117543676849.000 |
| redis | net_output_bytes | 691070857000.000 |
| redis | ops_per_sec | 8152.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34352.000 |
| redis | used_memory_bytes | 102480368.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
