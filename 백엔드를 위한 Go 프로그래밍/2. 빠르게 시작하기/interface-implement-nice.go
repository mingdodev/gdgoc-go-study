package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Dog struct {
	Name  string
	Owner *Person
	Age   int
}

func (person *Person) incrementAge() {
	person.Age++
}

func (person *Person) getAge() int {
	return person.Age
}

func (dog *Dog) incrementAge() {
	dog.Age++
}

func (dog *Dog) getAge() int {
	return dog.Age
}

type Living interface {
	incrementAge()
	getAge() int
}

func incrementAndPrintAge(being Living) {
	being.incrementAge()
	fmt.Println(being.getAge())
}

func main() {
	kate := Person{"kate", 19}
	mong := Dog{"mong", &kate, 3}

	incrementAndPrintAge(&kate)
	incrementAndPrintAge(&mong)
}
