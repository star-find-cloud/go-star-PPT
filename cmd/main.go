package main

import (
	"fmt"
	"localhost/business_code_test/config"
	"localhost/business_code_test/pkg/database/mysql"
)

func main() {
	var c = config.GetConfig()
	fmt.Println(c)

	var db1 = mysql.GetMasterDB()
	var db2 = mysql.GetSlaveDB1()
	var db3 = mysql.GetSlaveDB2()
	fmt.Println(db1, db2, db3)
}
