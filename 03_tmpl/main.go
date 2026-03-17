package main

import (
	"03_tmpl/router"
)

func main() {
	router := router.InitRouter()
	_ = router.Run()
}
