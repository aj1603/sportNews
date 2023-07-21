package main

import (
	"fmt"
	"sport_news/config"
	db "sport_news/database"
	app "sport_news/src"
	"sport_news/src/tools"
)

func main() {
	config.Init_config()
	db.Init_db()
	tools.Init_queries()
	router := app.Init_app()
	address := fmt.Sprintf("%v:%v", config.ENV.HOST, config.ENV.PORT)
	fmt.Println(address)
	router.Run(address)
}
