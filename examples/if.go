package examples

func RunIf(number int, secondNumber int) {
	// Условие if
	if number < secondNumber {
		println("первое число меньше второго")
	} else if number > secondNumber {
		println("второе число меньше первого")
	} else {
		println("числа равны")
	}
}
