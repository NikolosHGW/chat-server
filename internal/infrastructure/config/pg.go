package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type pg struct {
	DatabaseDSN string `env:"DATABASE_DSN"`
}

// NewPG - конструктор переменных для постгреса.
func NewPG() (*pg, error) {
	c := new(pg)
	err := env.Parse(c)
	if err != nil {
		return nil, fmt.Errorf("не удалось спарсить env для pg: %w", err)
	}

	return c, nil
}

func (c pg) GetDatabaseDSN() string {
	return c.DatabaseDSN
}

// PublicPG - алиас для структуры постгрес-конфига.
// Сделано в экспериментальных целях, чтоб пробросить
// сервис провайдер чёткую структуру без менющихся интерфейсов.
type PublicPG = pg
