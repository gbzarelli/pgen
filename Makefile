dependencies:
	go get
	go mod tidy

build:
	go build

run:
	go build
	go run main.go
