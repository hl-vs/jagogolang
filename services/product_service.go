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

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) Create(data *models.Product) error {
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
