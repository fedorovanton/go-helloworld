package examples

import "fmt"

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

// создание новой структуры
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

func RunStruct() {
	user1 := constructUser("Vasya", "Male", 23, 75, 185)
	user2 := User{"Petya", 29, "Male", 83, 180}
	fmt.Printf("%+v\n", user1)
	user1.printUserInfo()
	user2.setAgeUser(99)
	user2.printUserInfo()
}
