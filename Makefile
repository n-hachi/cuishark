GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
BINARY_NAME=cuishark
DEBUG_BINARY_NAME=dbg_cuishark
MAIN=cmd/cuishark/main.go

.PHONY: build clean test
build:
	$(GOBUILD) -o $(BINARY_NAME) -v $(MAIN)
debug:
	$(GOBUILD) -tags debug -o $(DEBUG_BINARY_NAME) -v $(MAIN)
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
test:
	$(GOTEST) ./...
vtest:
	$(GOTEST) -v ./...
vet:
	$(GOVET) ./...
