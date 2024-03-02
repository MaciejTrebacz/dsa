package main

import "fmt"

func main() {
	numbers := []int{1, 10, 100}
	fmt.Println(sumup(numbers[0], 10, 44, 55))

	// calling variadic function

	fmt.Println(sumup(1, numbers...))
}

func sumup(numbers int, numbers2 ...int) int {
	sum := 0

	for _, val := range numbers2 {
		sum += val
	}
	return sum
}
