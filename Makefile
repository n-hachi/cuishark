GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=cuishark
MAIN=cmd/cuishark/main.go

.PHONY: build clean test
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
test:
	$(GOTEST) ./...
vtest:
	$(GOTEST) -v ./...
vet:
	$(GOVET) ./...
