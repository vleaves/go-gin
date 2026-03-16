package main

import (
	"05_router/router"
)

func main() {
	router := router.SetupRouter()
	_ = router.Run()
}
