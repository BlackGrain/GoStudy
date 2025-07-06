package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	v2 := m["k2"]
	fmt.Println("v2:", v2)

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("maps.Equal(m, n) :", maps.Equal(n2, n))

	for key, value := range n2 {
		fmt.Println("key:", key, "value:", value)
	}

	for _, value := range n {
		fmt.Println("value:", value)
	}
}
