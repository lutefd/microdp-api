package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("DefaultValues", func(t *testing.T) {
		os.Clearenv()
		config, err := Load()
		if err != nil {
			t.Fatalf("Load() error = %v", err)
		}
		if config.DBPath != "test.db" {
			t.Errorf("Expected default DBPath to be 'microd.db', got %s", config.DBPath)
		}
		if config.Port != 8080 {
			t.Errorf("Expected default Port to be 8080, got %d", config.Port)
		}
	})

	t.Run("CustomValues", func(t *testing.T) {
		os.Clearenv()
		os.Setenv("DB_PATH", "/custom/path.db")
		os.Setenv("PORT", "9090")
		config, err := Load()
		if err != nil {
			t.Fatalf("Load() error = %v", err)
		}
		if config.DBPath != "/custom/path.db" {
			t.Errorf("Expected DBPath to be '/custom/path.db', got %s", config.DBPath)
		}
		if config.Port != 9090 {
			t.Errorf("Expected Port to be 9090, got %d", config.Port)
		}
	})

	t.Run("InvalidPort", func(t *testing.T) {
		os.Clearenv()
		os.Setenv("PORT", "invalid")
		_, err := Load()
		if err == nil {
			t.Errorf("Expected error for invalid PORT, got nil")
		}
	})
}
