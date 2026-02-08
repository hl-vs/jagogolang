package services

import (
	"kasir-online/models"
	"kasir-online/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem, useLock bool) (*models.Transaction, error) {
	return s.repo.CreateTransactionV2(items)
}
