package model

import (
	"github.com/jinzhu/gorm"
)

/*
title
description
dimImg
textContent
lable//可以有多个
category
*/
type Text struct {
	gorm.Model
	Title       string `json:"title" xorm:"varchar(1120)"` //标题
	TextId      int    `json:"id" gorm:"-"`
	Description string `json:"description" gorm:"type:varchar(750)"`  //文章描述
	Dimimg      string `xorm:"type:varchar(30)" json:"dimImg"`        //模糊图地址
	Textcontent string `xorm:"type:varchar(8000)" json:"textContent"` //文章内容
	Lable       string ` json:"lable"`                                //标签
	Category    string ` json:"category"`                             //分类
	Clicknub    int    `xorm:"int(20)" json:"clicknub"`               //点击次数
	Weight      int    `xorm:"int(2)" json:"weight"`                  //权重
	Like        int    `json:"like"`                                  //点赞
	Oppose      int    `json:"oppose"`                                //反感
}
