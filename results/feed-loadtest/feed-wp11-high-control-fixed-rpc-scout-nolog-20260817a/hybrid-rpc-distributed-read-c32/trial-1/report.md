# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-control-fixed-rpc-scout-nolog-20260817a`
- 开始时间：2026-08-17T01:06:21+08:00
- 采样时长：10.0077578s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 63726 | 63726 | 0 | 0 | 6370.75 | 4.841 | 6.460 | 7.108 | 8.601 | 14.348 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 25687 | 0.403 |
| mysql | 0 | 0.000 |
| redis | 191178 | 3.000 |
| relation | 63726 | 1.000 |

- Cold compute：63726（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 63726 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 63726 | 0.790 |
| counter | 25687 | 1.104 |
| hydrate | 63726 | 0.867 |
| inbox | 63726 | 0.889 |
| merge_dedup | 63726 | 0.002 |
| relation | 63726 | 1.741 |
| route | 63726 | 0.453 |
| total | 63726 | 4.761 |

## Redis 本轮边界增量

- Commands：629334；input：53467482 bytes；output：119290109 bytes
- Hits/Misses：991122/122041；run hit rate：89.04%
- Evicted/Rejected：0/0；ops/s max：65287；safety epoch：590 -> 590

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.109 |
| client:loadtest | cpu_percent_total | 97.737 |
| client:loadtest | logical_cpus | 16.000 |
| mysql | questions | 27279658.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 41.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 84.364 |
| process:counter | cpu_seconds_total | 124.703 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54013952.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 1.563 |
| process:gateway | cpu_seconds_total | 425.453 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 49029120.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 425.020 |
| process:knowpost | cpu_seconds_total | 1300.719 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 87216128.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 310.322 |
| process:relation | cpu_seconds_total | 1056.141 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 58925056.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.500 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35942400.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 156.000 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 52301824.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 271805567.000 |
| redis | connected_clients | 187.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 590.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 721664.000 |
| redis | keyspace_hits | 1101612208.000 |
| redis | keyspace_misses | 13285058.000 |
| redis | net_input_bytes | 51327046594.000 |
| redis | net_output_bytes | 204568230635.000 |
| redis | ops_per_sec | 65287.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53170.000 |
| redis | used_memory_bytes | 136648744.000 |

## 缺失指标

- docker
- docker_state
- kafka
- docker_state_initial
- docker_state_final

## 说明

- SLA values are reference lines, not pass/fail gates.
