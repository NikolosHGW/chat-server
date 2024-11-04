package config

import (
	"fmt"
	"net"

	"github.com/caarlos0/env"
)

type grpc struct {
	GRPCHost string `env:"GRPC_HOST"`
	GRPCPort string `env:"GRPC_PORT"`
}

// NewGRPC - конструктор для переменных для gRPC.
func NewGRPC() (*grpc, error) {
	c := new(grpc)
	err := env.Parse(c)
	if err != nil {
		return nil, fmt.Errorf("не удалось спарсить env для grpc: %w", err)
	}

	return c, nil
}

func (c grpc) GetRunAddress() string {
	return net.JoinHostPort(c.GRPCHost, c.GRPCPort)
}

type PublicGRPC = grpc
