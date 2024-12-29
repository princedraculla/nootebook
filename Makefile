build:
	@go build -o bin/app.go

run: build
	@./bin/app.go