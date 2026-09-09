# Research: load-test observation gate before Stage 5

- Query: Which historical performance/load tests are still executable and strong evidence that the Stage 4 topology runs correctly, and how should they become the observation gate before legacy API retirement?
- Scope: internal
- Date: 2026-09-09

## Findings

### Recommendation

Use the current Feed harness as the main Stage 5 observation gate, but use a
compressed set of stable points rather than replaying every exploratory matrix.
The harness is unusually strong runtime evidence because it crosses service
discovery, the migrated User/Relation/Counter/KnowPost processes, Kafka fanout,
Redis/MySQL, and either direct RPC or the real Gateway HTTP/JWT path. The local
launcher also builds the migrated `services/*/cmd/*` entry points before each
observation run (`scripts/start-feed-local-services.ps1:35-43,251-280`).

The gate must have two independent parts:

1. **Validity gate (hard, already automated):** every selected trial exits zero,
   writes `complete=true`, has zero failed requests/timeouts, and has no hard
   missing metric. Hard failures include service/container identity changes or
   restarts, unhealthy containers, Feed degradation, load-generator saturation,
   Redis eviction/rejection or identity/reset anomalies, Kafka recovery failure,
   missing MySQL/Redis/Prometheus/process evidence, and Cursor correctness errors
   (`cmd/loadtest/monitor.go:512-566,833-864`; `cmd/loadtest/feed_loadtest.go:1030-1068`).
2. **Performance non-regression gate (new policy, not currently automated):**
   compare three-trial medians with the matching historical card below. The
   current code explicitly says SLA values are reference lines rather than
   pass/fail gates (`cmd/loadtest/feed_loadtest.go:1042-1044`), so the repository
   has no existing universal numeric regression threshold. For this directory
   migration, a pragmatic proposed review threshold is median successful QPS no
   worse than 20% and median P95 no worse than 25% versus the matching historical
   topology/cardinality/cache-state/entry baseline. A failure should trigger one
   clean rerun after checking host contention; it should not be averaged away.

Do not require a Control/Treatment A/B replay. Stage 4 changed process ownership,
not Feed feature behavior. Re-running only the current `treatment` stable points
directly tests non-regression and avoids doubling the observation time. Historical
guidance also says future formal performance validation should use RPC stable
points plus only a necessary adjacent point, and not repeat the mutation Gateway
matrix (`docs/superpowers/reports/2026-08-27-feed-wp11-mutation-rpc-ab.md:104-121`).

### Currently runnable surface

The authoritative executable is:

```powershell
go run ./cmd/loadtest -f cmd/loadtest/load_test.yaml <flags>
```

Its current phases are `setup`, `resume-setup`, `setup-cursor-deep`,
`seed-cursor-deep`, `smoke`, `reset-feed`, `seed-feed`,
`create-mutation-checkpoint`, `verify-mutation-checkpoint`,
`restore-mutation-checkpoint`, `run`, `compare`, `snapshot`, `cleanup`, and
`all`; current scenarios and flags are defined at
`cmd/loadtest/feed_loadtest.go:408-438`.

Current orchestration scripts:

- `cmd/loadtest/run_feed_matrix.ps1`: pure reads through RPC and optionally
  Gateway; covers `hot-read`, `distributed-read`, `deep-page`, and optional
  `burst-read`/`steady-read` (`run_feed_matrix.ps1:229-252`).
- `cmd/loadtest/run_feed_mutation_matrix.ps1`: checkpointed `publish`,
  `mixed-90-10`, and `mixed-80-20`; refuses incomplete/extra/missing trial
  reports and verifies the checkpoint again after the matrix
  (`run_feed_mutation_matrix.ps1:101-175,256-286`).
- `cmd/loadtest/run_feed_cursor_ab.ps1`: RPC/Gateway page-number and Cursor
  scenarios; builds a current binary first and rejects any incomplete report
  (`run_feed_cursor_ab.ps1:62-130`).
- `scripts/start-feed-local-services.ps1` / `stop-feed-local-services.ps1`:
  hybrid local topology with existing middleware images and locally built
  project processes; PID, executable path, start time, and port ownership are
  checked before stop/start actions (`start-feed-local-services.ps1:33-43,240-280`;
  `stop-feed-local-services.ps1:31-69`).

