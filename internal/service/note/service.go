package note

import (
	"github.com/plusik10/note-service-api/internal/repository"
	"github.com/plusik10/note-service-api/internal/service"
)

var _ service.NoteService = (*noteService)(nil)

type noteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) service.NoteService {
	return &noteService{repo: repo}
}
