package compose

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/zhiguang/zhiguang-go/deploy/topology"
	"gopkg.in/yaml.v3"
)

func TestCoreDockerfileBuildsOnlyActiveServices(t *testing.T) {
	path := filepath.Join("..", "..", "Dockerfile.core")
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read core Dockerfile: %v", err)
	}

	text := string(content)
	required := []string{
		"FROM alpine:3.22 AS builder",
		"ARG GO_VERSION=1.25.8",
		"ARG TARGETARCH",
		"apk add --no-cache ca-certificates git",
		`GO_SHA256=ceb5e041bbc3893846bd1614d76cb4681c91dadee579426cf21a63f2d7e03be6`,
		`GO_SHA256=7d137f59f66bb93f40a6b2b11e713adc2a9d0c8d9ae581718e3fad19e5295dc7`,
		`https://go.dev/dl/go${GO_VERSION}.linux-${TARGETARCH}.tar.gz`,
		`echo "$GO_SHA256  /tmp/go.tar.gz" | sha256sum -c -`,
		`test "$(go env GOVERSION)" = "go${GO_VERSION}"`,
		"ARG GOPROXY=https://goproxy.cn",
		"openssl",
		"./services/gateway",
		"/out/user-rpc ./services/user/cmd/user",
		"./services/storage/cmd/storage",
		"./services/counter/cmd/counter",
		"./services/knowpost/cmd/knowpost",
		"./services/relation/cmd/relation",
		"./services/search/cmd/search",
	}
	for _, snippet := range required {
		if !strings.Contains(text, snippet) {
			t.Errorf("core Dockerfile missing %q", snippet)
		}
	}

	for _, forbidden := range []string{"./services/llm", "./services/agent", "./services/user/rpc", "/out/user-merged", "COPY certs/"} {
		if strings.Contains(text, forbidden) {
			t.Errorf("core Dockerfile must not build sealed service %q", forbidden)
		}
	}
}

func TestDockerIgnoreExcludesLocalBuildArtifacts(t *testing.T) {
	content, err := os.ReadFile(filepath.Join("..", "..", ".dockerignore"))
	if err != nil {
		t.Fatalf("read .dockerignore: %v", err)
	}

	for _, pattern := range []string{".gocache/", "**/*.exe", "certs/", "var/"} {
		if !containsLine(string(content), pattern) {
			t.Errorf(".dockerignore missing local artifact pattern %q", pattern)
		}
	}
}

func TestDevComposeDefinesCoreServicesWithoutSealedOrHostInfrastructure(t *testing.T) {
	services := loadDevComposeServices(t)

	required := []string{
		"host-infra-check",
		"jwt-cert-init",
		"etcd",
		"zookeeper",
		"kafka",
		"canal-server",
		"elasticsearch",
	}
	for _, name := range required {
		if _, ok := services[name]; !ok {
			t.Errorf("dev Compose missing service %q", name)
		}
	}

	for _, name := range []string{"mysql", "redis", "llm", "agent"} {
		if _, ok := services[name]; ok {
			t.Errorf("dev Compose must not define service %q", name)
		}
	}

	for _, name := range defaultLocalServiceIDs(t) {
		service := services[name]
		if got := stringValue(service["image"]); got != "zhiguang-go-core:dev" {
			t.Errorf("%s image = %q, want zhiguang-go-core:dev", name, got)
		}
		build := mapValue(t, name+".build", service["build"])
		if got := stringValue(build["dockerfile"]); got != "Dockerfile.core" {
			t.Errorf("%s Dockerfile = %q, want Dockerfile.core", name, got)
		}
	}
}

func TestFullComposeContainsManifestDefaultLocalServices(t *testing.T) {
	services := loadComposeServices(t, "docker-compose.full.yml")
	for _, name := range defaultLocalServiceIDs(t) {
		if _, ok := services[name]; !ok {
			t.Errorf("full Compose missing default local topology service %q", name)
		}
	}
}

