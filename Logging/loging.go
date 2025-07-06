package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("logging with flags")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("logging with shortflags")

	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "ohmy:", log.LstdFlags|log.Lmicroseconds)
	buflog.Println("Buffer log1")
	buflog.Println("Buffer log2")
	fmt.Println("string(buf) :", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi here")

	myslog.Info("hello again", "key", "val", "age", 25)
}
