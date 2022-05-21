package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		upperCase()
		time.Sleep(time.Second)

		wg.Done()
	}()

	go func() {
		lowerCase()
		time.Sleep(time.Second)

		wg.Done()
	}()

	wg.Wait()
	fmt.Println("waiting to finish")
	fmt.Println("-----------------------------")
}

func lowerCase() {
	fmt.Println("lowercase")
}

func upperCase() {
	fmt.Println("UPPERCASE")
}
