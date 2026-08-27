# Feed 压测报告：hybrid / gateway / page50-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:14:12+08:00
- 采样时长：1m0.0123895s
- 并发：16
- 重复轮次：3 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 29859 | 29859 | 0 | 0 | 497.56 | 31.204 | 42.173 | 45.645 | 52.613 | 73.887 |

## 分页正确性与准备成本

- 模式：`page`；目标页：50
- 准备请求：0；准备耗时：0s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.008 |
| mysql | 5316 | 0.178 |
| redis | 65034 | 2.178 |
| relation | 240 | 0.008 |

- Cold compute：29859（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 30456180 | 1020.000 |
| merge_candidates | 44758641 | 1499.000 |
| redis_commands | 179154 | 6.000 |
| redis_members | 44788500 | 1500.000 |
| redis_roundtrips | 29859 | 1.000 |
| tie_members | 0 | 0.000 |

| page cache source | requests |
|---|---:|
| bypass | 29859 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| bigv_pipeline | 29859 | 6.869 |
| counter | 240 | 5.545 |
| hydrate | 29859 | 23.550 |
| inbox | 29859 | 6.866 |
| merge_dedup | 29859 | 0.171 |
| relation | 240 | 6.366 |
| route | 29859 | 0.373 |
| total | 29859 | 31.154 |

## Redis 本轮边界增量

- Commands：236822；input：1082532990 bytes；output：6861119333 bytes
- Hits/Misses：30634971/7817；run hit rate：99.97%
- Evicted/Rejected：0/0；ops/s max：4486；safety epoch：3695 -> 3695

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 0.823 |
| client:loadtest | cpu_percent_total | 13.174 |
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
| docker:zg-canal | cpu_percent | 2.430 |
| docker:zg-canal | memory_percent | 4.220 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.410 |
| docker:zg-es | memory_percent | 12.340 |
| docker:zg-es | pids | 151.000 |
| docker:zg-etcd | cpu_percent | 4.540 |
| docker:zg-etcd | memory_percent | 0.260 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 133.890 |
| docker:zg-kafka | memory_percent | 7.590 |
| docker:zg-kafka | pids | 126.000 |
| docker:zg-zk | cpu_percent | 44.910 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 286006.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 23.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 5.417 |
| process:counter | cpu_seconds_total | 155.156 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 47415296.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 54.229 |
| process:gateway | cpu_seconds_total | 3260.531 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51027968.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 266.248 |
| process:knowpost | cpu_seconds_total | 10116.438 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 84029440.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 2.324 |
| process:relation | cpu_seconds_total | 26.438 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48087040.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.775 |
| process:search | cpu_seconds_total | 3.734 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38162432.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 32.658 |
| process:user-storage | cpu_seconds_total | 1181.422 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 51884032.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 181532165.000 |
| redis | connected_clients | 103.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3695.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596598.000 |
| redis | keyspace_hits | 1474425664.000 |
| redis | keyspace_misses | 334739.000 |
| redis | net_input_bytes | 62787440247.000 |
| redis | net_output_bytes | 387226977038.000 |
| redis | ops_per_sec | 4486.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28491.000 |
| redis | used_memory_bytes | 102786600.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
