// Package topology owns the repository's machine-readable runtime topology.
package topology

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const supportedSchemaVersion = 1

var supportedRoles = map[string]struct{}{
	"gateway":        {},
	"merged-service": {},
	"rpc":            {},
	"standalone-api": {},
	"worker":         {},
}

type Manifest struct {
	SchemaVersion            int                 `json:"schemaVersion"`
	DefaultDependencyProfile string              `json:"defaultDependencyProfile"`
	DockerDependencyProfile  string              `json:"dockerDependencyProfile"`
	DependencyProfiles       []DependencyProfile `json:"dependencyProfiles"`
	Services                 []Service           `json:"services"`
}

type DependencyProfile struct {
	ID                string     `json:"id"`
	Description       string     `json:"description"`
	ComposeFile       string     `json:"composeFile,omitempty"`
	ComposeServices   []string   `json:"composeServices,omitempty"`
	RequiredEndpoints []Endpoint `json:"requiredEndpoints"`
}

type Endpoint struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Service struct {
	ID                 string   `json:"id"`
	Role               string   `json:"role"`
	Run                Run      `json:"run"`
	LogFile            string   `json:"logFile"`
	DefaultLocal       bool     `json:"defaultLocal"`
	DependencyProfiles []string `json:"dependencyProfiles"`
	Ports              []Port   `json:"ports"`
}

type Run struct {
	Package string   `json:"package"`
	Config  string   `json:"config"`
	Args    []string `json:"args"`
}

type Port struct {
	Name           string `json:"name"`
	Number         int    `json:"number"`
	Protocol       string `json:"protocol"`
	ManagedOnStart bool   `json:"managedOnStart"`
}

func Load(path string) (Manifest, error) {
	f, err := os.Open(path)
	if err != nil {
		return Manifest{}, fmt.Errorf("open topology manifest: %w", err)
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	decoder.DisallowUnknownFields()

	var manifest Manifest
	if err := decoder.Decode(&manifest); err != nil {
		return Manifest{}, fmt.Errorf("decode topology manifest: %w", err)
	}
	if err := ensureJSONEOF(decoder); err != nil {
		return Manifest{}, err
	}
	return manifest, nil
}

func ensureJSONEOF(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); errors.Is(err, io.EOF) {
		return nil
	} else if err != nil {
		return fmt.Errorf("decode trailing topology data: %w", err)
	}
	return errors.New("topology manifest contains multiple JSON values")
}

