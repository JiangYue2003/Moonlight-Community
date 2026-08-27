# Feed 压测报告：hybrid / rpc / page20-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:04:13+08:00
- 采样时长：1m0.0130281s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 70331 | 70331 | 0 | 0 | 1172.03 | 11.947 | 20.923 | 22.820 | 26.969 | 44.477 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 1817 | 0.026 |
| redis | 142479 | 2.026 |
| relation | 240 | 0.003 |

- Cold compute：70331（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 29539020 | 420.000 |
| merge_candidates | 59148371 | 841.000 |
| redis_commands | 421986 | 6.000 |
| redis_members | 64774851 | 921.000 |
| redis_roundtrips | 70331 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 70331 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 70331 | 6.407 |
| counter | 240 | 5.447 |
| hydrate | 70331 | 6.628 |
| inbox | 70331 | 6.405 |
| merge_dedup | 70331 | 0.081 |
| relation | 240 | 6.039 |
| route | 70331 | 0.176 |
| total | 70331 | 13.403 |

## Redis 本轮边界增量

- Commands：513852；input：1067565320 bytes；output：7568031538 bytes
- Hits/Misses：29966031/2449；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：9298；safety epoch：3640 -> 3640

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.559 |
| client:loadtest | cpu_percent_total | 24.943 |
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
| docker:zg-canal | cpu_percent | 2.180 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.650 |
| docker:zg-es | memory_percent | 12.220 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.510 |
| docker:zg-etcd | memory_percent | 0.290 |
| docker:zg-etcd | pids | 22.000 |
| docker:zg-kafka | cpu_percent | 153.590 |
| docker:zg-kafka | memory_percent | 7.130 |
| docker:zg-kafka | pids | 99.000 |
| docker:zg-zk | cpu_percent | 46.120 |
| docker:zg-zk | memory_percent | 1.090 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 122218.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.424 |
| process:counter | cpu_seconds_total | 51.734 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46632960.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.010 |
| process:gateway | cpu_seconds_total | 258.266 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44331008.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 253.723 |
| process:knowpost | cpu_seconds_total | 3059.516 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72880128.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 4.645 |
| process:relation | cpu_seconds_total | 6.438 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48558080.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.984 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38096896.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 96.422 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43896832.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 65778773.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3640.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597381.000 |
| redis | keyspace_hits | 460510900.000 |
| redis | keyspace_misses | 113445.000 |
| redis | net_input_bytes | 19796496795.000 |
| redis | net_output_bytes | 147134758085.000 |
| redis | ops_per_sec | 9298.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 24292.000 |
| redis | used_memory_bytes | 103194520.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
