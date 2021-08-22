package main

import "fmt"

func main() {
	colors_v1 := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}
	var colors_v2 map[string]string
	colors_v3 := make(map[string]string)

	fmt.Println(colors_v1)
	fmt.Println(colors_v2)
	fmt.Println(colors_v3)

	colors_v1["default"] = "-1"
	delete(colors_v1, "default")
	fmt.Println(colors_v1)

	printMap(colors_v1)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex of", color, "is", hex)
	}
}
