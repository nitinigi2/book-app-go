package utilities

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVaribles() map[string]string {

	myEnv, err := godotenv.Read()

	if err != nil {
		log.Fatal(err)
	}

	return myEnv
}
