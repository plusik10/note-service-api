package note_v1

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/plusik10/note-service-api/internal/converter"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
)

func (n *Note) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := n.noteService.Create(ctx, converter.ToNoteInfo(req.NoteInfo))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
