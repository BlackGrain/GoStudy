package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("end")
	os.Exit(3)
}
