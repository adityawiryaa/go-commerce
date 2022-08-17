package config

import (
	"errors"
	"fmt"
	"os"
)

type ProductConfig struct {
	RpcPort string
}

func InitProductConfig() *ProductConfig {
	return &ProductConfig{
		RpcPort: os.Getenv("RPC_PORT"),
	}
}

func (c *ProductConfig) required(key string, value string) error {
	if value == "" {
		errorMsg := fmt.Sprintf("config %s is required", key)
		return errors.New(errorMsg)
	}
	return nil
}
