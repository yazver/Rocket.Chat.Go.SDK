.PHONY: mod test fmt

SRCS = $(shell git ls-files '*.go' | grep -v '^vendor/')

## Add missing and remove unused modules 
mod:
	go mod tidy

## Test the Code.
test: mod gen
	go test ./...	

## Format the Code.
fmt:
	gofmt -s -l -w $(SRCS)

## Lint the Code.
lint:
	golangci-lint run -v --out-format=tab --timeout 10m0s
