package configs

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port1    string `envconfig:"PORT1" required:"true"`
	Port2    string `envconfig:"PORT2" required:"true"`
	FilePath string `envconfig:"FILE_PATH" required:"true"`
}

func NewConfig() (*Config, error) {
	var cnf Config
	if err := envconfig.Process("", &cnf); err != nil {
		return nil, err
	}

	return &cnf, nil
}
