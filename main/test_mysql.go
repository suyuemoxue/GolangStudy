package main

import (
	"database/sql"
	"fmt"
	"time"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:sxj045115@tcp(127.0.0.1:3306)/chess?charset=utf8mb4"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sql.Open("mysql", "root:sxj045115@/chess")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Println(db)
	err = initDB()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("连接成功")
	}
}
