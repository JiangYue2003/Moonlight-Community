# Repository topology v2 implementation plan

## Stage 1: topology foundation and repository hygiene

- Add the authoritative runtime manifest and validator tests.
- Make Windows local startup resolve its core service list from the manifest.
- Document dependency profiles and topology ownership.
- Consolidate future logs/PIDs/caches under `var/`; keep compatibility for
  existing ignored locations during this stage.
- Verify validator, script parsing, Go tests/build, selected dependency startup,
  and local core-service startup.

## Stage 2: gateway transport split

- Split the monolithic gateway router by bounded context without changing routes.
- Keep one composition root and route-parity tests.
- Verify gateway tests plus HTTP smoke tests against locally running RPCs.

## Stage 3: one bounded-context pilot

- Migrate one low-coupling context (`media/storage`) into the target layout.
- Prove that `cmd` and context-level `internal` remove unnecessary public app
  adapters while preserving RPC behavior.
- Record the repeatable move pattern before applying it elsewhere.

## Stage 4: remaining contexts

- Migrate identity, engagement, social, content, discovery, intelligence, and
  assistant separately, in dependency order.
- Run full verification after every context.

## Stage 5: legacy retirement

- Complete one release-cycle observation gate.
- Tag the rollback baseline.
- Delete unused legacy HTTP entry points and compatibility bootstraps.
- Run full tests, build, Compose smoke tests, and documentation drift checks.

## Validation commands

```powershell
go test ./...
go vet ./...
go build ./...
docker compose -f deploy/compose/docker-compose.dev.yml config
docker compose -f deploy/compose/docker-compose.dev.yml up -d mysql redis etcd kafka
docker compose -f deploy/compose/docker-compose.dev.yml ps
scripts/start-all.ps1
```

HTTP smoke checks use unauthenticated health/debug endpoints first, followed by
existing authenticated workflow tests where fixtures are available.

## Rollback points

- Commit after every stage.
- No stage may combine directory movement with contract/schema changes.
- If runtime smoke tests fail, revert only the latest stage and preserve logs in
  `var/log/dev` for diagnosis.
