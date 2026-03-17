package main

import (
	"04_form/router"
)

func main() {
	router := router.InitRouter()
	_ = router.Run()
}
