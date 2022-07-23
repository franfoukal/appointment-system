package sqlconnector

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
)

func InitDBProductiveConnection() (*sql.DB, error) {
	connectionString := buildConnection()
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, fmt.Errorf("faliled to open connection with productive database")
	}

	// Check that the database is available and accessible
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping productive MySQL: %w", err)
	}

	logger.Infof("Successfully connected to productive database!")

	return db, nil
}

func InitDBLocalConnection() (*sql.DB, error) {
	connectionString := buildConnection()
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		return nil, fmt.Errorf("faliled to open connection with local database")
	}

	// Check that the database is available and accessible
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping local MySQL: %w", err)
	}

	logger.Infof("Successfully connected to local database!")

	return db, nil
}

func buildConnection() string {
	dbUsername := os.Getenv("DB_MYSQL_USERNAME")
	dbPassword := os.Getenv("DB_MYSQL_PASSWORD")
	dbHost := os.Getenv("DB_MYSQL_HOST")
	schemaName := os.Getenv("DB_MYSQL_SCHEMA_NAME")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUsername, dbPassword, dbHost, schemaName)
}
