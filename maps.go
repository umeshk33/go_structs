package main

//Map in Go is a collection of key,value pairs -> diction in python
import "fmt"

func mapFunc() {
	colors := map[string]string{
		"red":    "#ff000",
		"yellow": "#fe000",
		"white":  "#ababab",
		"blue":   "#232322",
	}
	fmt.Println(colors)

	//another way of init
	var colors1 map[string]string
	fmt.Println(colors1) //will print empty map

	//another way of initializing map
	colors2 := make(map[string]string)
	colors2["green"] = "#ffffff"
	fmt.Println(colors2)

	//delete an entry in map
	delete(colors, "red")
	fmt.Println(colors)

	printMap(colors)
}

// iterate over maps
func printMap(c map[string]string) {

	for c, hex := range c {
		fmt.Println("Hex code for the color", c, "is", hex)
	}
}
