package main

import "fmt"

func main() {

	//empty map
	var colors map[string]string

	fmt.Println(colors)

	// another way to declare map
	numbers := map[int]string{
		1: "one",
		2: "two",
	}

	//add an element to map
	numbers[3] = "three"

	//get an element from map
	fmt.Println(numbers[3])

	//delete an element from map
	delete(numbers, 3)

	//iterate over a map
	printMap(numbers)
}

func printMap(n map[int]string) {
	for key, value := range n {
		fmt.Println("Key is", key, " value is ", value)
	}
}
