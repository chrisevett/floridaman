package database

import (
	"fmt"

	i "github.com/chrisevett/floridaman/pkg/incident"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbClient() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=postgres dbname=floridaman port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open connection")
	}
	fmt.Println("connection established")

	fmt.Println("Migrating database")

	err = db.AutoMigrate(&i.Incident{}, &i.IncidentState{})
	if err != nil {
		fmt.Println("Failed to migrate ")
	}

	return db, err

}
