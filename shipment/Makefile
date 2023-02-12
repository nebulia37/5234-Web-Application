GOCMD=go
GOTEST=$(GOCMD) test
BINARY_NAME=shipment-management

all: build

.PHONY: build

build:
	mkdir -p bin
	cp config.json bin/
	$(GOCMD) build -o bin/$(BINARY_NAME) .

clean: 
	rm -f ./bin/$(BINARY_NAME)
	rm -f ./bin/config.json
