BIN		= ./bin
NAME	= adcli

run:
	go run ./

build:
	go build -o $(BIN)/$(NAME) ./

clean:
	rm -rf $(BIN)

fmt:
	go fmt ./...
	go vet ./...

fmt_failed:
	./scripts/fmt_failed.sh

test:
	go test

lint:
	golangci-lint run --enable-all

cover:
	go test -coverprofile coverage.out

coverweb: cover
	go tool cover -html=coverage.out

check: fmt lint cover

dep:
	go mod download
