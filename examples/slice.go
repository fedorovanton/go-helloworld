package examples

import "fmt"

func RunSlice() {
	slice1 := []int{10, 11, 12}
	fmt.Println(slice1)

	slice2 := make([]int, 1)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice2 = append(slice2, 13) // добавление одного элемента
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice2 = append(slice2, 14, 15) // добавление двух элементов
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice3 := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}
	for i := range slice3 {
		fmt.Println(slice3[i])
	}
}
