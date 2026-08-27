# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:42:48+08:00
- 采样时长：1m0.0510007s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 698670 | 698670 | 0 | 0 | 11643.61 | 3.165 | 4.872 | 6.343 | 9.917 | 63.945 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1057628 | 1.514 |
| counter | 147872 | 0.212 |
| mysql | 24 | 0.000 |
| redis | 1108599 | 1.587 |
| relation | 147875 | 0.212 |

- Cold compute：224109（0.321 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 0 |
| l1_fresh | 209779 |
| l1_stale | 39430 |
| l2_fresh | 266839 |
| l2_stale | 161093 |
| miss | 21529 |

- L1+L2 Fresh ratio：68.22%
- Refresh max：queue=1024 active=32 pending=1056

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 224116 | 1.842 |
| counter | 147872 | 2.144 |
| hydrate | 224109 | 1.702 |
| inbox | 224116 | 1.841 |
| merge_dedup | 224116 | 0.002 |
| relation | 147875 | 2.329 |
| route | 224125 | 3.037 |
| total | 698670 | 2.501 |

### 非成功 outcome（本轮已降级）

| kind | name | outcome | calls | mean(ms) |
|---|---|---|---:|---:|
| dependency | cache.page_cache_refresh_enqueue | error | 50691 | 0.000 |
| dependency | redis.page_cache_l2_set | error | 12320 | 0.000 |

## Redis 本轮边界增量

- Commands：3515580；input：475587736 bytes；output：794939239 bytes
- Hits/Misses：4636769/374579；run hit rate：92.53%
- Evicted/Rejected：0/0；ops/s max：67195；safety epoch：573 -> 573

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.302 |
| client:loadtest | cpu_percent_total | 132.830 |
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
| docker:zg-canal | cpu_percent | 1.630 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.830 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.440 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 235.760 |
| docker:zg-kafka | memory_percent | 8.810 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 69.100 |
| docker:zg-zk | memory_percent | 1.750 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 24989067.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 8.000 |
| process:counter | cpu_percent | 103.629 |
| process:counter | cpu_seconds_total | 870.625 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54784000.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 1840.969 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 45211648.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 487.381 |
| process:knowpost | cpu_seconds_total | 6026.906 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 248991744.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 170.674 |
| process:relation | cpu_seconds_total | 1162.250 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 63111168.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.773 |
| process:search | cpu_seconds_total | 1.625 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36196352.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.773 |
| process:user-storage | cpu_seconds_total | 457.062 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 43622400.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 248563799.000 |
| redis | connected_clients | 337.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 573.000 |
| redis | hit_rate | 0.991 |
| redis | keys | 817916.000 |
| redis | keyspace_hits | 1057498039.000 |
| redis | keyspace_misses | 10133999.000 |
| redis | net_input_bytes | 48617592992.000 |
| redis | net_output_bytes | 198219079453.000 |
| redis | ops_per_sec | 67195.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 51807.000 |
| redis | used_memory_bytes | 238398640.000 |

## 缺失指标

- feed_degraded

## 说明

- SLA values are reference lines, not pass/fail gates.
