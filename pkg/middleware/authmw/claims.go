package authmw

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
)

// tokenDuration - wrapper for limiting duration options
type tokenDuration time.Duration

const (
	// AccessDuration - validity duration for an access token
	AccessDuration = tokenDuration(time.Minute * 5)
	// RefreshDuration - validity duration for an access token
	RefreshDuration tokenDuration = tokenDuration(time.Hour * 24 * 7 * 3)

	// RefreshCookieName - name of a cookie that holds refresh token string
	RefreshCookieName = "knizhly-refresh-token"
	// AccessCookieName - name of a cookie that holds access token string
	AccessCookieName = "knizhly-access-token"
	// LoginPagePath
	LoginPagePath = "/auth/login"
)

// IssueToken - returns jwt-token string with a given duration, userId and signing key
func IssueToken(id uuid.UUID, d tokenDuration, skey string) (string, error) {
	iat := time.Now()
	eat := iat.Add(time.Duration(d))

	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iat": iat.Unix(),
			"exp": eat.Unix(),
			"uid": id,
		})

	return t.SignedString([]byte(skey))
}

// ParseToken - checks if given token is valid using given signing key
// and if it is, attempts to parse it and extract user ID from it
func ParseToken(ts, skey string) (uuid.UUID, bool) {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (any, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// keyfunc is required to return a signing key
		return []byte(skey), nil
	})

	// errors are irrelevant - if they exist it just means that
	// the token is invalid and we just early return
	if err != nil {
		return uuid.Nil, false
	}

	if !token.Valid {
		return uuid.Nil, false
	}

	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return uuid.Nil, false
	}

	// истек - не валидны
	if time.Until(exp.Time) < 0 {
		return uuid.Nil, false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, false
	}

	uidStr, ok := claims["uid"].(string)
	if !ok {
		return uuid.Nil, false
	}

	uid, _ := uuid.FromString(uidStr)

	return uid, true
}

// WriteTokenPair - writes a pair of access and refresh jwt-tokens to a context's cookie
func WriteTokenPair(c *fiber.Ctx, uid uuid.UUID, skey string) error {
	newAccessT, err := IssueToken(uid, AccessDuration, skey)
	if err != nil {
		return fmt.Errorf("issue access token: %w", err)
	}

	newRefreshT, err := IssueToken(uid, RefreshDuration, skey)
	if err != nil {
		return fmt.Errorf("issue refresh token: %w", err)
	}

	c.Cookie(&fiber.Cookie{
		Name:  AccessCookieName,
		Value: newAccessT,
		//HTTPOnly: true,
		Expires: time.Now().Local().Add(time.Duration(AccessDuration)),
		Path:    "/",
		Secure:  true,
	})

	c.Cookie(&fiber.Cookie{
		Name:  RefreshCookieName,
		Value: newRefreshT,
		//HTTPOnly: true,
		Expires: time.Now().Add(time.Duration(RefreshDuration)),
		Path:    "/",
		Secure:  true,
	})

	return nil
}
