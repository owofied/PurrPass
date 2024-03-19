package config_test

import (
	"errors"
	"os"
	"testing"

	"github.com/owofied/PurrPass/internal/config"
)

func TestInitializeConfig(t *testing.T) {
	tempDir := t.TempDir()

	// Use a temporary directory for testing
	// This variable is defined in config.go at line 40
	config.UserConfigDirFunc = func() (string, error) {
		return tempDir, nil
	}

	d, _ := config.UserConfigDirFunc()

	_ = os.Remove(d + "/PurrPass/config.json")
	_ = os.Remove(d + "/PurrPass")

	config.InitializeConfig()

	dir, err := os.Stat(d + "/PurrPass")
	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("Config directory not created at %s/PurrPass\n", d)
	}

	permissions := dir.Mode().Perm()
	if permissions != 0o755 {
		t.Errorf("Config directory permissions not set to 755\n")
	}

	f, err := os.Stat(d + "/PurrPass/config.json")
	if errors.Is(err, os.ErrNotExist) {
		t.Errorf("Config file not created at %s/PurrPass/config.json\n", d)
	}

	permissions = f.Mode().Perm()
	if permissions != 0o600 {
		t.Errorf("Config file permissions not set to 600\n")
	}
}
