package initDB

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(192.168.123.99:34298)/ginhello")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}
