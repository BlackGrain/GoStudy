package main

import (
	"fmt"
	"os"
)

func main() {
	//var s, sep string
	// echo v1
	/*	for i := 1; i < len(os.Args); i++ { // 从1开始是为了跳过程序名
			s = s + sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)*/

	// echo v2
	/*	for _, arg := range os.Args[1:] {
			s = s + sep + arg
			sep = " "
		}
		fmt.Println(s)*/

	// echo v3
	//fmt.Println(strings.Join(os.Args[1:], " "))

	// echo v4
	fmt.Println(os.Args[1:])
}
