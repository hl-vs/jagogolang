package services

import (
	"kasir-online/models"
	"kasir-online/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) Create(data *models.Category) (*int, error) {
	return s.repo.Create(data)
}

func (s *CategoryService) GetByID(id int) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(id int, category *models.Category) error {
	return s.repo.Update(id, category)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
