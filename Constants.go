package main

import (
	"fmt"
	"math"
)

const coString string = "String."

func main() {
	fmt.Println(coString)
	const a = 50000000000
	// 数值型常量会在需要的地方推理出 所需要的类型
	const b = 4e20 / a
	fmt.Println(a, b)
	fmt.Println(int64(b))
	fmt.Println(math.Sin(a))
}
