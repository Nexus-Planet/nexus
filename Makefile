SRC_DIR=./cmd/...
BUILD_DIR=./bin

all: build

build:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/linux/ $(SRC_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/windows/ $(SRC_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/macos/ $(SRC_DIR)
	@echo "build complete..."

clean:
	@rm -rf $(BUILD_DIR)
	@echo "cleared all build files..."
