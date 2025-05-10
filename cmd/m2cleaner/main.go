package main

import (
	"flag"
	"fmt"
)

var tuiMode bool

func init() {
    flag.BoolVar(&tuiMode, "tui", false, "Run with TUI interface")
    flag.BoolVar(&tuiMode, "t", false, "Run with TUI interface")
}

func main() {
    flag.Parse()

    if tuiMode {
        fmt.Println("Running TUI version of m2cleaner")
    } else {
        fmt.Println("Running CLI version of m2cleaner")
    }
}