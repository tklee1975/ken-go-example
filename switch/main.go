package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	checkType := func(i interface{}) string {
		switch i.(type) {
		case bool:
			return "bool"
		case int:
			return "int"
		case string:
			return "string"
		default:
			return "unknown"
		}
	}

	fmt.Printf("%v\n", checkType(true))
	fmt.Printf("%v\n", checkType(1))
	fmt.Printf("%v\n", checkType("hello"))
}
