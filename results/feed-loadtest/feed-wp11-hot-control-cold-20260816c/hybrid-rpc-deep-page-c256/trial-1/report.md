# Feed 压测报告：hybrid / rpc / deep-page-c256

- Run ID：`feed-wp11-hot-control-cold-20260816c`
- 开始时间：2026-08-16T16:46:13+08:00
- 采样时长：1m0.0784969s
- 并发：256
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 175429 | 175429 | 0 | 0 | 2920.68 | 74.672 | 144.342 | 174.410 | 244.042 | 515.292 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 169 | 0.001 |
| redis | 526456 | 3.001 |
| relation | 175429 | 1.000 |

- Cold compute：175429（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 175429 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 175429 | 0.217 |
| counter | 240 | 0.660 |
| hydrate | 175429 | 0.372 |
| inbox | 175429 | 0.237 |
| merge_dedup | 175429 | 0.014 |
| relation | 175429 | 86.282 |
| route | 175429 | 0.004 |
| total | 175429 | 87.153 |

## Redis 本轮边界增量

- Commands：1468340；input：445198489 bytes；output：2068296684 bytes
- Hits/Misses：11233376/175599；run hit rate：98.46%
- Evicted/Rejected：0/0；ops/s max：25980；safety epoch：425 -> 425

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.415 |
| client:loadtest | cpu_percent_total | 70.637 |
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
| docker:zg-canal | memory_percent | 7.840 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.790 |
| docker:zg-es | memory_percent | 11.850 |
| docker:zg-es | pids | 149.000 |
| docker:zg-etcd | cpu_percent | 4.980 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 178.520 |
| docker:zg-kafka | memory_percent | 8.460 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 53.850 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3101.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3101.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 16215550.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 74.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 12.369 |
| process:counter | cpu_seconds_total | 169.453 |
| process:counter | pid | 28056.000 |
| process:counter | process_start_ms | 1786867267666.000 |
| process:counter | rss_bytes | 43671552.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.768 |
| process:gateway | cpu_seconds_total | 0.703 |
| process:gateway | pid | 21316.000 |
| process:gateway | process_start_ms | 1786867287364.000 |
| process:gateway | rss_bytes | 37044224.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 325.680 |
| process:knowpost | cpu_seconds_total | 6504.766 |
| process:knowpost | pid | 21296.000 |
| process:knowpost | process_start_ms | 1786867278931.000 |
| process:knowpost | rss_bytes | 122953728.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 274.784 |
| process:relation | cpu_seconds_total | 6027.547 |
| process:relation | pid | 26828.000 |
| process:relation | process_start_ms | 1786867272170.000 |
| process:relation | rss_bytes | 101572608.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 1.062 |
| process:search | pid | 1848.000 |
| process:search | process_start_ms | 1786867283582.000 |
| process:search | rss_bytes | 37060608.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.774 |
| process:user-storage | cpu_seconds_total | 1.234 |
| process:user-storage | pid | 29400.000 |
| process:user-storage | process_start_ms | 1786867263806.000 |
| process:user-storage | rss_bytes | 35237888.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 132216456.000 |
| redis | connected_clients | 304.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 425.000 |
| redis | hit_rate | 0.998 |
| redis | keys | 697528.000 |
| redis | keyspace_hits | 788115527.000 |
| redis | keyspace_misses | 1402586.000 |
| redis | net_input_bytes | 31886790026.000 |
| redis | net_output_bytes | 144258281052.000 |
| redis | ops_per_sec | 25980.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 23212.000 |
| redis | used_memory_bytes | 107875960.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
