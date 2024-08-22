package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) outputMap() {
	fmt.Println(m)
}

func mapPractice() {
	urlCollection := map[string]string{
		"Google": "http://www.google.com",
		"AWS":    "http://aws.amazon",
		"Lamda":  "http://lamda.com",
	}
	fmt.Println(urlCollection)

	urlCollection["Linkedin"] = "http://www.linkedin.com"
	fmt.Println(urlCollection)

	delete(urlCollection, "Lamda")
	fmt.Println(urlCollection)

	/*
		make function makes space for the give size
		otherwise for dynamic array or map
		it will be creating a new space whenever we append or add one
	*/
	// courseRating := make(map[string]float64, 3)
	courseRating := make(floatMap, 3)
	courseRating["1"] = 1.4
	courseRating["2"] = 1.2
	courseRating["3"] = 0.2
	courseRating.outputMap()
}
