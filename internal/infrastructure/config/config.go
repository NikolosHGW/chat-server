package config

import (
	"flag"
	"fmt"

	"github.com/joho/godotenv"
)

// Load - подгружает .env файл с помощью флага -c.
func Load() error {
	var envPath string
	flag.StringVar(&envPath, "c", ".env", "path to config file .env")

	err := godotenv.Load(envPath)
	if err != nil {
		return fmt.Errorf("не удалось загрузить файл .env, который находится по пути %s: %w", envPath, err)
	}

	return nil
}
