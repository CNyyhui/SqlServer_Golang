package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

// 定义基础连接结构
const ()

func add_databases() {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 创建表
	createTableSQL := `CREATE TABLE TestTable (
						ID INT PRIMARY KEY,
						Name NVARCHAR(50) NOT NULL,
						Job NVARCHAR(10)
					);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating table: ", err.Error())
	}
	fmt.Println("Table created!")
}
