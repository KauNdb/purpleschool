package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email string
	Pass  string
	Addr  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error config")
	}

	return &Config{
		Email: os.Getenv("EMAIL"),
		Pass:  os.Getenv("PASS"),
		Addr:  os.Getenv("ADDR"),
	}
}
