package serve

import (
	"superboker2/model"
	"superboker2/query"
	"superboker2/uilt"

	"github.com/jinzhu/gorm"
)

type InSevLable interface {
	List(*query.Pagequery) ([]*model.Lable, error)
	GetTotal() (total int, err error) //获取一共的记录数
	Add(*model.Lable) (bool, error)
	Exist(*model.Lable) *model.Lable       //判断是否存在
	ExistByCategoryID(id int) *model.Lable //判断是否存在
	Delete(id int) (bool, error)
	Edit(*model.Lable) (bool, error) //修改
}
type SevLable struct {
	DB *gorm.DB
}

func (s *SevLable) List(page *query.Pagequery) ([]*model.Lable, error) {
	list := make([]*model.Lable, 0)
	li, off := uilt.Page(page.Size, page.Page)
	db := s.DB
	if err := db.Order("id desc").Limit(li).Offset(off).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (s *SevLable) GetTotal() (total int, err error) {
	db := s.DB
	var ms []model.Lable
	if err = db.Find(&ms).Count(&total).Error; err != nil {
		return 0, err
	}
	return
}
func (s *SevLable) Add(m *model.Lable) (bool, error) {
	db := s.DB
	err := db.Create(m).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func (s *SevLable) Exist(m *model.Lable) *model.Lable {
	temp := new(model.Lable)
	db := s.DB
	err := db.Where("name=?", m.Name).Find(temp).Error
	if err != nil {
		return nil
	}
	return temp
}
func (s *SevLable) ExistByCategoryID(id int) *model.Lable {
	db := s.DB
	m := new(model.Lable)
	err := db.Where("id=?", id).Find(m).Error
	if err != nil {
		return nil
	}
	return m
}
func (s *SevLable) Delete(id int) (bool, error) {
	db := s.DB
	if err := db.Delete(&model.Lable{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (s *SevLable) Edit(m *model.Lable) (bool, error) {
	db := s.DB
	if err := db.Model(m).Where("id=?", m.LableId).Update(m).Error; err != nil {
		return false, err
	}
	return true, nil
}
