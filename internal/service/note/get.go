package note

import (
	"context"

	"github.com/plusik10/note-service-api/internal/model"
)

func (s *noteService) Get(ctx context.Context, id int64) (*model.Note, error) {
	note, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return note, nil
}