func Validate(repoRoot string, manifest Manifest) error {
	if manifest.SchemaVersion != supportedSchemaVersion {
		return fmt.Errorf("schemaVersion must be %d", supportedSchemaVersion)
	}

	profiles := make(map[string]struct{}, len(manifest.DependencyProfiles))
	for i, profile := range manifest.DependencyProfiles {
		label := fmt.Sprintf("dependencyProfiles[%d]", i)
		if profile.ID == "" {
			return fmt.Errorf("%s.id is required", label)
		}
		if _, exists := profiles[profile.ID]; exists {
			return fmt.Errorf("duplicate dependency profile id %q", profile.ID)
		}
		profiles[profile.ID] = struct{}{}

		if (profile.ComposeFile == "") != (len(profile.ComposeServices) == 0) {
			return fmt.Errorf("dependency profile %q must set composeFile and composeServices together", profile.ID)
		}
		if profile.ComposeFile != "" {
			if err := validateRepoPath(repoRoot, profile.ComposeFile, false); err != nil {
				return fmt.Errorf("dependency profile %q composeFile: %w", profile.ID, err)
			}
			if err := validateComposeServices(repoRoot, profile.ComposeFile, profile.ComposeServices); err != nil {
				return fmt.Errorf("dependency profile %q: %w", profile.ID, err)
			}
		}

		endpoints := make(map[string]struct{}, len(profile.RequiredEndpoints))
		for j, endpoint := range profile.RequiredEndpoints {
			if endpoint.Name == "" || endpoint.Host == "" {
				return fmt.Errorf("%s.requiredEndpoints[%d] requires name and host", label, j)
			}
			if endpoint.Port < 1 || endpoint.Port > 65535 {
				return fmt.Errorf("%s.requiredEndpoints[%d] has invalid port %d", label, j, endpoint.Port)
			}
			key := fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
			if _, exists := endpoints[key]; exists {
				return fmt.Errorf("dependency profile %q has duplicate endpoint %s", profile.ID, key)
			}
			endpoints[key] = struct{}{}
		}
	}

	for field, profileID := range map[string]string{
		"defaultDependencyProfile": manifest.DefaultDependencyProfile,
		"dockerDependencyProfile":  manifest.DockerDependencyProfile,
	} {
		if _, exists := profiles[profileID]; !exists {
			return fmt.Errorf("%s references unknown profile %q", field, profileID)
		}
	}

	serviceIDs := make(map[string]struct{}, len(manifest.Services))
	ports := make(map[int]string)
	defaultLocalCount := 0
	for i, service := range manifest.Services {
		label := fmt.Sprintf("services[%d]", i)
		if service.ID == "" {
			return fmt.Errorf("%s.id is required", label)
		}
		if _, exists := serviceIDs[service.ID]; exists {
			return fmt.Errorf("duplicate service id %q", service.ID)
		}
		serviceIDs[service.ID] = struct{}{}

		if _, supported := supportedRoles[service.Role]; !supported {
			return fmt.Errorf("service %q has unsupported role %q", service.ID, service.Role)
		}
		if err := validateRepoPath(repoRoot, service.Run.Package, true); err != nil {
			return fmt.Errorf("service %q package: %w", service.ID, err)
		}
		if err := validateRepoPath(repoRoot, service.Run.Config, false); err != nil {
			return fmt.Errorf("service %q config: %w", service.ID, err)
		}
		if service.LogFile == "" || filepath.Base(service.LogFile) != service.LogFile {
			return fmt.Errorf("service %q logFile must be a file name", service.ID)
		}
		if service.DefaultLocal {
			defaultLocalCount++
		}

		profileRefs := make(map[string]struct{}, len(service.DependencyProfiles))
		for _, profileID := range service.DependencyProfiles {
			if _, exists := profiles[profileID]; !exists {
				return fmt.Errorf("service %q references unknown dependency profile %q", service.ID, profileID)
			}
			if _, duplicate := profileRefs[profileID]; duplicate {
				return fmt.Errorf("service %q repeats dependency profile %q", service.ID, profileID)
			}
			profileRefs[profileID] = struct{}{}
		}
		if len(profileRefs) == 0 {
			return fmt.Errorf("service %q requires at least one dependency profile", service.ID)
		}

		portNames := make(map[string]struct{}, len(service.Ports))
		for j, port := range service.Ports {
			if port.Name == "" || port.Protocol == "" {
				return fmt.Errorf("%s.ports[%d] requires name and protocol", label, j)
			}
			if port.Number < 1 || port.Number > 65535 {
				return fmt.Errorf("service %q port %q has invalid number %d", service.ID, port.Name, port.Number)
			}
			if _, exists := portNames[port.Name]; exists {
				return fmt.Errorf("service %q has duplicate port name %q", service.ID, port.Name)
			}
			portNames[port.Name] = struct{}{}
			if owner, exists := ports[port.Number]; exists {
				return fmt.Errorf("port %d is shared by services %q and %q", port.Number, owner, service.ID)
			}
			ports[port.Number] = service.ID
		}
	}
	if defaultLocalCount == 0 {
		return errors.New("at least one service must have defaultLocal enabled")
	}

	return nil
}

func validateRepoPath(repoRoot, path string, wantDir bool) error {
	if path == "" {
		return errors.New("path is required")
	}
	clean := filepath.Clean(filepath.FromSlash(path))
	if filepath.IsAbs(clean) || clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return fmt.Errorf("path %q must stay relative to the repository root", path)
	}

	info, err := os.Stat(filepath.Join(repoRoot, clean))
	if err != nil {
		return fmt.Errorf("path %q: %w", path, err)
	}
	if info.IsDir() != wantDir {
		if wantDir {
			return fmt.Errorf("path %q is not a directory", path)
		}
		return fmt.Errorf("path %q is not a file", path)
	}
	return nil
}

func validateComposeServices(repoRoot, composeFile string, serviceNames []string) error {
	content, err := os.ReadFile(filepath.Join(repoRoot, filepath.FromSlash(composeFile)))
	if err != nil {
		return fmt.Errorf("read composeFile: %w", err)
	}

	var document struct {
		Services map[string]any `yaml:"services"`
	}
	if err := yaml.Unmarshal(content, &document); err != nil {
		return fmt.Errorf("decode composeFile: %w", err)
	}

	seen := make(map[string]struct{}, len(serviceNames))
	for _, serviceName := range serviceNames {
		if serviceName == "" {
			return errors.New("composeServices contains an empty service name")
		}
		if _, duplicate := seen[serviceName]; duplicate {
			return fmt.Errorf("composeServices repeats service %q", serviceName)
		}
		seen[serviceName] = struct{}{}
		if _, exists := document.Services[serviceName]; !exists {
			return fmt.Errorf("compose service %q is not defined in %q", serviceName, composeFile)
		}
	}
	return nil
}
