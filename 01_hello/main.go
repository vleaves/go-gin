package main

import (
	"01_hello/router"
)

func main() {
	router := router.SetupRouter()
	_ = router.Run()
}
