package config

import (
	"os"

	env "github.com/joho/godotenv"
)

type config struct {
	HOST      string
	PORT      string
	GIN_MODE  string
	DB_URI    string
	QR_FOLDER string
	DIR       string
}

var ENV config

func Init_config() {
	env.Load()
	ENV.HOST = os.Getenv("HOST")
	ENV.PORT = os.Getenv("PORT")
	ENV.GIN_MODE = os.Getenv("GIN_MODE")
	ENV.DB_URI = os.Getenv("DB_URI")
	ENV.QR_FOLDER = os.Getenv("QR_FOLDER")
	ENV.DIR = os.Getenv("DIR")
}