func TestDevComposePublishesHostReachableRPCAndWaitsForHealth(t *testing.T) {
	services := loadDevComposeServices(t)

	expectedPorts := map[string][]string{
		"gateway":  {"127.0.0.1:8080:8080"},
		"user":     {"127.0.0.1:9002:9002"},
		"storage":  {"127.0.0.1:9013:9013"},
		"counter":  {"127.0.0.1:9003:9003"},
		"knowpost": {"127.0.0.1:9004:9004", "127.0.0.1:16064:6064"},
		"relation": {"127.0.0.1:9006:9006", "127.0.0.1:16066:6066"},
		"search":   {"127.0.0.1:9017:9017"},
	}
	for name, wantPorts := range expectedPorts {
		service := services[name]
		gotPorts := stringSlice(t, name+".ports", service["ports"])
		for _, want := range wantPorts {
			if !contains(gotPorts, want) {
				t.Errorf("%s ports %v missing %q", name, gotPorts, want)
			}
		}
		if _, ok := service["healthcheck"]; !ok {
			t.Errorf("%s missing healthcheck", name)
		}
		if name == "knowpost" || name == "relation" {
			if len(gotPorts) != len(wantPorts) {
				t.Errorf("%s ports = %v, want exact set %v", name, gotPorts, wantPorts)
			}
			for _, port := range gotPorts {
				if !strings.HasPrefix(port, "127.0.0.1:") {
					t.Errorf("%s publishes non-loopback port %q", name, port)
				}
			}
		}
	}

	for _, name := range []string{"user", "storage", "counter", "knowpost", "relation", "search"} {
		environment := mapValue(t, name+".environment", services[name]["environment"])
		if got := stringValue(environment["POD_IP"]); got != "127.0.0.1" {
			t.Errorf("%s POD_IP = %q, want 127.0.0.1", name, got)
		}
	}
	knowPostEnvironment := mapValue(t, "knowpost.environment", services["knowpost"]["environment"])
	if got := stringValue(knowPostEnvironment["FEED_RELATION_EPOCH_ENABLED"]); got != "${FEED_RELATION_EPOCH_ENABLED:-false}" {
		t.Errorf("knowpost relation epoch flag = %q, want opt-in Compose environment expansion", got)
	}
	if got := stringValue(knowPostEnvironment["FEED_CONTENT_SAFETY_EPOCH_ENABLED"]); got != "${FEED_CONTENT_SAFETY_EPOCH_ENABLED:-false}" {
		t.Errorf("knowpost content safety epoch flag = %q, want opt-in Compose environment expansion", got)
	}
	if got := stringValue(knowPostEnvironment["FEED_ROUTE_SNAPSHOT_ENABLED"]); got != "${FEED_ROUTE_SNAPSHOT_ENABLED:-false}" {
		t.Errorf("knowpost route snapshot flag = %q, want opt-in Compose environment expansion", got)
	}
	if got := stringValue(knowPostEnvironment["FEED_COMBINED_PIPELINE_ENABLED"]); got != "${FEED_COMBINED_PIPELINE_ENABLED:-false}" {
		t.Errorf("knowpost combined pipeline flag = %q, want opt-in Compose environment expansion", got)
	}
	if got := stringValue(knowPostEnvironment["FEED_PAGE_CACHE_MODE"]); got != "${FEED_PAGE_CACHE_MODE:-off}" {
		t.Errorf("knowpost page cache mode = %q, want opt-in Compose environment expansion", got)
	}

	gatewayDependsOn := mapValue(t, "gateway.depends_on", services["gateway"]["depends_on"])
	for _, dependency := range []string{"user", "storage", "counter", "knowpost", "relation", "search"} {
		condition := mapValue(t, "gateway.depends_on."+dependency, gatewayDependsOn[dependency])
		if got := stringValue(condition["condition"]); got != "service_healthy" {
			t.Errorf("gateway dependency %s condition = %q, want service_healthy", dependency, got)
		}
	}

	knowPostDependsOn := mapValue(t, "knowpost.depends_on", services["knowpost"]["depends_on"])
	relationCondition := mapValue(t, "knowpost.depends_on.relation", knowPostDependsOn["relation"])
	if got := stringValue(relationCondition["condition"]); got != "service_healthy" {
		t.Errorf("knowpost dependency relation condition = %q, want service_healthy", got)
	}
}

