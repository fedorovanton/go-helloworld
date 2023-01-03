package main

import (
	"log"
	"net/http"
	"strings"
)

// @see https://proglib.io/p/samouchitel-dlya-nachinayushchih-kak-osvoit-go-s-nulya-za-30-minut-2021-07-12

func main() {
	println("Привет Ева")

	var number, secondNumber int
	number = 10
	println(number)
	secondNumber = 15
	println(secondNumber)
	c := "Строка"
	println(c)

	// Условие if
	if number < secondNumber {
		println("первое число меньше второго")
	} else if number > secondNumber {
		println("второе число меньше первого")
	} else {
		println("числа равны")
	}

	// Условие switch
	switch number {
	case 10:
		println("число =10")
	case 15:
		println("число =15")
	}

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

	// Массив1
	// var numbers[3] int
	// numbers[0] = 10
	// numbers[1] = 11
	// numbers[2] = 12

	// Массив2
	// var numbers [3]int = [3]int{10, 11, 12}

	// Массив3
	// numbers := [3]int{10, 11, 12}

	// Слайс
	numbers := []int{10, 11, 12}
	numbers = append(numbers, 13)     // добавление одного элемента
	numbers = append(numbers, 14, 15) // добавление двух элементов

	sum := Sum(1, 2)
	println(sum) // 3

	sum, division := SumAndDivide(4, 2)
	println(sum)      // 6
	println(division) // 2

	upper := strings.ToUpper("заглавные буквы")
	println(upper) // Output: ЗАГЛАВНЫЕ БУКВЫ

	http.HandleFunc("/", HelloWeb)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Println("listen and serve:", err)
	}
}

func НазваниеФункции(первыйАргумент int, второйАргумент int) (возвращаемоеЗначение int) {
	// тело функции
	return 0
}

func Sum(a, b int) int {
	return a + b
}

func SumAndDivide(a, b int) (int, int) {
	return a + b, a / b
}

func HelloWeb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Web!"))
}
