package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

// 定义连接字符串
func connect_sqlserver() {

	// 打开数据库连接
	db, err := sql.Open("sqlserver", connectString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 测试是否连接成功
	err = db.Ping()
	if err == nil {
		fmt.Println("Connect to SQLServer")
	} else {
		fmt.Println("Error connecting:", err)
	}

}
