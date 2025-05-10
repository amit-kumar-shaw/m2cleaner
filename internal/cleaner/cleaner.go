package cleaner

import (
	"fmt"
	"path/filepath"
	"os"
)

var deps = make(map[string]struct{})

func expandPath(path string) (string, error) {
    if path[:2] == "~/" {
        home, err := os.UserHomeDir()
        if err != nil {
            return "", err
        }
        return filepath.Join(home, path[2:]), nil
    }
    return path, nil
}

func RunCleaner(m2path string, projectPaths string) {
	fmt.Printf("M2 Path: %s\n", m2path)
	fmt.Printf("Project Paths: %s\n", projectPaths)

	projectPath, err := expandPath(projectPaths)
	if err != nil {
		fmt.Println("Error resolving project path:", err)
		os.Exit(1)
	}

	filepath.WalkDir(projectPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !d.IsDir() && filepath.Base(path) == "pom.xml" {
			project := ParsePom(path)
			if project != nil {
				fmt.Printf("\nDependencies of %s\n",path)
				WhitelistDeps(*project)
			}
		}

		return nil
	})
}

func WhitelistDeps(proj Project) {
	for _, dep := range proj.Dependencies {
		fmt.Printf("%s %s %s\n",dep.Group, dep.Artifact, dep.Version)
		// deps[dep.group] = struct{}{}
	}
}
