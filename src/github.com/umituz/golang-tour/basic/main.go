package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

const Pi = 3.14

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

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var ii int
	var f float64
	var bb bool
	var s string
	fmt.Printf("%v %v %v %q\n", ii, f, bb, s)
	var x, y int = 3, 4
	var ff float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(ff)
	fmt.Println(x, y, z)
	v := 42 // change me!
	fmt.Printf("v is of type %T\n", v)
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
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

var c, python, java bool
var k, j int = 1, 2

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
