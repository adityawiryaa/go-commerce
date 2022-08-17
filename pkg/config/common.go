package config

import (
	"errors"
	"fmt"
	"os"
)

func Required(key string) error {
	if os.Getenv(key) == "" {
		errorMsg := fmt.Sprintf("config %s is required", key)
		return errors.New(errorMsg)
	}
	return nil
}
