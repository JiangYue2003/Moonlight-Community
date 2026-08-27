# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:41:38+08:00
- 采样时长：1m0.0507403s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 716622 | 716622 | 0 | 0 | 11942.89 | 3.154 | 4.773 | 6.184 | 9.578 | 34.952 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1073195 | 1.498 |
| counter | 148174 | 0.207 |
| mysql | 22 | 0.000 |
| redis | 1133010 | 1.581 |
| relation | 148186 | 0.207 |

- Cold compute：225470（0.315 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 226480 |
| l1_stale | 32665 |
| l2_fresh | 292463 |
| l2_stale | 144509 |
| miss | 20505 |

- L1+L2 Fresh ratio：72.42%
- Refresh max：queue=1024 active=32 pending=1056

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 225476 | 1.825 |
| counter | 148174 | 2.138 |
| hydrate | 225467 | 1.685 |
| inbox | 225476 | 1.824 |
| merge_dedup | 225476 | 0.002 |
| relation | 148186 | 2.303 |
| route | 225478 | 3.001 |
| total | 716622 | 2.436 |

### 非成功 outcome（本轮已降级）

| kind | name | outcome | calls | mean(ms) |
|---|---|---|---:|---:|
| dependency | cache.page_cache_refresh_enqueue | error | 32364 | 0.000 |

## Redis 本轮边界增量

- Commands：3560709；input：479282401 bytes；output：801014548 bytes
- Hits/Misses：4673199/375115；run hit rate：92.57%
- Evicted/Rejected：0/0；ops/s max：65839；safety epoch：572 -> 572

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.160 |
| client:loadtest | cpu_percent_total | 130.567 |
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
| docker:zg-canal | cpu_percent | 1.750 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.180 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.170 |
| docker:zg-etcd | memory_percent | 0.300 |
| docker:zg-etcd | pids | 25.000 |
| docker:zg-kafka | cpu_percent | 245.780 |
| docker:zg-kafka | memory_percent | 8.760 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 66.280 |
| docker:zg-zk | memory_percent | 1.770 |
| docker:zg-zk | pids | 106.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 24825926.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 27.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 92.810 |
| process:counter | cpu_seconds_total | 823.703 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 55123968.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 1840.969 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 45228032.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 504.977 |
| process:knowpost | cpu_seconds_total | 5756.406 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 252456960.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 166.627 |
| process:relation | cpu_seconds_total | 1085.344 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 63012864.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 1.578 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36220928.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 457.016 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 43601920.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 244711594.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 572.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 816499.000 |
| redis | keyspace_hits | 1052435362.000 |
| redis | keyspace_misses | 9725642.000 |
| redis | net_input_bytes | 48098403345.000 |
| redis | net_output_bytes | 197351567549.000 |
| redis | ops_per_sec | 65839.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51737.000 |
| redis | used_memory_bytes | 237360496.000 |

## 缺失指标

- feed_degraded

## 说明

- SLA values are reference lines, not pass/fail gates.
