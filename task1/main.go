package main

import "fmt"

type FizzBuzz struct {
	Name string
}

var fizz, buzz FizzBuzz

func main() {
	fizz = FizzBuzz{"Fizz"}
	buzz = FizzBuzz{"Buzz"}
	printFizzBuzz(100)
}

func printFizzBuzz(number int) {
	for i := 1; i <= number; i++ {

		if i%3 == 0 {
			fmt.Printf(fizz.Name)
		}
		if i%5 == 0 {
			fmt.Printf(buzz.Name)
		}

		if needPrintNumber(i) {
			fmt.Printf("%d", i)
		}

		fmt.Printf("\n")

	}
}

func needPrintNumber(i int) bool {
	return i%3 != 0 && i%5 != 0
}
