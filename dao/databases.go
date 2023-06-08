package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	InitMySQL()

}

func InitMySQL() {
	var err error
	dsn := "root:xl123321@tcp(127.0.0.1:3306)/picking?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
}

func Close() {
	if db != nil {
		db.Close()
	}
}
