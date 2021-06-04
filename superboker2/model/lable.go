package model

import (
	"github.com/jinzhu/gorm"
)

type Lable struct {
	gorm.Model
	LableId int    `json:"id" gorm:"-"`
	Name    string `json:"name" gorm:"column:name"`
}
