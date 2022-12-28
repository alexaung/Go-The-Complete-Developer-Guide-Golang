package main

import "fmt"

func main() {
	//var colors map[string]string

	//colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#FF0000",
		"white": "#FFFFFF",
		"green": "#00FF00",
	}

	///delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
