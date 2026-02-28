package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/nexus-planet/nexus/internal/api"
	"github.com/nexus-planet/nexus/internal/config"
)

var (
	JwtToken *jwtauth.JWTAuth
)

func init() {
	secret := os.Getenv("JWT_SECRET")

	JwtToken = jwtauth.New("HS256", []byte(secret), nil)
}

func MakeToken(data ...string) string {
	_, tokenString, err := JwtToken.Encode(api.M{"email": data[0], "expirationDate": time.Now().Add(config.JwtTokenExpirationDate)})
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: FAILED TO MAKE TOKEN")
		return ""
	}

	return tokenString
}
