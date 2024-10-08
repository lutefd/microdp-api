package server

import (
	"context"
	"microd-api/internal/config"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	cfg := &config.Config{
		DBPath: ":memory:",
		Port:   8080,
	}

	server, err := NewServer(cfg)
	if err != nil {
		t.Fatalf("NewServer() error = %v", err)
	}
	if server == nil {
		t.Fatal("NewServer() returned nil server")
	}
	if server.Addr != ":8080" {
		t.Errorf("NewServer() wrong address, got = %v, want %v", server.Addr, ":8080")
	}
}

func TestServer_Run(t *testing.T) {
	cfg := &config.Config{
		DBPath: ":memory:",
		Port:   8080,
	}

	t.Run("NormalShutdown", func(t *testing.T) {
		server, _ := NewServer(cfg)
		ctx, cancel := context.WithCancel(context.Background())
		errCh := make(chan error)

		go func() {
			errCh <- server.Run(ctx)
		}()

		time.Sleep(100 * time.Millisecond)

		cancel()

		select {
		case err := <-errCh:
			if err != nil {
				t.Errorf("Server.Run() error = %v", err)
			}
		case <-time.After(5 * time.Second):
			t.Error("Server.Run() didn't shut down in time")
		}
	})

	t.Run("ShutdownOnSIGTERM", func(t *testing.T) {
		server, _ := NewServer(cfg)
		ctx := context.Background()
		errCh := make(chan error)

		go func() {
			errCh <- server.Run(ctx)
		}()

		time.Sleep(100 * time.Millisecond)

		p, _ := os.FindProcess(os.Getpid())
		p.Signal(syscall.SIGTERM)

		select {
		case err := <-errCh:
			if err != nil {
				t.Errorf("Server.Run() error = %v", err)
			}
		case <-time.After(5 * time.Second):
			t.Error("Server.Run() didn't shut down in time")
		}
	})
}
func TestServer_GracefulShutdown(t *testing.T) {
	cfg := &config.Config{
		DBPath: ":memory:",
		Port:   8080,
	}

	server, _ := NewServer(cfg)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			t.Errorf("ListenAndServe() error = %v", err)
		}
	}()

	time.Sleep(100 * time.Millisecond)

	err := server.GracefulShutdown(context.Background())
	if err != nil {
		t.Errorf("Server.GracefulShutdown() error = %v", err)
	}
}

func TestServer_Close(t *testing.T) {
	cfg := &config.Config{
		DBPath: ":memory:",
		Port:   8080,
	}

	server, _ := NewServer(cfg)

	err := server.Close()
	if err != nil {
		t.Errorf("Server.Close() error = %v", err)
	}
}
