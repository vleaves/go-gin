package model

import (
	"09_orm/initDB"
	"log"
)

type Article struct {
	Id      int    `json:"id" example:"1"`
	Type    string `json:"type" example:"js"`
	Content string `json:"content" example:"hello js"`
}

func (article Article) TableName() string {
	return "article"
}

func (article Article) Insert() int {
	create := initDB.Db.Create(&article)
	if create.Error != nil {
		log.Panicln("文章添加失败", create.Error)
	}
	return 1
}

func (article Article) FindAll() []Article {
	var articles []Article
	result := initDB.Db.Find(&articles)
	if result.Error != nil {
		log.Panicln("查询数据失败")
	}

	return articles
}

func (article Article) FindById() Article {
	result := initDB.Db.Find(&article, article.Id)
	if result.Error != nil {
		log.Panicln("查询数据失败")
	}

	return article
}

func (article Article) DeleteOne() {
	result := initDB.Db.Delete(article)
	if result.Error != nil {
		log.Panicln("数据发生错误，无法删除")
	}
}
