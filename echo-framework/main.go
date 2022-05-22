package main

import (
	"golangconcepts/echo-framework/api/routers"
)

func main() {

	e := routers.New()

	e.Start(":8000")

}
