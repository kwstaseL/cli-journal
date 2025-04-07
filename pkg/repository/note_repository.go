package repository

import (
	"fmt"

	logger "github.com/kwstaseL/cli-journal/pkg/logging"
	"github.com/kwstaseL/cli-journal/pkg/model"
	"gorm.io/gorm"
)

type NoteRepository interface {
	CreateNote(note model.Note) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	err := db.AutoMigrate(&model.Note{})
	if err != nil {
		logger.LogError("Failed to auto-migrate database", err)
		panic(fmt.Sprintf("Failed to auto-migrate database: %v", err))
	}
	logger.LogInfo("Database auto-migration successful")
	logger.LogInfo("Note repository initialized")

	return &noteRepository{db: db}
}

func (n *noteRepository) CreateNote(note model.Note) error {
	result := n.db.Create(&note)
	if result.Error != nil {
		logger.LogError("Failed to create note", result.Error)
		return fmt.Errorf("failed to create note: %w", result.Error)
	}
	logger.LogInfo("Note created successfully", note)
	return nil
}
