package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	printMap(colours)

	colours1 := make(map[int]string)
	colours1[10] = "#ffffff"
	fmt.Println(colours1)

	// delete(colours, "red")
	// fmt.Println(colours)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Printf("%v color has hex of %v\n", k, v)
	}
}
