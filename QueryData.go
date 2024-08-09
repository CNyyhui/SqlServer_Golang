package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func Query_Data() {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 查询数据
	querySQL := `SELECT ID, Name, Job FROM TestTable;`
	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatal("Error querying data: ", err.Error())
	}
	defer rows.Close()

	fmt.Println("Query results:")
	for rows.Next() {
		var id int
		var name string
		var job string
		err := rows.Scan(&id, &name, &job)
		if err != nil {
			log.Fatal("Error scanning row: ", err.Error())
		}
		fmt.Printf("ID: %d, Name: %s, Job: %s\n", id, name, job)
	}

}
