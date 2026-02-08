package services

import (
	"kasir-online/models"
	"kasir-online/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(name string) ([]models.Product, error) {
	return s.repo.GetAll(name)
}

func (s *ProductService) Create(data *models.Product) (*int, error) {
	return s.repo.Create(data)
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(id int, product *models.Product) error {
	return s.repo.Update(id, product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}

// v1
func (s *ProductService) GetAll_V1() ([]models.Product, error) {
	return s.repo.GetAll_V1()
}

func (s *ProductService) Create_V1(data *models.Product) error {
	return s.repo.Create_V1(data)
}

func (s *ProductService) GetByID_V1(id int) (*models.Product, error) {
	return s.repo.GetByID_V1(id)
}

func (s *ProductService) Update_V1(id int, product *models.Product) error {
	return s.repo.Update_V1(id, product)
}

func (s *ProductService) Delete_V1(id int) error {
	return s.repo.Delete_V1(id)
}
