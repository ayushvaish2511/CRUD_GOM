package models

import (
	"curd_gin/database"
	"fmt"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string         `json:"title" validate:"required"`
	Body    string         `json:"body" validate:"required"`
	TagList pq.StringArray `json:"tagList" gorm:"type:text[]"`
}

func Migrate() error {
	// AutoMigrate other columns
	if err := database.DB.AutoMigrate(&Article{}); err != nil {
		return fmt.Errorf("error during AutoMigrate: %w", err)
	}

	// Manually run SQL command to add the "text[]" type to the "tag_list" column
	if err := database.DB.Exec("ALTER TABLE articles ADD COLUMN IF NOT EXISTS tag_list text[];").Error; err != nil {
		return fmt.Errorf("error adding column: %w", err)
	}

	return nil
}

