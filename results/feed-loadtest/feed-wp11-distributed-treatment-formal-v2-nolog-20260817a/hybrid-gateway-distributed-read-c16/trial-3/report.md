# Feed 压测报告：hybrid / gateway / distributed-read-c16

- Run ID：`feed-wp11-distributed-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:04:01+08:00
- 采样时长：1m0.0218996s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 85121 | 85121 | 0 | 0 | 1418.31 | 9.110 | 22.163 | 25.246 | 32.255 | 63.470 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 94998 | 1.116 |
| counter | 9133 | 0.107 |
| mysql | 2 | 0.000 |
| redis | 94415 | 1.109 |
| relation | 9133 | 0.107 |

- Cold compute：17441（0.205 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 42173 |
| l1_stale | 273 |
| l2_fresh | 35603 |
| l2_stale | 7072 |
| miss | 0 |

- L1+L2 Fresh ratio：91.37%
- Refresh max：queue=0 active=14 pending=14

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 17441 | 6.979 |
| counter | 9133 | 8.494 |
| hydrate | 17441 | 6.856 |
| inbox | 17441 | 6.977 |
| merge_dedup | 17441 | 0.003 |
| relation | 9133 | 11.209 |
| route | 17441 | 10.346 |
| total | 85121 | 6.361 |

## Redis 本轮边界增量

- Commands：290417；input：49467656 bytes；output：104925417 bytes
- Hits/Misses：457828/10586；run hit rate：97.74%
- Evicted/Rejected：0/0；ops/s max：8507；safety epoch：553 -> 553

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 7.092 |
| client:loadtest | cpu_percent_total | 113.474 |
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
| docker:zg-canal | cpu_percent | 0.170 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 3.220 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.550 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 187.930 |
| docker:zg-kafka | memory_percent | 8.760 |
| docker:zg-kafka | pids | 125.000 |
| docker:zg-zk | cpu_percent | 50.710 |
| docker:zg-zk | memory_percent | 1.450 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23425497.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 26.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 57.347 |
| process:counter | cpu_seconds_total | 246.297 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 54255616.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 446.638 |
| process:gateway | cpu_seconds_total | 1098.422 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 53346304.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 305.767 |
| process:knowpost | cpu_seconds_total | 2968.641 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 103993344.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 81.853 |
| process:relation | cpu_seconds_total | 264.656 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 57704448.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.772 |
| process:search | cpu_seconds_total | 0.672 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36376576.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 90.751 |
| process:user-storage | cpu_seconds_total | 278.641 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 60071936.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 211217910.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 553.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 705241.000 |
| redis | keyspace_hits | 1010905154.000 |
| redis | keyspace_misses | 6167879.000 |
| redis | net_input_bytes | 43636667851.000 |
| redis | net_output_bytes | 190501406386.000 |
| redis | ops_per_sec | 8507.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49480.000 |
| redis | used_memory_bytes | 123897184.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
