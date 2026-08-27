# Feed 压测报告：hybrid / rpc / page20-c16

- Run ID：`feed-cursor-deep-v1-keypoints-runtime-formal-20260817a`
- 开始时间：2026-08-17T19:14:24+08:00
- 采样时长：1m0.0116957s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 71035 | 71035 | 0 | 0 | 1183.78 | 11.798 | 20.616 | 22.600 | 27.340 | 53.940 |

## 分页正确性与准备成本

- 模式：`page`；目标页：20
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Redis 本轮边界增量

- Commands：519133；input：1078282203 bytes；output：7643768244 bytes
- Hits/Misses：30265559/2825；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：9379；safety epoch：3739 -> 3739

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.572 |
| client:loadtest | cpu_percent_total | 25.151 |
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
| docker:zg-canal | cpu_percent | 2.250 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.320 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.590 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 155.140 |
| docker:zg-kafka | memory_percent | 7.630 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 13.570 |
| docker:zg-zk | memory_percent | 1.160 |
| docker:zg-zk | pids | 87.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 464252.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 4.649 |
| process:counter | cpu_seconds_total | 252.125 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47124480.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 5478.406 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44298240.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 289.095 |
| process:knowpost | cpu_seconds_total | 14990.781 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 77471744.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.325 |
| process:relation | cpu_seconds_total | 42.531 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49336320.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 5.078 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38006784.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.004 |
| process:user-storage | cpu_seconds_total | 1947.203 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 46075904.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 244988151.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3739.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596008.000 |
| redis | keyspace_hits | 2185225742.000 |
| redis | keyspace_misses | 594601.000 |
| redis | net_input_bytes | 91837925103.000 |
| redis | net_output_bytes | 551928493239.000 |
| redis | ops_per_sec | 9379.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 32103.000 |
| redis | used_memory_bytes | 102339880.000 |

## 缺失指标

- feed_metrics

## 说明

- SLA values are reference lines, not pass/fail gates.
