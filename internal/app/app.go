package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/plusik10/note-service-api/internal/app/api/note_v1"
	"github.com/plusik10/note-service-api/internal/config"
	desc "github.com/plusik10/note-service-api/pkg/note_v1"
	"google.golang.org/grpc"
)

type App struct {
	noteImpl        *note_v1.Note
	serviceProvider *serviceProvider
	pathConfig      string
	grpcServer      *grpc.Server
	mux             *runtime.ServeMux
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{pathConfig: pathConfig}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	if err := a.runGRPC(wg); err != nil {
		return err
	}
	if err := a.runPublicHTTP(wg); err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func (a *App) runGRPC(wg *sync.WaitGroup) error {
	list, err := net.Listen("tcp", a.serviceProvider.GetConfig().GRPC.Port)
	if err != nil {
		return err
	}

	go func() {
		defer wg.Done()

		if err := a.grpcServer.Serve(list); err != nil {
			log.Fatalf("failing to process gRPC server: %v", err)
		}
	}()
	log.Printf("Run gRPC Server from %s port \n", a.serviceProvider.GetConfig().GRPC.Port)

	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	httpPort := a.serviceProvider.GetConfig().HTTP.Port
	go func() {
		defer wg.Done()
		if err := http.ListenAndServe(httpPort, a.mux); err != nil {
			log.Fatalf("failed to process muxer :%v", err)
		}
	}()

	log.Printf("Run public http handler on %s port", httpPort)

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// initServiceProvider initializes the service provider
func (a *App) initServiceProvider(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	a.serviceProvider = newServiceProvider(cfg)

	return nil
}

// initServer initializes the server
func (a *App) initServer(ctx context.Context) error {
	a.noteImpl = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterNoteV1Server(a.grpcServer, a.noteImpl)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	grpcPort := a.serviceProvider.GetConfig().GRPC.Port
	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, a.mux, grpcPort, opts)
	if err != nil {
		return err
	}

	return nil
}
