# Feed 压测报告：hybrid / gateway / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 开始时间：2026-08-17T20:04:33+08:00
- 采样时长：1m0.2861008s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 107800 | 107800 | 0 | 0 | 1788.37 | 8.487 | 12.530 | 14.642 | 18.737 | 35.305 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：2156/107800
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 236 | 0.002 |
| mysql | 1066 | 0.010 |
| redis | 111022 | 1.030 |
| relation | 236 | 0.002 |

- Cold compute：107800（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2304764 | 21.380 |
| merge_candidates | 6258868 | 58.060 |
| redis_commands | 1519980 | 14.100 |
| redis_members | 7865088 | 72.960 |
| redis_roundtrips | 213444 | 1.980 |
| tie_members | 1897280 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 107800 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2156 | 2.852 |
| counter | 236 | 3.280 |
| cursor_decode | 105644 | 0.011 |
| cursor_seek | 105644 | 5.054 |
| hydrate | 107800 | 2.667 |
| inbox | 2156 | 2.851 |
| merge_dedup | 107800 | 0.006 |
| relation | 236 | 4.075 |
| route | 107800 | 0.046 |
| total | 107800 | 7.745 |

## Redis 本轮边界增量

- Commands：1661376；input：249208996 bytes；output：714705265 bytes
- Hits/Misses：3830539/1375；run hit rate：99.96%
- Evicted/Rejected：0/0；ops/s max：30420；safety epoch：3775 -> 3775

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.382 |
| client:loadtest | cpu_percent_total | 54.117 |
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
| docker:zg-canal | cpu_percent | 0.270 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.780 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.490 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 24.000 |
| docker:zg-kafka | cpu_percent | 154.860 |
| docker:zg-kafka | memory_percent | 7.650 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 49.060 |
| docker:zg-zk | memory_percent | 1.180 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 575815.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 11.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 10.054 |
| process:counter | cpu_seconds_total | 341.719 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46583808.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 223.199 |
| process:gateway | cpu_seconds_total | 7113.469 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51372032.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 176.703 |
| process:knowpost | cpu_seconds_total | 19998.234 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71995392.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 7.745 |
| process:relation | cpu_seconds_total | 54.891 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48455680.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 6.453 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38064128.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 75.813 |
| process:user-storage | cpu_seconds_total | 2536.531 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 59277312.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 325725142.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3775.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596556.000 |
| redis | keyspace_hits | 2999401984.000 |
| redis | keyspace_misses | 777337.000 |
| redis | net_input_bytes | 125722742501.000 |
| redis | net_output_bytes | 733374349823.000 |
| redis | ops_per_sec | 30420.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 35112.000 |
| redis | used_memory_bytes | 102784488.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
