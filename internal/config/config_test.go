//nolint:goconst // Nuh uh
package config_test

import (
	"encoding/json"
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

func TestGetConfig(t *testing.T) {
	tempDir := t.TempDir()

	// Use a temporary directory for testing
	// This variable is defined in config.go at line 40
	config.UserConfigDirFunc = func() (string, error) {
		return tempDir, nil
	}

	config.InitializeConfig()

	cfg := config.GetConfig()
	file, err := os.ReadFile(tempDir + "/PurrPass/config.json")
	if err != nil {
		t.Errorf("Failed to read config file at %s/PurrPass/config.json\n", tempDir)
	}

	var uMarshal *config.Config
	err = json.Unmarshal(file, &uMarshal)
	if err != nil {
		t.Error("Failed to unmarshal config")
	}

	// I'm sure there is a better way to do this, but I'm too lazy lol
	if cfg.Server.Port != uMarshal.Server.Port {
		t.Errorf("Returned port %s doesn't match port in config file %s\n", cfg.Server.Port, uMarshal.Server.Port)
	}
	if cfg.Database.Host != uMarshal.Database.Host {
		t.Errorf("Returned database host %s doesn't match database host in config file %s\n", cfg.Database.Host, uMarshal.Database.Host)
	}
	if cfg.Database.Port != uMarshal.Database.Port {
		t.Errorf("Returned database port %s doesn't match database port in config file %s\n", cfg.Database.Port, uMarshal.Database.Port)
	}
	if cfg.Database.User != uMarshal.Database.User {
		t.Errorf("Returned database user %s doesn't match database user in config file %s\n", cfg.Database.User, uMarshal.Database.User)
	}
	if cfg.Database.Password != uMarshal.Database.Password {
		t.Errorf("Returned database password %s doesn't match database password in config file %s\n", cfg.Database.Password, uMarshal.Database.Password)
	}
	if cfg.Database.DBName != uMarshal.Database.DBName {
		t.Errorf("Returned database name %s doesn't match database name in config file %s\n", cfg.Database.DBName, uMarshal.Database.DBName)
	}
}
