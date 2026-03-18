package initDB

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", "root:123456@tcp(192.168.123.99:34298)/ginhello")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
}
