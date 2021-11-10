package main

import "fmt"

func main() {
	// Define map
	//names := make(map[string]string)

	//  Assign key value
	//names["umit"] = "umit@gmail.com"
	//names["Kenan"] = "kenan@gmail.com"
	//names["Umut"] = "umut@gmail.com"

	names := map[string]string{"umit":"umit@gmail.com","Kenan":"kenan@gmail.com","umut":"umut@gmail.com"}

	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(names["umut"])

	names["King"] = "king@gmail.com"

	delete(names,"umut")

	fmt.Println(names)

}
