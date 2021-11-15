package main

import (
	"fmt"
	"math"
	"math/rand"
)

var c, python, java bool
var k, j int = 1, 2

func main() {

	fmt.Println("Hello Go!")
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("PI number is", math.Pi)
	fmt.Println(add(42, 13))
	fmt.Println(add2(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)
	var c2, python2, java2 = true, false, "no!"
	fmt.Println(k, j, c2, python2, java2)
}

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return x,y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
