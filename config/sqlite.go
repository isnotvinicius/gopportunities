package config

import (
	"os"

	"github.com/isnotvinicius/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/main.db"

	// Check if the database file exists
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating it...")

		// Create the database directory
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		// Create the database file
		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		// Close the file to avoid errors
		file.Close()
	}

	// Create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Migrate the schema created before
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Failed to migrate schema: %v", err)
		return nil, err
	}

	return db, nil
}
