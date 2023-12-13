package note

import (
	"context"
)

func (s *noteService) Delete(ctx context.Context, id int64) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
