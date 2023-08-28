package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrMissingIDParam    = fiber.NewError(http.StatusBadRequest, "id parameter is missing")
	ErrBadUUID           = fiber.NewError(http.StatusBadRequest, "bad uuid")
	ErrUserAlreadyExists = fiber.NewError(http.StatusConflict, "this username is already taken away")
)
