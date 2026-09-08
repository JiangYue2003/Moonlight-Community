# Runtime topology

`services.json` is the authoritative local runtime inventory. It owns service
IDs, Go package/config paths, declared ports, start order, and dependency
profiles.

Validate it from the repository root:

```powershell
go run ./deploy/topology/cmd/validate
```

The array order is the safe local start order. `scripts/start-all.ps1` reads the
entries marked `defaultLocal`; do not add a second hand-maintained service list.
Compose remains hand-authored in this stage because it also owns health checks,
volumes, and container-specific configuration. The current schema models only
dependency Compose names; application-service mapping and its drift check must
be added explicitly before application topology is validated against Compose.

`host-core` validates dependencies already listening on the standard local
ports. `compose-core` keeps the host MySQL and Redis instances, starts etcd,
ZooKeeper/Kafka, and Elasticsearch from `docker-compose.full.yml`, and leaves
the Go application processes on the host.
