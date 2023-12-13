package note_v1

import (
	"github.com/plusik10/note-service-api/internal/service"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

var _ desc.NoteV1Server = (*Note)(nil)

type Note struct {
	desc.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewNote(service service.NoteService) *Note {
	return &Note{noteService: service}
}
