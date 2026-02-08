package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type TransactionService struct {
	repo *repositories.TransactionRepository
}

func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) GetTodayReport() (*models.SalesReport, error) {
	return s.repo.GetTodayReport()
}

func (s *TransactionService) GetReportByDateRange(startDate, endDate string) (*models.SalesReport, error) {
	return s.repo.GetReportByDateRange(startDate, endDate)
}
