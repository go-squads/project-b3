APP_NAME := reuni-server
BIN := bin

all:

dep:
	@echo "Downloading Dependencies"
	@dep ensure -v

test:
	@echo "Testing the application"
	@go test ./... -v

build: dep | test
	@echo "Building Binary to $(BIN)/$(APP_NAME)" 
	@go build -o $(BIN)/$(APP_NAME)
run:
	@$(BIN)/$(APP_NAME)
install:
	@go install . -o
