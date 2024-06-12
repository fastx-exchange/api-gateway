package seeders

import (
	"gorm.io/gorm"
	"log"
)

type BaseSeeder struct{}

func (s *BaseSeeder) Seed(db *gorm.DB) error {
	log.Printf("Seeder for %s not implemented\n", s.TableName())
	return nil
}

func (s *BaseSeeder) TableName() string {
	return ""
}
