package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	userName  string = "root"
	password  string = "Java_W_M"
	ipAddress string = "119.45.219.104"
	port      int    = 3306
	dbName    string = "resource-pool"
	charset   string = "utf8"
)

func connectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}

func queryData(Db *sqlx.DB) {
	rows, err := Db.Query("select * from resource_pool")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for rows.Next() {
		//定义变量接收查询数据
		var uid int
		var create_time, username, password, department, email string

		err := rows.Scan(&uid, &create_time, &username, &password, &department, &email)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(uid, create_time, username, password, department, email)
	}

	//关闭结果集（释放连接）
	_ = rows.Close()
}

func main() {
	var Db = connectMysql()
	queryData(Db)
	defer Db.Close()
}
