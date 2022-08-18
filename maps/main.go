package main

import "fmt"

func main() {
	// Alternative ways to init a map
	//var colors map[string]string
	//colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#f43fs4",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
