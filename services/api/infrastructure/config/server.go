package config

import (
	"errors"
	"fmt"
	"os"
)

type ApiConfig struct {
	AppPort     string
	RpcPort     string
	ServicePort ServicePort
}

type ServicePort struct {
	User        string
	Transaction string
	Product     string
}

func InitApiConfig() *ApiConfig {
	return &ApiConfig{
		AppPort: os.Getenv("APP_PORT"),
		RpcPort: os.Getenv("RPC_PORT"),
		ServicePort: ServicePort{
			User:    os.Getenv("USER_PORT"),
			Product: os.Getenv("PRODUCT_PORT"),
		},
	}
}

func (c *ApiConfig) required(key string, value string) error {
	if value == "" {
		errorMsg := fmt.Sprintf("config %s is required", key)
		return errors.New(errorMsg)
	}
	return nil
}
