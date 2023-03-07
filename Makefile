BIN=./bin
NAME=adcli

run:
	go run .

build:
	go build -o $(BIN)/$(NAME) . 

clean:
	rm -rf $(BIN)