Current verification performed during this research:

- `go test ./cmd/loadtest` passed (cached) on 2026-09-09.
- PowerShell parser accepted the five scripts above.
- Start/stop `-WhatIf` completed and showed the migrated local command paths and
  ports without changing containers, processes, data, or files.

The following historical paths are not runnable acceptance tests:

- `cmd/loadtest/README.md`, `QUICKSTART.md`, and `run_test.ps1/.sh` use the old
  CLI. The current guide explicitly marks them historical
  (`cmd/loadtest/FEED_MATRIX_GUIDE.md:3-6`).
- `docs/loadtest-knowpost-feed.md` references `.tmp_seed/k6/*` and monitor Go
  files which are not present in the checkout. Its July public-Feed figures are
  historical context only, not a Stage 5 command.
- Early publish/mixed results before trial-level checkpoint/restore are
  non-comparable probes; only the 2026-08-27 mutation reports are formal evidence
  (`docs/superpowers/reports/2026-08-17-feed-hybrid-comprehensive-performance-report.md:438-440`).

### Prerequisites and side effects

#### Runtime prerequisites

- Run from the repository root with Go and PowerShell available.
- Host MySQL and Redis must be reachable at `127.0.0.1:3306` and `:6379`.
- Docker must already have usable images for etcd, ZooKeeper, Kafka, Canal, and
  Elasticsearch. The local launcher uses `--no-build --pull never`; it stops the
  Compose application containers but preserves middleware containers/volumes
  (`scripts/start-feed-local-services.ps1:220-245`).
- The load generator discovers RPCs through etcd at `127.0.0.1:12379` and uses
  the service keys in `cmd/loadtest/load_test.yaml:59-94`.
- Signed Gateway runs require `certs/jwt_private.pem` and
  `certs/jwt_public.pem`; the request still traverses real JWT verification
  (`cmd/loadtest/feed_loadtest.go:428-430,475-489`).
- The local ports listed at `scripts/start-feed-local-services.ps1:35-43` must
  be free and not in Windows excluded ranges. Startup fails closed otherwise.
- Formal local runs require `benchmark-error-only`, which preserves errors but
  disables high-volume access/stat/SQL-info logging
  (`cmd/loadtest/FEED_EVOLUTION_GUIDE.md:16-18,130-147`).

#### Data and machine mutations

- `setup` creates owned benchmark users, follow edges, warmup posts, Outbox/Kafka
  events, and Redis Feed state. Presets require 1,225 users for `hot`,
  `distributed`, and `cursor-deep`, or 20,025 for `high`
  (`cmd/loadtest/topology.go:30-58`; `feed_loadtest.go:601-645`).
- `smoke` is not read-only: it publishes one normal and one BigV post, waits for
  Feed visibility/Kafka drain, and writes smoke evidence
  (`cmd/loadtest/smoke.go:40-98`).
- Pure read trials do not reset/seed/publish the base dataset, but explicit cache
  states advance the global Feed safety epoch and populate/expire caches.
  `cold`, `l1-warm`, `l2-warm`, and `expire-together` semantics are documented at
  `cmd/loadtest/FEED_EVOLUTION_GUIDE.md:43-54`.
- Cursor setup creates 1,500 published posts (20 normal authors x 50 plus five
  BigV authors x 100) and validates Redis, MySQL, Kafka and the oracle before
  setting the manifest ready (`FEED_EVOLUTION_GUIDE.md:190-215`).
- Mutation trials create real MySQL posts and Outbox rows, Kafka events, Redis
  Feed state, and safety-epoch changes. Each warmup and measured trial restores
  the exact checkpoint, and the final report remains incomplete until every
  generated post is covered and restoration succeeds
  (`cmd/loadtest/FEED_MUTATION_GUIDE.md:5-17`). Kafka history is immutable and
  therefore remains after restoration.
- `cleanup` requires exact RunID confirmation and drained Kafka. Dry-run performs
  ownership/SQL/key checks only. Real cleanup deletes manifest-owned users,
  relations, login logs, posts, Outbox and derived Redis keys in scoped batches;
  it does not remove Elasticsearch documents, Kafka history, reports, manifests,
  or process-local L1 entries (`cmd/loadtest/cleanup.go:48-119,436-457`).
