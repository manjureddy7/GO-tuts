package main

import "fmt"

func main() {

	// A Map is similar to what object in javascript
	// All the keys and values are static typed
	// All the keys must be of same type
	// Similarly all the values must be of same type
	// Lets create a Map here

	// A  color map saying
	// [string] -> keys of type string & string -> values of type string

	// colors := map[string]string{
	// 	"red":   "#344646d",
	// 	"green": "#6337383",
	// }

	// Other different ways to initialise map are
	// 1. using var

	// var colorsVar map

	// 2. using make

	colors := make(map[string]string)

	colors["red"] = "Im red"
	colors["green"] = "Im green"
	colors["white"] = "Im white"
	colors["blue"] = "Im blue"
	// colors[0] = "Im red"   -> This will give error as key must be string

	// How to delete keys in map
	// delete(colors, "blue")

	fmt.Println(colors)
	iterateOverMap(colors)

	// Struct vs Map, where to use whicj
}

// Lets make a function that will iterate a MAP
func iterateOverMap(mapArg map[string]string) {
	for key, value := range mapArg {
		fmt.Println("Key is", key+" "+"Value is", value)
	}
}
