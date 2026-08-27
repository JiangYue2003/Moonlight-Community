# Feed 压测报告：hybrid / rpc / page5-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:51:42+08:00
- 采样时长：1m0.0187198s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 141751 | 141751 | 0 | 0 | 2362.17 | 12.260 | 20.010 | 22.744 | 27.436 | 49.127 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 1207 | 0.009 |
| redis | 284709 | 2.009 |
| relation | 240 | 0.002 |

- Cold compute：141751（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 17010120 | 120.000 |
| merge_candidates | 34161991 | 241.000 |
| redis_commands | 850506 | 6.000 |
| redis_members | 88027371 | 621.000 |
| redis_roundtrips | 141751 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 141751 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 141751 | 6.707 |
| counter | 240 | 6.303 |
| hydrate | 141751 | 6.375 |
| inbox | 141751 | 6.705 |
| merge_dedup | 141751 | 0.032 |
| relation | 240 | 6.339 |
| route | 141751 | 0.121 |
| total | 141751 | 13.286 |

## Redis 本轮边界增量

- Commands：1011792；input：661325912 bytes；output：6493062111 bytes
- Hits/Misses：17866862/1348；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：18320；safety epoch：3630 -> 3630

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.753 |
| client:loadtest | cpu_percent_total | 44.049 |
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
| docker:zg-canal | cpu_percent | 2.000 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 6.710 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.470 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 134.690 |
| docker:zg-kafka | memory_percent | 7.520 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 43.910 |
| docker:zg-zk | memory_percent | 1.220 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 108686.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.875 |
| process:counter | cpu_seconds_total | 34.969 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 45068288.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.141 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 45371392.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 280.369 |
| process:knowpost | cpu_seconds_total | 1876.328 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 74362880.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 3.578 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 47726592.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.991 |
| process:search | cpu_seconds_total | 1.688 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 42332160.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 95.422 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44326912.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 41155381.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3630.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597901.000 |
| redis | keyspace_hits | 303601103.000 |
| redis | keyspace_misses | 96500.000 |
| redis | net_input_bytes | 12726218889.000 |
| redis | net_output_bytes | 98386569841.000 |
| redis | ops_per_sec | 18320.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23541.000 |
| redis | used_memory_bytes | 103391768.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
