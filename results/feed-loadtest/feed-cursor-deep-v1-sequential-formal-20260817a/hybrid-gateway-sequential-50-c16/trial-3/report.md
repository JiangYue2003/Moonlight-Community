# Feed 压测报告：hybrid / gateway / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-formal-20260817a`
- 开始时间：2026-08-17T18:37:45+08:00
- 采样时长：1m0.2282675s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 108550 | 108550 | 0 | 0 | 1802.54 | 8.429 | 12.616 | 14.541 | 19.073 | 45.293 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：2171/108550
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 237 | 0.002 |
| mysql | 1015 | 0.009 |
| redis | 111736 | 1.029 |
| relation | 237 | 0.002 |

- Cold compute：108550（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2320799 | 21.380 |
| merge_candidates | 6302413 | 58.060 |
| redis_commands | 1530555 | 14.100 |
| redis_members | 7919808 | 72.960 |
| redis_roundtrips | 214929 | 1.980 |
| tie_members | 1910480 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 108550 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 2171 | 2.796 |
| counter | 237 | 3.206 |
| cursor_decode | 106379 | 0.011 |
| cursor_seek | 106379 | 4.961 |
| hydrate | 108550 | 2.617 |
| inbox | 2171 | 2.795 |
| merge_dedup | 108550 | 0.006 |
| relation | 237 | 4.039 |
| route | 108550 | 0.046 |
| total | 108550 | 7.603 |

## Redis 本轮边界增量

- Commands：1672441；input：250910019 bytes；output：719666722 bytes
- Hits/Misses：3857156/1389；run hit rate：99.96%
- Evicted/Rejected：0/0；ops/s max：32837；safety epoch：3713 -> 3713

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.413 |
| client:loadtest | cpu_percent_total | 54.610 |
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
| docker:zg-canal | cpu_percent | 0.160 |
| docker:zg-canal | memory_percent | 4.240 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.550 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.800 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 145.960 |
| docker:zg-kafka | memory_percent | 7.150 |
| docker:zg-kafka | pids | 98.000 |
| docker:zg-zk | cpu_percent | 45.230 |
| docker:zg-zk | memory_percent | 1.310 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 344446.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 19.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 11.602 |
| process:counter | cpu_seconds_total | 207.359 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46133248.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 229.917 |
| process:gateway | cpu_seconds_total | 4718.406 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51556352.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 173.709 |
| process:knowpost | cpu_seconds_total | 12427.531 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71012352.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.323 |
| process:relation | cpu_seconds_total | 34.406 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49111040.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.188 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38043648.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 81.859 |
| process:user-storage | cpu_seconds_total | 1674.172 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 57475072.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 230335181.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3713.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596552.000 |
| redis | keyspace_hits | 1661955682.000 |
| redis | keyspace_misses | 430998.000 |
| redis | net_input_bytes | 72738867426.000 |
| redis | net_output_bytes | 423551508026.000 |
| redis | ops_per_sec | 32837.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 29904.000 |
| redis | used_memory_bytes | 102167576.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
