package main

import (
	"fmt"
	"strings"
)

func main() {
	//Samples()

	Slice()

}

func Slice() {
	message := "Hello World"

	runes := []rune(message)
	subString1 := string(runes[0:4])
	fmt.Println("Rune Substring : ", subString1)

	subString2 := message[0:5]
	fmt.Println("ASCII Substring : ", subString2)

	subString3 := message[1:len(subString2)]
	fmt.Println("Sub string3 :",subString3)
}

func Samples() {
	fmt.Println(strings.ToTitle("Tüm harfleri büyük harfe çevirir"))

	fmt.Println(strings.Title("kelimelerin ilk harfini büyütür"))

	fmt.Println(strings.ToLower("Tüm Harfleri KÜÇÜK HARFE çevirir"))

	message := `Çoklu 
satırlarda string işlemler için
		normal tırnaklar yerine 
yamuk tırnak işareti kullanılır.`

	fmt.Println(message)
	//fmt.Printf("%s", message)
}
