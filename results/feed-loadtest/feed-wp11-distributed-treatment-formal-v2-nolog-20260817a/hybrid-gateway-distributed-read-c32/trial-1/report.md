# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-distributed-treatment-formal-v2-nolog-20260817a`
- 开始时间：2026-08-17T00:06:00+08:00
- 采样时长：1m0.0376458s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 95817 | 95817 | 0 | 0 | 1596.15 | 13.383 | 42.499 | 47.176 | 58.813 | 112.909 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| cache | 98144 | 1.024 |
| counter | 9027 | 0.094 |
| mysql | 3 | 0.000 |
| redis | 96549 | 1.008 |
| relation | 9027 | 0.094 |

- Cold compute：17614（0.184 / successful read）

| page cache source | requests |
|---|---:|
| l1_fresh | 49729 |
| l1_stale | 786 |
| l2_fresh | 38790 |
| l2_stale | 6512 |
| miss | 0 |

- L1+L2 Fresh ratio：92.38%
- Refresh max：queue=51 active=32 pending=83

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 17614 | 16.360 |
| counter | 9027 | 18.237 |
| hydrate | 17614 | 16.204 |
| inbox | 17614 | 16.358 |
| merge_dedup | 17614 | 0.003 |
| relation | 9027 | 20.339 |
| route | 17614 | 19.937 |
| total | 95817 | 14.243 |

## Redis 本轮边界增量

- Commands：288132；input：49621529 bytes；output：107399435 bytes
- Hits/Misses：464309/10496；run hit rate：97.79%
- Evicted/Rejected：0/0；ops/s max：14039；safety epoch：554 -> 554

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 8.094 |
| client:loadtest | cpu_percent_total | 129.502 |
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
| docker:zg-es | cpu_percent | 2.470 |
| docker:zg-es | memory_percent | 12.040 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 5.940 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 195.730 |
| docker:zg-kafka | memory_percent | 8.770 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 54.080 |
| docker:zg-zk | memory_percent | 1.690 |
| docker:zg-zk | pids | 105.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 23438417.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 51.000 |
| mysql | threads_running | 6.000 |
| process:counter | cpu_percent | 49.430 |
| process:counter | cpu_seconds_total | 271.672 |
| process:counter | pid | 10668.000 |
| process:counter | process_start_ms | 1786894716793.000 |
| process:counter | rss_bytes | 55078912.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 479.763 |
| process:gateway | cpu_seconds_total | 1342.812 |
| process:gateway | pid | 7032.000 |
| process:gateway | process_start_ms | 1786894734161.000 |
| process:gateway | rss_bytes | 54714368.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 297.264 |
| process:knowpost | cpu_seconds_total | 3140.281 |
| process:knowpost | pid | 31560.000 |
| process:knowpost | process_start_ms | 1786894724886.000 |
| process:knowpost | rss_bytes | 106094592.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 95.472 |
| process:relation | cpu_seconds_total | 301.078 |
| process:relation | pid | 2532.000 |
| process:relation | process_start_ms | 1786894720562.000 |
| process:relation | rss_bytes | 57942016.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 0.719 |
| process:search | pid | 30524.000 |
| process:search | process_start_ms | 1786894729124.000 |
| process:search | rss_bytes | 36384768.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 129.374 |
| process:user-storage | cpu_seconds_total | 331.562 |
| process:user-storage | pid | 18428.000 |
| process:user-storage | process_start_ms | 1786894713048.000 |
| process:user-storage | rss_bytes | 54194176.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 211624918.000 |
| redis | connected_clients | 273.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 554.000 |
| redis | hit_rate | 0.994 |
| redis | keys | 706294.000 |
| redis | keyspace_hits | 1011455251.000 |
| redis | keyspace_misses | 6183977.000 |
| redis | net_input_bytes | 43700090206.000 |
| redis | net_output_bytes | 190624587698.000 |
| redis | ops_per_sec | 14039.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 49599.000 |
| redis | used_memory_bytes | 125179096.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
