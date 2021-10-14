package main

import (
	"github.com/galamarv/uppercase-services/router"
)

func main() {
	e := router.New()
	e.Logger.Fatal(e.Start(":8081"))

}
