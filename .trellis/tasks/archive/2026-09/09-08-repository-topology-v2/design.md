# Repository topology v2 design

## Architecture direction

The repository remains a single Go module. Business ownership is expressed by
bounded-context directories; go-zero's `handler`/`logic`/`svc`/`model` convention
is retained at transport edges instead of becoming the whole system architecture.

```text
services/<context>/
  cmd/<process>/              executable entry points
  internal/
    domain/                   business invariants, when they exist
    application/              commands, queries, orchestration
    adapter/                  MySQL, Redis, Kafka, ES, OSS implementations
    transport/                gRPC, HTTP, worker adapters
    bootstrap/                config and component lifecycle

contracts/proto/              cross-service source contracts
platform/                     domain-neutral infrastructure packages
deploy/topology/              authoritative runtime manifest and validator
var/                          ignored repository-local runtime state
```

Initial context mapping:

| Current owner | Target context |
|---|---|
| `user`, `auth`, `profile` | `identity` |
| `knowpost` | `content` |
| `relation` | `social` |
| `counter` | `engagement` |
| `search` | `discovery` |
| `llm` | `intelligence` |
| `agent` | `assistant` |
| `storage` | `media` |

Names are migration destinations, not an instruction to rename everything in
one change. The first stage establishes the manifest and hygiene rules without
moving business packages.

## Runtime topology

`deploy/topology/services.json` is the source of truth. It records stable service
IDs, roles, Go commands, config paths, ports, dependency profiles, and whether a
component is part of the default local stack. A small Go validation command reads
the manifest and checks referenced paths, unique ports/IDs, and supported roles.

Startup scripts may read the manifest directly. Compose files may remain hand
authored initially, but CI validates their declared application services against
the manifest until generation is justified.

## Compatibility boundaries

These remain invariant throughout migration:

- public HTTP methods, paths, auth requirements, payloads, and status codes;
- protobuf package/service/message definitions;
- SQL schemas and transaction boundaries;
- Redis keys, Kafka topics, consumer groups, and event payloads;
- process shutdown semantics and default feature flags.

Directory moves use compatibility packages only when needed to keep stages
buildable. Such shims must have an owner and removal stage.

## Legacy API retirement

1. Keep legacy API code disabled by default.
2. Add parity tests at the gateway/RPC boundary.
3. Run the new topology for one release cycle.
4. Tag the last compatible commit.
5. Remove unused API trees and their bootstrap adapters.

## Artifact policy

- `var/log`, `var/run`, `var/cache`, and `bin` are disposable and ignored.
- Credentials remain external; only documentation/examples are tracked.
- Generated protobuf code remains tracked until a separate generation-contract
  change is approved.
- Elasticsearch IK binaries are build inputs, not temporary files; later work
  must pin their source/checksum or build them into an image.

## Rollout and rollback

Each stage is a separate commit. A stage must build and test before the next
begins. Rollback is `git revert` of that stage; no database or message migration
is needed because contracts are unchanged.
