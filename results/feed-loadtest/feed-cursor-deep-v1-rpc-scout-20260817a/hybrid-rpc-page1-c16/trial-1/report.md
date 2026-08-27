# Feed 压测报告：hybrid / rpc / page1-c16

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:16:44+08:00
- 采样时长：10.0042379s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 52151 | 52151 | 0 | 0 | 5214.52 | 2.785 | 4.566 | 5.076 | 6.213 | 10.083 |

## 分页正确性与准备成本

- 模式：`page`；目标页：1
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.001 |
| mysql | 0 | 0.000 |
| redis | 104302 | 2.000 |
| relation | 40 | 0.001 |

- Cold compute：52151（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 52151 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 52151 | 1.512 |
| counter | 40 | 1.482 |
| hydrate | 52151 | 1.308 |
| inbox | 52151 | 1.511 |
| merge_dedup | 52151 | 0.009 |
| relation | 40 | 1.845 |
| route | 52151 | 0.016 |
| total | 52151 | 2.866 |

## Redis 本轮边界增量

- Commands：371202；input：97080420 bytes；output：882902747 bytes
- Hits/Misses：2400201/15；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：39501；safety epoch：3586 -> 3586

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.646 |
| client:loadtest | cpu_percent_total | 74.343 |
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
| docker:zg-canal | memory_percent | 4.190 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.550 |
| docker:zg-es | memory_percent | 11.950 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 2.240 |
| docker:zg-etcd | memory_percent | 0.280 |
| docker:zg-etcd | pids | 19.000 |
| docker:zg-kafka | cpu_percent | 24.510 |
| docker:zg-kafka | memory_percent | 7.020 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 0.150 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 82882.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.507 |
| process:counter | cpu_seconds_total | 31.109 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 43474944.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 37736448.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 257.133 |
| process:knowpost | cpu_seconds_total | 263.859 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 69365760.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.774 |
| process:relation | cpu_seconds_total | 16.609 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 50745344.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 2.172 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42872832.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.203 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40468480.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 5862231.000 |
| redis | connected_clients | 107.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3586.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 597447.000 |
| redis | keyspace_hits | 3365055.000 |
| redis | keyspace_misses | 29226.000 |
| redis | net_input_bytes | 554812879.000 |
| redis | net_output_bytes | 1295190658.000 |
| redis | ops_per_sec | 39501.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21393.000 |
| redis | used_memory_bytes | 101319424.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
