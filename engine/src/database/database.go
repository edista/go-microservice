package database

import (
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	engine *gorm.DB
}

func NewDatabase(uri string) (db *Database, err error) {
	pdb, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{
		engine: pdb,
	}, nil
}

func (db *Database) WithTransaction(ctx context.Context) *gorm.DB {
	transaction := db.engine.WithContext(ctx)
	return transaction
}
