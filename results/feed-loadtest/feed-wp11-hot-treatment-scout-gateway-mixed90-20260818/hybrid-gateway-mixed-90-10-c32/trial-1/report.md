# Feed 压测报告：hybrid / gateway / mixed-90-10-c32

- Run ID：`feed-wp11-hot-treatment-scout-gateway-mixed90-20260818`
- 开始时间：2026-08-18T22:45:29+08:00
- 采样时长：0s
- 并发：32
- 重复轮次：1 / 1
- 报告完整：false

| stage | total | success | failed | timeout | QPS | P50(ms) | P90(ms) | P95(ms) | P99(ms) | Max(ms) |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---:|

## Mutation checkpoint

- Checkpoint：`ca37fe4c148024bb3d8cb46573e4616ebb1e93f5a6c903e67a1336f6dd31b5d7`；baseline：`2bdf1b5730304cd6cef84adbed8ff3e6541d89af27b522951096fda1b86cee54`
- 创建时间：2026-08-18T07:15:15.962547Z；工具：`feed-loadtest-mutation-v1`
- 成功发布帖子：0
- 测量前恢复：complete=true；耗时=4.9680869s；删除帖子/Outbox=0/0；safety epoch=8251
- 预热后恢复：complete=true；耗时=7.6596186s；删除帖子/Outbox=32/64；safety epoch=8316

## 缺失指标

- mutation_preparation

## 说明

- Mutation preparation error: prepare mutation c32 trial 1: mutation warmup: warmup c32 trial 1 contained failed or empty stages
