package model

type Result struct {

	Result bool

	Code int

	Message string

	Data interface{}
}

func NewDefaultResult() Result {
	return Result{
		Result:  true,
		Code:    200,
		Message: "成功",
	}
}
