# Feed 压测报告：hybrid / gateway / page5-c16

- Run ID：`feed-cursor-deep-v1-gateway-page5-recheck-a`
- 开始时间：2026-08-17T19:04:46+08:00
- 采样时长：10.0068748s
- 并发：16
- 重复轮次：1 / 1
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 19868 | 19868 | 0 | 0 | 1985.64 | 7.471 | 11.520 | 12.547 | 14.967 | 21.988 |

## 分页正确性与准备成本

- 模式：`page`；目标页：5
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 40 | 0.002 |
| mysql | 122 | 0.006 |
| redis | 39858 | 2.006 |
| relation | 40 | 0.002 |

- Cold compute：19868（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2384160 | 120.000 |
| merge_candidates | 4788188 | 241.000 |
| redis_commands | 119208 | 6.000 |
| redis_members | 12338028 | 621.000 |
| redis_roundtrips | 19868 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 19868 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 19868 | 3.920 |
| counter | 40 | 2.692 |
| hydrate | 19868 | 3.240 |
| inbox | 19868 | 3.919 |
| merge_dedup | 19868 | 0.028 |
| relation | 40 | 2.888 |
| route | 19868 | 0.061 |
| total | 19868 | 7.294 |

## Redis 本轮边界增量

- Commands：144016；input：92839156 bytes；output：910122313 bytes
- Hits/Misses：2504378/242；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：14678；safety epoch：3731 -> 3731

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 2.528 |
| client:loadtest | cpu_percent_total | 40.441 |
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
| docker:zg-canal | cpu_percent | 0.140 |
| docker:zg-canal | memory_percent | 4.270 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.330 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 5.090 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 121.860 |
| docker:zg-kafka | memory_percent | 7.110 |
| docker:zg-kafka | pids | 94.000 |
| docker:zg-zk | cpu_percent | 31.600 |
| docker:zg-zk | memory_percent | 1.120 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 454045.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 32.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 6.890 |
| process:counter | cpu_seconds_total | 238.203 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 44224512.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 129.854 |
| process:gateway | cpu_seconds_total | 5397.562 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 50823168.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 212.635 |
| process:knowpost | cpu_seconds_total | 14449.172 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 73334784.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 0.000 |
| process:relation | cpu_seconds_total | 41.000 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 49496064.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.000 |
| process:search | cpu_seconds_total | 4.750 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38035456.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 54.749 |
| process:user-storage | cpu_seconds_total | 1916.922 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 50728960.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 241517888.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3731.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596192.000 |
| redis | keyspace_hits | 2071429950.000 |
| redis | keyspace_misses | 576058.000 |
| redis | net_input_bytes | 87683019313.000 |
| redis | net_output_bytes | 522381032376.000 |
| redis | ops_per_sec | 14678.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 31475.000 |
| redis | used_memory_bytes | 101468840.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
