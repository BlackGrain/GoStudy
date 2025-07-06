package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("foo", "1")
	fmt.Println(os.Getenv("foo"))
	fmt.Println(os.Getenv("bar"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], ":", pair[1])
	}
}
