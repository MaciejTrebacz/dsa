package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 1, 2}
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled, tripled)

	transformerFn1 := getTransformerFunction(&numbers)
	transformerFn2 := getTransformerFunction(&moreNumbers)
	transformedNumbers := transformNumbers(&numbers, transformerFn1)
	moreTransformNumbers := transformNumbers(&moreNumbers, transformerFn2)
	fmt.Println(transformedNumbers, moreTransformNumbers)
}

// using types
type transformFn func(int) int

// taking func as parameter
func transformNumbers(numbers *[]int, transform transformFn) []int {
	var dNumbers []int
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
