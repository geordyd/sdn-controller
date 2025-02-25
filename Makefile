run:
	go run .

build:
	go build -o bin/ .
test:
	go test -timeout 30s -v ./tests/...
