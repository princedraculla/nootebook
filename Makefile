build:
	@go build -o bin/app.go

run: build
	@./bin/app.go

http:
	@go run main.go http