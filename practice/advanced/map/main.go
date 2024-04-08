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

	if colors2 != nil {
		fmt.Println(colors2)
	} else {
		fmt.Println("Empty Map", colors2)
	}
}

func getMap() map[string]string {
	return map[string]string{}
}
