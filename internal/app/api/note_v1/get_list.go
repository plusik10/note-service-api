package note_v1

import (
	"context"

	"github.com/plusik10/note-service-api/internal/converter"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (n *Note) GetList(ctx context.Context, req *emptypb.Empty) (*desc.GetListResponse, error) {
	notes, err := n.noteService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return converter.ToDescGetListResponse(notes), nil
}
