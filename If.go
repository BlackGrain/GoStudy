package main

import "fmt"

func main() {
	// 定义语句可以写在 条件语句前面
	if a := 9; a < 9 {
		fmt.Println("a<9")
	} else if a == 9 {
		fmt.Println("a=9")
	} else {
		fmt.Println("a<9")
	}
}
