# Repository topology v2

## Goal

Reduce repository and runtime-topology complexity without changing public HTTP,
RPC, persistence, cache, or message behavior. Establish a domain-oriented
service layout, one authoritative runtime manifest, and explicit local-artifact
boundaries so future migrations are incremental and reversible.

## Background

- The runtime has converged toward one main gateway plus domain RPC/worker
  processes, while legacy HTTP `api/` trees remain as rollback code.
- Service lists currently drift between `README.md`, `scripts/start-all.ps1`,
  and Compose variants.
- Merged processes use public `app` adapters to cross sibling Go `internal/`
  boundaries.
- Local caches, binaries, logs, generated front-end output, and tool state are
  spread across the repository root and subtrees.

## Requirements

- Preserve all externally observable HTTP routes, RPC contracts, table schemas,
  Redis key formats, Kafka topics/events, and default feature behavior.
- Use bounded-context ownership for services while retaining go-zero conventions
  inside transport adapters; do not mechanically force DDD layers where no
  domain logic exists.
- Define one machine-readable runtime manifest as the authoritative source for
  service names, commands, ports, configs, roles, and dependency profiles.
- Make local scripts and documentation consume or validate against that manifest.
- Consolidate repository-local ephemeral state under documented ignored roots.
- Keep required build inputs, including the Elasticsearch IK plugin, reproducible
  and distinct from disposable local artifacts.
- Retire legacy HTTP `api/` code only after one release cycle on the new topology;
  retain a Git tag as the long-term rollback point.
- Execute migration in independently testable stages on `chore/topology-v2`.

## Out of Scope

- Product behavior changes or endpoint redesign.
- Proto, database-schema, Redis-key, and Kafka-event migrations.
- Merging Agent into the main gateway.
- Changing deployment infrastructure from Compose to Kubernetes.
- A big-bang rename of every bounded context in the first stage.

## Acceptance Criteria

- [ ] A clean checkout contains no runtime-generated binaries, logs, PID files,
      caches, or front-end dependency/build directories.
- [ ] Runtime topology has one authoritative manifest and automated drift checks.
- [ ] Local startup can select a dependency profile and launch the same core
      services described by the manifest.
- [ ] `go test ./...`, `go vet ./...`, and `go build ./...` pass after each code
      migration stage.
- [ ] Selected Compose dependencies become healthy and migrated services start
      locally with their configured health/debug endpoints reachable.
- [ ] Existing gateway route and RPC contract tests demonstrate behavior parity.
- [ ] Legacy HTTP entry points are disabled first, observed for one release cycle,
      and removed only in a later stage with an explicit rollback tag.
- [ ] The target directory map and migration/rollback rules are documented.

## Key Decisions

- Migration style: incremental strangler migration, not a big-bang rewrite.
- Architecture style: pragmatic DDD boundaries with go-zero transport adapters.
- Legacy rollback: retain for one release cycle, then delete; Git tag remains.
- Workspace isolation: implementation happens in a dedicated Git worktree.
