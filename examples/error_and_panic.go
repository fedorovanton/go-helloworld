package examples

import (
	"errors"
	"fmt"
)

type Age int

func has18Age(age Age) (string, error) {
	if age < 18 {
		return "", errors.New("Ошибка: Возраст меньше 18!")
	}

	return "Возраст больше или равен 18", nil
}

func RunErrorAndPanic() {
	// Функция с 2 возвращаемым значением (результат и error)
	res, err := has18Age(17)
	// res, err := has18Age(19)
	if err != nil {
		fmt.Println(err)

		panic("Опачки! Тут у нас паника :)")
	} else {
		fmt.Println(res)
	}
}
