# Feed 压测报告：hybrid / rpc / sequential-page-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T18:53:35+08:00
- 采样时长：1m0.8048184s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 60750 | 60750 | 0 | 0 | 999.18 | 32.599 | 50.462 | 56.734 | 69.489 | 118.059 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：1215/60750
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 235 | 0.004 |
| mysql | 9091 | 0.150 |
| redis | 130591 | 2.150 |
| relation | 235 | 0.004 |

- Cold compute：60750（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 32197500 | 530.000 |
| merge_candidates | 55866915 | 919.620 |
| redis_commands | 364500 | 6.000 |
| redis_members | 61895745 | 1018.860 |
| redis_roundtrips | 60750 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 60750 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 60750 | 10.001 |
| counter | 235 | 8.461 |
| hydrate | 60750 | 21.066 |
| inbox | 60750 | 9.998 |
| merge_dedup | 60750 | 0.107 |
| relation | 235 | 9.634 |
| route | 60750 | 0.309 |
| total | 60750 | 31.626 |

## Redis 本轮边界增量

- Commands：450517；input：1157869683 bytes；output：7872833967 bytes
- Hits/Misses：32558013/11165；run hit rate：99.97%
- Evicted/Rejected：0/0；ops/s max：8131；safety epoch：3722 -> 3722

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.343 |
| client:loadtest | cpu_percent_total | 21.483 |
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
| docker:zg-canal | cpu_percent | 2.170 |
| docker:zg-canal | memory_percent | 4.260 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 0.600 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.480 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 150.050 |
| docker:zg-kafka | memory_percent | 7.600 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 26.410 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 400966.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 22.000 |
| mysql | threads_running | 4.000 |
| process:counter | cpu_percent | 3.099 |
| process:counter | cpu_seconds_total | 227.594 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46817280.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 0.774 |
| process:gateway | cpu_seconds_total | 5120.141 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 44785664.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 269.541 |
| process:knowpost | cpu_seconds_total | 13676.219 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 95756288.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.323 |
| process:relation | cpu_seconds_total | 38.406 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48463872.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.547 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38047744.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 2.323 |
| process:user-storage | cpu_seconds_total | 1811.375 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 44138496.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 238679424.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3722.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 1902039616.000 |
| redis | keyspace_misses | 507612.000 |
| redis | net_input_bytes | 81628419052.000 |
| redis | net_output_bytes | 480964999896.000 |
| redis | ops_per_sec | 8131.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 30854.000 |
| redis | used_memory_bytes | 104637808.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
