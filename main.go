package main

import (
	"fmt"

	"portfolio/config"
	"portfolio/router"
)

func main() {
	var AppConfig config.ConfigurationStruct

	config.InitalConfig(&AppConfig)
	router.InitRouter(&AppConfig)
	fmt.Println(AppConfig)

}
