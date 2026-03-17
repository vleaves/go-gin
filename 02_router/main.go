package main

import (
	"02_router/router"
)

func main() {
	router := router.SetupRouter()
	_ = router.Run()
}
