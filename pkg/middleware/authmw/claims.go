package authmw

import (
	"fmt"
	"time"

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
)

// IssueToken - returns jwt-token string with a given duration, userId and signing key
func IssueToken(id uuid.UUID, d tokenDuration, skey string) (string, error) {
	iat := time.Now()
	eat := iat.Add(time.Duration(d))

	t := jwt.NewWithClaims(
		jwt.SigningMethodES256,
		jwt.MapClaims{
			"iat": iat.Unix(),
			"exp": eat.Unix(),
			"uid": id,
		})

	return t.SignedString(skey)
}

// VerifyToken - checks if given token is valid using given signing key
func VerifyToken(ts, skey string) bool {
	token, err := jwt.Parse(ts, func(token *jwt.Token) (any, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// keyfunc is required to return a signing key
		return skey, nil
	})

	// errors are irrelevant - if they exist it just means that
	// the token is invalid and we just early return
	if err != nil {
		return false
	}

	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return false
	}

	// истек - не валидны
	if time.Until(exp.Time) < 0 {
		return false
	}

	return true
}
