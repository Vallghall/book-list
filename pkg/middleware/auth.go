package middleware

import (
	"net/http"

	"github.com/Vallghall/book-list/pkg/middleware/authmw"
	"github.com/gofiber/fiber/v2"
)

const (
	redirectPath = "/"
)

// WithJWTAuth - middleware decorator that handles
// user authentication on handlers, validates
// JWT-token and adds authenticated userID into
// context's locals
func WithJWTAuth(skey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// retrieving access token
		token := c.Cookies(authmw.AccessCookieName)
		if token == "" {
			token = c.GetReqHeaders()["Authentication"]
		}

		uid, ok := authmw.ParseToken(token, skey)
		if ok {
			c.Locals("user-id", uid)
			return c.Next()
		}

		// retrieving refresh token
		token = c.Cookies(authmw.RefreshCookieName)
		if token == "" {
			return c.Redirect(redirectPath, http.StatusTemporaryRedirect)
		}

		uid, ok = authmw.ParseToken(token, skey)
		if !ok {
			return c.Redirect(redirectPath, http.StatusTemporaryRedirect)
		}

		if err := authmw.WriteTokenPair(c, uid, skey); err != nil {
			return fiber.ErrInternalServerError
		}

		return c.Next()
	}
}
