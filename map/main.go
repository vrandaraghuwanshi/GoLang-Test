package main

import "fmt"

func main(){
	// var colors map[string]string
	colors := map[string]string{
		"red" : "#ff0000",
		"green" : "4bf745",
	}
	 
	printMap(colors)
	// colors[10] = "#fffff"

	// delete(colors, 10)

	// fmt.Println(colors)


}

func printMap(c map[string]string){
	for color, hex := range c{
	fmt.Println("Hex code for", color, "is", hex)
	}
}