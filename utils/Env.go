package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EviromentVaribles struct {
	GithubSecret      string
	GithubCleintId    string
	PostgressHost     string
	PostgressUuser    string
	PostgressPassword string
	PostgressDbName   string
	PostgressPort     int
	PostgressSSLMODE  string
	ServerPort        string
}

func LoadEnv() (EviromentVaribles, error) {
	var EnvKeys EviromentVaribles

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("An error occured loading the .env: %s", err)
		return EnvKeys, err
	}
	EnvKeys.GithubCleintId = os.Getenv("GITHUB_CLIENT_ID")
	EnvKeys.GithubSecret = os.Getenv("GITUHB_SECRET")
	EnvKeys.PostgressHost = os.Getenv("POSTGRESS_HOST")
	EnvKeys.PostgressUuser = os.Getenv("POSTGRESS_USER")
	EnvKeys.PostgressPassword = os.Getenv("POSTGRESS_PASSWORD")
	EnvKeys.PostgressDbName = os.Getenv("POSTGRESS_DBNAME")
	EnvKeys.ServerPort = os.Getenv("SERVER_PORT")
	PostgressPort, _ := strconv.Atoi(os.Getenv("POSTGRESS_PORT"))
	EnvKeys.PostgressPort = (PostgressPort)
	EnvKeys.PostgressSSLMODE = os.Getenv("POSTGRESS_SSLMODE")
	fmt.Println("Finished env")
	return EnvKeys, nil
}
