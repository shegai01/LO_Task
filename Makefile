.phony: run
PORT:=8080
run:
	go run ./cmd/app/main.go -port=8080
stop:
	@fuser -k ${PORT}/tcp
build:
	go build ./cmd/app/main.go