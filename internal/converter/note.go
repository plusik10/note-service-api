package converter

import (
	"database/sql"

	"github.com/plusik10/note-service-api/internal/model"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToNoteInfo(info *desc.NoteInfo) *model.NoteInfo {
	return &model.NoteInfo{
		Text:   info.GetText(),
		Author: info.GetAuthor(),
		Title:  info.GetTitle(),
	}
}

func ToDescNoteInfo(info *model.NoteInfo) *desc.NoteInfo {
	return &desc.NoteInfo{
		Text:   info.Text,
		Author: info.Author,
		Title:  info.Title,
	}
}

func ToUpdateNoteInfo(info *desc.UpdateRequest) *model.UpdateNoteInfo {
	noteInfo := model.UpdateNoteInfo{}
	if info.UpdateRequestInfo.Author != nil {
		noteInfo.Author = sql.NullString{String: info.UpdateRequestInfo.Author.Value, Valid: true}
	}
	if info.UpdateRequestInfo.Title != nil {
		noteInfo.Title = sql.NullString{String: info.UpdateRequestInfo.Title.Value, Valid: true}
	}
	if info.UpdateRequestInfo.Text != nil {
		noteInfo.Text = sql.NullString{String: info.UpdateRequestInfo.Text.Value, Valid: true}
	}

	return &noteInfo
}

func ToDescGetListResponse(info []*model.Note) *desc.GetListResponse {
	descNotes := make([]*desc.Note, 0, len(info))
	for _, i := range info {
		descNotes = append(descNotes, ToDescNote(i))
	}

	return &desc.GetListResponse{
		Notes: descNotes,
	}
}

func ToDescNote(info *model.Note) *desc.Note {
	var updateAt *timestamppb.Timestamp
	if info.UpdatedAt != nil {
		if info.UpdatedAt.Valid {
			updateAt = timestamppb.New(info.UpdatedAt.Time)
		}
	}
	return &desc.Note{
		Id:        info.ID,
		NoteInfo:  ToDescNoteInfo(&info.NoteInfo),
		CreatedAt: timestamppb.New(info.CreatedAt),
		UpdatedAt: updateAt,
	}
}
