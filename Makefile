BIN ?= bitclock-go

test:
	go test ./...

run:
	go run ./cmd/bitclock-go

build:
	go build -o $(BIN) ./cmd/bitclock-go

clean:
	if [ -f $(BIN) ]; then rm $(BIN); fi
