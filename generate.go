package main

import (
	"fmt"
	"github.com/akyoto/autoimport"
	"os"
	"path"
	"sort"
	"strings"
)

func main() {
	home, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	// Find packages
	var filteredPackages []*autoimport.Package
	projectsPath := path.Join(home, "projects")
	packageLists := autoimport.GetPackagesInDirectory(projectsPath, projectsPath)

	for _, packageList := range packageLists {
		for _, pkg := range packageList {
			if !pkg.IsModuleRoot {
				continue
			}

			filteredPackages = append(filteredPackages, pkg)
		}
	}

	// Sort by import path
	sort.Slice(filteredPackages, func(i int, j int) bool {
		return filteredPackages[i].ImportPath < filteredPackages[j].ImportPath
	})

	// Header
	fmt.Println(strings.ReplaceAll(`
		| Package | Path | Report | Tests | Coverage |
		|---------|------|--------|-------|----------|`,
		"\t",
		"",
	))

	// Package lines
	for _, pkg := range filteredPackages {
		generate(pkg)
	}

	// URLs
	for _, pkg := range filteredPackages {
		generateURLs(pkg)
	}
}

func generate(pkg *autoimport.Package) {
	fmt.Printf(
		"| %s | [%s](https://github.com/%s) |[![Report][report-image-%s]][report-url-%s] | [![Tests][tests-image-%s]][tests-url-%s] | [![Coverage][codecov-image-%s]][codecov-url-%s] |\n",
		pkg.Name,
		pkg.ImportPath,
		pkg.ImportPath,
		pkg.Name,
		pkg.Name,
		pkg.Name,
		pkg.Name,
		pkg.Name,
		pkg.Name,
	)
}

func generateURLs(pkg *autoimport.Package) {
	fmt.Printf(
		strings.ReplaceAll(`
			[report-image-%s]: https://goreportcard.com/badge/github.com/%s
			[report-url-%s]: https://goreportcard.com/report/github.com/%s
			[tests-image-%s]: https://cloud.drone.io/api/badges/%s/status.svg
			[tests-url-%s]:  https://cloud.drone.io/%s
			[codecov-image-%s]: https://codecov.io/gh/%s/graph/badge.svg
			[codecov-url-%s]: https://codecov.io/gh/%s`,
			"\t",
			"",
		),
		pkg.Name, pkg.ImportPath,
		pkg.Name, pkg.ImportPath,
		pkg.Name, pkg.ImportPath,
		pkg.Name, pkg.ImportPath,
		pkg.Name, pkg.ImportPath,
		pkg.Name, pkg.ImportPath,
	)

	fmt.Println()
}
