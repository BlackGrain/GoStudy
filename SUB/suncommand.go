package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	foocmd := flag.NewFlagSet("foo", flag.ExitOnError)

	fooEnable := foocmd.Bool("enable", false, "enable")
	fooName := foocmd.String("name", "李宏伟", "name")

	barcmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barcmd.Int("level", 0, "level")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		foocmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo' called")
		fmt.Println("	enable: ", *fooEnable)
		fmt.Println("	name: ", *fooName)
		fmt.Println("	tail: ", foocmd.Args())
	case "bar":
		barcmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar' called")
		fmt.Println("	level: ", *barLevel)
		fmt.Println("	tail: ", barcmd.Args())
	}
}