- If Windows Redis reports `MISCONF`, affected reports are invalid. Temporarily
  disabling RDB/write protection changes Redis runtime configuration and is only
  allowed after recording the old values; restore them and run SET/GET/DEL after
  testing (`cmd/loadtest/FEED_EVOLUTION_GUIDE.md:149-162`).
- The launcher writes ignored state/binaries/config/logs under `.tmp/feed-local`.
  Reports, manifests, and checkpoints are written under `results/feed-loadtest`
  and are not ignored, so new observation IDs will appear as worktree changes.

### Coverage map

| Concern | Current executable coverage | What a pass proves | Limitation |
|---|---|---|---|
| Direct RPC | `-entry rpc` | etcd discovery and the migrated KnowPost/Relation/Counter/User RPC chain work under load | Does not include HTTP/JWT/JSON |
| Gateway | `-entry gateway -gateway-auth-mode signed` | real Gateway route, JWT, HTTP/JSON, and Gateway-to-RPC hop work | Gateway has high host-scheduling variance; keep it separate from RPC |
| Hot | `hot` manifest + `hot-read` + `l1-warm` | 20-reader L1 target state and refresh path | Not a cold-path claim |
| Cold | `high` + `cold`, or explicit `-cache-state cold` | cold compute, Relation/Redis/hydration path | Incrementing safety epoch is a controlled cache mutation |
| Distributed | `distributed` + `distributed-read` + `l2-warm` | 1,200-reader L2 target state | The 102K historical RPC value is a cache upper-bound card |
| High cardinality | `high` + `cold` | 20,000-reader locality loss, refresh pressure and memory behavior | Dataset setup can dominate wall time |
| Deep page | `deep-page`, plus Cursor page20/page50/sequential scenarios | legacy page-number cost and current Cursor correctness/cost | Requires the dedicated Cursor manifest for meaningful depth |
| Mixed read/write | `publish`, `mixed-90-10`, `mixed-80-20` with checkpoint | complete four-step publish flow, read isolation, Kafka recovery and data restoration | Current formal historical baseline is RPC-only by deliberate scope decision |
| Burst | `burst-read` (10s) and `steady-read` (30s) | idle-to-load request behavior at a fixed distributed target | No dedicated historical numeric pass threshold |
| Restart/recovery | report monitor + explicit stop/start/smoke cycle | no accidental restart during measurement; services recover after a controlled lifecycle restart | Harness does not inject a service restart during load |
| Resources | built-in monitor | Docker/local CPU/RSS, Redis ops and identity, MySQL status, Kafka lag, Feed Prometheus, client CPU, binary/Git identity | Snapshot alone is not load-window evidence |

`snapshot` records container health/restart count, runtime strategy and Kafka
offset/lag, while every trial records the correlated samples. A snapshot alone
must not replace trial evidence (`cmd/loadtest/FEED_MATRIX_GUIDE.md:125-145`).

### Historical baselines worth retaining

All values below are three-trial medians on the 16-logical-CPU Windows host,
Hybrid strategy, local business services plus Compose middleware, and
`benchmark-error-only`. They are reference cards, not portable production SLAs.

| Card | Historical treatment median | Evidence |
|---|---:|---|
| Hot RPC, c32, L1-warm | 8,222.49 QPS; P95 5.309ms | comprehensive report lines 318-322 |
| Hot Gateway, c32, L1-warm | 4,930.08 QPS; P95 9.835ms | same |
| Hot deep-page RPC, c32 | 2,932.98 QPS; P95 13.990ms | same |
| Distributed RPC, c64, L2-warm | 102,075.49 QPS; P95 1.16ms | capacity report lines 39-50 |
| Distributed Gateway, c16, L2-warm | 1,529.10 QPS; P95 26.34ms | same; historically noisy |
| High RPC, c16, cold | 9,477.51 QPS; P95 3.73ms | same |
| High Gateway, c8, cold | 1,407.80 QPS; P95 15.16ms | same; no proven cache gain |
| Cursor RPC page50, c16 | page-number 592.4 QPS/P95 35.66ms; Cursor 7,621.7/3.36ms | Cursor report lines 41-64 |
| Cursor Gateway page50, c16 | page-number 489.2 QPS/P95 49.02ms; Cursor 2,180.4/11.37ms | same |
| Cursor sequential 50, RPC | page-number 991.0 QPS/P95 28.02ms; Cursor 4,974.6/5.42ms | Cursor report lines 67-78 |
| Mutation RPC publish c2 | 6.33 complete workflows/s; P95 388.98ms | mutation report lines 8-25 |
| Mutation RPC 90/10 c4 | total 61.37 QPS; read P95 2.64ms; publish P95 779.37ms | same |
| Mutation RPC 80/20 c4 | total 30.05 QPS; read P95 3.19ms; publish P95 828.74ms | same |

