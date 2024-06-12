package main

import (
	"fastx-api/config"
	"fastx-api/database/seeders"
	"fastx-api/src/pkg/database"
	"gorm.io/gorm"
	"log"
)

func runSeeders(db *gorm.DB, seeders []seeders.Seeder) {
	for _, seeder := range seeders {
		log.Printf("Seeding %s...\n", seeder.TableName())
		if err := seeder.Seed(db); err != nil {
			log.Fatalf("Failed to seed %s: %v", seeder.TableName(), err)
		}
	}
}

func main() {
	config.LoadConfig()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	seeders := []seeders.Seeder{
		&seeders.UserSeeder{},
	}

	runSeeders(db, seeders)
}
