package repository

import (
	"context"

	"github.com/plusik10/note-service-api/internal/model"
)

type NoteRepository interface {
	Create(ctx context.Context, info *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
	GetAll(ctx context.Context) ([]*model.Note, error)
	Update(ctx context.Context, id int64, info *model.UpdateNoteInfo) error
	Delete(ctx context.Context, id int64) error
}
