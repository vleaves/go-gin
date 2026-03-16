package main

import (
	"06_tmpl/router"
)

func main() {
	router := router.InitRouter()
	_ = router.Run()
}
