package main

import (
	"errors"
	"fmt"
	"helloworld/shape"
	"net/http"
	"strings"
)

func main() {
	// Обработка возникающих panic
	defer handlerPanic()

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
	// numbers := []int{10, 11, 12}
	numbers := make([]int, 1)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	numbers = append(numbers, 13) // добавление одного элемента
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	numbers = append(numbers, 14, 15) // добавление двух элементов
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}
	for i := range messages {
		fmt.Println(messages[i])
	}

	// Функция с 1 возвращаемым значением
	sum := Sum(1, 2)
	println(sum) // 3

	// Функция с 2 возвращаемыми значениями
	sum, division := SumAndDivide(4, 2)
	println(sum)      // 6
	println(division) // 2

	// Отложенный вызов функции (после main())
	defer printMessageLast()

	// Функция с 2 возвращаемым значением (результат и error)
	// res, err := has18Age(17)
	res, err := has18Age(19)
	if err != nil {
		fmt.Println(err)

		panic("Опачки! Ошибочка")
	} else {
		fmt.Println(res)
	}

	upper := strings.ToUpper("заглавные буквы")
	println(upper) // Output: ЗАГЛАВНЫЕ БУКВЫ

	// Ссылки, указатели (работа с памятью)
	message := "Скоро я стану ниндзя!"
	println(message)
	changeMessage(&message)
	println(message)

	// Мапэ (map)
	users := map[string]int{
		"Ivanov":  10,
		"Sidorov": 15,
		"Petrov":  20,
	}
	// users := make(map[string]int)
	fmt.Println(users)
	age, exists := users["Ivanov"]
	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("Ivanov нет в списке")
	}
	users["Seregov"] = 25
	delete(users, "Petrov")
	for key, value := range users {
		fmt.Println(key, value)
	}

	// Web Server
	// http.HandleFunc("/", HelloWeb)
	// err := http.ListenAndServe("localhost:9999", nil)
	// if err != nil {
	// 	log.Println("listen and serve:", err)
	// }

	// Работа со структурами (struct)
	user1 := constructUser("Vasya", "Male", 23, 75, 185)
	user2 := User{"Petya", 29, "Male", 83, 180}
	fmt.Printf("%+v\n", user1)
	user1.printUserInfo()
	user2.setAgeUser(99)
	user2.printUserInfo()

	// Работа с интерфейсами
	square := shape.NewSquare(5)
	// circle := Circle{8}
	shape.PrintShapeArea(square)
	// printShapeArea(circle)

	shape.PrintInterface(square)
	// printInterface(circle)
}

func constructUser(name, sex string, age Age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    Age(age),
		weight: weight,
		height: height,
	}
}

// (u User) - это ресивер (reciever) структуры (User) для метода printUserInfo
func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

// (u User) - это поинт ресивер (point reciever) структуры (User) для метода printUserInfo
func (u *User) setAgeUser(age Age) {
	u.age = age
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

func changeMessage(message *string) {
	*message += " (из функции printMessage())"
}

func Sum(a, b int) int {
	return a + b
}

func SumAndDivide(a, b int) (int, int) {
	return a + b, a / b
}

func has18Age(age Age) (string, error) {
	if age < 18 {
		return "", errors.New("error: Еще не имеет 18 лет!")
	}

	return "Все ОК", nil
}

func HelloWeb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Web!"))
}

func printMessageLast() {
	fmt.Println("Эта функция сработала в конце, хотя ее вызвали в середине")
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

// @see https://youtu.be/h0zxh2TPN_I?t=10240
