# Feed 压测报告：hybrid / gateway / sequential-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-formal-20260817a`
- 开始时间：2026-08-17T18:39:00+08:00
- 采样时长：1m0.449591s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 97850 | 97850 | 0 | 0 | 1618.88 | 19.013 | 27.454 | 30.915 | 40.700 | 81.344 |

## 分页正确性与准备成本

- 模式：`cursor-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1957/97850
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 241 | 0.002 |
| mysql | 1498 | 0.015 |
| redis | 101305 | 1.035 |
| relation | 241 | 0.002 |

- Cold compute：97850（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2092033 | 21.380 |
| merge_candidates | 5681171 | 58.060 |
| redis_commands | 1379685 | 14.100 |
| redis_members | 7139136 | 72.960 |
| redis_roundtrips | 193743 | 1.980 |
| tie_members | 1722160 | 17.600 |

| page cache source | requests |
|---|---:|
| bypass | 97850 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 1957 | 6.240 |
| counter | 241 | 6.725 |
| cursor_decode | 95893 | 0.011 |
| cursor_seek | 95893 | 12.065 |
| hydrate | 97850 | 6.164 |
| inbox | 1957 | 6.238 |
| merge_dedup | 97850 | 0.007 |
| relation | 241 | 7.984 |
| route | 97850 | 0.138 |
| total | 97850 | 18.274 |

## Redis 本轮边界增量

- Commands：1497280；input：225531852 bytes；output：648411278 bytes
- Hits/Misses：3477093/2012；run hit rate：99.94%
- Evicted/Rejected：0/0；ops/s max：28851；safety epoch：3714 -> 3714

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 3.402 |
| client:loadtest | cpu_percent_total | 54.436 |
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
| docker:zg-canal | cpu_percent | 2.080 |
| docker:zg-canal | memory_percent | 4.240 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.370 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.550 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 109.650 |
| docker:zg-kafka | memory_percent | 7.350 |
| docker:zg-kafka | pids | 116.000 |
| docker:zg-zk | cpu_percent | 46.720 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 346757.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 26.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 6.962 |
| process:counter | cpu_seconds_total | 209.391 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46280704.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 236.920 |
| process:gateway | cpu_seconds_total | 4855.656 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52490240.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 174.866 |
| process:knowpost | cpu_seconds_total | 12536.766 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 72159232.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.094 |
| process:relation | cpu_seconds_total | 34.734 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49164288.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 4.219 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38023168.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 82.632 |
| process:user-storage | cpu_seconds_total | 1721.859 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 57143296.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 232153405.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3714.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596552.000 |
| redis | keyspace_hits | 1666171533.000 |
| redis | keyspace_misses | 433557.000 |
| redis | net_input_bytes | 73012559364.000 |
| redis | net_output_bytes | 424337769252.000 |
| redis | ops_per_sec | 28851.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 29980.000 |
| redis | used_memory_bytes | 103231256.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
