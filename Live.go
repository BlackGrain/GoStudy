package main

import "fmt"

func main() {
	arr1 := make([]int, 3, 5)
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr2 := []int{0, 0, 0, 0}
	fmt.Println(arr1)
	fmt.Println(arr2)
	arr2 = append(arr1, 4)
	arr1[1] = 9
	fmt.Println(arr1)
	fmt.Println(arr2)

}
