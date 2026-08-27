# Feed 压测报告：hybrid / gateway / cursor-page20-c32

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:36:40+08:00
- 采样时长：10.013438s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 17667 | 17667 | 0 | 0 | 1764.53 | 17.639 | 24.227 | 26.752 | 36.415 | 51.976 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：505.5637ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 71 | 0.004 |
| redis | 17738 | 1.004 |
| relation | 40 | 0.002 |

- Cold compute：17667（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 371007 | 21.000 |
| merge_candidates | 1130688 | 64.000 |
| redis_commands | 247338 | 14.000 |
| redis_members | 1183689 | 67.000 |
| redis_roundtrips | 35334 | 2.000 |
| tie_members | 53001 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 17667 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 40 | 3.859 |
| cursor_decode | 17667 | 0.012 |
| cursor_seek | 17667 | 11.007 |
| hydrate | 17667 | 5.702 |
| merge_dedup | 17667 | 0.007 |
| relation | 40 | 6.486 |
| route | 17667 | 0.105 |
| total | 17667 | 16.839 |

## Redis 本轮边界增量

- Commands：268198；input：40353329 bytes；output：111765190 bytes
- Hits/Misses：619511/95；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：27536；safety epoch：3616 -> 3616

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.725 |
| client:loadtest | cpu_percent_total | 59.607 |
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
| docker:zg-canal | cpu_percent | 1.850 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 80.000 |
| docker:zg-es | cpu_percent | 1.310 |
| docker:zg-es | memory_percent | 12.080 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.320 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 87.120 |
| docker:zg-kafka | memory_percent | 7.410 |
| docker:zg-kafka | pids | 124.000 |
| docker:zg-zk | cpu_percent | 45.040 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 94898.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 30.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 8.627 |
| process:counter | cpu_seconds_total | 9.156 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45297664.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 207.355 |
| process:gateway | cpu_seconds_total | 187.891 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52137984.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 172.596 |
| process:knowpost | cpu_seconds_total | 206.875 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 67887104.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.104 |
| process:relation | cpu_seconds_total | 1.094 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47882240.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.438 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 41852928.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 81.683 |
| process:user-storage | cpu_seconds_total | 69.922 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 52928512.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 19728447.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3616.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597683.000 |
| redis | keyspace_hits | 106172524.000 |
| redis | keyspace_misses | 70543.000 |
| redis | net_input_bytes | 5021874593.000 |
| redis | net_output_bytes | 28172374055.000 |
| redis | ops_per_sec | 27536.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22589.000 |
| redis | used_memory_bytes | 103369312.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
