package main

import (
	"embed"
	"github.com/Vallghall/book-list/configs"
	"log"

	"github.com/Vallghall/book-list/pkg/api"
)

//go:embed script/migrations/*.sql
var migrations embed.FS

func main() {
	c, err := configs.Bootstrap(migrations)
	if err != nil {
		log.Fatalln(err)
	}

	app := api.InitApp(c)
	log.Fatalln(app.Listen(c.AppConf.Port))
}
