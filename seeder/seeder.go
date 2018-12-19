package seeder

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Seeder model
type Seeder struct {
	DB *gorm.DB
}

// NewSeeder creates a new instance of seeder
func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{
		DB: db,
	}
}

// Run Start database seeds.
func (s *Seeder) Run() error {
	log.Println("Generating brands...")

	return nil
}
