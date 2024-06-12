package seeders

import (
	"fastx-api/src/models"
	"gorm.io/gorm"
	"log"
)

type UserSeeder struct {
	BaseSeeder
}

func (s *UserSeeder) Seed(db *gorm.DB) error {
	users := []models.User{
		{Name: "John Doe", Email: "john@example.com", Password: "123123"},
		{Name: "Jane Doe", Email: "jane@example.com", Password: "123123"},
	}

	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	log.Println("User seeder ran successfully!")
	return nil
}

func (s *UserSeeder) TableName() string {
	return "users"
}
