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
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err == nil {
		fmt.Println("Connect to SQLServer")
	} else {
		fmt.Println("Error connecting:", err)
	}

	/*
		// 删除数据
		deleteSQL := `DELETE FROM TestTable WHERE ID = @ID;`
		_, err = db.Exec(deleteSQL, sql.Named("ID", 1))
		if err != nil {
			log.Fatal("Error deleting data: ", err.Error())
		}
		fmt.Println("Data deleted!")

		// 删除表
		dropTableSQL := `DROP TABLE TestTable;`
		_, err = db.Exec(dropTableSQL)
		if err != nil {
			log.Fatal("Error dropping table: ", err.Error())
		}
		fmt.Println("Table dropped!")
	*/
}
