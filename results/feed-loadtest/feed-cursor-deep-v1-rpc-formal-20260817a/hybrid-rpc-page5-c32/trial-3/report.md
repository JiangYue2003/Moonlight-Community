# Feed 压测报告：hybrid / rpc / page5-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T16:54:12+08:00
- 采样时长：1m0.0174085s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 142775 | 142775 | 0 | 0 | 2379.31 | 12.166 | 19.945 | 22.602 | 27.173 | 51.562 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 1461 | 0.010 |
| redis | 287011 | 2.010 |
| relation | 240 | 0.002 |

- Cold compute：142775（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 17133000 | 120.000 |
| merge_candidates | 34408775 | 241.000 |
| redis_commands | 856650 | 6.000 |
| redis_members | 88663275 | 621.000 |
| redis_roundtrips | 142775 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 142775 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 142775 | 6.669 |
| counter | 240 | 5.935 |
| hydrate | 142775 | 6.316 |
| inbox | 142775 | 6.668 |
| merge_dedup | 142775 | 0.032 |
| relation | 240 | 6.029 |
| route | 142775 | 0.120 |
| total | 142775 | 13.190 |

## Redis 本轮边界增量

- Commands：1019249；input：666135786 bytes；output：6539910132 bytes
- Hits/Misses：17995514/1625；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：18308；safety epoch：3632 -> 3632

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.846 |
| client:loadtest | cpu_percent_total | 45.534 |
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
| docker:zg-canal | cpu_percent | 1.860 |
| docker:zg-canal | memory_percent | 4.200 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.510 |
| docker:zg-es | memory_percent | 12.210 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.640 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 25.000 |
| docker:zg-kafka | cpu_percent | 118.740 |
| docker:zg-kafka | memory_percent | 7.490 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 41.320 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 113050.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 33.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.725 |
| process:counter | cpu_seconds_total | 37.672 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46714880.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 258.156 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44752896.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 264.577 |
| process:knowpost | cpu_seconds_total | 2177.438 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 75317248.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.150 |
| process:relation | cpu_seconds_total | 4.328 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48168960.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 1.766 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 41844736.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.548 |
| process:user-storage | cpu_seconds_total | 95.672 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43991040.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 43562379.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3632.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597784.000 |
| redis | keyspace_hits | 345978270.000 |
| redis | keyspace_misses | 100432.000 |
| redis | net_input_bytes | 14295399445.000 |
| redis | net_output_bytes | 113787419187.000 |
| redis | ops_per_sec | 18308.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23691.000 |
| redis | used_memory_bytes | 103314936.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
