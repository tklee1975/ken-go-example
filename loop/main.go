package main

import "fmt"

func main() {
	fmt.Println("for 1 to 3")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i=%d\n", i)
	}

	//
	fmt.Println("----------------")

	//
	k := 3
	for {
		fmt.Printf("k=%d\n", k)
		if k == 0 {
			break
		}
		k--
	}

	fmt.Println("----------------")

	list := []int{11, 12, 13, 14, 15}
	//
	for i, v := range list {
		fmt.Printf("i=%d, v=%d\n", i, v)
	}

}
