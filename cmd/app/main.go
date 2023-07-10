package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Vallghall/book-list/pkg/api"
)

func main() {
	app := api.InitApp()
	port := os.Getenv("PORT")
	fmt.Println(port)

	log.Fatalln(app.Listen(port))
}
