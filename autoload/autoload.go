package autoload

import (
	"log"

	"github.com/morf1lo/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}
