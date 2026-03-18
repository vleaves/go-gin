package main

import (
	"10_jwt/router"
)

func main() {
	router := router.InitRouter()
	_ = router.Run()
}
