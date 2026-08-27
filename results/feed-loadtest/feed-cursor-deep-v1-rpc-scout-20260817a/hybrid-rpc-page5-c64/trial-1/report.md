# Feed 压测报告：hybrid / rpc / page5-c64

- Run ID：`feed-cursor-deep-v1-rpc-scout-20260817a`
- 开始时间：2026-08-17T16:18:18+08:00
- 采样时长：10.0143509s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 23029 | 23029 | 0 | 0 | 2299.85 | 26.100 | 33.225 | 36.832 | 54.661 | 77.350 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 0 | 0.000 |
| redis | 46058 | 2.000 |
| relation | 40 | 0.002 |

- Cold compute：23029（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 23029 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 23029 | 13.677 |
| counter | 40 | 5.740 |
| hydrate | 23029 | 13.598 |
| inbox | 23029 | 13.675 |
| merge_dedup | 23029 | 0.025 |
| relation | 40 | 6.245 |
| route | 23029 | 0.230 |
| total | 23029 | 27.573 |

## Redis 本轮边界增量

- Commands：163482；input：107341304 bytes；output：1054891094 bytes
- Hits/Misses：2902901/12；run hit rate：100.00%
- Evicted/Rejected：0/0；ops/s max：17638；safety epoch：3591 -> 3591

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.052 |
| client:loadtest | cpu_percent_total | 48.836 |
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
| docker:zg-es | cpu_percent | 3.860 |
| docker:zg-es | memory_percent | 11.960 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.460 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 15.000 |
| docker:zg-kafka | cpu_percent | 151.160 |
| docker:zg-kafka | memory_percent | 7.500 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 0.160 |
| docker:zg-zk | memory_percent | 0.910 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 83777.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 31.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 2.322 |
| process:counter | cpu_seconds_total | 33.062 |
| process:counter | pid | 3628.000 |
| process:counter | process_start_ms | 1786953404373.000 |
| process:counter | rss_bytes | 46854144.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 0.266 |
| process:gateway | pid | 27004.000 |
| process:gateway | process_start_ms | 1786953427247.000 |
| process:gateway | rss_bytes | 38326272.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 243.834 |
| process:knowpost | cpu_seconds_total | 401.219 |
| process:knowpost | pid | 27376.000 |
| process:knowpost | process_start_ms | 1786953415366.000 |
| process:knowpost | rss_bytes | 81436672.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 16.719 |
| process:relation | pid | 26688.000 |
| process:relation | process_start_ms | 1786953408504.000 |
| process:relation | rss_bytes | 55173120.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 2.219 |
| process:search | pid | 4292.000 |
| process:search | process_start_ms | 1786953421788.000 |
| process:search | rss_bytes | 42909696.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 1.234 |
| process:user-storage | pid | 26536.000 |
| process:user-storage | process_start_ms | 1786953400518.000 |
| process:user-storage | rss_bytes | 40906752.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 7523051.000 |
| redis | connected_clients | 160.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3591.000 |
| redis | hit_rate | 0.999 |
| redis | keys | 597292.000 |
| redis | keyspace_hits | 21353191.000 |
| redis | keyspace_misses | 31157.000 |
| redis | net_input_bytes | 1243796821.000 |
| redis | net_output_bytes | 7860832354.000 |
| redis | ops_per_sec | 17638.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 21487.000 |
| redis | used_memory_bytes | 104582248.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
