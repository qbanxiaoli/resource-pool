package dao

import (
	"container/list"
	"fmt"
	"github.com/jmoiron/sqlx"
	"resource-pool/model"
)

func GetResourceById(Db *sqlx.DB, id int) *model.ResourceTotal {
	rows, err := Db.Query("select * from resource_total where id = ?", id)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}
	resourceTotal := &model.ResourceTotal{}
	for rows.Next() {
		err := rows.Scan(&resourceTotal.Id, &resourceTotal.AverageCpu, &resourceTotal.AverageDisk, &resourceTotal.AverageMemory, &resourceTotal.MaxCpu,
			&resourceTotal.MaxMemory, &resourceTotal.Name, &resourceTotal.ReadAverage, &resourceTotal.WriteAverage, &resourceTotal.SpeedAverage)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	//关闭结果集（释放连接）
	rows.Close()
	return resourceTotal
}

func GetResourceName(Db *sqlx.DB, currentPage int, pageSize int) *list.List {
	rows, err := Db.Query("select name from resource_total order by id desc limit ? offset ?", pageSize, (currentPage-1)*pageSize)
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}
	nameList := list.New()
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Println(err.Error())
		}
		nameList.PushBack(name)
	}
	//关闭结果集（释放连接）
	rows.Close()
	return nameList
}
