build:
	@go build -o bin/first-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/first-go