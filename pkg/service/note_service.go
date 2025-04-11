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
	logIfError(err, "Error creating note")
	logger.LogDebug("Note created successfully: %s", note.String())
	return nil
}


func (n *noteService) ListFrequentNotes(limit int) ([]model.Note, error) {
	frequentNotes, err := n.repo.ListFrequentNotes(limit)
	logIfError(err, "Error when trying to list frequent notes")
	logger.LogDebug("Notes acquired %s", frequentNotes)
	return frequentNotes, err
}

func (n *noteService) ListNotesBy(filters model.NoteFilters) ([]model.Note, error) {
	filteredNotes, err := n.repo.ListNotesBy(filters)
	logIfError(err, "Error when trying to filter notes")
	return filteredNotes, nil
}

func logIfError(err error, context string) error {
	if err != nil {
		logger.LogError("%s: %v", context, err)
	}
	return err
}
