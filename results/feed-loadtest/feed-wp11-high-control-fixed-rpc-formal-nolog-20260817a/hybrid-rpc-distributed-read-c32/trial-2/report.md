# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-control-fixed-rpc-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:13:23+08:00
- 采样时长：1m0.0300428s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 371486 | 371486 | 0 | 0 | 6191.11 | 4.937 | 6.842 | 7.540 | 9.151 | 16.203 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 149288 | 0.402 |
| mysql | 79 | 0.000 |
| redis | 1114537 | 3.000 |
| relation | 371486 | 1.000 |

- Cold compute：371486（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 371486 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 371486 | 0.810 |
| counter | 149288 | 1.173 |
| hydrate | 371486 | 0.895 |
| inbox | 371486 | 0.920 |
| merge_dedup | 371486 | 0.002 |
| relation | 371486 | 1.781 |
| route | 371486 | 0.480 |
| total | 371486 | 4.907 |

## Redis 本轮边界增量

- Commands：3665907；input：311579889 bytes；output：695135768 bytes
- Hits/Misses：5775873/711620；run hit rate：89.03%
- Evicted/Rejected：0/0；ops/s max：72807；safety epoch：599 -> 599

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 5.435 |
| client:loadtest | cpu_percent_total | 86.962 |
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
| docker:zg-canal | cpu_percent | 3.120 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 4.480 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 6.440 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 260.430 |
| docker:zg-kafka | memory_percent | 8.790 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 68.510 |
| docker:zg-zk | memory_percent | 1.550 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 28429352.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 36.000 |
| mysql | threads_running | 9.000 |
| process:counter | cpu_percent | 103.830 |
| process:counter | cpu_seconds_total | 300.062 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54554624.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 2.340 |
| process:gateway | cpu_seconds_total | 494.953 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 48603136.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 393.271 |
| process:knowpost | cpu_seconds_total | 2014.406 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 86515712.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 286.731 |
| process:relation | cpu_seconds_total | 1572.562 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 60149760.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.551 |
| process:search | cpu_seconds_total | 0.688 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 36036608.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 183.922 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 50245632.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 283712657.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 599.000 |
| redis | hit_rate | 0.987 |
| redis | keys | 713018.000 |
| redis | keyspace_hits | 1119747289.000 |
| redis | keyspace_misses | 15486376.000 |
| redis | net_input_bytes | 52319144788.000 |
| redis | net_output_bytes | 206730966001.000 |
| redis | ops_per_sec | 72807.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53642.000 |
| redis | used_memory_bytes | 127706992.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
