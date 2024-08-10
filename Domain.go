package main

// use in func Connect_Sqlserver
var server string = "localhost"   //your server name
var port int = 1433               //your sqlserver port default 1433
var user string = "sa"            //your adminstrator name
var password string = "zxc123!@#" //password for this server
var database string = "Test"      //use database name
var tableName = "TestTable"       //use table name

// String for connection
var connectString string = "server=localhost;user id=sa;password=zxc123!@#;port=1433;database=Test"

//**** InsertData ****

// use in func Insert_Data
// // The number of data that needs to be inserted
var need_insert int = 3

// // The var u need insert
var insert_var1 = 5      //ID
var insert_var2 = "沐雨橙风" //名字
var insert_var3 = "枪炮师"  //职业
var insert_var4 = "1"    //是否有效

// use in func Insert_Data_ColumnName
// 新增有效否属性的列
var newColumnName = "Useful"
var newColumnType = "bit"

var teststring1 = 6
var teststring2 = "君莫笑"
var teststring3 = "散人"
var teststring4 = "1"
