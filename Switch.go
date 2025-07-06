package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekend")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before 12")
	default:
		fmt.Println("It's after 12")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			println("I'm a bool")
		case int:
			println("I'm a int")
		case string:
			println("I'm a string")

		default:
			println("unkown: ", t)
		}
	}

	whatAmI(true)
	whatAmI("hey")
	whatAmI(1)
}
