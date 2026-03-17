package main

import (
	"05_sql/router"
)

func main() {
	router := router.InitRouter()
	_ = router.Run()
}
