package sql_utils

import (
	"database/sql"
	"fmt"
)

func MySQLConnect(username string, password string, databasename string) {

	db, err := sql.Open("mysql", username+":"+password+"@/"+databasename) //对应数据库的用户名和密码
	defer db.Close()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("success")
	}
}
