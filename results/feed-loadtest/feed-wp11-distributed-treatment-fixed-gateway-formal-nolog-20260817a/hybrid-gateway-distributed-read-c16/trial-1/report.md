# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-fixed-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:26:52+08:00
- 采样时长：1m0.0179097s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 96552 | 96552 | 0 | 0 | 1608.90 | 6.367 | 22.627 | 26.338 | 33.405 | 69.013 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 96349 | 0.998 |
| counter | 9075 | 0.094 |
| mysql | 3 | 0.000 |
| redis | 95743 | 0.992 |
| relation | 9075 | 0.094 |

- Cold compute：17479（0.181 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 52341 |
| l1_stale | 299 |
| l2_fresh | 37343 |
| l2_stale | 6569 |
| miss | 0 |

- L1+L2 Fresh ratio：92.89%
- Refresh max：queue=0 active=18 pending=18

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 17479 | 7.374 |
| counter | 9075 | 9.013 |
| hydrate | 17479 | 7.227 |
| inbox | 17479 | 7.373 |
| merge_dedup | 17479 | 0.004 |
| relation | 9075 | 11.383 |
| route | 17479 | 10.623 |
| total | 96552 | 5.826 |

## Redis 本轮边界增量

- Commands：293378；input：49771683 bytes；output：106720841 bytes
- Hits/Misses：460824/10518；run hit rate：97.77%
- Evicted/Rejected：0/0；ops/s max：14134；safety epoch：607 -> 607

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.182 |
| client:loadtest | cpu_percent_total | 114.914 |
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
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.050 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 5.160 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 180.640 |
| docker:zg-kafka | memory_percent | 8.830 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 50.360 |
| docker:zg-zk | memory_percent | 1.830 |
| docker:zg-zk | pids | 107.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29213267.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 58.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 61.257 |
| process:counter | cpu_seconds_total | 57.062 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 53796864.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 516.116 |
| process:gateway | cpu_seconds_total | 225.938 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 51556352.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 272.475 |
| process:knowpost | cpu_seconds_total | 1204.734 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 101474304.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 120.963 |
| process:relation | cpu_seconds_total | 63.531 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 56606720.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.776 |
| process:search | cpu_seconds_total | 0.125 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35553280.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 138.470 |
| process:user-storage | cpu_seconds_total | 53.297 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 51290112.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 293774970.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 607.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 709218.000 |
| redis | keyspace_hits | 1134313958.000 |
| redis | keyspace_misses | 16956090.000 |
| redis | net_input_bytes | 53274985104.000 |
| redis | net_output_bytes | 208803965228.000 |
| redis | ops_per_sec | 14134.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54451.000 |
| redis | used_memory_bytes | 126532032.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
