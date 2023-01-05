package examples

func RunPointers() {
	// Ссылки, указатели (работа с памятью)
	message := "Скоро я стану ниндзя!"
	println(message)
	changeMessage(&message)
	println(message)
}

func changeMessage(message *string) {
	*message += " (из функции printMessage())"
}
