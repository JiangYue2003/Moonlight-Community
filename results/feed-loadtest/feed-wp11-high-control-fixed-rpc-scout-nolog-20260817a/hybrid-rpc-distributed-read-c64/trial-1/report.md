# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-high-control-fixed-rpc-scout-nolog-20260817a`
- 开始时间：2026-08-17T01:06:35+08:00
- 采样时长：10.0117983s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 66279 | 66279 | 0 | 0 | 6624.00 | 9.309 | 13.071 | 14.358 | 17.642 | 25.468 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 26186 | 0.395 |
| mysql | 0 | 0.000 |
| redis | 198837 | 3.000 |
| relation | 66279 | 1.000 |

- Cold compute：66279（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 66279 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 66279 | 1.816 |
| counter | 26186 | 2.338 |
| hydrate | 66279 | 1.934 |
| inbox | 66279 | 1.988 |
| merge_dedup | 66279 | 0.002 |
| relation | 66279 | 2.682 |
| route | 66279 | 0.932 |
| total | 66279 | 9.374 |

## Redis 本轮边界增量

- Commands：646155；input：55126072 bytes；output：123888275 bytes
- Hits/Misses：1028330/126931；run hit rate：89.01%
- Evicted/Rejected：0/0；ops/s max：69365；safety epoch：591 -> 591

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.204 |
| client:loadtest | cpu_percent_total | 99.258 |
| client:loadtest | logical_cpus | 16.000 |
| mysql | questions | 27359970.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 43.000 |
| mysql | threads_running | 8.000 |
| process:counter | cpu_percent | 92.183 |
| process:counter | cpu_seconds_total | 133.703 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54218752.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.562 |
| process:gateway | cpu_seconds_total | 425.531 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 49029120.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 415.641 |
| process:knowpost | cpu_seconds_total | 1345.734 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 88510464.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 336.490 |
| process:relation | cpu_seconds_total | 1089.578 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 60010496.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.500 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35942400.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 1.562 |
| process:user-storage | cpu_seconds_total | 156.016 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 45359104.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 272589800.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 591.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 721298.000 |
| redis | keyspace_hits | 1102855788.000 |
| redis | keyspace_misses | 13438867.000 |
| redis | net_input_bytes | 51393987864.000 |
| redis | net_output_bytes | 204718256938.000 |
| redis | ops_per_sec | 69365.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53184.000 |
| redis | used_memory_bytes | 138602536.000 |

## 缺失指标

- docker
- docker_state
- kafka
- docker_state_initial
- docker_state_final

## 说明

- SLA values are reference lines, not pass/fail gates.
