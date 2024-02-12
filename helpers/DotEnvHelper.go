package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

type dotEnvHelper struct {

}

func NewDotEnvHelper() * dotEnvHelper {
	return &dotEnvHelper{}
}

func (dotEnvHelper) GoDotEnvVariable(key string) string {
  err := godotenv.Load(".env")

  if err != nil {
    panic(err)
  }

  return os.Getenv(key)
}