# Feed 压测报告：hybrid / gateway / sequential-page-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:58:40+08:00
- 采样时长：1m1.7142417s
- 并发：32
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 41550 | 41550 | 0 | 0 | 673.29 | 44.241 | 77.387 | 87.197 | 105.857 | 163.159 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：831/41550
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 236 | 0.006 |
| mysql | 9276 | 0.223 |
| redis | 92376 | 2.223 |
| relation | 236 | 0.006 |

- Cold compute：41550（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 22021500 | 530.000 |
| merge_candidates | 38210211 | 919.620 |
| redis_commands | 249300 | 6.000 |
| redis_members | 42333633 | 1018.860 |
| redis_roundtrips | 41550 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 41550 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 41550 | 14.681 |
| counter | 236 | 12.219 |
| hydrate | 41550 | 30.548 |
| inbox | 41550 | 14.678 |
| merge_dedup | 41550 | 0.095 |
| relation | 236 | 13.766 |
| route | 41550 | 0.679 |
| total | 41550 | 46.142 |

## Redis 本轮边界增量

- Commands：315284；input：793102530 bytes；output：5383992475 bytes
- Hits/Misses：22265531/12461；run hit rate：99.94%
- Evicted/Rejected：0/0；ops/s max：6165；safety epoch：3726 -> 3726

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.402 |
| client:loadtest | cpu_percent_total | 22.432 |
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
| docker:zg-canal | cpu_percent | 2.160 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 0.440 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.810 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 163.050 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 44.040 |
| docker:zg-zk | memory_percent | 1.070 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 431442.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 2.325 |
| process:counter | cpu_seconds_total | 232.406 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46256128.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 103.493 |
| process:gateway | cpu_seconds_total | 5288.484 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52989952.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 220.086 |
| process:knowpost | cpu_seconds_total | 14183.859 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 90193920.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 3.160 |
| process:relation | cpu_seconds_total | 39.719 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49074176.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.609 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38055936.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 30.982 |
| process:user-storage | cpu_seconds_total | 1877.078 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 52445184.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 240372523.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3726.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2021617440.000 |
| redis | keyspace_misses | 543657.000 |
| redis | net_input_bytes | 85882278756.000 |
| redis | net_output_bytes | 509879181418.000 |
| redis | ops_per_sec | 6165.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31161.000 |
| redis | used_memory_bytes | 104974344.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
