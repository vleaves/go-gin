package main

import (
	"04_hello/router"
)

func main() {
	router := router.SetupRouter()
	_ = router.Run()
}
