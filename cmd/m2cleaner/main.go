package main

import (
	"flag"
	"fmt"
    "github.com/amit-kumar-shaw/m2cleaner/internal/cleaner"
)

var tuiMode bool
var m2Path string
var projectPaths string

func init() {
    flag.BoolVar(&tuiMode, "tui", false, "Run with TUI interface")
    flag.BoolVar(&tuiMode, "t", false, "Run with TUI interface")
	flag.StringVar(&m2Path, "repo", "~/.m2/repository", "M2 repository path")
	flag.StringVar(&m2Path, "r", "~/.m2/repository", "M2 repository path")
	flag.StringVar(&projectPaths, "project", "~", "M2 repository path")
	flag.StringVar(&projectPaths, "p", "~", "M2 repository path")
}

func main() {
    flag.Parse()

    if tuiMode {
        fmt.Println("Running TUI version of m2cleaner")
    } else {
        fmt.Println("Running CLI version of m2cleaner")
        cleaner.RunCleaner(m2Path, projectPaths)
    }
}