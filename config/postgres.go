package config

import (
	"github.com/vpcmps/finanseed/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	dsn := "host=localhost user=postgres password=admin123 dbname=finanseed port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing postgres: %s", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Transaction{})

	if err != nil {
		logger.Errorf("Error migrating database: %s", err)
		return nil, err
	}
	return db, nil
}