func TestDevComposeGeneratesJwtCertsOutsideImageAndScopesThemToUser(t *testing.T) {
	services := loadDevComposeServices(t)

	certInit := services["jwt-cert-init"]
	if got := stringValue(certInit["image"]); got != "zhiguang-go-core:dev" {
		t.Errorf("jwt-cert-init image = %q, want zhiguang-go-core:dev", got)
	}
	command := strings.Join(stringSlice(t, "jwt-cert-init.command", certInit["command"]), "\n")
	for _, snippet := range []string{"openssl genpkey", "openssl pkey", "jwt_private.pem", "jwt_public.pem"} {
		if !strings.Contains(command, snippet) {
			t.Errorf("jwt-cert-init command missing %q", snippet)
		}
	}
	if !contains(
		stringSlice(t, "jwt-cert-init.volumes", certInit["volumes"]),
		"zg_dev_jwt_certs:/app/certs",
	) {
		t.Errorf("jwt-cert-init must write the managed JWT certificate volume")
	}

	user := services["user"]
	if !contains(
		stringSlice(t, "user.volumes", user["volumes"]),
		"zg_dev_jwt_certs:/app/certs:ro",
	) {
		t.Errorf("user must mount the managed JWT certificate volume read-only")
	}
	userDependsOn := mapValue(t, "user.depends_on", user["depends_on"])
	certCondition := mapValue(t, "user.depends_on.jwt-cert-init", userDependsOn["jwt-cert-init"])
	if got := stringValue(certCondition["condition"]); got != "service_completed_successfully" {
		t.Errorf("user jwt-cert-init condition = %q, want service_completed_successfully", got)
	}

	for _, name := range []string{"gateway", "storage", "counter", "knowpost", "relation", "search"} {
		if volumes, ok := services[name]["volumes"]; ok {
			for _, volume := range stringSlice(t, name+".volumes", volumes) {
				if strings.Contains(volume, "zg_dev_jwt_certs") {
					t.Errorf("%s must not receive the JWT private-key volume", name)
				}
			}
		}
	}
}

func TestDevComposeMakesDependenciesHealthyAndPublishesOnlyToLoopback(t *testing.T) {
	services := loadDevComposeServices(t)

	canal := services["canal-server"]
	health := mapValue(t, "canal-server.healthcheck", canal["healthcheck"])
	healthCommand := strings.Join(stringSlice(t, "canal-server.healthcheck.test", health["test"]), "\n")
	for _, snippet := range []string{"canal.pid", "the canal server is running now", "start successful"} {
		if !strings.Contains(healthCommand, snippet) {
			t.Errorf("canal-server healthcheck missing %q", snippet)
		}
	}

	expectedPorts := map[string][]string{
		"etcd":          {"127.0.0.1:12379:2379"},
		"zookeeper":     {"127.0.0.1:2181:2181"},
		"kafka":         {"127.0.0.1:9092:9092"},
		"canal-server":  {"127.0.0.1:11111:11111"},
		"elasticsearch": {"127.0.0.1:9200:9200"},
		"adminer":       {"127.0.0.1:8090:8080"},
		"kibana":        {"127.0.0.1:5601:5601"},
	}
	for name, wantPorts := range expectedPorts {
		gotPorts := stringSlice(t, name+".ports", services[name]["ports"])
		for _, want := range wantPorts {
			if !contains(gotPorts, want) {
				t.Errorf("%s ports %v missing %q", name, gotPorts, want)
			}
		}
		for _, port := range gotPorts {
			if !strings.HasPrefix(port, "127.0.0.1:") {
				t.Errorf("%s publishes non-loopback port %q", name, port)
			}
		}
	}
}

