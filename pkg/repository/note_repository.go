package repository

import (
	"fmt"
	"strings"

	logger "github.com/kwstaseL/cli-journal/pkg/logging"
	"github.com/kwstaseL/cli-journal/pkg/model"
	"gorm.io/gorm"
)

type NoteRepository interface {
	CreateNote(note model.Note) error
	ListFrequentNotes(limit int) ([]model.Note, error)
	ListNotesBy(filters model.NoteFilters) ([]model.Note, error)
	GetNoteById(id string) (*model.Note, error)
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	err := db.AutoMigrate(&model.Note{})
	if err != nil {
		panic(fmt.Sprintf("Failed to auto-migrate database: %v", err))
	}
	logger.LogInfo("Database auto-migration successful")
	logger.LogInfo("Note repository initialized")

	return &noteRepository{db: db}
}

func (n *noteRepository) CreateNote(note model.Note) error {
	result := n.db.Create(&note)
	if result.Error != nil {
		return fmt.Errorf("failed to create note: %w", result.Error)
	}
	logger.LogInfo("Note created successfully", note)
	return nil
}

func (n *noteRepository) ListFrequentNotes(limit int) ([]model.Note, error) {
    var frequentNotes []model.Note
    err := n.db.Order("created_at DESC").
                Limit(limit).
                Find(&frequentNotes).Error
    return frequentNotes, err
}


func (n *noteRepository) ListNotesBy(filters model.NoteFilters) ([]model.Note, error) {
    var filteredNotes []model.Note
    query := n.db.Model(&model.Note{})

    if filters.Tags != "" {
        tags := strings.Split(filters.Tags, ",")
        for _, tag := range tags {
            query = query.Where("tags LIKE ?", "%"+strings.TrimSpace(tag)+"%")
        }
    }

    if filters.Category != "" {
        query = query.Where("category LIKE ?", "%"+strings.TrimSpace(filters.Category)+"%")
    }

    if filters.SearchTerm != "" {
        query = query.Where("header LIKE ? OR body LIKE ?", "%"+strings.TrimSpace(filters.SearchTerm)+"%", "%"+strings.TrimSpace(filters.SearchTerm)+"%")
    }

    query.Debug().Find(&filteredNotes)
    err := query.Find(&filteredNotes).Error
    return filteredNotes, err
}

func (n *noteRepository) GetNoteById(id string) (*model.Note, error) {
	var note model.Note
	result := n.db.First(&note, "id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("note with ID %s not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve note: %w", result.Error)
	}
	return &note, nil
}
