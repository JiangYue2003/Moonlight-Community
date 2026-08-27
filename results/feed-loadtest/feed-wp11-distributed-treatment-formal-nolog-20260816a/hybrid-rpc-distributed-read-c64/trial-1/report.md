# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-distributed-treatment-formal-nolog-20260816a`
- 开始时间：2026-08-16T23:50:25+08:00
- 采样时长：1m0.1602936s
- 并发：64
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 5907278 | 5907278 | 0 | 0 | 98453.61 | 0.528 | 1.077 | 1.167 | 1.695 | 25.040 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 146728 | 0.025 |
| counter | 10815 | 0.002 |
| mysql | 1 | 0.000 |
| redis | 145476 | 0.025 |
| relation | 10815 | 0.002 |

- Cold compute：22263（0.004 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 5827339 |
| l1_stale | 0 |
| l2_fresh | 79939 |
| l2_stale | 0 |
| miss | 0 |

- L1+L2 Fresh ratio：100.00%
- Refresh max：queue=0 active=2 pending=2

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 22263 | 0.167 |
| counter | 10815 | 0.507 |
| hydrate | 22263 | 0.185 |
| inbox | 22263 | 0.166 |
| merge_dedup | 22263 | 0.005 |
| relation | 10815 | 1.132 |
| route | 22263 | 0.808 |
| total | 5907278 | 0.008 |

## Redis 本轮边界增量

- Commands：457696；input：68741547 bytes；output：167472163 bytes
- Hits/Misses：626562/12668；run hit rate：98.02%
- Evicted/Rejected：0/0；ops/s max：13999；safety epoch：545 -> 545

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 24.024 |
| client:loadtest | cpu_percent_total | 384.390 |
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
| docker:zg-canal | cpu_percent | 0.200 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.410 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.920 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 302.270 |
| docker:zg-kafka | memory_percent | 8.760 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 69.480 |
| docker:zg-zk | memory_percent | 1.670 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23344463.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 64.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 23.988 |
| process:counter | cpu_seconds_total | 89.797 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54239232.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 195.672 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 46985216.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 598.756 |
| process:knowpost | cpu_seconds_total | 1559.188 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 108052480.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 35.506 |
| process:relation | cpu_seconds_total | 76.688 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 57315328.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.777 |
| process:search | cpu_seconds_total | 0.359 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36257792.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 59.594 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 41869312.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 208263411.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 545.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 701108.000 |
| redis | keyspace_hits | 1007161103.000 |
| redis | keyspace_misses | 6064161.000 |
| redis | net_input_bytes | 43197615354.000 |
| redis | net_output_bytes | 189616557336.000 |
| redis | ops_per_sec | 13999.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 48664.000 |
| redis | used_memory_bytes | 122379864.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
