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

func (person Person) getAge() int {
	return person.Age
}

func (dog *Dog) incrementAge() {
	dog.Age++
}

func (dog Dog) getAge() int {
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

	incrementAndPrintAge(kate) // Error: cannot use kate (variable of type Person) as Living value in argument to
	incrementAndPrintAge(mong) // Error: cannot use mong (variable of type Dog) as Living value in argument to

	/** 에러가 발생하는 이유는, Living 인터페이스를 구현한 타입이 각자 포인터 리시버와 값 리시버를 사용하고 있기 때문입니다.
	인터페이스를 구현하는 타입이 포인터 리시버 메서드를 가지고 있다면, 그 타입의 포인터만이 해당 인터페이스를 구현한 것으로 간주됩니다.
	반대로, 값 리시버 메서드를 가지고 있다면, 그 타입의 값만이 해당 인터페이스를 구현한 것으로 간주됩니다.
	따라서, Living 인터페이스를 구현하는 타입이 포인터 리시버와 값 리시버를 혼합해서 사용하고 있기 때문에, Living 인터페이스를 구현한 타입으로 인식되지 않는 것입니다.
	이를 해결하려면, Living 인터페이스를 구현하는 타입에서 사용하는 메서드의 리시버 유형을 일관되게 맞춰야 합니다.
	예를 들어, Person과 Dog 모두 포인터 리시버 또는 값 리시버로 통일하면 됩니다.
	*/
}
