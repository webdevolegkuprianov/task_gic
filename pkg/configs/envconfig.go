package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port     string `envconfig:"PORT" required:"true"`
	FilePath string `envconfig:"FILE_PATH" required:"true"`
}

func NewConfig() (*Config, error) {
	var cnf Config
	if err := envconfig.Process("", &cnf); err != nil {
		return nil, err
	}

	return &cnf, nil
}
