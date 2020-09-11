package db

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"resource-pool/dao"
	"resource-pool/model"
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

func GetResourceTotal() model.Result {
	var Db = connectMysql()
	defer Db.Close()
	data := dao.GetResourceTotal(Db)
	result := model.NewDefaultResult()
	result.Data = data
	str, err := json.Marshal(*data)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(string(str))
	fmt.Println(result.Data)
	fmt.Println(result)
	return result
}
