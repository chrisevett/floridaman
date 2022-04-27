// todo: this probably goes in a utils or something

package database

import (
	"fmt"

	i "github.com/chrisevett/floridaman/pkg/incident"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// this is going to init the db client and also migrated the database
func DbClient() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=floridaman port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to open connection")
	}
	fmt.Println("connection established")

	fmt.Println("Migrating database")

	err = db.AutoMigrate(&i.AffectedSystem{}, &i.Severity{}, &i.Incident{}, &i.IncidentEvent{},
		&i.IncidentEventType{}, &i.IncidentState{}, &i.ChecklistItem{},
		&i.ChecklistTemplate{}, &i.Checklist{})

	if err != nil {
		fmt.Println("Failed to migrate ")
	}

	if err != nil {
		fmt.Println("Failed to establish database connection: ", err)
		panic("Fatal")
	}

	// add static data

	// fuck it lets truncate for now

	// todo : only do this if our entrypoint is cli
	// todo: find a less tarded way to pass around this database client
	return db, err

}
