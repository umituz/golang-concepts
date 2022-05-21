package main

import (
	"fmt"
	"time"
)

func main() {

	//g := runtime.GOMAXPROCS(0)
	// runtime.GOMAXPROCS(1)

	fmt.Println("works")
	time.Sleep(time.Second)

	go func() {
		fmt.Println("go routine example 1")
		time.Sleep(time.Second)
	}()

	go exampleFunc()
	time.Sleep(time.Second)

}

func exampleFunc() {
	fmt.Println("example function")
	time.Sleep(time.Second)
}
