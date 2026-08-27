# Feed 压测报告：hybrid / rpc / page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-formal-20260817a`
- 开始时间：2026-08-17T19:11:53+08:00
- 采样时长：1m0.0088794s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 71330 | 71330 | 0 | 0 | 1188.76 | 11.764 | 20.607 | 22.498 | 26.749 | 53.687 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 1657 | 0.023 |
| redis | 144317 | 2.023 |
| relation | 240 | 0.003 |

- Cold compute：71330（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 29958600 | 420.000 |
| merge_candidates | 59988530 | 841.000 |
| redis_commands | 427980 | 6.000 |
| redis_members | 65694930 | 921.000 |
| redis_roundtrips | 71330 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 71330 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 71330 | 6.320 |
| counter | 240 | 5.669 |
| hydrate | 71330 | 6.529 |
| inbox | 71330 | 6.318 |
| merge_dedup | 71330 | 0.080 |
| relation | 240 | 6.299 |
| route | 71330 | 0.176 |
| total | 71330 | 13.215 |

## Redis 本轮边界增量

- Commands：520339；input：1082555605 bytes；output：7675642819 bytes
- Hits/Misses：30392083/1971；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：9271；safety epoch：3737 -> 3737

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.663 |
| client:loadtest | cpu_percent_total | 26.611 |
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
| docker:zg-canal | cpu_percent | 2.080 |
| docker:zg-canal | memory_percent | 4.270 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.060 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.690 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 109.250 |
| docker:zg-kafka | memory_percent | 7.500 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 44.720 |
| docker:zg-zk | memory_percent | 1.350 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 457218.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.028 |
| process:counter | cpu_seconds_total | 249.359 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46202880.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.548 |
| process:gateway | cpu_seconds_total | 5478.391 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45019136.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 279.624 |
| process:knowpost | cpu_seconds_total | 14691.391 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 76566528.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.322 |
| process:relation | cpu_seconds_total | 41.969 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49008640.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 5.016 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38027264.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 1947.062 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 46379008.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 243766978.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3737.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596041.000 |
| redis | keyspace_hits | 2114553184.000 |
| redis | keyspace_misses | 586523.000 |
| redis | net_input_bytes | 89319126451.000 |
| redis | net_output_bytes | 534079392966.000 |
| redis | ops_per_sec | 9271.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31952.000 |
| redis | used_memory_bytes | 101988848.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
