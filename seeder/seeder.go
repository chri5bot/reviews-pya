package seeder

import (
	"log"

	"github.com/chri5bot/reviews-pya/seeder/defaults"
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
	log.Println("Generating reviews...")
	// TODO: Generate and Insert Variation Options

	log.Println("Inserting reviews...")
	log.Println("Reviews... ", len(defaults.Reviews))

	for _, r := range defaults.Reviews {
		if err := s.DB.Create(&r).Error; err != nil {
			return err
		}
	}

	return nil
}
