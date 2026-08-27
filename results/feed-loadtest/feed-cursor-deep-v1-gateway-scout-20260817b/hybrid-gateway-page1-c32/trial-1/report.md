# Feed 压测报告：hybrid / gateway / page1-c32

- Run ID：`feed-cursor-deep-v1-gateway-scout-20260817b`
- 开始时间：2026-08-17T16:34:07+08:00
- 采样时长：10.0095716s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 21053 | 21053 | 0 | 0 | 2103.51 | 14.828 | 19.123 | 21.687 | 31.208 | 46.057 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 0 | 0.000 |
| redis | 42106 | 2.000 |
| relation | 40 | 0.002 |

- Cold compute：21053（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 842120 | 40.000 |
| merge_candidates | 1705293 | 81.000 |
| redis_commands | 126318 | 6.000 |
| redis_members | 5179038 | 246.000 |
| redis_roundtrips | 21053 | 1.000 |

| page cache source | requests |
|---|---:|
| bypass | 21053 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 21053 | 7.002 |
| counter | 40 | 4.399 |
| hydrate | 21053 | 6.818 |
| inbox | 21053 | 7.000 |
| merge_dedup | 21053 | 0.012 |
| relation | 40 | 5.670 |
| route | 21053 | 0.110 |
| total | 21053 | 13.965 |

## Redis 本轮边界增量

- Commands：150140；input：39185246 bytes；output：356441241 bytes
- Hits/Misses：969686/12；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：15752；safety epoch：3608 -> 3608

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.273 |
| client:loadtest | cpu_percent_total | 68.372 |
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
| docker:zg-canal | cpu_percent | 0.120 |
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.230 |
| docker:zg-es | memory_percent | 12.060 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 0.590 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 129.780 |
| docker:zg-kafka | memory_percent | 7.460 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 49.730 |
| docker:zg-zk | memory_percent | 0.950 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 93219.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 3.125 |
| process:counter | cpu_seconds_total | 4.828 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 42528768.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 239.914 |
| process:gateway | cpu_seconds_total | 57.797 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51761152.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 173.758 |
| process:knowpost | cpu_seconds_total | 55.453 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 65224704.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.775 |
| process:relation | cpu_seconds_total | 0.312 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 43741184.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.219 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 40972288.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 89.746 |
| process:user-storage | cpu_seconds_total | 20.047 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 49516544.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 17929480.000 |
| redis | connected_clients | 100.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3608.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597251.000 |
| redis | keyspace_hits | 89404787.000 |
| redis | keyspace_misses | 62935.000 |
| redis | net_input_bytes | 4322917227.000 |
| redis | net_output_bytes | 23487587515.000 |
| redis | ops_per_sec | 15752.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 22436.000 |
| redis | used_memory_bytes | 102240240.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
