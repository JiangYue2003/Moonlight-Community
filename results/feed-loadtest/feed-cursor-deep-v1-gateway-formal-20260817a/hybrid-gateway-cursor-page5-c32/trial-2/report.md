# Feed 压测报告：hybrid / gateway / cursor-page5-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:54:02+08:00
- 采样时长：1m0.0180775s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 75569 | 75569 | 0 | 0 | 1259.23 | 24.632 | 34.556 | 39.120 | 52.116 | 88.725 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：5
- 准备请求：80；准备耗时：166.1236ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 99 | 0.001 |
| redis | 75668 | 1.001 |
| relation | 240 | 0.003 |

- Cold compute：75569（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 1586949 | 21.000 |
| merge_candidates | 11335350 | 150.000 |
| redis_commands | 1284673 | 17.000 |
| redis_members | 17834284 | 236.000 |
| redis_roundtrips | 151138 | 2.000 |
| tie_members | 9521694 | 126.000 |

| page cache source | requests |
|---|---:|
| bypass | 75569 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 8.602 |
| cursor_decode | 75569 | 0.011 |
| cursor_seek | 75569 | 15.958 |
| hydrate | 75569 | 7.986 |
| merge_dedup | 75569 | 0.009 |
| relation | 240 | 8.899 |
| route | 75569 | 0.249 |
| total | 75569 | 24.216 |

## Redis 本轮边界增量

- Commands：1375707；input：195262443 bytes；output：1015844246 bytes
- Hits/Misses：2878988/100；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：26308；safety epoch：3679 -> 3679

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.509 |
| client:loadtest | cpu_percent_total | 40.144 |
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
| docker:zg-canal | cpu_percent | 2.170 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.340 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.590 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 170.730 |
| docker:zg-kafka | memory_percent | 7.570 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 48.950 |
| docker:zg-zk | memory_percent | 1.010 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 235301.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 4.670 |
| process:counter | cpu_seconds_total | 124.922 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45875200.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 165.760 |
| process:gateway | cpu_seconds_total | 2008.656 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52056064.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 159.688 |
| process:knowpost | cpu_seconds_total | 8258.062 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71266304.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.721 |
| process:relation | cpu_seconds_total | 19.844 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48562176.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.779 |
| process:search | cpu_seconds_total | 3.297 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38154240.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 56.514 |
| process:user-storage | cpu_seconds_total | 728.219 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 58232832.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 164700399.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3679.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596886.000 |
| redis | keyspace_hits | 1165626648.000 |
| redis | keyspace_misses | 259332.000 |
| redis | net_input_bytes | 50924604353.000 |
| redis | net_output_bytes | 313748413163.000 |
| redis | ops_per_sec | 26308.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 27281.000 |
| redis | used_memory_bytes | 106152816.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
