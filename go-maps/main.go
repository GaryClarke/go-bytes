package main

import "fmt"

func main() {
	// Challenge
	//Write a program that:
	//
	//Creates a map of country codes to country names (e.g., "uk": "United Kingdom")
	countries := map[string]string{
		"uk": "United Kingdom",
		"fr": "France",
		"de": "Germany",
	}
	//Adds one new country
	countries["us"] = "United States"
	//Updates one existing country name
	countries["fr"] = "French Republic"
	//Deletes one country
	delete(countries, "de")
	//Prints out the final map using fmt.Println()
	fmt.Println(countries)
	//Try using map[string]string and call the variable countries.

}
