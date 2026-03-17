package main

import (
	"06_file/router"
)

func main() {
	router := router.InitRouter()
	_ = router.Run()
}
