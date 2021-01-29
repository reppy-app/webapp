package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL address to comunicate with api.
	APIURL = ""
	// Port that the app is running.
	Port = 0
	// HashKey is used to auth the cookie.
	HashKey []byte
	// BlockKey is used to crypt cookie data.
	BlockKey []byte
)

// Load init the environment variables.
func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
