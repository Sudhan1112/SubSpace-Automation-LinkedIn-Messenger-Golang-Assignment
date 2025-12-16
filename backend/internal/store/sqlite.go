package store

import (
	"log"

	"github.com/glebarez/sqlite"
	"github.com/sudhan/browser-automation/internal/models"
	"gorm.io/gorm"
)

type Store struct {
	DB *gorm.DB
}

func NewStore(dbParams string) *Store {
	db, err := gorm.Open(sqlite.Open(dbParams), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&models.ProfileActivity{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return &Store{DB: db}
}

// SaveActivity logs an action (Connection Request, Message, etc.)
func (s *Store) SaveActivity(activity models.ProfileActivity) error {
	result := s.DB.Create(&activity)
	return result.Error
}

// GetActivities returns all logged activities
func (s *Store) GetActivities() ([]models.ProfileActivity, error) {
	var activities []models.ProfileActivity
	result := s.DB.Find(&activities)
	return activities, result.Error
}
