package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App    `yaml:"app"`
		HTTP   `yaml:"http"`
		Log    `yaml:"logger"`
		Trello `yaml:"trello"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// TrelloApi -.
	Trello struct {
		PublicKey   string `env-required:"true" yaml:"trello_developer_public_key"   env:"TRELLO_DEVELOPER_PUBLIC_KEY"`
		MemberToken string `env-required:"true" yaml:"trello_member_token"   env:"TRELLO_MEMBER_TOKEN"`
		Username    string `env-required:"true" yaml:"trello_username"   env:"TRELLO_USERNAME"`
		MemberBoard string `env-required:"true" yaml:"trello_board"   env:"TRELLO_board"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Trello.PublicKey == "" {
		cfg.Trello.PublicKey = os.Getenv("trello_developer_public_key")
	}

	if cfg.Trello.MemberToken == "" {
		cfg.Trello.MemberToken = os.Getenv("trello_member_token")
	}

	if cfg.Trello.Username == "" {
		cfg.Trello.Username = os.Getenv("trello_username")
	}

	if cfg.Trello.MemberBoard == "" {
		cfg.Trello.MemberBoard = os.Getenv("trello_board")
	}

	return cfg, nil
}
