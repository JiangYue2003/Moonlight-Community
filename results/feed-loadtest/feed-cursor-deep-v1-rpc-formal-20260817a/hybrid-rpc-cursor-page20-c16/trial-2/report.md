# Feed 压测报告：hybrid / rpc / cursor-page20-c16

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:11:44+08:00
- 采样时长：1m0.0191131s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 284687 | 284687 | 0 | 0 | 4744.64 | 3.171 | 4.805 | 5.468 | 6.896 | 13.720 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：20
- 准备请求：380；准备耗时：316.7684ms（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 167 | 0.001 |
| redis | 284854 | 1.001 |
| relation | 240 | 0.001 |

- Cold compute：284687（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 5978427 | 21.000 |
| merge_candidates | 18219968 | 64.000 |
| redis_commands | 3985618 | 14.000 |
| redis_members | 19074029 | 67.000 |
| redis_roundtrips | 569374 | 2.000 |
| tie_members | 854061 | 3.000 |

| page cache source | requests |
|---|---:|
| bypass | 284687 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 1.417 |
| cursor_decode | 284687 | 0.009 |
| cursor_seek | 284687 | 2.097 |
| hydrate | 284687 | 1.023 |
| merge_dedup | 284687 | 0.007 |
| relation | 240 | 2.242 |
| route | 284687 | 0.012 |
| total | 284687 | 3.152 |

## Redis 本轮边界增量

- Commands：4306589；input：649396801 bytes；output：1800610687 bytes
- Hits/Misses：9971407/167；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：80505；safety epoch：3646 -> 3646

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.531 |
| client:loadtest | cpu_percent_total | 72.503 |
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
| docker:zg-es | cpu_percent | 2.400 |
| docker:zg-es | memory_percent | 12.220 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.420 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 158.550 |
| docker:zg-kafka | memory_percent | 7.530 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 42.660 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 145234.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 18.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.966 |
| process:counter | cpu_seconds_total | 59.828 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45899776.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.344 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44060672.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 245.912 |
| process:knowpost | cpu_seconds_total | 3926.469 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 69349376.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.885 |
| process:relation | cpu_seconds_total | 7.891 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47464448.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 2.188 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38113280.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.775 |
| process:user-storage | cpu_seconds_total | 96.844 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44023808.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 78360926.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3646.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597217.000 |
| redis | keyspace_hits | 625184212.000 |
| redis | keyspace_misses | 138287.000 |
| redis | net_input_bytes | 26359646322.000 |
| redis | net_output_bytes | 187037562176.000 |
| redis | ops_per_sec | 80505.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 24743.000 |
| redis | used_memory_bytes | 102737816.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
