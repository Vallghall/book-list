package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	mw "github.com/Vallghall/book-list/pkg/middleware"
	"github.com/Vallghall/book-list/pkg/middleware/authmw"
	"github.com/Vallghall/book-list/pkg/types"
	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	usernamePattern = regexp.MustCompile(`^\w+$`)
)

// CreateUser - handles user creation
func CreateUser(c *fiber.Ctx) error {
	uc := new(types.UserCreation)
	err := c.BodyParser(uc)
	if err != nil {
		return fiber.ErrBadRequest
	}

	if !validUserName(uc.UserName) {
		return fiber.ErrBadRequest // improve with htmx
	}

	repo := mw.Repo(c)
	u, _ := repo.FindUserByUsername(uc.UserName)
	if u != nil {
		return ErrUserAlreadyExists
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(uc.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("%w: hashing: %v,", fiber.ErrInternalServerError, err)
	}
	uc.Password = string(hash)

	user, err := repo.CreateUser(uc)
	if err != nil {
		return err // TODO: improve
	}

	return c.JSON(user)
}

// GetUser - handles getting user by id
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return ErrMissingIDParam
	}

	uid, err := uuid.FromString(id)
	if err != nil {
		return ErrBadUUID
	}

	repo := mw.Repo(c)

	user, err := repo.GetUser(uid)
	if err != nil {
		return err // TODO: improve
	}

	return c.JSON(user)
}

// Login - handles user authentication and authorization
func Login(skey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		repo := mw.Repo(c)

		user, err := repo.FindUserByUsername(username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(map[string]string{
					"msg": "login or password is incorrect",
				})
			}

			return fiber.ErrInternalServerError
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			return c.JSON(map[string]string{
				"msg": "login or password is incorrect",
			})
		}

		err = authmw.WriteTokenPair(c, user.ID, skey)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(map[string]string{
			"msg": "ok",
		})
	}
}

// validUserName - helper for username validation
func validUserName(name string) bool {
	return usernamePattern.MatchString(name)
}
