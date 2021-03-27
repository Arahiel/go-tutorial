package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i < 11; i++ {
		numbers = append(numbers, i)
	}

	for i, number := range numbers {
		fmt.Printf("%v is %v\n", i, toString(isEven(number)))
	}
}

func isEven(number int) bool {
	return number%2 == 0
}

func toString(isEven bool) string {
	if isEven {
		return "even"
	} else {
		return "odd"
	}
}