Existing Cursor decision thresholds are also useful correctness/performance
invariants: page50 members must fall by at least 70%; page20/page50 P95 must
improve by over 30% without P99/CPU-per-request/allocation/GC regression;
sequential members/details/CPU-per-request must fall by over 50% without total
time or QPS regression; Cursor page50 P95 must be no more than 1.5x page5
(`docs/superpowers/reports/2026-08-17-feed-cursor-pagination-ab.md:94-102`).

### Proposed compressed observation run

Use fresh `ReportRunId` values. For strongest evidence, create fresh dataset
RunIDs; for an initial fast compatibility pass, an existing manifest may be used
only if `smoke` and its identity/cardinality checks pass. Existing manifests do
not prove that their underlying database state is still pristine, so results
from reused datasets must not be called a strict new baseline.

1. **Preflight and runtime start**

   ```powershell
   go test ./cmd/loadtest ./deploy/topology ./deploy/compose
   ./scripts/start-feed-local-services.ps1 `
     -Strategy hybrid -FeaturePreset treatment `
     -LoggingPreset benchmark-error-only
   ```

   Then run `-phase smoke` for the chosen hot RunID and one short signed Gateway
   read. Capture `-phase snapshot`. This gives fast go/no-go evidence before any
   long matrix.

2. **Selected pure-read cards**

   Run three trials with `-warmup 10s -duration 60s -requests 0` at only these
   points: hot RPC c32, hot Gateway c32, hot deep-page RPC c32, distributed RPC
   c64, high RPC c16, and high Gateway c8. Use `l1-warm`, `l2-warm`, and `cold`
   respectively. This distinguishes RPC/Gateway, hot/distributed/high, warm/cold,
   and legacy deep-page behavior without the default four-concurrency sweep.
   Add `burst-read` 10s and `steady-read` 30s at one distributed RPC concurrency
   as validity cards, not capacity baselines.

3. **Cursor deep-page card**

   Restart with `-FeaturePreset cursor-ab`. At c16, run RPC
   `page50,cursor-page50,sequential-page-50,sequential-50,same-second-cursor`,
   then Gateway `page50,cursor-page50`, with a unique report namespace per
   invocation. This retains the deepest correctness and allocation/GC signals
   while skipping noisy shallow-page and adjacent-overload sweeps.

4. **Mutation/recovery card**

   Restart with `treatment`; create a new checkpoint tied to the current binary.
   Run RPC-only `publish c2`, `mixed-90-10 c4`, and `mixed-80-20 c4`, each 10s
   warmup + 60s measurement x 3. Reuse the same checkpoint but separate fresh
   report namespaces. Do not reuse the 2026-08-27 checkpoint because checkpoint
   validation binds to manifest, Redis identity, and subject binary identity
   (`cmd/loadtest/mutation_checkpoint_runtime.go:444-550`).

5. **Lifecycle recovery and cleanup**

   Capture a final snapshot, stop the local processes, start them once more with
   the same treatment preset, rerun the hot smoke/short Gateway read, and stop.
   This is the explicit restart/recovery check missing from the load harness.
   For fresh datasets, execute cleanup dry-run, review counts, then exact RunID
   cleanup. Preserve raw reports and a concise gate summary.

