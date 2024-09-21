package repository

import (
	"github.com/google/uuid"
	"github.com/the-great-checkout/transaction-logs-crud/internal/database"
	"github.com/the-great-checkout/transaction-logs-crud/internal/entity"
	"gorm.io/gorm"
)

type TransactionLogRepository struct {
	db *gorm.DB
}

func NewTransactionLogRepository(postgresDB database.Postgres) *TransactionLogRepository {
	return &TransactionLogRepository{db: postgresDB.DB}
}

// Create saves a new transaction log in the database
func (r *TransactionLogRepository) Create(log *entity.TransactionLog) error {
	return r.db.Create(log).Error
}

// FindAllByTransactionID retrieves all logs associated with a specific transaction ID
func (r *TransactionLogRepository) FindAllByTransactionID(transactionID uuid.UUID) ([]entity.TransactionLog, error) {
	var logs []entity.TransactionLog
	if err := r.db.Where("transaction_id = ?", transactionID).Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
