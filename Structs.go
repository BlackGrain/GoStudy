package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name}
	p.age = age
	return &p
}

func main() {
	fmt.Println(person{name: "张三", age: 12})
	fmt.Println(newPerson("Bob", 20))
	fmt.Println(person{name: "李四"})
	fmt.Println(&person{
		name: "李霞龙", age: 26,
	})

	s := person{name: "李显龙", age: 18}
	fmt.Println(s.name)
	sp := &s
	fmt.Println(sp.age)
	sp.age = 51
	fmt.Println(s.age)

	dog := struct {
		name string
		age  int
	}{
		"利好",
		12,
	}

	fmt.Println(dog.age)

}
