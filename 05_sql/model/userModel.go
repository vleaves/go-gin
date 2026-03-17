package model

import (
	"05_sql/initDB"
	"log"
)

type UserModel struct {
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}

func (user *UserModel) Save() int64 {
	result, e := initDB.Db.Exec("insert into ginhello.user (email, password) values (?,?)", user.Email, user.Password)
	if e != nil {
		log.Panicln("user insert error", e.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}