The selected measured windows are about 60-70 minutes. Allow roughly 90 minutes
including cache preparation, checkpoint restore, process restarts, comparisons,
and snapshots when existing datasets are valid. Fresh setup is intentionally not
given a fixed estimate: the repository records no trustworthy setup duration,
and creating 20,025 high-cardinality users or 1,500 Cursor posts depends heavily
on the current MySQL/Kafka/host state.

For comparison, the default read script is at least about 73 minutes per
cardinality before setup (20 read cells x 3 x 70s, plus supplemental trials), the
default two-entry mutation matrix has 72 trials and historically needs roughly
90-100 seconds per trial, and the full Cursor formal default with c16+c32 has 120
trials (at least 140 minutes). Replaying all of those would add little topology
confidence relative to the compressed stable-point gate.

### Pass decision for entering Stage 5

Enter Stage 5 only when all of the following are recorded for the same Stage 4
commit and local topology fingerprint:

- Every selected formal report is present and `complete=true`; no failed request,
  timeout, Feed degradation, process/container restart, unhealthy dependency,
  client saturation, Redis eviction/rejection, or unresolved hard metric gap.
- Kafka final lag is zero. Every mutation trial has complete pre-warmup,
  post-warmup and post-measurement restore evidence, and a final independent
  checkpoint verification succeeds.
- Cursor reports have zero oracle mismatch, duplicate, loop, or early termination
  and retain the published Cursor thresholds above.
- Current three-trial medians meet the proposed QPS/P95 non-regression band, or
  any exception is explained by a concrete environment fingerprint difference
  and confirmed by a clean rerun. Do not compare RPC to Gateway, cold to warm,
  or different cardinalities/topologies.
- The controlled stop/start cycle returns to a complete smoke plus signed
  Gateway read, and the stop script releases the recorded processes/ports.
- Cleanup preview and exact cleanup complete for all new datasets; Redis runtime
  persistence settings are restored if they were changed.
- Separate smoke evidence exists for LLM and Agent. The Feed harness does not
  launch or call either service, so it cannot by itself close observation for
  those two Stage 4 migrations.

## Files Found

- `cmd/loadtest/feed_loadtest.go` - current CLI, phases, execution and report completion.
- `cmd/loadtest/run_feed_matrix.ps1` - current pure-read RPC/Gateway matrix.
- `cmd/loadtest/run_feed_mutation_matrix.ps1` - current checkpointed write/mixed matrix.
- `cmd/loadtest/run_feed_cursor_ab.ps1` - current deep-pagination matrix.
- `scripts/start-feed-local-services.ps1` - current mixed local topology launcher.
- `scripts/stop-feed-local-services.ps1` - PID/executable/start-time-safe stop path.
- `cmd/loadtest/FEED_EVOLUTION_GUIDE.md` - current topology, cache and execution guide.
- `cmd/loadtest/FEED_MATRIX_GUIDE.md` - current read matrix, artifacts and cleanup guide.
- `cmd/loadtest/FEED_MUTATION_GUIDE.md` - current mutation isolation and pass conditions.
- `docs/superpowers/reports/2026-08-17-feed-hybrid-comprehensive-performance-report.md` - consolidated hot/deep-page history and evidence grading.
- `docs/superpowers/reports/2026-08-17-feed-hybrid-wp11-capacity.md` - distributed/high formal baselines.
- `docs/superpowers/reports/2026-08-17-feed-cursor-pagination-ab.md` - deep Cursor baselines and thresholds.
- `docs/superpowers/reports/2026-08-27-feed-wp11-mutation-rpc-ab.md` - latest valid mutation baseline.
- `results/feed-loadtest/manifests/*.json` - historical dataset manifests.
- `results/feed-loadtest/<report-run-id>/**` - raw historical reports and comparisons.

## Result Artifacts

- Dataset manifest: `results/feed-loadtest/manifests/<run-id>.json`
- Mutation checkpoint: `results/feed-loadtest/checkpoints/<run-id>.json` (or explicit path)
- Smoke: `results/feed-loadtest/<run-id>/smoke-<strategy>.json`
- Trial: `results/feed-loadtest/<report-run-id>/<strategy>-<entry>-<scenario>-c<N>/trial-<N>/{report.json,report.md,stages.csv}`
- Comparison: `results/feed-loadtest/<report-run-id>/{comparison.json,comparison.csv,comparison.md}`
- Mutation expected grid: `results/feed-loadtest/<report-run-id>/expected-matrix.json`
- Snapshot: `results/feed-loadtest/<run-id>/environment-snapshot.json`
- Cleanup: `results/feed-loadtest/<run-id>/{cleanup-preview.json,cleanup.json}`
- Local runtime identity/logs: `.tmp/feed-local/{state.json,bin/,config/,logs/}`
- Optional profiles: the path passed through `-pprof-dir`

