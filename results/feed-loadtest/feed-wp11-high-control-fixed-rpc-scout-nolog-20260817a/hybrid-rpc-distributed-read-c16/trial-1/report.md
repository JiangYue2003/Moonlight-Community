# Feed 压测报告：hybrid / rpc / distributed-read-c16

- Run ID：`feed-wp11-high-control-fixed-rpc-scout-nolog-20260817a`
- 开始时间：2026-08-17T01:06:07+08:00
- 采样时长：10.0056001s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 60267 | 60267 | 0 | 0 | 6025.66 | 2.623 | 3.701 | 4.164 | 4.856 | 10.105 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 22930 | 0.380 |
| mysql | 0 | 0.000 |
| redis | 180801 | 3.000 |
| relation | 60267 | 1.000 |

- Cold compute：60267（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 60267 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 60267 | 0.349 |
| counter | 22930 | 0.621 |
| hydrate | 60267 | 0.379 |
| inbox | 60267 | 0.373 |
| merge_dedup | 60267 | 0.002 |
| relation | 60267 | 1.045 |
| route | 60267 | 0.243 |
| total | 60267 | 2.410 |

## Redis 本轮边界增量

- Commands：589136；input：50420772 bytes；output：112650660 bytes
- Hits/Misses：930680/115409；run hit rate：88.97%
- Evicted/Rejected：0/0；ops/s max：64290；safety epoch：589 -> 589

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.813 |
| client:loadtest | cpu_percent_total | 109.001 |
| client:loadtest | logical_cpus | 16.000 |
| mysql | questions | 27202592.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 34.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 134.425 |
| process:counter | cpu_seconds_total | 116.516 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 53940224.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.562 |
| process:gateway | cpu_seconds_total | 425.391 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 50819072.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 422.087 |
| process:knowpost | cpu_seconds_total | 1256.641 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 83632128.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 306.251 |
| process:relation | cpu_seconds_total | 1022.438 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 58126336.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.500 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35901440.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 156.000 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 52281344.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 271042669.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 589.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 722038.000 |
| redis | keyspace_hits | 1100414418.000 |
| redis | keyspace_misses | 13137485.000 |
| redis | net_input_bytes | 51262277241.000 |
| redis | net_output_bytes | 204424042430.000 |
| redis | ops_per_sec | 64290.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53156.000 |
| redis | used_memory_bytes | 135380800.000 |

## 缺失指标

- docker
- docker_state
- kafka
- docker_state_initial
- docker_state_final

## 说明

- SLA values are reference lines, not pass/fail gates.
