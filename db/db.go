package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"resource-pool/dao"
	"resource-pool/model"
)

var (
	userName  = "root"
	password  = "Java_W_M"
	ipAddress = "119.45.219.104"
	port      = 3306
	dbName    = "resource-pool"
	charset   = "utf8"
)

func connectMysql() *sqlx.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]\n", err.Error())
	}
	return Db
}

func GetResourceById(id int) model.Result {
	var Db = connectMysql()
	defer Db.Close()
	data := dao.GetResourceById(Db, id)
	result := model.NewDefaultResult()
	result.Data = data
	return result
}

func GetResourceName(currentPage int, pageSize int) model.Result {
	var Db = connectMysql()
	defer Db.Close()
	data := dao.GetResourceName(Db, currentPage, pageSize)
	result := model.NewDefaultResult()
	var str []interface{}
	for i := data.Front(); i != nil; i = i.Next() {
		str = append(str, i.Value)
	}
	result.Data = str
	return result
}
