package note

import (
	"context"

	"github.com/plusik10/note-service-api/internal/model"
)

func (s *noteService) Create(ctx context.Context, info *model.NoteInfo) (int64, error) {
	id, err := s.repo.Create(ctx, info)
	if err != nil {
		return 0, err
	}

	return id, nil
}
