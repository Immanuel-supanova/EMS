package databases

import (
	"log"
	"os"

	"github.com/immanuel-supanova/EMS/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error

	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to db")
	} else {
		log.Println("Success connecting to db")
	}

}

func SyncDatabase() {
	DB.AutoMigrate(&models.Employee{})
	DB.AutoMigrate(&models.Job{})
	DB.AutoMigrate(&models.Contact{})
	DB.AutoMigrate(&models.Profile{})
	DB.AutoMigrate(&models.Enquiry{})
	DB.AutoMigrate(&models.Raise{})
	DB.AutoMigrate(&models.Leave{})
	DB.AutoMigrate(&models.Question{})
	DB.AutoMigrate(&models.Advance{})

}
