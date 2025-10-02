package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func nameAndAge(uid int) User {
	switch uid {
	case 0:
		return User{"Baheer", 24}
	case 1:
		return User{"Tanmay", 25}
	default:
		return User{"Unknown", -1}
	}
}

func main() {
	user := nameAndAge(0)
	fmt.Println("User age: ")
	fmt.Println(user.Age)
}
