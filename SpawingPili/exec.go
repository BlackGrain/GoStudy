package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {
	dateCmd := exec.Command("date")
	dateOut, _ := dateCmd.Output()
	fmt.Println(string(dateOut))

	dateCmd = exec.Command("date", "-x")
	dateOut, err := dateCmd.Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		case *exec.Error:
			fmt.Println("ExitError: ", err)
		default:
			panic(err)
		}
	}

	grepCommand := exec.Command("grep", "hello")
	inpi, _ := grepCommand.StdinPipe()
	outpi, _ := grepCommand.StdoutPipe()
	grepCommand.Start()
	inpi.Write([]byte("1. hello grep\n 2. 你好 \n 3. hello china"))
	inpi.Close()
	fmt.Println("> grep hello")

	grepBytes, _ := io.ReadAll(outpi)
	grepCommand.Wait()
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))

}
