package settings

import (
	"github.com/joho/godotenv"
)

func Config() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return err
}
