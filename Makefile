build:
	@go build -o bin/remitly cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/remitly