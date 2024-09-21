package mapper

import (
	"github.com/the-great-checkout/transaction-logs-crud/internal/dto"
	"github.com/the-great-checkout/transaction-logs-crud/internal/entity"
)

type TransactionLogMapper struct{}

func NewTransactionLogMapper() *TransactionLogMapper {
	return &TransactionLogMapper{}
}

func (m *TransactionLogMapper) ToEntity(log dto.TransactionLog) entity.TransactionLog {
	return entity.TransactionLog{
		TransactionID: log.TransactionID,
		Timestamp:     log.Timestamp,
		Status:        log.Status,
		Value:         log.Value,
	}
}

func (m *TransactionLogMapper) ToDTO(log entity.TransactionLog) dto.TransactionLog {
	return dto.TransactionLog{
		TransactionID: log.TransactionID,
		Timestamp:     log.Timestamp,
		Status:        log.Status,
		Value:         log.Value,
	}
}
