package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/otiai10/primes"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage:", os.Args[0], "‹number›")
		os.Exit(1)
	}

	number, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	f := primes.Factorize(int64(number))
	fmt.Println("primes:", len(f.Powers()) == 1)
}
