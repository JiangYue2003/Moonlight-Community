# Feed 压测报告：hybrid / gateway / page50-c32

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:17:57+08:00
- 采样时长：1m0.0434974s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 22757 | 22757 | 0 | 0 | 379.02 | 83.679 | 103.925 | 110.782 | 125.337 | 169.100 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.011 |
| mysql | 10252 | 0.450 |
| redis | 55766 | 2.450 |
| relation | 240 | 0.011 |

- Cold compute：22757（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 23212140 | 1020.000 |
| merge_candidates | 34112743 | 1499.000 |
| redis_commands | 136542 | 6.000 |
| redis_members | 34135500 | 1500.000 |
| redis_roundtrips | 22757 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 22757 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22757 | 16.894 |
| counter | 240 | 11.413 |
| hydrate | 22757 | 63.702 |
| inbox | 22757 | 16.890 |
| merge_dedup | 22757 | 0.166 |
| relation | 240 | 14.361 |
| route | 22757 | 1.426 |
| total | 22757 | 82.417 |

## Redis 本轮边界增量

- Commands：186724；input：826944005 bytes；output：5227703437 bytes
- Hits/Misses：23340571/15520；run hit rate：99.93%
- Evicted/Rejected：0/0；ops/s max：3452；safety epoch：3698 -> 3698

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.225 |
| client:loadtest | cpu_percent_total | 19.595 |
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
| docker:zg-canal | cpu_percent | 2.530 |
| docker:zg-canal | memory_percent | 4.220 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.330 |
| docker:zg-es | memory_percent | 12.350 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.180 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 132.900 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 43.520 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 322626.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 33.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.644 |
| process:counter | cpu_seconds_total | 158.094 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46116864.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 68.961 |
| process:gateway | cpu_seconds_total | 3344.297 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52273152.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 195.895 |
| process:knowpost | cpu_seconds_total | 10435.547 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 99983360.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.326 |
| process:relation | cpu_seconds_total | 27.641 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49025024.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.796 |
| process:search | cpu_seconds_total | 3.859 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38166528.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 25.594 |
| process:user-storage | cpu_seconds_total | 1205.875 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 52637696.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 182154998.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3698.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596583.000 |
| redis | keyspace_hits | 1549309485.000 |
| redis | keyspace_misses | 395395.000 |
| redis | net_input_bytes | 65444156975.000 |
| redis | net_output_bytes | 403999636979.000 |
| redis | ops_per_sec | 3452.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28716.000 |
| redis | used_memory_bytes | 104749144.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
