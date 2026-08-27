# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-fixed-gateway-formal-nolog-20260817a`
- 开始时间：2026-08-17T01:29:16+08:00
- 采样时长：1m0.0257915s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 91775 | 91775 | 0 | 0 | 1529.10 | 7.734 | 21.867 | 25.114 | 31.492 | 69.815 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 96775 | 1.054 |
| counter | 9100 | 0.099 |
| mysql | 5 | 0.000 |
| redis | 96179 | 1.048 |
| relation | 9100 | 0.099 |

- Cold compute：17604（0.192 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 47504 |
| l1_stale | 308 |
| l2_fresh | 37163 |
| l2_stale | 6800 |
| miss | 0 |

- L1+L2 Fresh ratio：92.25%
- Refresh max：queue=0 active=19 pending=19

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 17604 | 7.025 |
| counter | 9100 | 8.726 |
| hydrate | 17604 | 6.984 |
| inbox | 17604 | 7.023 |
| merge_dedup | 17604 | 0.004 |
| relation | 9100 | 10.850 |
| route | 17604 | 10.153 |
| total | 91775 | 6.115 |

## Redis 本轮边界增量

- Commands：294398；input：50075863 bytes；output：107127898 bytes
- Hits/Misses：463330/10559；run hit rate：97.77%
- Evicted/Rejected：0/0；ops/s max：10328；safety epoch：609 -> 609

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 6.810 |
| client:loadtest | cpu_percent_total | 108.964 |
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
| docker:zg-canal | cpu_percent | 0.200 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.820 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.880 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 191.140 |
| docker:zg-kafka | memory_percent | 8.820 |
| docker:zg-kafka | pids | 127.000 |
| docker:zg-zk | cpu_percent | 11.390 |
| docker:zg-zk | memory_percent | 1.720 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 29235859.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 55.910 |
| process:counter | cpu_seconds_total | 110.797 |
| process:counter | pid | 32692.000 |
| process:counter | process_start_ms | 1786900892409.000 |
| process:counter | rss_bytes | 53600256.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 409.415 |
| process:gateway | cpu_seconds_total | 618.266 |
| process:gateway | pid | 27008.000 |
| process:gateway | process_start_ms | 1786900909619.000 |
| process:gateway | rss_bytes | 52367360.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 267.904 |
| process:knowpost | cpu_seconds_total | 1522.094 |
| process:knowpost | pid | 5676.000 |
| process:knowpost | process_start_ms | 1786900901519.000 |
| process:knowpost | rss_bytes | 99127296.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 88.525 |
| process:relation | cpu_seconds_total | 139.812 |
| process:relation | pid | 6032.000 |
| process:relation | process_start_ms | 1786900896944.000 |
| process:relation | rss_bytes | 56139776.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.281 |
| process:search | pid | 31360.000 |
| process:search | process_start_ms | 1786900905870.000 |
| process:search | rss_bytes | 35573760.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 107.070 |
| process:user-storage | cpu_seconds_total | 152.016 |
| process:user-storage | pid | 27996.000 |
| process:user-storage | process_start_ms | 1786900887854.000 |
| process:user-storage | rss_bytes | 49979392.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 294454306.000 |
| redis | connected_clients | 216.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 609.000 |
| redis | hit_rate | 0.985 |
| redis | keys | 709964.000 |
| redis | keyspace_hits | 1135355400.000 |
| redis | keyspace_misses | 16984507.000 |
| redis | net_input_bytes | 53388953609.000 |
| redis | net_output_bytes | 209033962798.000 |
| redis | ops_per_sec | 10328.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 54595.000 |
| redis | used_memory_bytes | 128070560.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
