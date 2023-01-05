package examples

import "fmt"

func RunMap() {
	// Мапэ (map)
	users1 := map[string]int{
		"Ivanov":  10,
		"Sidorov": 15,
		"Petrov":  20,
	}
	fmt.Println(users1)

	users2 := make(map[string]int)
	fmt.Println(users2)

	// Существует ли элемент по ключу
	age, exists := users1["Ivanov"]
	if exists {
		fmt.Println(age)
	} else {
		fmt.Println("Ivanov нет в списке")
	}

	// Добавить новый элемент
	users1["Seregov"] = 25
	delete(users1, "Petrov")
	for key, value := range users1 {
		fmt.Println(key, value)
	}
}
