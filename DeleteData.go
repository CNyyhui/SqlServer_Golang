package main

import (
	"database/sql"
	"fmt"
	"log"
)

func Delete_Data() {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connectString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 删除数据
	deleteSQL := `DELETE FROM TestTable WHERE ID = @ID;`
	_, err = db.Exec(deleteSQL, sql.Named("ID", 1))
	if err != nil {
		log.Fatal("Error deleting data: ", err.Error())
	}
	fmt.Println("Data deleted!")
}

func Delete_Table() {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connectString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 删除表
	dropTableSQL := `DROP TABLE TestTable000;`
	_, err = db.Exec(dropTableSQL)
	if err != nil {
		log.Fatal("Error dropping table: ", err.Error())
	}
	fmt.Println("Table dropped!")

}
