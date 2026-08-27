# Feed 压测报告：hybrid / gateway / distributed-read-c32

- Run ID：`feed-wp11-high-control-fixed-gateway-scout-valid-nolog-20260817a`
- 开始时间：2026-08-17T01:10:46+08:00
- 采样时长：10.0134369s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 18866 | 18866 | 0 | 0 | 1884.30 | 16.797 | 22.780 | 24.897 | 30.301 | 41.614 |

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 13770 | 0.730 |
| mysql | 0 | 0.000 |
| redis | 56598 | 3.000 |
| relation | 18866 | 1.000 |

- Cold compute：18866（1.000 / successful read）

| page cache source | requests |
|---|---:|
| bypass | 18866 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 18866 | 2.865 |
| counter | 13770 | 3.580 |
| hydrate | 18866 | 2.873 |
| inbox | 18866 | 2.952 |
| merge_dedup | 18866 | 0.002 |
| relation | 18866 | 4.272 |
| route | 18866 | 2.626 |
| total | 18866 | 15.610 |

## Redis 本轮边界增量

- Commands：216413；input：16673067 bytes；output：36071672 bytes
- Hits/Misses：322747/36166；run hit rate：89.92%
- Evicted/Rejected：0/0；ops/s max：23802；safety epoch：597 -> 597

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.131 |
| client:loadtest | cpu_percent_total | 50.089 |
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
| docker:zg-canal | cpu_percent | 1.600 |
| docker:zg-canal | memory_percent | 7.830 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.810 |
| docker:zg-es | memory_percent | 12.020 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.460 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 17.000 |
| docker:zg-kafka | cpu_percent | 117.540 |
| docker:zg-kafka | memory_percent | 8.330 |
| docker:zg-kafka | pids | 95.000 |
| docker:zg-zk | cpu_percent | 41.060 |
| docker:zg-zk | memory_percent | 1.550 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 3104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 3104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 27672662.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 104.571 |
| process:counter | cpu_seconds_total | 202.109 |
| process:counter | pid | 17060.000 |
| process:counter | process_start_ms | 1786899211176.000 |
| process:counter | rss_bytes | 54464512.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 194.211 |
| process:gateway | cpu_seconds_total | 494.078 |
| process:gateway | pid | 8504.000 |
| process:gateway | process_start_ms | 1786899226856.000 |
| process:gateway | rss_bytes | 52473856.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 305.139 |
| process:knowpost | cpu_seconds_total | 1573.219 |
| process:knowpost | pid | 6176.000 |
| process:knowpost | process_start_ms | 1786899219072.000 |
| process:knowpost | rss_bytes | 82636800.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 213.789 |
| process:relation | cpu_seconds_total | 1243.516 |
| process:relation | pid | 15340.000 |
| process:relation | process_start_ms | 1786899214850.000 |
| process:relation | rss_bytes | 58064896.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 0.625 |
| process:search | pid | 19300.000 |
| process:search | process_start_ms | 1786899223294.000 |
| process:search | rss_bytes | 35987456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 78.781 |
| process:user-storage | cpu_seconds_total | 183.844 |
| process:user-storage | pid | 10160.000 |
| process:user-storage | process_start_ms | 1786899207729.000 |
| process:user-storage | rss_bytes | 48816128.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 276091609.000 |
| redis | connected_clients | 190.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 597.000 |
| redis | hit_rate | 0.988 |
| redis | keys | 715894.000 |
| redis | keyspace_hits | 1107921714.000 |
| redis | keyspace_misses | 14037684.000 |
| redis | net_input_bytes | 51676611222.000 |
| redis | net_output_bytes | 205312432601.000 |
| redis | ops_per_sec | 23802.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 53435.000 |
| redis | used_memory_bytes | 130186056.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