func TestDevComposeRunsInternalEtcdForContainersAndHostLoadTest(t *testing.T) {
	services := loadDevComposeServices(t)
	etcd := services["etcd"]

	if got := stringValue(etcd["image"]); got != "bitnamilegacy/etcd:3.5.21-debian-12-r6" {
		t.Errorf("etcd image = %q, want bitnamilegacy/etcd:3.5.21-debian-12-r6", got)
	}
	environment := mapValue(t, "etcd.environment", etcd["environment"])
	expectedEnvironment := map[string]string{
		"ETCD_LISTEN_CLIENT_URLS":          "http://0.0.0.0:2379",
		"ETCD_ADVERTISE_CLIENT_URLS":       "http://etcd:2379,http://127.0.0.1:12379",
		"ETCD_INITIAL_ADVERTISE_PEER_URLS": "http://etcd:2380",
	}
	for key, want := range expectedEnvironment {
		if got := stringValue(environment[key]); got != want {
			t.Errorf("etcd environment %s = %q, want %q", key, got, want)
		}
	}
	if _, ok := etcd["healthcheck"]; !ok {
		t.Errorf("etcd missing healthcheck")
	}

	hostCheckCommand := strings.Join(
		stringSlice(t, "host-infra-check.command", services["host-infra-check"]["command"]),
		"\n",
	)
	if strings.Contains(hostCheckCommand, "etcd:2379") {
		t.Errorf("host-infra-check must not require the retired host etcd")
	}

	for _, name := range []string{"gateway", "user", "storage", "counter", "knowpost", "relation", "search"} {
		dependsOn := mapValue(t, name+".depends_on", services[name]["depends_on"])
		condition := mapValue(t, name+".depends_on.etcd", dependsOn["etcd"])
		if got := stringValue(condition["condition"]); got != "service_healthy" {
			t.Errorf("%s etcd condition = %q, want service_healthy", name, got)
		}
	}

	loadTestConfig := loadProjectYAML(t, "cmd/loadtest/load_test.yaml")
	for _, clientName := range []string{"KnowPostRpc", "RelationRpc", "CounterRpc", "UserRpc"} {
		client := nestedMap(t, "cmd/loadtest/load_test.yaml", loadTestConfig, clientName)
		etcdConfig := mapValue(t, clientName+".Etcd", client["Etcd"])
		hosts := stringSlice(t, clientName+".Etcd.Hosts", etcdConfig["Hosts"])
		if len(hosts) != 1 || hosts[0] != "127.0.0.1:12379" {
			t.Errorf("%s Etcd hosts = %v, want [127.0.0.1:12379]", clientName, hosts)
		}
	}
}

func TestDevComposeMapsHostInfrastructureForApplicationContainers(t *testing.T) {
	services := loadDevComposeServices(t)
	wantHosts := []string{
		"mysql=host-gateway",
		"redis=host-gateway",
	}

	for _, name := range []string{
		"host-infra-check",
		"gateway",
		"user",
		"storage",
		"counter",
		"knowpost",
		"relation",
		"search",
	} {
		gotHosts := stringSlice(t, name+".extra_hosts", services[name]["extra_hosts"])
		for _, want := range wantHosts {
			if !contains(gotHosts, want) {
				t.Errorf("%s extra_hosts %v missing %q", name, gotHosts, want)
			}
		}
		if contains(gotHosts, "etcd=host-gateway") {
			t.Errorf("%s must resolve etcd through the Compose service", name)
		}
	}
}

