package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
)

//

func Insert_Data2() {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connectString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 插入数据                    修改下列属性，并且给予一个自增变量进行标识
	insertSQL := `INSERT INTO TestTable (ID, Name) VALUES (@ID, @Name);`
	_, err = db.Exec(insertSQL, sql.Named("ID", insert_var1), sql.Named("Name", insert_var2))
	if err != nil {
		log.Fatal("Error inserting data: ", err.Error())
	}
	fmt.Println("Data inserted!")
}

/*
此处可以新增插入数量对应等组件
对映不同的 need_insert 变量
可以简化对应的方法数量
*/
func Insert_Data3() {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connectString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 插入数据                    修改下列属性，并且给予一个自增变量进行标识
	/*
		小小尝试
		//need_insert_var1 := "ID"
		//need_insert_var2 := "Name"
		//need_insert_var3 := "Job"
		//insertSQLstatement := fmt.Sprint("`INSERT INTO ", tableName, " (", need_insert_var1, ",", need_insert_var2, ",", need_insert_var3, ") VALUES (@", need_insert_var1, ", @", need_insert_var2, ", @", need_insert_var3, ");`")
		//fmt.Println(insertSQLstatement)
	*/
	insertSQL := `INSERT INTO TestTable (ID, Name,Job) VALUES (@ID, @Name,@Job);`
	/*
		小小尝试
		//sqlNamed1 := fmt.Sprint("sql.Named(\"", need_insert_var1, "\",\"", insert_var1, "\")")
		//sqlNamed2 := fmt.Sprint("sql.Named(\"", need_insert_var2, "\",\"", insert_var2, "\")")
		//sqlNamed3 := fmt.Sprint("sql.Named(\"", need_insert_var3, "\",\"", insert_var3, "\")")
		//fmt.Println(sqlNamed1)
		//fmt.Println(sqlNamed2)
		//fmt.Println(sqlNamed3)
		//fmt.Println(insertSQL, sqlNamed1, sqlNamed2, sqlNamed3)
	*/
	_, err = db.Exec(insertSQL, sql.Named("ID", insert_var1), sql.Named("Name", insert_var2), sql.Named("Job", insert_var3))
	//_, err = db.Exec(insertSQL, sqlNamed1, sqlNamed2, sqlNamed3)
	if err != nil {
		log.Fatal("Error inserting data: ", err.Error())
	}
	fmt.Println("Data inserted!")
}

func Insert_Data4() {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connectString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 插入数据                    修改下列属性，并且给予一个自增变量进行标识
	insertSQL := `INSERT INTO TestTable (ID, Name,Job,Useful) VALUES (@ID, @Name,@Job,@Useful);`
	_, err = db.Exec(insertSQL, sql.Named("ID", insert_var1), sql.Named("Name", insert_var2), sql.Named("Job", insert_var3), sql.Named("Useful", insert_var4))
	if err != nil {
		log.Fatal("Error inserting data: ", err.Error())
	}
	fmt.Println("Data inserted!")
}

// 传参的插入值 ！！！ 需要更改
func Insert_Data_Company(string, string, string) {
	// 打开数据库连接
	db, err := sql.Open("sqlserver", connectString)
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

func Insert_New_Column() {

	// 打开数据库连接
	db, err := sql.Open("sqlserver", connectString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	// 插入新属性
	sqlStatement := fmt.Sprint("ALTER TABLE ", tableName, " ADD ", newColumnName, " ", newColumnType)
	fmt.Println(sqlStatement)

	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("Add new column failed ", err)
	}
	fmt.Println("Add seccessful")
}
