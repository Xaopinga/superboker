package serve

import (
	"fmt"
	"superboker2/model"
	"superboker2/query"
	"superboker2/uilt"

	"github.com/jinzhu/gorm"
)

type InSevText interface {
	Add(*model.Text) (bool, error)
	List(*query.Pagequery) ([]*model.Text, error)
	Edit(*model.Text) (bool, error)
	Del(int) (bool, error) //删除
	GetTotal() (total int, err error)
	Exist(*model.Text) *model.Text    //判断是否存在
	ExistByTextID(id int) *model.Text //根据id获取详细信息
}

type SevText struct {
	DB *gorm.DB
}

func (s *SevText) Add(m *model.Text) (bool, error) {
	db := s.DB
	if err := db.Create(m).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (s *SevText) List(page *query.Pagequery) ([]*model.Text, error) {
	list := make([]*model.Text, 0)
	li, off := uilt.Page(page.Size, page.Page)
	db := s.DB
	if err := db.Order("id desc").Limit(li).Offset(off).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (s *SevText) Edit(m *model.Text) (bool, error) {
	db := s.DB
	if err := db.Model(m).Where("id=?", m.TextId).Update(m).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (s *SevText) Del(id int) (bool, error) {
	db := s.DB
	if err := db.Delete(&model.Text{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (s *SevText) GetTotal() (total int, err error) {
	if err := s.DB.Model(&model.Text{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return
}
func (s *SevText) Exist(m *model.Text) *model.Text {
	if err := s.DB.Find(m, "id=?", m.ID).Error; err != nil {
		fmt.Println("查询text  err", err)
		return nil
	}
	return m
}
func (s *SevText) ExistByTextID(id int) *model.Text {
	db := s.DB
	m := new(model.Text)
	if err := db.Where("id=?", id).Find(m).Error; err != nil {
		fmt.Println("ID 查询text err", err)
		return nil
	}
	return m
}
