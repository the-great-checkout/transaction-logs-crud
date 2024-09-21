package service

import (
	"github.com/google/uuid"
	"github.com/the-great-checkout/transaction-logs-crud/internal/dto"
	"github.com/the-great-checkout/transaction-logs-crud/internal/entity"
)

type TransactionLogMapper interface {
	ToEntity(log dto.TransactionLog) entity.TransactionLog
	ToDTO(log entity.TransactionLog) dto.TransactionLog
}

type TransactionLogRepository interface {
	Create(log *entity.TransactionLog) error
	FindAllByTransactionID(transactionID uuid.UUID) ([]entity.TransactionLog, error)
}

type TransactionLogService struct {
	repository TransactionLogRepository
	mapper     TransactionLogMapper
}

func NewTransactionLogService(repository TransactionLogRepository, mapper TransactionLogMapper) *TransactionLogService {
	return &TransactionLogService{repository: repository, mapper: mapper}
}

// Create saves a new transaction log
func (s *TransactionLogService) Create(log dto.TransactionLog) error {
	entityLog := s.mapper.ToEntity(log)
	return s.repository.Create(&entityLog)
}

// FindAllByTransactionID retrieves all logs associated with a specific transaction ID
func (s *TransactionLogService) FindAllByTransactionID(transactionID uuid.UUID) ([]dto.TransactionLog, error) {
	logs, err := s.repository.FindAllByTransactionID(transactionID)
	if err != nil {
		return nil, err
	}

	var dtos []dto.TransactionLog
	for _, log := range logs {
		dtos = append(dtos, s.mapper.ToDTO(log))
	}

	return dtos, nil
}
