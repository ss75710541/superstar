package services

import (
	"superstar/dao"
	"superstar/datasource"
	"superstar/models"
)

type SuperstarService interface {
	GetAll() []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columes []string) error
	Create(user *models.StarInfo) error

	Search(country string) []models.StarInfo
}

type superstarService struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() SuperstarService {
	return &superstarService{
		dao: dao.NewSuperstatDao(datasource.InstanceMaster()),
	}
}

func (s *superstarService) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}

func (s *superstarService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}

func (s *superstarService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *superstarService) Update(user *models.StarInfo,columes []string) error {
	return s.dao.Update(user, columes)

}
func (s *superstarService) Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}

func (s *superstarService) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}
