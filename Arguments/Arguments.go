package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argWithoutProg := os.Args[1:]

	arg := argsWithProg[3]

	fmt.Println(arg)
	fmt.Println(argsWithProg)
	fmt.Println(argWithoutProg)
	fmt.Println(arg)
}
