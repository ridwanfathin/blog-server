BINARY=blog-server
all: test build run

test:
		go test ./...

build:
		go build -o ${BINARY} main/main.go

run:
		./${BINARY}

.PHONY: test build run