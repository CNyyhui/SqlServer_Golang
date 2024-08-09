package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

func Insert_Data() {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 插入数据
	insertSQL := `INSERT INTO TestTable (ID, Name,Job) VALUES (@ID, @Name,@Job);`
	_, err = db.Exec(insertSQL, sql.Named("ID", 3), sql.Named("Name", "爱吃饭"), sql.Named("Job", "厨师"))
	if err != nil {
		log.Fatal("Error inserting data: ", err.Error())
	}
	fmt.Println("Data inserted!")
}

func Insert_Data_Company(insert_var1 string, insert_var2 string, insert_var3 string) {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 插入数据                    修改下列属性，并且给予一个自增变量进行标识
	insertSQL := `INSERT INTO TestTable (ID, Name,Job) VALUES (@ID, @Name,@Job);`
	_, err = db.Exec(insertSQL, sql.Named("ID", insert_var1), sql.Named("Name", insert_var2), sql.Named("Job", insert_var3))
	if err != nil {
		log.Fatal("Error inserting data: ", err.Error())
	}
	fmt.Println("Data inserted!")
}