func TestDockerConfigsUseContainerEndpointsForActiveRPCClients(t *testing.T) {
	testCases := []struct {
		file     string
		path     []string
		endpoint string
	}{
		{
			file:     "services/gateway/etc/gateway-docker.yaml",
			path:     []string{"AuthRpc"},
			endpoint: "user:9002",
		},
		{
			file:     "services/gateway/etc/gateway-docker.yaml",
			path:     []string{"UserRpc"},
			endpoint: "user:9002",
		},
		{
			file:     "services/gateway/etc/gateway-docker.yaml",
			path:     []string{"StorageRpc"},
			endpoint: "storage:9013",
		},
		{
			file:     "services/gateway/etc/gateway-docker.yaml",
			path:     []string{"KnowPostRpc"},
			endpoint: "knowpost:9004",
		},
		{
			file:     "services/gateway/etc/gateway-docker.yaml",
			path:     []string{"RelationRpc"},
			endpoint: "relation:9006",
		},
		{
			file:     "services/gateway/etc/gateway-docker.yaml",
			path:     []string{"CounterRpc"},
			endpoint: "counter:9003",
		},
		{
			file:     "services/gateway/etc/gateway-docker.yaml",
			path:     []string{"UserCounterRpc"},
			endpoint: "counter:9003",
		},
		{
			file:     "services/gateway/etc/gateway-docker.yaml",
			path:     []string{"SearchRpc"},
			endpoint: "search:9017",
		},
		{
			file:     "services/knowpost/cmd/knowpost/etc/knowpost-docker.yaml",
			path:     []string{"Rpc", "UserCounterRpc"},
			endpoint: "counter:9003",
		},
		{
			file:     "services/knowpost/cmd/knowpost/etc/knowpost-docker.yaml",
			path:     []string{"Rpc", "CounterRpc"},
			endpoint: "counter:9003",
		},
		{
			file:     "services/knowpost/cmd/knowpost/etc/knowpost-docker.yaml",
			path:     []string{"Rpc", "RelationRpc"},
			endpoint: "relation:9006",
		},
		{
			file:     "services/relation/cmd/relation/etc/relation-docker.yaml",
			path:     []string{"Rpc", "UserRpc"},
			endpoint: "user:9002",
		},
		{
			file:     "services/relation/cmd/relation/etc/relation-docker.yaml",
			path:     []string{"Syncer", "UserCounterRpc"},
			endpoint: "counter:9003",
		},
		{
			file:     "services/search/cmd/search/etc/search-docker.yaml",
			path:     []string{"Rpc", "CounterRpc"},
			endpoint: "counter:9003",
		},
		{
			file:     "services/search/cmd/search/etc/search-docker.yaml",
			path:     []string{"Indexer", "KnowPostRpc"},
			endpoint: "knowpost:9004",
		},
	}

	configs := make(map[string]map[string]any)
	for _, testCase := range testCases {
		config, ok := configs[testCase.file]
		if !ok {
			config = loadProjectYAML(t, testCase.file)
			configs[testCase.file] = config
		}

		client := nestedMap(t, testCase.file, config, testCase.path...)
		endpoints := stringSlice(t, testCase.file+"."+strings.Join(testCase.path, ".")+".Endpoints", client["Endpoints"])
		if len(endpoints) != 1 || endpoints[0] != testCase.endpoint {
			t.Errorf(
				"%s %s endpoints = %v, want [%s]",
				testCase.file,
				strings.Join(testCase.path, "."),
				endpoints,
				testCase.endpoint,
			)
		}
		if _, ok := client["Etcd"]; ok {
			t.Errorf("%s %s must use direct endpoints, not Etcd", testCase.file, strings.Join(testCase.path, "."))
		}
	}
}

func TestStorageConfigsPreserveRpcIdentityAndMetrics(t *testing.T) {
	for _, file := range []string{
		"services/storage/cmd/storage/etc/storage.yaml",
		"services/storage/cmd/storage/etc/storage-docker.yaml",
	} {
		config := loadProjectYAML(t, file)
		if got := stringValue(config["Name"]); got != "storage.rpc" {
			t.Errorf("%s Name = %q, want storage.rpc", file, got)
		}
		if got := stringValue(config["ListenOn"]); got != "0.0.0.0:9013" {
			t.Errorf("%s ListenOn = %q, want 0.0.0.0:9013", file, got)
		}
		etcd := mapValue(t, file+".Etcd", config["Etcd"])
		if got := stringValue(etcd["Key"]); got != "storage.rpc" {
			t.Errorf("%s Etcd.Key = %q, want storage.rpc", file, got)
		}
		prometheus := mapValue(t, file+".Prometheus", config["Prometheus"])
		if got := intValue(t, file+".Prometheus.Port", prometheus["Port"]); got != 9106 {
			t.Errorf("%s Prometheus.Port = %d, want 9106", file, got)
		}
	}
}

func TestUserConfigsPreserveRpcIdentityAndMetrics(t *testing.T) {
	for _, file := range []string{
		"services/user/cmd/user/etc/user.yaml",
		"services/user/cmd/user/etc/user-docker.yaml",
	} {
		config := loadProjectYAML(t, file)
		if got := stringValue(config["Name"]); got != "user.rpc" {
			t.Errorf("%s Name = %q, want user.rpc", file, got)
		}
		if got := stringValue(config["ListenOn"]); got != "0.0.0.0:9002" {
			t.Errorf("%s ListenOn = %q, want 0.0.0.0:9002", file, got)
		}
		etcd := mapValue(t, file+".Etcd", config["Etcd"])
		if got := stringValue(etcd["Key"]); got != "user.rpc" {
			t.Errorf("%s Etcd.Key = %q, want user.rpc", file, got)
		}
		prometheus := mapValue(t, file+".Prometheus", config["Prometheus"])
		if got := intValue(t, file+".Prometheus.Port", prometheus["Port"]); got != 9102 {
			t.Errorf("%s Prometheus.Port = %d, want 9102", file, got)
		}
	}
}

