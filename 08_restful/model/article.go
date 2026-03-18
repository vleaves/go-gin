package model

import (
	"08_restful/initDB"
	"log"
)

type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}

func (article Article) Insert() int {
	result, e := initDB.Db.Exec("insert into `article` (type, content) values (?, ?);", article.Type, article.Content)
	if e != nil {
		log.Panicln("文章添加失败", e.Error())
	}
	i, _ := result.LastInsertId()
	return int(i)
}
