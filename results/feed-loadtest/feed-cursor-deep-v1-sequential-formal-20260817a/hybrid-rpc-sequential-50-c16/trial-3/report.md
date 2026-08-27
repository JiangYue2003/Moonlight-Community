# Feed 压测报告：hybrid / rpc / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-formal-20260817a`
- 开始时间：2026-08-17T18:30:13+08:00
- 采样时长：1m0.0691412s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 296250 | 296250 | 0 | 0 | 4933.34 | 3.107 | 4.782 | 5.496 | 7.106 | 15.482 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：5925/296250
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 952 | 0.003 |
| redis | 303127 | 1.023 |
| relation | 240 | 0.001 |

- Cold compute：296250（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6333825 | 21.380 |
| merge_candidates | 17200275 | 58.060 |
| redis_commands | 4177125 | 14.100 |
| redis_members | 21614400 | 72.960 |
| redis_roundtrips | 586575 | 1.980 |
| tie_members | 5214000 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 296250 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 5925 | 1.182 |
| counter | 240 | 1.413 |
| cursor_decode | 290325 | 0.009 |
| cursor_seek | 290325 | 2.040 |
| hydrate | 296250 | 0.985 |
| inbox | 5925 | 1.180 |
| merge_dedup | 296250 | 0.006 |
| relation | 240 | 2.114 |
| route | 296250 | 0.010 |
| total | 296250 | 3.038 |

## Redis 本轮边界增量

- Commands：4510553；input：681143817 bytes；output：1963107699 bytes
- Hits/Misses：10516860/1519；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：82133；safety epoch：3707 -> 3707

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.516 |
| client:loadtest | cpu_percent_total | 72.260 |
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
| docker:zg-canal | cpu_percent | 2.100 |
| docker:zg-canal | memory_percent | 4.230 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.630 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.470 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 206.570 |
| docker:zg-kafka | memory_percent | 7.150 |
| docker:zg-kafka | pids | 97.000 |
| docker:zg-zk | cpu_percent | 5.310 |
| docker:zg-zk | memory_percent | 1.160 |
| docker:zg-zk | pids | 101.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 332320.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.196 |
| process:counter | cpu_seconds_total | 189.734 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47226880.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4321.891 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 48222208.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 249.114 |
| process:knowpost | cpu_seconds_total | 11657.188 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71143424.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.095 |
| process:relation | cpu_seconds_total | 31.938 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48672768.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.781 |
| process:search | cpu_seconds_total | 4.109 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38035456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.770 |
| process:user-storage | cpu_seconds_total | 1533.578 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 48762880.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 208932952.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3707.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596552.000 |
| redis | keyspace_hits | 1612236311.000 |
| redis | keyspace_misses | 417601.000 |
| redis | net_input_bytes | 69513278824.000 |
| redis | net_output_bytes | 414271798342.000 |
| redis | ops_per_sec | 82133.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 29452.000 |
| redis | used_memory_bytes | 102228784.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
