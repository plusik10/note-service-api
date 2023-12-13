package note

import (
	"context"

	"github.com/plusik10/note-service-api/internal/model"
)

func (s *noteService) GetAll(ctx context.Context) ([]*model.Note, error) {
	notes, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return notes, nil
}
