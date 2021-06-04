package query

import (
	"superboker2/enum"
)

//请求
type Pagequery struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

//响应
type Entity struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Total     int         `json:"total"`
	TotalPage int         `json:"totalPage"`
	Data      interface{} `json:"data"`
}

func NewEntity() *Entity {
	return &Entity{
		Code:      int(enum.OperateFail),
		Msg:       enum.OperateFail.String(),
		Total:     0,
		TotalPage: 1,
		Data:      nil,
	}
}
func (e *Entity) OK(data interface{}) {
	e.Msg = enum.OperateOk.String()
	e.Code = int(enum.OperateOk)
	e.Data = data
}
