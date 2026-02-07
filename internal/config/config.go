package config

import (
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

var (
	ServerTimeout          = 60 * time.Second
	JwtTokenExpirationDate = 7 * 24 * time.Hour
	DataSourceName         = os.Getenv("DATABASE_URL")
	ServerPort             = 3000
	CustomDataSourceName   = ""
	CustomServerPort       = 0
)

type Config struct {
	ServerPort     int
	DataSourceName string
	ServerTimeout  time.Duration
}

// Loads the configurations
func Load() Config {
	LoadArgs()

	var dsn string
	var port int

	if CustomDataSourceName == "" {
		dsn = DataSourceName
	} else {
		dsn = CustomDataSourceName
	}

	if CustomServerPort == 0 {
		port = ServerPort
	} else {
		port = CustomServerPort
	}

	return Config{
		ServerPort:     port,
		DataSourceName: dsn,
		ServerTimeout:  ServerTimeout,
	}
}

// Loads the command line arguements
func LoadArgs() {
	flag.IntVarP(&CustomServerPort, "port", "p", 0, "Changes the default port for server")
	flag.StringVarP(&CustomDataSourceName, "database-url", "du", "", "Changes the default database url")
	flag.Bool("default", false, "Use default options from environment variables of system\ni.e:\nDATABASE_URL=<url>\nJWT_SECRET=<secret>")
	flag.Parse()
	if len(os.Args) < 2 {
		flag.Usage()
		return
	}

}
