package examples

func RunFor(number int, secondNumber int) {
	// Цикл for
	for i := 1; i <= 10; i++ {
		println(i)
	}

	// Цикл for без объявления начального инкремента
	for number < secondNumber {
		number = number + 1

		println(number)
	}

	// Цикл for с break
	i := 0
	for {
		println("Бесконечный цикл")
		if i == 10 {
			break
		}
		i++
	}
}
