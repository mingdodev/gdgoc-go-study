package main

import "fmt"

func squareIt(inputChan, outputChan chan int) {
	for x := range inputChan {
		outputChan <- x * x
	}
}

func main() {
	inputChan := make(chan int)
	outputChan := make(chan int, 5)

	go squareIt(inputChan, outputChan)
	for i := 0; i < 5; i++ {
		inputChan <- i
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-outputChan)
	}
	close(inputChan)
}
