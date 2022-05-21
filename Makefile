BIN ?= bitclock-go

.PHONY : all
all: clean build

clean:
	if [ -f $(BIN) ]; then rm $(BIN); fi

build:
	go build -o $(BIN) .

test:
	go test ./...

run:
	go run .
