package main

import "fmt"

func main() {
	// First way of creating a map
	colors := map[string]string{
		"red":    "#FF0000",
		"yellow": "FFFF00",
		"green":  "#00FF00",
	}

	// Add value to map
	colors["white"] = "#ffffff"

	// Change value
	colors["red"] = "New Red"

	// Delete key from map
	delete(colors, "red")

	fmt.Println("Colors:", colors)

	printMap(colors)

	// initialize a map without preset value
	var cars map[string]string
	fmt.Println("Cars: ", cars)

	// Or
	devices := make(map[string]string)
	fmt.Println("Devices:", devices)
}

func printMap(c map[string]string) {
	// Iterate over map
	for _, hex := range c {
		fmt.Println(hex)
	}
}
