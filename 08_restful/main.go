package main

import (
	_ "08_restful/docs"
	"08_restful/router"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name youngxhu
// @contact.url https://youngxhui.top
// @contact.email youngxhui@g mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080

func main() {
	router := router.InitRouter()
	_ = router.Run()
}
