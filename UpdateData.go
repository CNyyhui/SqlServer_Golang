package main

import (
	"database/sql"
	"fmt"
	"log"
)

func Update_Data() {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connectString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 更新数据
	updateSQL := `UPDATE TestTable SET Name = @Name WHERE ID = @ID;`
	_, err = db.Exec(updateSQL, sql.Named("ID", 1), sql.Named("Name", "炎雨辉"))
	if err != nil {
		log.Fatal("Error updating data: ", err.Error())
	}
	fmt.Println("Data updated!")
}
