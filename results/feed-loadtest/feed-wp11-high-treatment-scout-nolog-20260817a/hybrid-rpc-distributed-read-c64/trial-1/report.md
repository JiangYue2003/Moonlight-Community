# Feed 压测报告：hybrid / rpc / distributed-read-c64

- Run ID：`feed-wp11-high-treatment-scout-nolog-20260817a`
- 开始时间：2026-08-17T00:11:08+08:00
- 采样时长：10.0196699s
- 并发：64
- 重复轮次：1 / 1
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 92325 | 92325 | 0 | 0 | 9220.25 | 5.320 | 18.066 | 21.679 | 28.922 | 61.329 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 110425 | 1.196 |
| counter | 25285 | 0.274 |
| mysql | 40 | 0.000 |
| redis | 168547 | 1.826 |
| relation | 25290 | 0.274 |

- Cold compute：34481（0.373 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 24206 |
| l1_stale | 2795 |
| l2_fresh | 33710 |
| l2_stale | 11721 |
| miss | 19893 |

- L1+L2 Fresh ratio：62.73%
- Refresh max：queue=941 active=32 pending=973

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 34487 | 2.708 |
| counter | 25285 | 2.942 |
| hydrate | 34482 | 2.538 |
| inbox | 34487 | 2.707 |
| merge_dedup | 34487 | 0.002 |
| relation | 25290 | 3.084 |
| route | 34495 | 4.470 |
| total | 92325 | 6.681 |

### 非成功 outcome（本轮已降级）

| kind | name | outcome | calls | mean(ms) |
|---|---|---|---:|---:|
| dependency | cache.page_cache_refresh_enqueue | error | 6771 | 0.000 |

## Redis 本轮边界增量

- Commands：547398；input：73122460 bytes；output：104843823 bytes
- Hits/Misses：699995/76893；run hit rate：90.10%
- Evicted/Rejected：0/0；ops/s max：59029；safety epoch：559 -> 559

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.534 |
| client:loadtest | cpu_percent_total | 120.544 |
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
| docker:zg-canal | cpu_percent | 0.220 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.520 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.260 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 148.240 |
| docker:zg-kafka | memory_percent | 8.850 |
| docker:zg-kafka | pids | 123.000 |
| docker:zg-zk | cpu_percent | 0.130 |
| docker:zg-zk | memory_percent | 1.460 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23549401.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 44.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 90.454 |
| process:counter | cpu_seconds_total | 358.156 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 53788672.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1763.609 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 52359168.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 432.130 |
| process:knowpost | cpu_seconds_total | 3629.016 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 240336896.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 151.550 |
| process:relation | cpu_seconds_total | 433.078 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 56610816.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.812 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36397056.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 422.219 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 54501376.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 214133535.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 559.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 765769.000 |
| redis | keyspace_hits | 1014714869.000 |
| redis | keyspace_misses | 6464145.000 |
| redis | net_input_bytes | 44051536754.000 |
| redis | net_output_bytes | 191167731733.000 |
| redis | ops_per_sec | 59029.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49857.000 |
| redis | used_memory_bytes | 187821792.000 |

## 缺失指标

- feed_degraded

## 说明

- SLA values are reference lines, not pass/fail gates.
