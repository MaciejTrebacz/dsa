package main

import (
	"fmt"
)

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{1.99, 2.99, 3.99, 4.99}
	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(productNames[0])
	productNames[2] = "2 BOOK"

	//slices
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	// getting number of elements in array
	fmt.Println(len(prices))

	//
	fmt.Println(cap(prices))

	// getting from start of to end
	//          prices[:1]   prices[1:]
	fmt.Println(prices[:1])
	fmt.Println(prices[1:])

	dynamicPrices := []float64{10.99, 8.99}
	fmt.Println(dynamicPrices[0:1])
	prices[1] = 9.99
	prices[2] = 10.00
	newArray := append(dynamicPrices, 5.99)
	fmt.Println(newArray)

	discountPrices2 := []float64{101.99, 80.99, 20.59}
	dynamicPrices = append(dynamicPrices, discountPrices2...) // pull out elements from list and use them as separate elements

	// using make function

	userNames := make([]string, 2, 5)
	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	fmt.Println(userNames)
}
