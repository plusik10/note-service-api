package app

import (
	"context"
	"log"

	"github.com/plusik10/note-service-api/internal/config"
	"github.com/plusik10/note-service-api/internal/pkg/db"
	"github.com/plusik10/note-service-api/internal/repository"
	note2 "github.com/plusik10/note-service-api/internal/repository/note"
	"github.com/plusik10/note-service-api/internal/service"
	"github.com/plusik10/note-service-api/internal/service/note"
)

type serviceProvider struct {
	db     db.Client
	config *config.Config

	// repository
	noteRepository repository.NoteRepository

	//service
	noteService service.NoteService
}

func newServiceProvider(config *config.Config) *serviceProvider {
	return &serviceProvider{
		config: config,
	}
}

// GetDB returns the database
func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can't connect to db: %s", err.Error)
		}
		s.db = dbc
	}

	return s.db
}

// GetConfig
func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig()
		if err != nil {
			log.Fatalf("failed to create config: %s", err.Error())
		}
		s.config = cfg
	}
	return s.config
}

// GetNoteRepository returns a note repository
func (s *serviceProvider) GetNoteRepository(ctx context.Context) repository.NoteRepository {
	if s.noteRepository == nil {
		s.noteRepository = note2.NewNoteRepository(s.GetDB(ctx))
	}

	return s.noteRepository
}

// GetNoteService returns a service
func (s *serviceProvider) GetNoteService(ctx context.Context) service.NoteService {
	if s.noteService == nil {
		s.noteService = note.NewNoteService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}
