package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/nexus-planet/nexus-planet-api/internal/config"
)

var (
	JwtToken *jwtauth.JWTAuth
)

func init() {
	secret := os.Getenv("JWT_SECRET")

	JwtToken = jwtauth.New("HS256", []byte(secret), nil)
}

func MakeToken(data ...string) string {
	_, tokenString, err := JwtToken.Encode(map[string]interface{}{"email": data[0], "expirationDate": time.Now().Add(config.JwtTokenExpirationDate)})
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: FAILED TO MAKE TOKEN")
		return ""
	}

	return tokenString
}
