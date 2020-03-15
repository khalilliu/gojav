package model

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	orm.Debug = true
	db, _  = sql.Open("mysql", "root:liukang0901@tcp(127.0.0.1:3306)/spiders?charset=utf8")
}

