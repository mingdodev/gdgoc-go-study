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

	incrementAndPrintAge(&kate)
	incrementAndPrintAge(&mong)

	/** 에러가 발생하지 않는 이유는, Living 인터페이스를 구현한 타입이 모두 포인터 리시버를 사용할 수 있기 때문입니다.
	getAge 메서드는 값 리시버로 정의되어 있지만, 값 타입(Person)과 포인터 타입(*Person) 모두에서 호출될 수 있습니다.
	컴파일러가 포인터에서 값을 가져오는 역참조(dereferencing)를 자동으로 처리해주기 때문입니다.
	하지만 이는 좋은 코드는 아닙니다. 일관성있게 코드를 써야지!
	실제로 Struct Dog has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
	라는 경고도 뜹니다.
	*/
}
