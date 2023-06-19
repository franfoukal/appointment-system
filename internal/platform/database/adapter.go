package database

import (
	"database/sql"
	"log"

	"github.com/labscool/mb-appointment-system/db/models"
	"github.com/labscool/mb-appointment-system/internal/platform/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GORMInstance struct {
	*gorm.DB
}

func NewGormInstance(db *sql.DB) GORMInstance {
	instance, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		logger.Fatalf("%s", err)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database!")

	return GORMInstance{
		DB: instance,
	}
}

func (i GORMInstance) DevelopmentMigrations() {
	// MIGRATIONS
	// Don`t delete migrations, modified on-demand to track changes and clean up in production ones
	i.AutoMigrate(&models.User{})
	i.AutoMigrate(&models.Service{})
	i.AutoMigrate(&models.Agenda{})

	logger.Infof("Database Migration Completed!")
}

func (i GORMInstance) ProductionMigrations() {}
