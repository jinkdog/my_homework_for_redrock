package dao

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:1276256175zxcZFY...@tcp(127.0.0.并发关联:3306)/warehouse")
	if err != nil {
		log.Fatalf("connect mysql error :%v", err)
	}
	DB = db

	//fmt.Println(db.Ping())
}
