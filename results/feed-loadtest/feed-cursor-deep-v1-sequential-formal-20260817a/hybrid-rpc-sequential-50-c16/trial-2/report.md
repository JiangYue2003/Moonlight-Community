# Feed 压测报告：hybrid / rpc / sequential-50-c16

- Run ID：`feed-cursor-deep-v1-sequential-formal-20260817a`
- 开始时间：2026-08-17T18:28:58+08:00
- 采样时长：1m0.0796504s
- 并发：16
- 重复轮次：2 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 294650 | 294650 | 0 | 0 | 4905.80 | 3.124 | 4.784 | 5.485 | 7.023 | 14.215 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：5893/294650
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.001 |
| mysql | 953 | 0.003 |
| redis | 301496 | 1.023 |
| relation | 240 | 0.001 |

- Cold compute：294650（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 6299617 | 21.380 |
| merge_candidates | 17107379 | 58.060 |
| redis_commands | 4154565 | 14.100 |
| redis_members | 21497664 | 72.960 |
| redis_roundtrips | 583407 | 1.980 |
| tie_members | 5185840 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 294650 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 5893 | 1.186 |
| counter | 240 | 1.459 |
| cursor_decode | 288757 | 0.010 |
| cursor_seek | 288757 | 2.052 |
| hydrate | 294650 | 0.989 |
| inbox | 5893 | 1.184 |
| merge_dedup | 294650 | 0.006 |
| relation | 240 | 2.182 |
| route | 294650 | 0.011 |
| total | 294650 | 3.053 |

## Redis 本轮边界增量

- Commands：4486494；input：677501168 bytes；output：1952494258 bytes
- Hits/Misses：10460003/1614；run hit rate：99.98%
- Evicted/Rejected：0/0；ops/s max：80108；safety epoch：3706 -> 3706

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.327 |
| client:loadtest | cpu_percent_total | 69.231 |
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
| docker:zg-canal | memory_percent | 4.230 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.570 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.550 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 116.750 |
| docker:zg-kafka | memory_percent | 7.340 |
| docker:zg-kafka | pids | 117.000 |
| docker:zg-zk | cpu_percent | 45.920 |
| docker:zg-zk | memory_percent | 1.240 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 330447.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 25.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 7.729 |
| process:counter | cpu_seconds_total | 187.328 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46653440.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.000 |
| process:gateway | cpu_seconds_total | 4321.891 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 48218112.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 261.812 |
| process:knowpost | cpu_seconds_total | 11510.641 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 70553600.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.096 |
| process:relation | cpu_seconds_total | 31.594 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48586752.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.094 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38035456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 0.781 |
| process:user-storage | cpu_seconds_total | 1533.562 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 59912192.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 203606218.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3706.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596554.000 |
| redis | keyspace_hits | 1599822456.000 |
| redis | keyspace_misses | 415366.000 |
| redis | net_input_bytes | 68709001178.000 |
| redis | net_output_bytes | 411954541909.000 |
| redis | ops_per_sec | 80108.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 29377.000 |
| redis | used_memory_bytes | 102270064.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
