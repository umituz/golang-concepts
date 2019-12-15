package main

import (
	"fmt"
	"math"
)

func main() {
	/**
	@todo I am Ümit UZ and awesome!
	 */
	fmt.Println("Hello World Via GoLang")
	/*
	@todo math kütüphanesinden Pi fonksyionu çağırma işlemi gerçekleştirilir ve ekrana yazar
	 */
	fmt.Println(math.Pi)

	/**
	@todo fmt.Println iki değer döndürür. Biri int diğeri error.
	* integer olan bite sayısıdır
	* error da <nil> olarak döner, yani nothing anlamındadır
	 */
	message, error := fmt.Println("Hi There")
	fmt.Println(message)
	fmt.Println(error)

	/**
	@todo Dönen değerleri değişkenlere aktarmak gerekir ama kullanılmazsa uyar verir.
	* golang de bir değişken kullanılmaz ise hata verir, yani değişken tanımlanmış ise kullanılmak zorundadır.
	* Bunun için kullanılmayacak olan değişkeni _ işareti ile belirtmemiz gerekir.
	*/
	x, _ := fmt.Println("Hello")
	fmt.Println(x)
}