## External References

None. This gate is repository-specific and was derived from current executable
code/scripts plus checked-in historical raw reports. No external documentation
is required.

## Related Specs

- `.trellis/tasks/09-08-repository-topology-v2/prd.md:21-59` - behavior preservation, local startup, contract tests and one-cycle observation requirement.
- `.trellis/tasks/09-08-repository-topology-v2/design.md:53-72` - compatibility invariants and legacy retirement sequence.
- `.trellis/tasks/09-08-repository-topology-v2/implement.md:32-49` - Stage 5 gate and runtime validation commands.
- `.trellis/spec/backend/directory-structure.md:97-109` - required context-migration tests and live Gateway/RPC/port smoke.
- `.trellis/spec/backend/directory-structure.md:143-212` - topology manifest, local artifact and lifecycle contracts.

## Caveats / Not Found

- No current script injects a service crash/restart during active load. A restart
  invalidates a measured trial; recovery must be observed as a separate
  stop/start/smoke sequence.
- No existing code enforces absolute QPS or P95 regression thresholds. The
  proposed 20% QPS / 25% P95 band is a new acceptance policy and should be
  approved before execution, not presented as historical project policy.
- Fresh dataset setup wall times are not recorded. Only measured-window and
  mutation trial durations can be estimated defensibly.
- Historical result directories include invalid/scout runs. Use only the named
  formal namespaces and require current `complete=true` reports.
- The Feed suite starts Storage and Search but does not meaningfully load their
  business operations; it neither starts nor observes LLM or Agent. Separate
  runtime smokes remain necessary before declaring every Stage 4 context observed.
- No long load test was run during this research. Only the loadtest unit suite,
  PowerShell parsing, and side-effect-free launcher `-WhatIf` checks were run.

## Execution results (2026-09-09)

The observation run used the local Hybrid/Treatment topology and isolated
results under `.tmp/feed-observation`. The launcher initially needed
`-buildvcs=false`; the Kafka parser also needed to accept an uncommitted
partition with `CURRENT-OFFSET=-` as lag zero for an empty topic. Gateway load
generation required a bounded reusable HTTP transport; the previous default
client exhausted the Windows Go thread limit at c32.

Completed valid cards:

- hot RPC c32: median 95.72K QPS, P95 0.668ms, 3/3 complete;
- hot Gateway c32 signed: median 24.17K QPS, P95 2.316ms, 3/3 complete;
- hot deep-page RPC c32: median 15.58K QPS, P95 3.465ms, 3/3 complete;
- fresh distributed RPC c64: median 98.27K QPS, P95 1.240ms, 3/3 complete;
- fresh high Gateway c8 signed: median 3.41K QPS, P95 5.684ms, 3/3 complete.

High RPC c16 cold was valid at the request/health level but failed the proposed
non-regression band: the first three trials had median 5.66K QPS / 6.47ms P95,
and a clean rerun had median 4.56K QPS / 9.00ms P95. The cache Fresh ratio fell
from 0.58 to 0.458 across the clean rerun while Redis eviction/rejection and
client CPU remained zero/low. The user explicitly approved entering Stage 5
after confirming that the migrated topology was fully connected. Cursor and
mutation formal matrices were intentionally not started after this performance
residual was observed; they remain follow-up evidence rather than a claim that
the high RPC cold card passed.

The fresh distributed and high manifests were cleaned by exact RunID. Cleanup
needed bounded SQL batches and a 60-second maintenance-only MySQL timeout;
these fixes were verified by the high dataset cleanup (97,100 outbox, 97,000
relations in each direction, 20,025 users). Local business processes were then
stopped and all managed ports were released. Redis persistence settings were
unchanged.
