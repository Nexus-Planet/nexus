# The backend for nexus planet consisting of multiple microservices

## Requirements

To build nexus planet backend services you will need the following:

- Go 1.24.6 or later

Recommended but optional:

- make
- A PostgreSQL database engine
- reverse proxy such as nginx

## Get Started

### Build with make

```sh
$ git clone https://github.com/nexus-planet/nexus-planet-api
$ cd nexus-planet-api
# or make build
$ make
# run each service in bin directory
# example running auth service
$ ./bin/auth
```

### Build with go

```sh
$ git clone https://github.com/nexus-planet/nexus-planet-api
$ cd nexus-planet-api
$ go build ./bin/ ./cmd/...
# run each service in bin directory
# example running auth service
$ ./bin/auth
```

**More coming soon...**
