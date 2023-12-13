package note_v1

import (
	"context"

	"github.com/plusik10/note-service-api/internal/converter"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := n.noteService.Update(ctx, req.Id, converter.ToUpdateNoteInfo(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
