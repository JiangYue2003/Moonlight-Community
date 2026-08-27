# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-treatment-fixed-rpc-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:31:19+08:00
- 采样时长：1m0.0368577s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 466152 | 466152 | 0 | 0 | 7768.72 | 3.681 | 8.271 | 10.763 | 25.435 | 80.072 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 785027 | 1.684 |
| counter | 131432 | 0.282 |
| mysql | 35 | 0.000 |
| redis | 860165 | 1.845 |
| relation | 131435 | 0.282 |

- Cold compute：176842（0.379 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 108999 |
| l1_stale | 26909 |
| l2_fresh | 167761 |
| l2_stale | 136856 |
| miss | 25627 |

- L1+L2 Fresh ratio：59.37%
- Refresh max：queue=1024 active=32 pending=1056

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 176848 | 2.373 |
| counter | 131432 | 2.655 |
| hydrate | 176843 | 2.206 |
| inbox | 176848 | 2.373 |
| merge_dedup | 176848 | 0.002 |
| relation | 131436 | 2.862 |
| route | 176852 | 4.192 |
| total | 466152 | 3.865 |

### 非成功 outcome（本轮已降级）

| kind | name | outcome | calls | mean(ms) |
|---|---|---|---:|---:|
| dependency | cache.page_cache_refresh_enqueue | error | 40995 | 0.000 |

## Redis 本轮边界增量

- Commands：2828901；input：375659280 bytes；output：598192935 bytes
- Hits/Misses：3677644/319112；run hit rate：92.02%
- Evicted/Rejected：0/0；ops/s max：63683；safety epoch：610 -> 610

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.871 |
| client:loadtest | cpu_percent_total | 109.932 |
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
| docker:zg-canal | cpu_percent | 0.190 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.650 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.740 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 230.110 |
| docker:zg-kafka | memory_percent | 8.810 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 61.240 |
| docker:zg-zk | memory_percent | 1.580 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29379866.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 28.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 91.179 |
| process:counter | cpu_seconds_total | 164.172 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 54829056.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 618.281 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 51392512.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 428.987 |
| process:knowpost | cpu_seconds_total | 1764.719 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 243417088.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 153.443 |
| process:relation | cpu_seconds_total | 220.797 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 57298944.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 0.344 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35594240.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 152.031 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 49451008.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 297554288.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 610.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 740342.000 |
| redis | keyspace_hits | 1139289779.000 |
| redis | keyspace_misses | 17340435.000 |
| redis | net_input_bytes | 53796382276.000 |
| redis | net_output_bytes | 209662200360.000 |
| redis | ops_per_sec | 63683.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54718.000 |
| redis | used_memory_bytes | 160524568.000 |

## 缺失指标

- feed_degraded

## 说明

- SLA values are reference lines, not pass/fail gates.