func TestUserGenerationRetainsOnlyPublicRpcContracts(t *testing.T) {
	testCases := []struct {
		file     string
		required []string
	}{
		{
			file: "scripts/gen.sh",
			required: []string{
				"--zrpc_out=services/user/rpc",
				"rm -rf services/user/rpc/internal services/user/rpc/etc",
				"rm -f services/user/rpc/user.go",
				"-dir services/user/internal/adapter/model -c",
				"-dir services/user/internal/adapter/model_auth -c",
			},
		},
		{
			file: "scripts/gen.bat",
			required: []string{
				"--zrpc_out=services/user/rpc",
				"rmdir /s /q services\\user\\rpc\\internal",
				"del /q services\\user\\rpc\\user.go",
				"-dir services/user/internal/adapter/model -c",
				"-dir services/user/internal/adapter/model_auth -c",
			},
		},
	}

	for _, testCase := range testCases {
		content, err := os.ReadFile(filepath.Join("..", "..", filepath.FromSlash(testCase.file)))
		if err != nil {
			t.Fatalf("read %s: %v", testCase.file, err)
		}
		for _, snippet := range testCase.required {
			if !strings.Contains(string(content), snippet) {
				t.Errorf("%s missing %q", testCase.file, snippet)
			}
		}
		if strings.Contains(string(content), "-dir services/auth/rpc/internal/model") {
			t.Errorf("%s still generates login logs under the retired auth RPC runtime", testCase.file)
		}
	}
}

func loadDevComposeServices(t *testing.T) map[string]map[string]any {
	t.Helper()
	return loadComposeServices(t, "docker-compose.dev.yml")
}

func loadComposeServices(t *testing.T, path string) map[string]map[string]any {
	t.Helper()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}

	var document struct {
		Services map[string]map[string]any `yaml:"services"`
	}
	if err := yaml.Unmarshal(content, &document); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	return document.Services
}

func defaultLocalServiceIDs(t *testing.T) []string {
	t.Helper()

	repoRoot := filepath.Join("..", "..")
	manifest, err := topology.Load(filepath.Join(repoRoot, "deploy", "topology", "services.json"))
	if err != nil {
		t.Fatalf("load runtime topology: %v", err)
	}
	if err := topology.Validate(repoRoot, manifest); err != nil {
		t.Fatalf("validate runtime topology: %v", err)
	}

	var ids []string
	for _, service := range manifest.Services {
		if service.DefaultLocal {
			ids = append(ids, service.ID)
		}
	}
	return ids
}

func loadProjectYAML(t *testing.T, path string) map[string]any {
	t.Helper()

	content, err := os.ReadFile(filepath.Join("..", "..", filepath.FromSlash(path)))
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}

	var document map[string]any
	if err := yaml.Unmarshal(content, &document); err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	return document
}

func nestedMap(t *testing.T, name string, root map[string]any, path ...string) map[string]any {
	t.Helper()

	current := root
	for _, segment := range path {
		current = mapValue(t, name+"."+strings.Join(path, "."), current[segment])
	}
	return current
}

func mapValue(t *testing.T, name string, value any) map[string]any {
	t.Helper()

	result, ok := value.(map[string]any)
	if !ok {
		t.Fatalf("%s has type %T, want mapping", name, value)
	}
	return result
}

func stringSlice(t *testing.T, name string, value any) []string {
	t.Helper()

	values, ok := value.([]any)
	if !ok {
		t.Fatalf("%s has type %T, want sequence", name, value)
	}

	result := make([]string, 0, len(values))
	for _, value := range values {
		result = append(result, stringValue(value))
	}
	return result
}

func stringValue(value any) string {
	switch value := value.(type) {
	case string:
		return value
	default:
		return ""
	}
}

func intValue(t *testing.T, name string, value any) int {
	t.Helper()

	result, ok := value.(int)
	if !ok {
		t.Fatalf("%s has type %T, want integer", name, value)
	}
	return result
}

func contains(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func containsLine(text, want string) bool {
	for _, line := range strings.Split(text, "\n") {
		if strings.TrimSpace(line) == want {
			return true
		}
	}
	return false
}
