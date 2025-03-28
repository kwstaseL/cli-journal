package repository

import (
	"fmt"

	"github.com/kwstaseL/cli-journal/cmd/internal/model"
	"gorm.io/gorm"
)

type NoteRepository interface {
	CreateNote(note model.Note) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	// Auto-migrate the Note model to create the table if it doesn't exist
	err := db.AutoMigrate(&model.Note{})
	if err != nil {
		panic(fmt.Sprintf("Failed to auto-migrate database: %v", err))
	}
	return &noteRepository{db: db}
}

func (n *noteRepository) CreateNote(note model.Note) error {
	result := n.db.Create(note)
	return result.Error
}
