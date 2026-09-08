# Directory Structure

## Current Boundaries

The repository is one Go module. Runtime ownership starts under `services/`,
while domain-neutral packages live under `pkg/` or `common/`. A go-zero
transport keeps its generated `internal/handler`, `internal/logic`,
`internal/svc`, and `internal/model` conventions inside the owning service.

Repository-level operational files have separate owners:

```text
deploy/topology/   machine-readable process and dependency inventory
deploy/compose/    container topology and health checks
scripts/           local lifecycle commands
var/               ignored local binaries, caches, logs, and PID records
```

Do not move business packages merely to make the tree visually uniform. Move
one bounded context at a time, preserving HTTP, RPC, SQL, Redis, and Kafka
contracts throughout the migration.

## Scenario: Runtime Topology and Local Artifacts

### 1. Scope / Trigger

This contract applies when adding, removing, renaming, starting, or stopping a
process, changing a service port/config path, or adding repository-local runtime
output.

### 2. Signatures

```powershell
go run ./deploy/topology/cmd/validate
scripts/start-all.ps1 [-WithDocker] [-DependencyProfile <id>] [-ValidateOnly]
scripts/stop-all.ps1 [-PidFile <path>]
```

The validator accepts `-repo-root` and `-manifest`. Paths in the manifest are
repository-relative and use `/` separators.

### 3. Contracts

- `deploy/topology/services.json` is the authoritative local process inventory.
- A service entry owns its stable ID, role, Go package, config, CLI arguments,
  log filename, dependency profiles, ports, and default-local membership.
- `defaultLocal` service array order is startup order.
- A port with `managedOnStart: true` is both a pre-start ownership check and a
  readiness condition. Metrics/debug ports may be recorded without becoming a
  Windows startup blocker.
- `host-core` checks host dependencies. `compose-core` starts only its declared
  Compose dependencies and still checks any required host endpoints.
- New local state goes under `var/bin`, `var/cache`, `var/log`, or `var/run`.
  Required build inputs and credentials do not belong under `var/`.
- PID records contain the executable path and process start timestamp. Stop
  logic verifies both before terminating a current-format record.
- Legacy PID records without that identity evidence are never terminated
  automatically; a live record is retained for manual inspection.

### 4. Validation & Error Matrix

| Condition | Required result |
|---|---|
| Unknown JSON field, schema version, role, or dependency profile | Validator fails |
| Absolute, escaping, or missing package/config/Compose path | Validator fails |
| Duplicate service ID, profile ID, service port, or per-service port name | Validator fails |
| Required dependency endpoint is unreachable | Startup fails before building services |
| A managed service port is already owned | Startup fails without killing that process |
| A started process exits or does not own its managed ports | Startup rolls back recorded processes |
| PID executable or start time differs from the live process | Stop refuses the process and retains the record |
| A live legacy PID record has no executable/start-time identity | Stop refuses the process and retains the record |

### 5. Good/Base/Bad Cases

- Good: add a worker to `services.json`, point it at an existing package and
  config, declare its dependency profiles, then run topology and Compose tests.
- Base: a default-local service builds into `var/bin/dev`, writes logs under
  `var/log/dev`, and is stopped from `var/run/dev/pids.json`.
- Bad: add the same service separately to a PowerShell array, README list, and
  Compose contract test without deriving or checking it against the manifest.

### 6. Tests Required

- `go test ./deploy/topology/...` checks schema, paths, references, IDs, ports,
  and dependency-profile Compose service names.
- `go test ./deploy/compose` checks that Compose contains manifest-defined
  default-local services and excludes `var/` from build context.
- PowerShell parser validation covers `scripts/start-all.ps1` and
  `scripts/stop-all.ps1`.
- A local smoke test must start the default stack, verify each managed port,
  exercise at least one Gateway route, stop the stack, and assert all managed
  ports were released.

### 7. Wrong vs Correct

Wrong:

```powershell
$services = @("gateway", "counter", "search") # second topology list
Stop-Process -Id (Get-PortOwner 8080) -Force    # kills an unverified process
```

Correct:

```powershell
$topology = Get-Content deploy/topology/services.json -Raw | ConvertFrom-Json
$services = @($topology.services | Where-Object { $_.defaultLocal })
# Stop only a recorded PID whose executable path and start time still match.
```

## Scenario: Gateway Route Ownership

### 1. Scope / Trigger

This contract applies when adding, moving, renaming, or removing a Gateway HTTP
route or handler under `services/gateway/internal/handler`.

### 2. Signatures

```go
func NewEngine(sc *srv.ServiceContext) *gin.Engine
func register<Context>Routes(r *gin.Engine, sc *srv.ServiceContext)
```

`NewEngine` is the single composition root. It configures global middleware and
calls one registration function per bounded context. Context files own their
route groups, context-specific middleware, handlers, and response projections.

### 3. Contracts

- Public HTTP methods, paths, registration order, middleware, payloads, status
  codes, RPC calls, and response shapes remain compatibility contracts.
- Route files use bounded-context names (`identity`, `media`, `content`,
  `social`, `engagement`, `discovery`, and `intelligence`), not one file per
  endpoint.
- Cross-context parsing helpers may live in `shared.go` only when used by more
  than one context. Business behavior stays with its owning context.
- Adding a context requires one registration call in `NewEngine`; it must not
  create another engine or composition root.

### 4. Validation & Error Matrix

| Condition | Required result |
|---|---|
| HTTP method/path is added, removed, or changed | Route-parity expectation changes explicitly or the test fails |
| A `RequiredAuth` route loses or bypasses middleware | Missing-token request does not return 401 and the auth-parity test fails |
| A handler move changes imports, RPC types, or response mapping | Gateway tests, vet, or build fails |
| Two contexts register the same method/path | Gin registration panics during `NewEngine` construction |

### 5. Good/Base/Bad Cases

- Good: add a content route in `content.go`, register it through
  `registerContentRoutes`, and update route/auth parity assertions.
- Base: `router.go` creates the engine, applies global middleware, calls bounded
  context registrars, and contains no endpoint handler bodies.
- Bad: append unrelated route groups and handlers directly to `NewEngine`, or
  create a second Gateway engine for one context.

### 6. Tests Required

- `TestNewEngineRouteParity` asserts the complete sorted HTTP method/path set.
- `TestNewEngineRequiredAuthRouteParity` sends missing-token requests to every
  protected route and requires HTTP 401 before handler/RPC execution.
- Context-specific behavior tests remain next to the Gateway handler package.
- Run `go test ./services/gateway/...`, `go vet ./services/gateway/...`, and a
  local Gateway HTTP smoke test through live RPC dependencies.

### 7. Wrong vs Correct

Wrong:

```go
func NewEngine(sc *srv.ServiceContext) *gin.Engine {
	r := gin.New()
	r.POST("/api/v1/knowposts/drafts", createDraft(sc))
	// More endpoint bodies and route groups accumulate here.
	return r
}
```

Correct:

```go
func NewEngine(sc *srv.ServiceContext) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	registerContentRoutes(r, sc)
	return r
}
```
