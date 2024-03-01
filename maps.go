package main

import "fmt"

// type aliases
type floatMap map[string]float64

// function for my type aliases
func (m floatMap) cw() {
	fmt.Println(m)
}

func main() {
	websites := []string{"https://google.com", "https://aws.com"}
	websitesMap := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websitesMap)
	fmt.Println(websites)

	// adding
	websitesMap["LinkedIn"] = "https://linkedin.com"

	// deleting from maps
	delete(websitesMap, "Google")

	courseRatings := make(map[string]float64, 3)
	typeCourseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.9

	typeCourseRatings.cw()
	fmt.Println(courseRatings)
}
