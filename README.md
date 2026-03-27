# The backend for nexus planet consisting of multiple microservices

## Requirements

To build nexus planet backend services you will need the following:

- Go 1.24.6 or later

### Database (required)
- PostgreSQL (recommended for production)
- SQLite (partially supported; recommended for development/testing)
- MySQL (partially supported)

### Deployment (recommended)
- Reverse proxy such as nginx

### Development (optional)
- make (build automations)

## Get Started

### Build with make

```sh
$ git clone https://github.com/nexus-planet/nexus-planet-api
$ cd nexus-planet-api
# or make build
$ make
# run each service in bin directory
# example running auth service
# Note: Services may require environment variables (e.g., DB settings). Use --help to see options.
$ ./bin/auth
# view available commands and flags for each service
$ ./bin/auth --help
# or
$ ./bin/auth -h
```

### Build with go

```sh
$ git clone https://github.com/nexus-planet/nexus-planet-api
$ cd nexus-planet-api
$ go build ./bin/ ./cmd/...
# run each service in bin directory
# example running auth service
# Note: Services may require environment variables (e.g., DB settings). Use --help to see options.
$ ./bin/auth
# view available commands and flags for each service
$ ./bin/auth --help
# or
$ ./bin/auth -h
```

**More coming soon...**
