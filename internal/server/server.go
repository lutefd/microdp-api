package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"microd-api/internal/config"
	"microd-api/internal/controller"
	"microd-api/internal/repository"
	"microd-api/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Server struct {
	*http.Server
	db            *sql.DB
	apiRepository repository.APIRepository
	apiService    service.APIService
	apiController controller.APIController
}

func NewServer(cfg *config.Config) (*Server, error) {
	db, err := sql.Open("sqlite3", cfg.DBPath)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	apiRepo := repository.NewSQLiteAPIRepository(db)

	apiService := service.NewAPIService(apiRepo)

	apiController := controller.NewAPIController(apiService)

	s := &Server{
		Server: &http.Server{
			Addr:         fmt.Sprintf(":%d", cfg.Port),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		db:            db,
		apiRepository: apiRepo,
		apiService:    apiService,
		apiController: apiController,
	}

	s.Handler = s.RegisterRoutes()
	return s, nil
}

func (s *Server) Run(ctx context.Context) error {
	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server is listening on %s", s.Addr)
		serverErrors <- s.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		if err != http.ErrServerClosed {
			return fmt.Errorf("server error: %w", err)
		}
	case <-shutdown:
		log.Println("Starting shutdown...")
		return s.GracefulShutdown(ctx)
	case <-ctx.Done():
		log.Println("Context cancelled, starting shutdown...")
		return s.GracefulShutdown(context.Background())
	}

	return nil
}

func (s *Server) GracefulShutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := s.Shutdown(ctx); err != nil {
		return fmt.Errorf("server forced to shutdown: %w", err)
	}

	log.Println("Server exited properly")
	return nil
}

func (s *Server) Close() error {
	return s.db.Close()
}
