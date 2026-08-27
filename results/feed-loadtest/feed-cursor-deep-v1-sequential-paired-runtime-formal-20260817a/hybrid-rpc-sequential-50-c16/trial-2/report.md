# Feed 压测报告：hybrid / rpc / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-paired-runtime-formal-20260817a`
- 开始时间：2026-08-17T19:55:41+08:00
- 采样时长：1m0.0756824s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 302000 | 302000 | 0 | 0 | 5028.57 | 2.918 | 4.716 | 5.376 | 6.950 | 14.150 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：6040/302000
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 1036 | 0.003 |
| redis | 309076 | 1.023 |
| relation | 240 | 0.001 |

- Cold compute：302000（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6456760 | 21.380 |
| merge_candidates | 17534120 | 58.060 |
| redis_commands | 4258200 | 14.100 |
| redis_members | 22033920 | 72.960 |
| redis_roundtrips | 597960 | 1.980 |
| tie_members | 5315200 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 302000 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 6040 | 1.156 |
| counter | 240 | 1.367 |
| cursor_decode | 295960 | 0.009 |
| cursor_seek | 295960 | 2.003 |
| hydrate | 302000 | 0.965 |
| inbox | 6040 | 1.155 |
| merge_dedup | 302000 | 0.006 |
| relation | 240 | 2.119 |
| route | 302000 | 0.010 |
| total | 302000 | 2.980 |

## Redis 本轮边界增量

- Commands：4597313；input：694294554 bytes；output：2001212763 bytes
- Hits/Misses：10720971/1422；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：83029；safety epoch：3768 -> 3768

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.332 |
| client:loadtest | cpu_percent_total | 69.314 |
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
| docker:zg-canal | cpu_percent | 1.960 |
| docker:zg-canal | memory_percent | 4.290 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.660 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.520 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 148.550 |
| docker:zg-kafka | memory_percent | 7.640 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 43.890 |
| docker:zg-zk | memory_percent | 1.380 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 550012.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 7.744 |
| process:counter | cpu_seconds_total | 323.984 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46714880.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 3.096 |
| process:gateway | cpu_seconds_total | 6559.109 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44662784.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 231.295 |
| process:knowpost | cpu_seconds_total | 19121.078 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71450624.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 1.547 |
| process:relation | cpu_seconds_total | 52.125 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49745920.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.798 |
| process:search | cpu_seconds_total | 6.250 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38064128.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 3.116 |
| process:user-storage | cpu_seconds_total | 2342.797 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 47030272.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 312578676.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3768.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596563.000 |
| redis | keyspace_hits | 2875693611.000 |
| redis | keyspace_misses | 749483.000 |
| redis | net_input_bytes | 120521820367.000 |
| redis | net_output_bytes | 704967977156.000 |
| redis | ops_per_sec | 83029.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 34580.000 |
| redis | used_memory_bytes | 102252984.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
