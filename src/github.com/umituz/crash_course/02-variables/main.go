package main

import "fmt"

func main(){

	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	//Using variable
	//var name string = "Ümit Kenan"
	var name = "Ümit Kenan"

	fmt.Println(name)
	fmt.Printf("%T\n",name)

	//var age int32 = 25
	var age = 25

	fmt.Println(age)
	fmt.Printf("%T\n",age)

	const isCool = true
	fmt.Println(isCool)
	fmt.Printf("%T\n",isCool)

	var size float32 = 2.3
	fmt.Println(size)
	fmt.Printf("%T\n",size)

	// Shorthand
	fullName := "Ümit Kenan UZ"
	email := "umitkenanuz@gmail.com"
	fmt.Println(fullName,email)
	fmt.Printf("%T\n",fullName)
	fmt.Printf("%T\n",email)

	fullName2, email2 := "Ümit Kenan UZ", "umitkenanuz@gmail.com"
	fmt.Println(fullName2,email2)

}
