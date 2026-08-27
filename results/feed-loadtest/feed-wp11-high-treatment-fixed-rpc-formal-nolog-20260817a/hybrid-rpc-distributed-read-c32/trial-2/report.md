# Feed 压测报告：hybrid / rpc / distributed-read-c32

- Run ID：`feed-wp11-high-treatment-fixed-rpc-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:32:29+08:00
- 采样时长：1m0.0532101s
- 并发：32
- 重复轮次：2 / 3
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 743085 | 743085 | 0 | 0 | 12384.04 | 3.138 | 4.587 | 5.956 | 9.206 | 28.824 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 1103316 | 1.485 |
| counter | 149389 | 0.201 |
| mysql | 30 | 0.000 |
| redis | 1162753 | 1.565 |
| relation | 149397 | 0.201 |

- Cold compute：231559（0.312 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 243828 |
| l1_stale | 30278 |
| l2_fresh | 307841 |
| l2_stale | 140722 |
| miss | 20416 |

- L1+L2 Fresh ratio：74.24%
- Refresh max：queue=1024 active=32 pending=1056

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 231561 | 1.781 |
| counter | 149389 | 2.083 |
| hydrate | 231559 | 1.645 |
| inbox | 231561 | 1.780 |
| merge_dedup | 231561 | 0.002 |
| relation | 149397 | 2.272 |
| route | 231568 | 2.888 |
| total | 743085 | 2.342 |

### 非成功 outcome（本轮已降级）

| kind | name | outcome | calls | mean(ms) |
|---|---|---|---:|---:|
| dependency | cache.page_cache_refresh_enqueue | error | 23843 | 0.000 |

## Redis 本轮边界增量

- Commands：3637394；input：491582507 bytes；output：822028739 bytes
- Hits/Misses：4783272/381929；run hit rate：92.61%
- Evicted/Rejected：0/0；ops/s max：66488；safety epoch：611 -> 611

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.555 |
| client:loadtest | cpu_percent_total | 136.884 |
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
| docker:zg-es | cpu_percent | 4.240 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 6.040 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 235.520 |
| docker:zg-kafka | memory_percent | 8.820 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 64.610 |
| docker:zg-zk | memory_percent | 1.580 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29545744.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 35.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 102.159 |
| process:counter | cpu_seconds_total | 210.297 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 55312384.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 618.281 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 49025024.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 474.020 |
| process:knowpost | cpu_seconds_total | 2042.875 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 247607296.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 176.098 |
| process:relation | cpu_seconds_total | 301.000 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 57421824.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.769 |
| process:search | cpu_seconds_total | 0.359 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35598336.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.000 |
| process:user-storage | cpu_seconds_total | 152.031 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 49451008.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 301558221.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 611.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 758894.000 |
| redis | keyspace_hits | 1144540331.000 |
| redis | keyspace_misses | 17759038.000 |
| redis | net_input_bytes | 54335708117.000 |
| redis | net_output_bytes | 210564409602.000 |
| redis | ops_per_sec | 66488.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54788.000 |
| redis | used_memory_bytes | 179515712.000 |

## 缺失指标

- feed_degraded

## 说明

- SLA values are reference lines, not pass/fail gates.
