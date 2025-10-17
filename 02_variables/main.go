package main

import "fmt"

func main() {
	// Different ways to declare variables
	var name string = "Go"
	var version = 1.21
	age := 14
	isAwesome := true

	fmt.Println("Language:", name)
	fmt.Println("Version:", version)
	fmt.Println("Age:", age, "years")
	fmt.Println("Is awesome:", isAwesome)

	// Constants
	const pi = 3.14159
	fmt.Println("Pi:", pi)
}
