package dao

import (
	"container/list"
	"fmt"
	"github.com/jmoiron/sqlx"
	"resource-pool/model"
)

func GetResourceTotal(Db *sqlx.DB) *list.List {
	rows, err := Db.Query("select * from resource_total limit 10")
	if err != nil {
		fmt.Printf(err.Error())
		return nil
	}
	resourceTotalList := list.New()
	for rows.Next() {
		resourceTotal := &model.ResourceTotal{}
		err := rows.Scan(&resourceTotal.Id, &resourceTotal.AverageCpu, &resourceTotal.AverageDisk, &resourceTotal.AverageMemory, &resourceTotal.MaxCpu,
			&resourceTotal.MaxMemory, &resourceTotal.Name, &resourceTotal.ReadAverage, &resourceTotal.WriteAverage, &resourceTotal.SpeedAverage)
		if err != nil {
			fmt.Println(err.Error())
		}
		resourceTotalList.PushBack(resourceTotal)
	}
	//关闭结果集（释放连接）
	rows.Close()
	return resourceTotalList
}
