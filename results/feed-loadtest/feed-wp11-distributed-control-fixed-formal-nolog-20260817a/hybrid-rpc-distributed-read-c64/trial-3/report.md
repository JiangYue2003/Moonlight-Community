# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-control-fixed-formal-nolog-20260817a`
- 开始时间：2026-08-17T00:59:05+08:00
- 采样时长：1m0.0312641s
- 并发：64
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 407890 | 407890 | 0 | 0 | 6797.75 | 9.050 | 12.783 | 14.057 | 16.799 | 30.229 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 14399 | 0.035 |
| mysql | 247 | 0.001 |
| redis | 1223917 | 3.001 |
| relation | 407890 | 1.000 |

- Cold compute：407890（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 407890 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 407890 | 1.807 |
| counter | 14399 | 2.531 |
| hydrate | 407890 | 2.143 |
| inbox | 407890 | 2.214 |
| merge_dedup | 407890 | 0.003 |
| relation | 407890 | 2.841 |
| route | 407890 | 0.093 |
| total | 407890 | 9.121 |

## Redis 本轮边界增量

- Commands：3214961；input：372336341 bytes；output：1118916326 bytes
- Hits/Misses：7567651/442004；run hit rate：94.48%
- Evicted/Rejected：0/0；ops/s max：60976；safety epoch：585 -> 585

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.520 |
| client:loadtest | cpu_percent_total | 88.313 |
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
| docker:zg-canal | cpu_percent | 0.350 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 1.020 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.720 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 234.380 |
| docker:zg-kafka | memory_percent | 8.780 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 70.190 |
| docker:zg-zk | memory_percent | 1.530 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 26837683.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 40.000 |
| mysql | threads_running | 10.000 |
| process:counter | cpu_percent | 30.936 |
| process:counter | cpu_seconds_total | 40.375 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54022144.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.547 |
| process:gateway | cpu_seconds_total | 52.891 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 48533504.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 377.334 |
| process:knowpost | cpu_seconds_total | 760.953 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 75702272.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 312.476 |
| process:relation | cpu_seconds_total | 635.156 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 60432384.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 2.304 |
| process:search | cpu_seconds_total | 0.234 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35975168.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 21.125 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 38809600.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 267536213.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 585.000 |
| redis | hit_rate | 0.989 |
| redis | keys | 736167.000 |
| redis | keyspace_hits | 1093666801.000 |
| redis | keyspace_misses | 12683624.000 |
| redis | net_input_bytes | 50908793811.000 |
| redis | net_output_bytes | 203475484681.000 |
| redis | ops_per_sec | 60976.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 52784.000 |
| redis | used_memory_bytes | 152399000.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
