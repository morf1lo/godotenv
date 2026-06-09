package safe

import (
	"log"

	"github.com/morf1lo/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %s", err.Error())
	}
}
