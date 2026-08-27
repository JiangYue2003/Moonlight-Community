# Feed 压测报告：hybrid / rpc / page20-c32

- Run ID：`feed-cursor-deep-v1-rpc-formal-20260817a`
- 开始时间：2026-08-17T17:06:43+08:00
- 采样时长：1m0.0149004s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 70519 | 70519 | 0 | 0 | 1175.12 | 24.544 | 40.068 | 46.985 | 57.972 | 102.253 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.003 |
| mysql | 4042 | 0.057 |
| redis | 145080 | 2.057 |
| relation | 240 | 0.003 |

- Cold compute：70519（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 29617980 | 420.000 |
| merge_candidates | 59306479 | 841.000 |
| redis_commands | 423114 | 6.000 |
| redis_members | 64947999 | 921.000 |
| redis_roundtrips | 70519 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 70519 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 70519 | 13.142 |
| counter | 240 | 10.205 |
| hydrate | 70519 | 13.197 |
| inbox | 70519 | 13.139 |
| merge_dedup | 70519 | 0.079 |
| relation | 240 | 11.069 |
| route | 70519 | 0.437 |
| total | 70519 | 26.968 |

## Redis 本轮边界增量

- Commands：511340；input：1070545870 bytes；output：7587751906 bytes
- Hits/Misses：30043603/4951；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：9146；safety epoch：3642 -> 3642

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.676 |
| client:loadtest | cpu_percent_total | 26.816 |
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
| docker:zg-canal | cpu_percent | 2.390 |
| docker:zg-canal | memory_percent | 4.210 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.630 |
| docker:zg-es | memory_percent | 12.220 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.500 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 145.320 |
| docker:zg-kafka | memory_percent | 7.570 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 45.440 |
| docker:zg-zk | memory_percent | 0.960 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 131714.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 3.101 |
| process:counter | cpu_seconds_total | 53.844 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46608384.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.775 |
| process:gateway | cpu_seconds_total | 258.328 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44023808.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 257.705 |
| process:knowpost | cpu_seconds_total | 3349.016 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 77012992.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.897 |
| process:relation | cpu_seconds_total | 6.891 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48021504.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.031 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38088704.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.576 |
| process:user-storage | cpu_seconds_total | 96.625 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 43991040.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 66985942.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3642.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 597297.000 |
| redis | keyspace_hits | 530690430.000 |
| redis | keyspace_misses | 123779.000 |
| redis | net_input_bytes | 22297867628.000 |
| redis | net_output_bytes | 164859293233.000 |
| redis | ops_per_sec | 9146.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 24442.000 |
| redis | used_memory_bytes | 104454312.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
