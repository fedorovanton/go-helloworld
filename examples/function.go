package examples

func Sum(a, b int) int {
	return a + b
}

func SumAndDivide(a, b int) (int, int) {
	return a + b, a / b
}

func RunFunction() {
	// Функция с 1 возвращаемым значением
	sum := Sum(1, 2)
	println(sum) // 3

	// Функция с 2 возвращаемыми значениями
	sum, division := SumAndDivide(4, 2)
	println(sum)      // 6
	println(division) // 2
}
