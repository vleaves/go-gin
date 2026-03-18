package main

import (
	"07_middleware/router"
)

func main() {
	router := router.InitRouter()
	_ = router.Run()
}
