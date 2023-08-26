package main

import (
	"log"

	"github.com/Vallghall/book-list/configs"

	"github.com/Vallghall/book-list/pkg/api"
)

func main() {
	c, err := configs.Bootstrap()
	if err != nil {
		log.Fatalln(err)
	}

	app := api.InitApp(c)
	log.Fatalln(app.Listen(c.AppConf.Port))
}
