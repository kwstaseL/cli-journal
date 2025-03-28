package service

import (
	"github.com/kwstaseL/cli-journal/cmd/internal/model"
	"github.com/kwstaseL/cli-journal/cmd/internal/repository"
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
	return n.repo.CreateNote(note)
}
