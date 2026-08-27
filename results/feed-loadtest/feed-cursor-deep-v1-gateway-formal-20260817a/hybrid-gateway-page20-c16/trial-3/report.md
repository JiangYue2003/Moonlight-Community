# Feed 压测报告：hybrid / gateway / page20-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:59:02+08:00
- 采样时长：1m0.0126604s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 62945 | 62945 | 0 | 0 | 1048.94 | 13.790 | 22.647 | 25.100 | 30.280 | 48.871 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.004 |
| mysql | 2216 | 0.035 |
| redis | 128106 | 2.035 |
| relation | 240 | 0.004 |

- Cold compute：62945（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 26436900 | 420.000 |
| merge_candidates | 52936745 | 841.000 |
| redis_commands | 377670 | 6.000 |
| redis_members | 57972345 | 921.000 |
| redis_roundtrips | 62945 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 62945 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 62945 | 7.076 |
| counter | 240 | 5.850 |
| hydrate | 62945 | 6.978 |
| inbox | 62945 | 7.074 |
| merge_dedup | 62945 | 0.085 |
| relation | 240 | 6.483 |
| route | 62945 | 0.218 |
| total | 62945 | 14.476 |

## Redis 本轮边界增量

- Commands：461746；input：955609589 bytes；output：6773244614 bytes
- Hits/Misses：26819262/2781；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：8433；safety epoch：3683 -> 3683

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.341 |
| client:loadtest | cpu_percent_total | 21.454 |
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
| docker:zg-canal | cpu_percent | 0.190 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.440 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.820 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 192.410 |
| docker:zg-kafka | memory_percent | 7.580 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 47.920 |
| docker:zg-zk | memory_percent | 1.090 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 244995.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 21.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 3.867 |
| process:counter | cpu_seconds_total | 130.203 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46825472.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 92.084 |
| process:gateway | cpu_seconds_total | 2253.859 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51097600.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 241.432 |
| process:knowpost | cpu_seconds_total | 8736.391 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 76275712.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.873 |
| process:relation | cpu_seconds_total | 21.500 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48390144.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.551 |
| process:search | cpu_seconds_total | 3.438 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38162432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 44.964 |
| process:user-storage | cpu_seconds_total | 823.672 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 56893440.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 168058890.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3683.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597055.000 |
| redis | keyspace_hits | 1258976889.000 |
| redis | keyspace_misses | 274065.000 |
| redis | net_input_bytes | 54374301329.000 |
| redis | net_output_bytes | 337702108810.000 |
| redis | ops_per_sec | 8433.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27581.000 |
| redis | used_memory_bytes | 102648040.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
