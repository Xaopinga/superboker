package serve

import (
	"superboker2/model"
	"superboker2/query"
	"superboker2/uilt"

	"github.com/jinzhu/gorm"
)

type InSevCategory interface {
	List(*query.Pagequery) ([]*model.Category, error)
	GetTotal() (total int, err error) //获取一共的记录数
	Add(*model.Category) (bool, error)
	Exist(*model.Category) *model.Category    //判断是否存在
	ExistByCategoryID(id int) *model.Category //判断是否存在
	Delete(id int) (bool, error)
	Edit(*model.Category) (bool, error) //修改
	ListSonTotal() (total int, err error)
	ListSon(*query.Pagequery) ([]*model.Category, error)
}
type SevCategory struct {
	DB *gorm.DB
}

func (s *SevCategory) ListSon(page *query.Pagequery) ([]*model.Category, error) {
	list := make([]*model.Category, 0)
	li, off := uilt.Page(page.Size, page.Page)
	db := s.DB
	if err := db.Where("parent_id > ?", 0).Order("id desc").Limit(li).Offset(off).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (s *SevCategory) ListSonTotal() (total int, err error) {
	db := s.DB
	var ms []model.Category
	if err = db.Where("parent_id > ?", 0).Find(&ms).Count(&total).Error; err != nil {
		return 0, err
	}
	return
}

func (s *SevCategory) List(page *query.Pagequery) ([]*model.Category, error) {
	list := make([]*model.Category, 0)
	li, off := uilt.Page(page.Size, page.Page)
	db := s.DB
	if err := db.Order("id desc").Limit(li).Offset(off).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (s *SevCategory) GetTotal() (total int, err error) { //获取一共的记录数
	db := s.DB
	var ms []model.Category
	if err = db.Find(&ms).Count(&total).Error; err != nil {
		return 0, err
	}
	return
}

func (s *SevCategory) Add(m *model.Category) (bool, error) {
	db := s.DB
	err := db.Create(m).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *SevCategory) Exist(m *model.Category) *model.Category { //判断是否存在
	temp := new(model.Category)
	db := s.DB
	err := db.Where("name=?", m.Name).Find(temp).Error
	if err != nil {
		return nil
	}
	return temp
}
func (s *SevCategory) ExistByCategoryID(id int) *model.Category {
	db := s.DB
	m := new(model.Category)
	err := db.Where("id=?", id).Find(m).Error
	if err != nil {
		return nil
	}
	return m
}

func (s *SevCategory) Delete(id int) (bool, error) {
	db := s.DB
	if err := db.Delete(&model.Category{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (s *SevCategory) Edit(m *model.Category) (bool, error) {
	db := s.DB
	if err := db.Model(m).Where("id=?", m.CategoryId).Update(m).Error; err != nil {
		return false, err
	}
	return true, nil
}
