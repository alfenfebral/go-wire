package pkg_config

import (
	"go-clean-architecture/utils"

	"github.com/joho/godotenv"
)

func NewConfig() error {
	err := godotenv.Load()
	if err != nil {
		utils.CaptureError(err)
	}

	return nil
}
