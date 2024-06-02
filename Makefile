build:
	@go build -o bin/go-rest cmd/main.go

test:
	@go test -v ./..
	
run: build
	@./bin/go-rest
