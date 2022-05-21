BIN ?= bitclock-go

test:
	go test ./...

run:
	go run .

build:
	go build -o $(BIN) .

clean:
	if [ -f $(BIN) ]; then rm $(BIN); fi
