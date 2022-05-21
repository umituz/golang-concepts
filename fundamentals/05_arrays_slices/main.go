package main

import "fmt"

func main() {

	var fruitArray [2]string
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"
	fmt.Println(fruitArray)
	fmt.Println(fruitArray[0])
	fmt.Println(fruitArray[1])

	carArr := [2]string{"Lamborghini", "BMW"}
	fmt.Println(carArr)
	fmt.Println(carArr[0])
	fmt.Println(carArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[0])
	fmt.Println(fruitSlice[1])
	fmt.Println(fruitSlice[2])

	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])

}
