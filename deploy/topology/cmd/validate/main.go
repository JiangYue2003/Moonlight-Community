package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zhiguang/zhiguang-go/deploy/topology"
)

func main() {
	repoRoot := flag.String("repo-root", ".", "repository root used to resolve manifest paths")
	manifestPath := flag.String("manifest", "deploy/topology/services.json", "topology manifest path")
	flag.Parse()

	path := *manifestPath
	if !filepath.IsAbs(path) {
		path = filepath.Join(*repoRoot, path)
	}
	manifest, err := topology.Load(path)
	if err == nil {
		err = topology.Validate(*repoRoot, manifest)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "topology validation failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("topology valid: %d services, %d dependency profiles\n", len(manifest.Services), len(manifest.DependencyProfiles))
}
