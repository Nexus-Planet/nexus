package config

import (
	"os"
	"time"
)

var (
	ServerTimeout          = 60 * time.Second
	JwtTokenExpirationDate = 7 * 24 * time.Hour
	DatabaseUrl            = os.Getenv("DATABASE_URL")
	ServerPort             = 3000
	CustomDatabaseUrl      = ""
	CustomServerPort       = 0
)
