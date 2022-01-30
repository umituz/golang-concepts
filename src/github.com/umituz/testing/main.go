package email

import "fmt"

func Greeting(name string) (output string) {
	output = "Hello " + name
	return
}

func main() {
	greeting := Greeting("Kenan")
	fmt.Println(greeting)
}
