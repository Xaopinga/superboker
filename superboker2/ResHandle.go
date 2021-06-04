package main

import (
	"fmt"
	"superboker2/handle"
	"superboker2/model"
	"superboker2/serve"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db             *gorm.DB
	LableHandle    *handle.HandleLable
	CategoryHandle *handle.HandleCategory
	UpfileHandle   = &handle.HandleUpfile{}
	TextHandle     *handle.HandleText
)

func InitDb() {
	fmt.Println("数据库init链接")
	var err error
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		BaseUsername,
		BasePassword,
		BaseHost,
		BaseName,
		true,
		"Local")
	Db, err = gorm.Open("mysql", config)
	if err != nil {
		panic(fmt.Sprint("数据库链接err", err))
	}
	Db.AutoMigrate(
		&model.Lable{},
		&model.Category{},
		&model.Text{},
	)
	Db.LogMode(true)
	fmt.Println("数据库链接成功")

}
func InitHandle() {
	LableHandle = &handle.HandleLable{
		Sev: &serve.SevLable{
			DB: Db,
		},
	}
	CategoryHandle = &handle.HandleCategory{
		Sev: &serve.SevCategory{
			DB: Db,
		},
	}
	TextHandle = &handle.HandleText{
		Sev: &serve.SevText{
			DB: Db,
		},
	}
}
func init() {
	InitDb()
	InitHandle()
}
