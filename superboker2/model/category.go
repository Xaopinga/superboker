package model

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name       string `json:"name"`
	CategoryId int    `json:"id" gorm:"-"`
	ParentId   int    `json:"parentId"`
}

type CategoryResult struct {
	C1CategoryID string `gorm:"c1_category_id"`
	C1Name       string `gorm:"column:c1_name"`
	C1Desc       string `gorm:"column:c1_desc"`
	C1Order      int    `gorm:"column:c1_order"`
	C1ParentId   int    `gorm:"column:c1_parent_id"`

	C2CategoryID string `gorm:"c2_category_id"`
	C2Name       string `gorm:"column:c2_name"`
	C2Order      int    `gorm:"column:c2_order"`
	C2ParentId   int    `gorm:"column:c2_parent_id"`
}
