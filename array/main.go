package main

import "fmt"

func main() {
	// Fix array
	var arr [5]int
	fmt.Printf("arr=%v\n", arr)

	// Loop
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	// Loop 2
	for v := range arr {
		fmt.Printf("%v ", v*10)
	}
	fmt.Println()

	// Define array
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("arr2=%v\n", arr2)
	fmt.Printf("arr2 (partial)=%v\n", arr2[2:4])

	//
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
