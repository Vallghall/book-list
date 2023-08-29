package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
)

// engine - creates template engine
func engine(path, tmplext string) fiber.Views {
	return django.New(path, tmplext)
}
