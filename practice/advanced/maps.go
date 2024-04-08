package main

import "fmt"

func main() {
	// var mapName map[keyType]valueType

	var colors, colors2 map[string]string

	colors = map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}

	fmt.Println(colors)

	colors2 = getMap()

	// check if the map is empty
	if len(colors2) > 0 {
		fmt.Println(colors2)
	} else {
		fmt.Println("Empty Map", colors2)
	}

	// create and initialize a map
	favouriteColors := make(map[string]string)

	fmt.Println(favouriteColors)

	// insert a value
	favouriteColors["white"] = "#ffffff"
	favouriteColors["red"] = "#ff0000"
	favouriteColors["green"] = "#00ff00"
	favouriteColors["blue"] = "#0000ff"

	fmt.Println(favouriteColors)

	// delete a color
	// delete(map, key)
	delete(favouriteColors, "white")

	fmt.Println(favouriteColors)

	// check if a key exists in the map
	value, exists := favouriteColors["red"]
	if exists {
		fmt.Println("Color:", value)
	} else {
		fmt.Println("Color does not exist")
	}

	// single line key check
	if color, exists := favouriteColors["pink"]; exists {
		fmt.Println("Color:", color)
	} else {
		fmt.Println("Color does not exist")
	}

	// Looping over a map
	fmt.Println("The colors:")
	for k, v := range favouriteColors {
		fmt.Println(k, ":", v)
	}

}

func getMap() map[string]string {
	return map[string]string{}
}
