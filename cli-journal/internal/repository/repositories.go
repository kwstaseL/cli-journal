package repository

import (
	"github.com/kwstaseL/cli-journal/internal/model"
	"gorm.io/gorm"
)

type NoteRepository interface {
	CreateNote(note model.Note) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (n *noteRepository) CreateNote(note model.Note) error {
	panic("unimplemented")
}
