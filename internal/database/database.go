package database

import (
	"github.com/the-great-checkout/transaction-logs-crud/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres(dsn, schemaName string) Postgres {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   schemaName,
			SingularTable: false,
		}})
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")

	err = db.AutoMigrate(&entity.TransactionLog{})
	if err != nil {
		panic(err)
	}

	return Postgres{
		db,
	}
}
