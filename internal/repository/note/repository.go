package note

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/plusik10/note-service-api/internal/model"
	"github.com/plusik10/note-service-api/internal/pkg/db"
	repo "github.com/plusik10/note-service-api/internal/repository"
	constantRepository "github.com/plusik10/note-service-api/internal/repository/note/const"
)

type repository struct {
	client db.Client
}

func NewNoteRepository(db db.Client) repo.NoteRepository {
	return &repository{
		client: db,
	}
}

func (r *repository) Get(ctx context.Context, id int64) (*model.Note, error) {
	query, arg, err := squirrel.
		Select(constantRepository.Id,
			constantRepository.Author,
			constantRepository.Title,
			constantRepository.Text,
			constantRepository.CreatedAt,
			constantRepository.UpdatedAt).
		From(constantRepository.NoteTable).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "GetNote",
		QueryRow: query,
	}

	var note model.Note
	err = r.client.DB().GetContext(ctx, &note, q, arg...)
	if err != nil {

		return nil, err
	}
	return &note, nil
}

func (r *repository) GetAll(ctx context.Context) ([]*model.Note, error) {
	query, arg, err := squirrel.Select(constantRepository.Id,
		constantRepository.Author,
		constantRepository.Title,
		constantRepository.Text,
		constantRepository.UpdatedAt,
		constantRepository.CreatedAt).
		From(constantRepository.NoteTable).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var notes []*model.Note

	q := db.Query{
		Name:     "GetAll",
		QueryRow: query,
	}

	err = r.client.DB().SelectContext(ctx, &notes, q, arg...)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *repository) Update(ctx context.Context, ID int64, info *model.UpdateNoteInfo) error {
	builder := squirrel.Update(constantRepository.NoteTable)
	if info.Author.Valid {
		builder = builder.Set(constantRepository.Author, info.Author.String)
	}
	if info.Title.Valid {
		builder = builder.Set(constantRepository.Title, info.Title.String)
	}
	if info.Text.Valid {
		builder = builder.Set(constantRepository.Text, info.Text.String)
	}
	builder = builder.Set(constantRepository.UpdatedAt, "now()")
	builder = builder.PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{constantRepository.Id: ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{Name: "Update", QueryRow: query}

	_, err = r.client.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	query, arg, err := squirrel.
		Delete(constantRepository.NoteTable).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "DeleteNote",
		QueryRow: query,
	}
	_, err = r.client.DB().ExecContext(ctx, q, arg...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Create(ctx context.Context, info *model.NoteInfo) (int64, error) {
	query, arg, err := squirrel.
		Insert(constantRepository.NoteTable).
		Columns(constantRepository.Author,
			constantRepository.Title,
			constantRepository.Text).
		PlaceholderFormat(squirrel.Dollar).
		Values(info.Author, info.Title, info.Text).Suffix("returning id").ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "Create",
		QueryRow: query,
	}
	row, err := r.client.DB().QueryContext(ctx, q, arg...)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer row.Close()
	var id int64
	if row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}
