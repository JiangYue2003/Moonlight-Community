# Feed 压测报告：hybrid / gateway / cursor-page50-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-fixed-formal-20260817a`
- 开始时间：2026-08-17T19:46:16+08:00
- 采样时长：1m0.0182769s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 94571 | 94571 | 0 | 0 | 1575.93 | 9.902 | 14.243 | 16.041 | 20.424 | 37.860 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：1.2987061s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 119 | 0.001 |
| redis | 94690 | 1.001 |
| relation | 240 | 0.003 |

- Cold compute：94571（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1985991 | 21.000 |
| merge_candidates | 2080562 | 22.000 |
| redis_commands | 1229423 | 13.000 |
| redis_members | 2269704 | 24.000 |
| redis_roundtrips | 189142 | 2.000 |
| tie_members | 189142 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 94571 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 3.349 |
| cursor_decode | 94571 | 0.011 |
| cursor_seek | 94571 | 5.457 |
| hydrate | 94571 | 2.957 |
| merge_dedup | 94571 | 0.005 |
| relation | 240 | 4.596 |
| route | 94571 | 0.070 |
| total | 94571 | 8.507 |

## Redis 本轮边界增量

- Commands：1354709；input：207405961 bytes；output：425252362 bytes
- Hits/Misses：3222566/359；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：26201；safety epoch：3762 -> 3762

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.003 |
| client:loadtest | cpu_percent_total | 64.043 |
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
| docker:zg-canal | cpu_percent | 2.220 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.780 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.710 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 150.260 |
| docker:zg-kafka | memory_percent | 7.650 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 46.790 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 530651.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 12.386 |
| process:counter | cpu_seconds_total | 306.859 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46497792.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 244.920 |
| process:gateway | cpu_seconds_total | 6392.234 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51499008.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 202.215 |
| process:knowpost | cpu_seconds_total | 18246.000 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72130560.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.870 |
| process:relation | cpu_seconds_total | 50.047 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49238016.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.558 |
| process:search | cpu_seconds_total | 6.047 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 86.657 |
| process:user-storage | cpu_seconds_total | 2284.906 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 55717888.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 298012958.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3762.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596563.000 |
| redis | keyspace_hits | 2731646545.000 |
| redis | keyspace_misses | 717476.000 |
| redis | net_input_bytes | 114515303688.000 |
| redis | net_output_bytes | 672078676644.000 |
| redis | ops_per_sec | 26201.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34015.000 |
| redis | used_memory_bytes | 101904984.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
