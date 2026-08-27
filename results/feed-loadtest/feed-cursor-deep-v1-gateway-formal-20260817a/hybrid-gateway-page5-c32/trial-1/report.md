# Feed 压测报告：hybrid / gateway / page5-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T17:45:15+08:00
- 采样时长：1m0.0274684s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 82459 | 82459 | 0 | 0 | 1373.83 | 24.320 | 31.312 | 35.264 | 52.959 | 106.278 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 1391 | 0.017 |
| redis | 166309 | 2.017 |
| relation | 240 | 0.003 |

- Cold compute：82459（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 9895080 | 120.000 |
| merge_candidates | 19872619 | 241.000 |
| redis_commands | 494754 | 6.000 |
| redis_members | 51207039 | 621.000 |
| redis_roundtrips | 82459 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 82459 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 82459 | 11.300 |
| counter | 240 | 9.857 |
| hydrate | 82459 | 10.632 |
| inbox | 82459 | 11.299 |
| merge_dedup | 82459 | 0.027 |
| relation | 240 | 10.938 |
| route | 82459 | 0.343 |
| total | 82459 | 22.346 |

## Redis 本轮边界增量

- Commands：592131；input：384934119 bytes；output：3777126547 bytes
- Hits/Misses：10395823/1473；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：11927；safety epoch：3672 -> 3672

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.076 |
| client:loadtest | cpu_percent_total | 33.214 |
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
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 5.130 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.770 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 122.790 |
| docker:zg-kafka | memory_percent | 7.550 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 48.410 |
| docker:zg-zk | memory_percent | 1.000 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 228373.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 3.113 |
| process:counter | cpu_seconds_total | 111.672 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46510080.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 137.795 |
| process:gateway | cpu_seconds_total | 1411.203 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51834880.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 174.112 |
| process:knowpost | cpu_seconds_total | 7660.094 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 74661888.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.591 |
| process:relation | cpu_seconds_total | 16.578 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49106944.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 3.078 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38137856.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 57.584 |
| process:user-storage | cpu_seconds_total | 513.031 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 54149120.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 155335795.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3672.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596935.000 |
| redis | keyspace_hits | 1124972271.000 |
| redis | keyspace_misses | 254660.000 |
| redis | net_input_bytes | 48903778255.000 |
| redis | net_output_bytes | 299151393911.000 |
| redis | ops_per_sec | 11927.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 26754.000 |
| redis | used_memory_bytes | 103470856.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
