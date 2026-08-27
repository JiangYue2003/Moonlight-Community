# Feed 压测报告：hybrid / gateway / cursor-page50-c16

- Run ID：`feed-cursor-deep-v1-gateway-formal-20260817a`
- 开始时间：2026-08-17T18:19:15+08:00
- 采样时长：1m0.0137664s
- 并发：16
- 重复轮次：1 / 3
- 报告完整：true

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| read | 118381 | 118381 | 0 | 0 | 1972.86 | 7.839 | 11.031 | 12.624 | 15.729 | 24.667 |

## 分页正确性与准备成本

- 模式：`cursor`；目标页：50
- 准备请求：980；准备耗时：1.3428519s（不计目标页 stage 延迟）
- 完成序列/读取页：0/0
- 重复/Oracle不匹配/游标循环/提前结束：0/0/0/0
- Oracle/Observed hash：`c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc` / `c816b4ef0d68bd218cd33d8e164cdbfd9e12bdf4438e24a1f3afa1b54c6acfcc`

## Feed 冷路径指标（本次运行增量）

| dependency | calls | calls / successful read |
|---|---:|---:|
| counter | 240 | 0.002 |
| mysql | 111 | 0.001 |
| redis | 118492 | 1.001 |
| relation | 240 | 0.002 |

- Cold compute：118381（1.000 / successful read）

### 分页工作量

| pagination work | total | per successful read |
|---|---:|---:|
| hydrate_ids | 2486001 | 21.000 |
| merge_candidates | 2604382 | 22.000 |
| redis_commands | 1538953 | 13.000 |
| redis_members | 2841144 | 24.000 |
| redis_roundtrips | 236762 | 2.000 |
| tie_members | 236762 | 2.000 |

| page cache source | requests |
|---|---:|
| bypass | 118381 |

- L1+L2 Fresh ratio：0.00%
- Refresh max：queue=0 active=0 pending=0

| feed stage | calls(all outcomes) | success mean(ms) |
|---|---:|---:|
| counter | 240 | 2.650 |
| cursor_decode | 118381 | 0.011 |
| cursor_seek | 118381 | 4.334 |
| hydrate | 118381 | 2.368 |
| merge_dedup | 118381 | 0.005 |
| relation | 240 | 3.623 |
| route | 118381 | 0.047 |
| total | 118381 | 6.771 |

## Redis 本轮边界增量

- Commands：1693205；input：259508555 bytes；output：532238022 bytes
- Hits/Misses：4032131/353；run hit rate：99.99%
- Evicted/Rejected：0/0；ops/s max：31869；safety epoch：3699 -> 3699

## 资源采样峰值

| source | metric | max |
|---|---|---:|
| client:loadtest | cpu_percent_normalized | 4.623 |
| client:loadtest | cpu_percent_total | 73.967 |
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
| docker:zg-canal | cpu_percent | 2.100 |
| docker:zg-canal | memory_percent | 4.220 |
| docker:zg-canal | pids | 78.000 |
| docker:zg-es | cpu_percent | 2.420 |
| docker:zg-es | memory_percent | 12.360 |
| docker:zg-es | pids | 153.000 |
| docker:zg-etcd | cpu_percent | 4.890 |
| docker:zg-etcd | memory_percent | 0.270 |
| docker:zg-etcd | pids | 16.000 |
| docker:zg-kafka | cpu_percent | 152.520 |
| docker:zg-kafka | memory_percent | 7.610 |
| docker:zg-kafka | pids | 121.000 |
| docker:zg-zk | cpu_percent | 48.690 |
| docker:zg-zk | memory_percent | 1.020 |
| docker:zg-zk | pids | 82.000 |
| kafka | current_offset_total | 4104.000 |
| kafka | lag_max | 0.000 |
| kafka | lag_total | 0.000 |
| kafka | log_end_offset_total | 4104.000 |
| kafka | partitions | 1.000 |
| mysql | questions | 323298.000 |
| mysql | slow_queries | 0.000 |
| mysql | threads_connected | 24.000 |
| mysql | threads_running | 3.000 |
| process:counter | cpu_percent | 13.955 |
| process:counter | cpu_seconds_total | 163.125 |
| process:counter | pid | 7060.000 |
| process:counter | process_start_ms | 1786955465447.000 |
| process:counter | rss_bytes | 46399488.000 |
| process:counter | running | 1.000 |
| process:gateway | cpu_percent | 270.739 |
| process:gateway | cpu_seconds_total | 3508.391 |
| process:gateway | pid | 27984.000 |
| process:gateway | process_start_ms | 1786955487497.000 |
| process:gateway | rss_bytes | 51572736.000 |
| process:gateway | running | 1.000 |
| process:knowpost | cpu_percent | 214.244 |
| process:knowpost | cpu_seconds_total | 10564.844 |
| process:knowpost | pid | 28168.000 |
| process:knowpost | process_start_ms | 1786955477430.000 |
| process:knowpost | rss_bytes | 71020544.000 |
| process:knowpost | running | 1.000 |
| process:relation | cpu_percent | 5.422 |
| process:relation | cpu_seconds_total | 28.078 |
| process:relation | pid | 15024.000 |
| process:relation | process_start_ms | 1786955470022.000 |
| process:relation | rss_bytes | 48959488.000 |
| process:relation | running | 1.000 |
| process:search | cpu_percent | 0.774 |
| process:search | cpu_seconds_total | 3.875 |
| process:search | pid | 21780.000 |
| process:search | process_start_ms | 1786955483082.000 |
| process:search | rss_bytes | 38166528.000 |
| process:search | running | 1.000 |
| process:user-storage | cpu_percent | 99.013 |
| process:user-storage | cpu_seconds_total | 1261.047 |
| process:user-storage | pid | 4640.000 |
| process:user-storage | process_start_ms | 1786955461190.000 |
| process:user-storage | rss_bytes | 64880640.000 |
| process:user-storage | running | 1.000 |
| redis | blocked_clients | 0.000 |
| redis | commands_total | 184171010.000 |
| redis | connected_clients | 104.000 |
| redis | evicted_keys | 0.000 |
| redis | feed_safety_epoch | 3699.000 |
| redis | hit_rate | 1.000 |
| redis | keys | 596581.000 |
| redis | keyspace_hits | 1554097317.000 |
| redis | keyspace_misses | 396125.000 |
| redis | net_input_bytes | 65752693582.000 |
| redis | net_output_bytes | 404635523930.000 |
| redis | ops_per_sec | 31869.000 |
| redis | rejected_connections | 0.000 |
| redis | uptime_seconds | 28794.000 |
| redis | used_memory_bytes | 102097056.000 |

## 说明

- SLA values are reference lines, not pass/fail gates.
