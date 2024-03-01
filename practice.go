package main

import "fmt"

func Practice() {
	var hobbies = []string{"girls", "bikes", "food"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	firstElement := hobbies[:2]
	fmt.Println(firstElement)

	firstElement = firstElement[1:3] // cannot use shortcut if we reassigning from existing slice SO IF YOU EXPANDING YOU HAVE TO SPECIFY

}
