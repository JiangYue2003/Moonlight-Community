package topology

import (
	"os"
	"path/filepath"
	"testing"
)

func TestStage5LegacyHTTPTreesAreRetired(t *testing.T) {
	repoRoot := filepath.Clean(filepath.Join("..", ".."))
	retired := []string{
		"services/agent/api",
		"services/auth/api",
		"services/counter/api",
		"services/knowpost/api",
		"services/llm/api",
		"services/profile/api",
		"services/relation/api",
		"services/search/api",
		"services/storage/api",
	}
	for _, rel := range retired {
		path := filepath.Join(repoRoot, filepath.FromSlash(rel))
		err := filepath.WalkDir(path, func(path string, entry os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !entry.IsDir() {
				t.Errorf("retired legacy HTTP tree still contains file: %s", filepath.ToSlash(path))
			}
			return nil
		})
		if err != nil && !os.IsNotExist(err) {
			t.Errorf("stat retired legacy HTTP tree %s: %v", rel, err)
		}
	}
}
