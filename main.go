package main

func main() {
	/**
	@todo I am Ümit UZ and awesome! :D
	*/

	//fmt.Println("Hello World Via GoLang")

	/*
		@todo math kütüphanesinden Pi fonksyionu çağırma işlemi gerçekleştirilir ve ekrana yazar
	*/

	//fmt.Println(math.Pi)

	/**
	@todo fmt.Println iki değer döndürür. Biri int diğeri error.
	* integer olan bite sayısıdır
	* error da <nil> olarak döner, yani nothing anlamındadır
	*/

	//message, error := fmt.Println("Hi There")
	//fmt.Println(message)
	//fmt.Println(error)

	/**
	@todo Dönen değerleri değişkenlere aktarmak gerekir ama kullanılmazsa uyar verir.
	* golang de bir değişken kullanılmaz ise hata verir, yani değişken tanımlanmış ise kullanılmak zorundadır.
	* Bunun için kullanılmayacak olan değişkeni _ işareti ile belirtmemiz gerekir.
	*/

	//x, _ := fmt.Println("Hello")
	//fmt.Println(x)

	/**
	* Değişken tanımlamak için var anahtar kelimesi kullanılır.
	 */

	//var language = "Go Lang"
	//fmt.Println(language)
	//var name, surname = "Ümit", "UZ"
	//fmt.Println(name, surname)
	//var dream string
	//dream = "Change The World"
	//fmt.Println(dream)

	/**
	* Kısa değişken tanımlamalarını döngü içerisinde kullanılabilir.
	* var ile kullanılamaz
	 */
	//hey := 15
	//for i := 0; i <= hey; i++ {
	//	fmt.Println(i)
	//}

	/**
	* Statically type
	* Bir değişkenin türünü belirledikten sonra değiştiremezsin dinamik olarak da türünü bulmak gerekmez
	* int, int8, int16, int32,uint
	* int eksi olabilir ama uint eksi olamaz.
	 */
	//var speedOfLight int8 = 123 // (8 byte)
	//var speedOfLight int16 = 12345 // (16 byte)
	//var speedOfLight int32 = 1234567890 // (32 byte)
	//fmt.Println(speedOfLight)
	//fmt.Printf("%T", speedOfLight) // %T Türünü verir
	//fmt.Println()
	//fmt.Printf("%b", speedOfLight) // %b byte değerini verir

	//var red uint = 255
	//red = red + 1 // 256

	//var red uint8 = 255
	//red = red + 1 // 0

	//fmt.Println(red)
}
