# Feed 压测报告：hybrid / gateway / page50-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:16:42+08:00
- 采样时长：1m0.0457048s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 19202 | 19202 | 0 | 0 | 319.80 | 100.716 | 128.514 | 136.779 | 153.040 | 233.239 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.012 |
| mysql | 9328 | 0.486 |
| redis | 47732 | 2.486 |
| relation | 240 | 0.012 |

- Cold compute：19202（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 19586040 | 1020.000 |
| merge_candidates | 28783798 | 1499.000 |
| redis_commands | 115212 | 6.000 |
| redis_members | 28803000 | 1500.000 |
| redis_roundtrips | 19202 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 19202 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 19202 | 19.629 |
| counter | 240 | 14.839 |
| hydrate | 19202 | 75.803 |
| inbox | 19202 | 19.626 |
| merge_dedup | 19202 | 0.159 |
| relation | 240 | 17.214 |
| route | 19202 | 2.018 |
| total | 19202 | 97.825 |

## Redis 本轮边界增量

- Commands：161239；input：698348603 bytes；output：4410747668 bytes
- Hits/Misses：19693199/15449；run hit rate：99.92%
- Evicted/Rejected：0/0；ops/s max：3071；safety epoch：3697 -> 3697

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.005 |
| client:loadtest | cpu_percent_total | 16.081 |
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
| docker:zg-canal | cpu_percent | 1.940 |
| docker:zg-canal | memory_percent | 4.220 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.640 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.140 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 119.360 |
| docker:zg-kafka | memory_percent | 7.340 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 0.110 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 309663.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 37.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.869 |
| process:counter | cpu_seconds_total | 157.125 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46067712.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 62.750 |
| process:gateway | cpu_seconds_total | 3313.609 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52711424.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 173.349 |
| process:knowpost | cpu_seconds_total | 10319.906 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 104992768.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.091 |
| process:relation | cpu_seconds_total | 27.234 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47972352.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.792 |
| process:search | cpu_seconds_total | 3.797 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38158336.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 20.328 |
| process:user-storage | cpu_seconds_total | 1196.797 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 51347456.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 181933596.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3697.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596586.000 |
| redis | keyspace_hits | 1522183693.000 |
| redis | keyspace_misses | 375853.000 |
| redis | net_input_bytes | 64482504296.000 |
| redis | net_output_bytes | 397923993755.000 |
| redis | ops_per_sec | 3071.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28641.000 |
| redis | used_memory_bytes | 105880088.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
