package examples

import "fmt"

func RunArray() {
	// Массив1
	var arr1 [3]int
	arr1[0] = 10
	arr1[1] = 11
	arr1[2] = 12
	fmt.Println(arr1)

	// Массив2
	var arr2 [3]int = [3]int{10, 11, 12}
	fmt.Println(arr2)

	// Массив3
	arr3 := [3]int{20, 21, 22}
	fmt.Println(arr3)
}
