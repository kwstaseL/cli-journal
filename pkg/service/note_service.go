package service

import (
	logger "github.com/kwstaseL/cli-journal/pkg/logging"
	"github.com/kwstaseL/cli-journal/pkg/model"
	"github.com/kwstaseL/cli-journal/pkg/repository"
)

type NoteService interface {
	CreateNewNote(note model.Note) error
	ListFrequentNotes(limit int) ([]model.Note, error)
	ListNotesBy(filters model.NoteFilters) ([]model.Note, error)
}

type noteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteService{repo: repo}
}

func (n *noteService) CreateNewNote(note model.Note) error {
	err := n.repo.CreateNote(
		note,
	)
	if err != nil {
		logger.LogError("Error creating note: %v", err)
        return err
    }
	logger.LogDebug("Note created successfully: %s", note.String())
	return nil
}


func (n *noteService) ListFrequentNotes(limit int) ([]model.Note, error) {
	frequentNotes, err := n.repo.ListFrequentNotes(limit)
	if err != nil {
		logger.LogError("Error while trying to list frequent notes: %v", err)
        return nil, err
    }
	logger.LogDebug("Notes acquired %s", frequentNotes)
	return frequentNotes, err
}

func (n *noteService) ListNotesBy(filters model.NoteFilters) ([]model.Note, error) {
	filteredNotes, err := n.repo.ListNotesBy(filters)
	if err != nil {
		logger.LogError("Error while trying to filter notes: %v", err)
        return nil, err
    }
	return filteredNotes, nil
}