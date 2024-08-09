package main

import (
	"fmt"
)

var server string = "localhost"
var port int = 1433
var user string = "sa"
var password string = "zxc123!@#"
var database string = "Test"
var connString string = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s", server, user, password, port, database)
var need_insert int = 3
