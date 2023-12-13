package note

import (
	"context"

	"github.com/plusik10/note-service-api/internal/model"
)

func (s *noteService) Update(ctx context.Context, id int64, info *model.UpdateNoteInfo) error {
	if err := s.repo.Update(ctx, id, info); err != nil {
		return err
	}

	return nil
}
