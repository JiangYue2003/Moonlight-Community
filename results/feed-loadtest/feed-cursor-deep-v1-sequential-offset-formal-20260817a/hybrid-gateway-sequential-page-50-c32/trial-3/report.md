# Feed 压测报告：hybrid / gateway / sequential-page-50-c32

- Run ID：`feed-cursor-deep-v1-sequential-offset-formal-20260817a`
- 开始时间：2026-08-17T19:01:15+08:00
- 采样时长：1m1.4319756s
- 并发：32
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 39500 | 39500 | 0 | 0 | 643.01 | 46.146 | 80.116 | 89.294 | 106.436 | 155.606 |

## 分页正确性与准备成本

- 模式：`page-sequential`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：790/39500
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 230 | 0.006 |
| mysql | 9135 | 0.231 |
| redis | 88135 | 2.231 |
| relation | 230 | 0.006 |

- Cold compute：39500（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 20935000 | 530.000 |
| merge_candidates | 36324990 | 919.620 |
| redis_commands | 237000 | 6.000 |
| redis_members | 40244970 | 1018.860 |
| redis_roundtrips | 39500 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 39500 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 39500 | 15.359 |
| counter | 230 | 12.566 |
| hydrate | 39500 | 31.718 |
| inbox | 39500 | 15.356 |
| merge_dedup | 39500 | 0.095 |
| relation | 230 | 14.789 |
| route | 39500 | 0.744 |
| total | 39500 | 48.050 |

## Redis 本轮边界增量

- Commands：300960；input：754238155 bytes；output：5118185639 bytes
- Hits/Misses：21165999/13016；run hit rate：99.94%
- Evicted/Rejected：0/0；ops/s max：5997；safety epoch：3728 -> 3728

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 1.307 |
| client:loadtest | cpu_percent_total | 20.907 |
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
| docker:zg-canal | cpu_percent | 1.800 |
| docker:zg-canal | memory_percent | 4.280 |
| docker:zg-canal | pids | 84.000 |
| docker:zg-es | cpu_percent | 2.220 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 152.000 |
| docker:zg-etcd | cpu_percent | 4.530 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 158.150 |
| docker:zg-kafka | memory_percent | 7.560 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 48.670 |
| docker:zg-zk | memory_percent | 1.280 |
| docker:zg-zk | pids | 104.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 453646.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 20.000 |
| mysql | threads_running | 5.000 |
| process:counter | cpu_percent | 3.868 |
| process:counter | cpu_seconds_total | 234.609 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46313472.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 89.788 |
| process:gateway | cpu_seconds_total | 5379.641 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 52846592.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 205.758 |
| process:knowpost | cpu_seconds_total | 14412.766 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 93720576.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.420 |
| process:relation | cpu_seconds_total | 40.812 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48898048.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 1.547 |
| process:search | cpu_seconds_total | 4.719 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38055936.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 40.250 |
| process:user-storage | cpu_seconds_total | 1909.000 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 52776960.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 241038440.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3728.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596568.000 |
| redis | keyspace_hits | 2067645962.000 |
| redis | keyspace_misses | 574247.000 |
| redis | net_input_bytes | 87523616941.000 |
| redis | net_output_bytes | 521009461098.000 |
| redis | ops_per_sec | 5997.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31315.000 |
| redis | used_memory_bytes | 104105632.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
