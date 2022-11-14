package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"portfolio/utils"
)

type ConfigurationStruct struct {
	GithubCleintId string
	ServerPort     string

	GithubSecret string
	DbConnection *sql.DB
}

func InitalConfig(c *ConfigurationStruct) error {
	EnvLoader, err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("An error occured loading enviroment: %s", err)
		return err
	}
	c.GithubCleintId = EnvLoader.GithubCleintId
	c.ServerPort = EnvLoader.ServerPort

	c.GithubSecret = EnvLoader.GithubSecret
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		EnvLoader.PostgressHost, EnvLoader.PostgressPort, EnvLoader.PostgressUuser, EnvLoader.PostgressPassword, EnvLoader.PostgressDbName)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	c.DbConnection = db
	return nil

}
