package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup1: %v\n", err)
			continue
		}
		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
