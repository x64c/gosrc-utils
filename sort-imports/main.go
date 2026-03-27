package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		stderr("usage: %s <file.go|directory>\n\n", os.Args[0])
		stderr("Sorts Go import declarations: two groups (stdlib, others), alphabetical within each.\n")
		stderr("Single import uses no parentheses. Multiple import blocks are merged into one.\n\n")
		stderr("Stdlib detection uses `go list std` which inherits your environment.\n")
		stderr("For experimental packages (e.g. encoding/json/v2), set GOEXPERIMENT accordingly.\n")
		stderr("Example: GOEXPERIMENT=jsonv2 %s ./src\n", os.Args[0])
		os.Exit(1)
	}

	if err := loadStdlib(); err != nil {
		stderr("failed to load stdlib list: %v\n", err)
		os.Exit(1)
	}

	target := os.Args[1]
	info, err := os.Stat(target)
	if err != nil {
		stderr("stat: %v\n", err)
		os.Exit(1)
	}

	if info.IsDir() {
		if err := ProcSrcTree(target); err != nil {
			stderr("error: %v\n", err)
			os.Exit(1)
		}
	} else {
		modified, err := ProcSrcFile(target)
		if err != nil {
			stderr("error: %v\n", err)
			os.Exit(1)
		}
		if modified {
			fmt.Println(target)
		}
	}
}
