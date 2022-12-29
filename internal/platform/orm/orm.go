package orm

import (
	"database/sql"
	"log"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(db *sql.DB) {
	Instance, dbError = gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if dbError != nil {
		logger.Fatalf("%s", dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}
func DevelopmentMigrations() {
	// MIGRATIONS
	// Don`t delete migrations, modified on-demand to track changes and clean up in production ones
	Instance.AutoMigrate(&models.User{})
	Instance.AutoMigrate(&models.Service{})
	Instance.AutoMigrate(&models.Agenda{})

	logger.Infof("Database Migration Completed!")
}

func ProductionMigrations() {}
