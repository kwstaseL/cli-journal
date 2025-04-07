package service

import (
	logger "github.com/kwstaseL/cli-journal/pkg/logging"
	"github.com/kwstaseL/cli-journal/pkg/model"
	"github.com/kwstaseL/cli-journal/pkg/repository"
)

type NoteService interface {
	CreateNewNote(note model.Note) error
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
	logger.LogInfo("Note created successfully: %s", note.String())
	return nil
}
