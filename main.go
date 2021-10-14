package main

import "uppercase/router"

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8081"))

}
