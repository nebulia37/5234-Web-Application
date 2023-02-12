GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=inventory-management

all: build

.PHONY: build

build:
	mkdir -p bin
	cp config.json bin/
	cp .env bin/
	cp -r migration bin/
	$(GOCMD) build -o bin/$(BINARY_NAME) .

clean: 
	rm -f ./bin/$(BINARY_NAME)
	rm -f ./bin/.env
	rm -f ./bin/config.json
	rm -rf ./bin/migration
	rm -rf ./client/build
