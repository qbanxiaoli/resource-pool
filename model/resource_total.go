package model

type ResourceTotal struct {
	// 自增主键
	Id int64
	// CPU使用平均值
	AverageCpu float64
	// 云盘使用率平均值
	AverageDisk float64
	// 内存使用率平均值
	AverageMemory float64
	// CPU使用率最大值
	MaxCpu float64
	// 内存使用率最大值
	MaxMemory float64
	// 名称
	Name string
	// 云硬盘读出平均值（KB/s）
	ReadAverage float64
	// 云硬盘写入平均值（KB/s）
	WriteAverage float64
	// 网络流出速率平均值（KB/s）
	SpeedAverage float64
}

