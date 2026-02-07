package config

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/nexus-planet/nexus-planet-api/internal/db"
	flag "github.com/spf13/pflag"
)

var (
	ServerTimeout          = 60 * time.Second
	JwtTokenExpirationDate = 7 * 24 * time.Hour
	ServerPort             = 3000
	DefaultDatabase        = db.Postgres
	DataSourceName         = os.Getenv("DATA_SOURCE_NAME")
	CustomDataSourceName   = ""
	CustomServerPort       = 0
	CustomDatabase         = ""
)

type Config struct {
	ServerPort     int
	Database       string
	DataSourceName string
	ServerTimeout  time.Duration
}

// Loads the configurations
func Load() Config {
	LoadArgs()

	var database string
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

	if CustomDatabase == "" {
		database = DefaultDatabase
	} else {
		database = CustomDatabase
	}

	return Config{
		ServerPort:     port,
		Database:       database,
		DataSourceName: dsn,
		ServerTimeout:  ServerTimeout,
	}
}

// Loads the command line arguements
func LoadArgs() {
	flag.IntVarP(&CustomServerPort, "port", "p", 0, "Changes the default port for server")
	flag.StringVarP(&CustomDatabase, "database", "b", "", "Changes the default database")
	flag.StringVar(&CustomDataSourceName, "dsn", "", "Changes the default database data source name")
	flag.BoolP("default", "d", false, "Use default options from environment variables of system i.e:\nDATA_SOURCE_NAME=<dsn>\nJWT_SECRET=<secret>")
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Printf("%s", color.RedString("NOTE:You need to setup environment variables for JWT_SECRET regardless if you are using default or custom options"))
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
		return
	}

}
