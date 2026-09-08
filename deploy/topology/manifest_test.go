package topology

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestRepositoryManifestIsValid(t *testing.T) {
	repoRoot := filepath.Clean(filepath.Join("..", ".."))
	manifest, err := Load(filepath.Join(repoRoot, "deploy", "topology", "services.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := Validate(repoRoot, manifest); err != nil {
		t.Fatal(err)
	}
}

func TestLoadRejectsUnknownFields(t *testing.T) {
	path := filepath.Join(t.TempDir(), "services.json")
	if err := os.WriteFile(path, []byte(`{"schemaVersion":1,"unexpected":true}`), 0o600); err != nil {
		t.Fatal(err)
	}

	_, err := Load(path)
	if err == nil || !strings.Contains(err.Error(), "unknown field") {
		t.Fatalf("Load() error = %v, want unknown field error", err)
	}
}

func TestRepositoryManifestOutboxGCUsesPositionalConfigArgument(t *testing.T) {
	repoRoot := filepath.Clean(filepath.Join("..", ".."))
	manifest, err := Load(filepath.Join(repoRoot, "deploy", "topology", "services.json"))
	if err != nil {
		t.Fatal(err)
	}

	for _, service := range manifest.Services {
		if service.ID != "outbox-gc" {
			continue
		}
		want := []string{service.Run.Config}
		if !reflect.DeepEqual(service.Run.Args, want) {
			t.Fatalf("outbox-gc args = %v, want positional config argument %v", service.Run.Args, want)
		}
		return
	}
	t.Fatal("outbox-gc is missing from the repository topology")
}

func TestValidateRejectsUnknownComposeService(t *testing.T) {
	repoRoot, manifest := testManifest(t)
	composePath := filepath.Join(repoRoot, "config", "compose.yml")
	if err := os.WriteFile(composePath, []byte("services:\n  etcd: {}\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	manifest.DependencyProfiles[0].ComposeFile = "config/compose.yml"
	manifest.DependencyProfiles[0].ComposeServices = []string{"missing"}

	err := Validate(repoRoot, manifest)
	if err == nil || !strings.Contains(err.Error(), `compose service "missing" is not defined`) {
		t.Fatalf("Validate() error = %v, want missing Compose service error", err)
	}
}

func TestValidateRejectsTopologyDrift(t *testing.T) {
	tests := []struct {
		name   string
		mutate func(*Manifest)
		want   string
	}{
		{name: "duplicate service id", mutate: func(m *Manifest) {
			duplicate := m.Services[0]
			duplicate.Ports = nil
			m.Services = append(m.Services, duplicate)
		}, want: "duplicate service id"},
		{name: "duplicate port", mutate: func(m *Manifest) {
			duplicate := m.Services[0]
			duplicate.ID = "other"
			m.Services = append(m.Services, duplicate)
		}, want: "port 8080 is shared"},
		{name: "unsupported role", mutate: func(m *Manifest) { m.Services[0].Role = "controller" }, want: "unsupported role"},
		{name: "missing config", mutate: func(m *Manifest) { m.Services[0].Run.Config = "config/missing.yaml" }, want: "missing.yaml"},
		{name: "unknown dependency profile", mutate: func(m *Manifest) { m.Services[0].DependencyProfiles = []string{"missing"} }, want: "unknown dependency profile"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoRoot, manifest := testManifest(t)
			tt.mutate(&manifest)
			err := Validate(repoRoot, manifest)
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("Validate() error = %v, want substring %q", err, tt.want)
			}
		})
	}
}

func testManifest(t *testing.T) (string, Manifest) {
	t.Helper()
	repoRoot := t.TempDir()
	if err := os.MkdirAll(filepath.Join(repoRoot, "services", "gateway"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(repoRoot, "config"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(repoRoot, "config", "gateway.yaml"), nil, 0o600); err != nil {
		t.Fatal(err)
	}

	return repoRoot, Manifest{
		SchemaVersion:            supportedSchemaVersion,
		DefaultDependencyProfile: "host",
		DockerDependencyProfile:  "host",
		DependencyProfiles: []DependencyProfile{
			{ID: "host", RequiredEndpoints: []Endpoint{{Name: "mysql", Host: "127.0.0.1", Port: 3306}}},
		},
		Services: []Service{
			{
				ID:                 "gateway",
				Role:               "gateway",
				Run:                Run{Package: "services/gateway", Config: "config/gateway.yaml"},
				LogFile:            "gateway.log",
				DefaultLocal:       true,
				DependencyProfiles: []string{"host"},
				Ports:              []Port{{Name: "http", Number: 8080, Protocol: "http", ManagedOnStart: true}},
			},
		},
	}
}
